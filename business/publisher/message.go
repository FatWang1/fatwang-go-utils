package pubsub

var (
	DefaultExpire = 300
)

type Message struct {
	Event     string
	Data      any
	Source    string
	TimeStamp string
	Expire    int
}

//
//type MessageBuilder struct {
//	options Message
//}
//
//func NewMessageBuilder() *MessageBuilder {
//	return &MessageBuilder{
//		options: Message{
//			Expire: DefaultExpire,
//		},
//	}
//}
//
//func (b *MessageBuilder) WithEvent(event string) *MessageBuilder {
//	b.options.Event = event
//	return b
//}
//
//func (b *MessageBuilder) WithData(data any) *MessageBuilder {
//	b.options.Data = data
//	return b
//}
//
//func (b *MessageBuilder) WithSource(source string) *MessageBuilder {
//	b.options.Source = source
//	return b
//}
//
//func (b *MessageBuilder) WithTimeStamp(timeStamp string) *MessageBuilder {
//	b.options.TimeStamp = timeStamp
//	return b
//}
//
//func (b *MessageBuilder) WithExpire(expire int) *MessageBuilder {
//	b.options.Expire = expire
//	return b
//}
//
//func (b *MessageBuilder) Build() *Message {
//	return &b.options
//}
