package pubsub

import (
	"sync"
	"time"
)

type (
	subscriber chan *Message         // 订阅者为一个通道
	topicFunc  func(v *Message) bool // 主题为一个过滤器
)

type Publisher struct {
	m           sync.RWMutex             // 读写锁
	buffer      int                      // 订阅队列的缓存长度
	subscribers map[subscriber]topicFunc // 订阅者信息
}

// NewPublisher 构建发送者对象
func NewPublisher(buffer int) *Publisher {
	return &Publisher{
		buffer:      buffer,
		subscribers: make(map[subscriber]topicFunc),
	}
}

// Subscribe 订阅全部主题
func (p *Publisher) Subscribe() chan *Message {
	return p.SubscribeTopic(nil)
}

// SubscribeTopic 订阅某一主题
func (p *Publisher) SubscribeTopic(topic topicFunc) chan *Message {
	ch := make(chan *Message, p.buffer)
	p.m.Lock()
	defer p.m.Unlock()
	p.subscribers[ch] = topic
	return ch
}

// Evict 退出主题
func (p *Publisher) Evict(sub chan *Message) {
	p.m.Lock()
	defer p.m.Unlock()
	if _, exists := p.subscribers[sub]; exists {
		delete(p.subscribers, sub)
		// 使用select避免重复关闭channel
		select {
		case <-sub:
			// channel已经关闭
		default:
			close(sub)
		}
	}
}

// Close 关闭所有订阅渠道
func (p *Publisher) Close() {
	p.m.Lock()
	defer p.m.Unlock()
	for sub := range p.subscribers {
		delete(p.subscribers, sub)
		// 使用select避免重复关闭channel
		select {
		case <-sub:
			// channel已经关闭
		default:
			close(sub)
		}
	}
}

// Publish 向所有满足条件的主题发送消息
func (p *Publisher) Publish(v *Message) {
	p.m.Lock()
	defer p.m.Unlock()
	var wg sync.WaitGroup
	for sub, topic := range p.subscribers {
		wg.Add(1)
		go p.SendTopic(sub, topic, v, &wg)
	}
	wg.Wait()
}

// SendTopic 向某一主题发送消息
func (p *Publisher) SendTopic(sub subscriber, topic topicFunc, v *Message, wg *sync.WaitGroup) {
	defer wg.Done()
	if topic != nil && !topic(v) {
		return
	}
	select {
	case sub <- v:
	case <-time.After(time.Duration(v.Expire) * time.Second):
		return
	}
}
