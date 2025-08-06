package pubsub

import (
	"sync"
	"testing"
	"time"
)

func TestNewPublisher(t *testing.T) {
	buffer := 10
	publisher := NewPublisher(buffer)
	
	if publisher == nil {
		t.Error("NewPublisher should not return nil")
	}
	
	if publisher.buffer != buffer {
		t.Errorf("Expected buffer size %d, got %d", buffer, publisher.buffer)
	}
	
	if len(publisher.subscribers) != 0 {
		t.Errorf("Expected empty subscribers map, got %d subscribers", len(publisher.subscribers))
	}
}

func TestPublisher_Subscribe(t *testing.T) {
	publisher := NewPublisher(10)
	
	sub := publisher.Subscribe()
	
	if sub == nil {
		t.Error("Subscribe should not return nil")
	}
	
	// Check that subscriber was added to the map
	publisher.m.RLock()
	defer publisher.m.RUnlock()
	
	if len(publisher.subscribers) != 1 {
		t.Errorf("Expected 1 subscriber, got %d", len(publisher.subscribers))
	}
}

func TestPublisher_SubscribeTopic(t *testing.T) {
	publisher := NewPublisher(10)
	
	// Subscribe with a topic function
	topicFunc := func(v *Message) bool {
		return v.Event == "test"
	}
	
	sub := publisher.SubscribeTopic(topicFunc)
	
	if sub == nil {
		t.Error("SubscribeTopic should not return nil")
	}
	
	// Check that subscriber was added to the map with the correct topic function
	publisher.m.RLock()
	defer publisher.m.RUnlock()
	
	if len(publisher.subscribers) != 1 {
		t.Errorf("Expected 1 subscriber, got %d", len(publisher.subscribers))
	}
	
	// Check that the topic function was stored correctly
	if publisher.subscribers[sub] == nil {
		t.Error("Topic function should not be nil")
	}
}

func TestPublisher_Evict(t *testing.T) {
	publisher := NewPublisher(10)
	sub := publisher.Subscribe()
	
	// Ensure subscriber is in the map
	publisher.m.RLock()
	if len(publisher.subscribers) != 1 {
		t.Errorf("Expected 1 subscriber before eviction, got %d", len(publisher.subscribers))
	}
	publisher.m.RUnlock()
	
	// Evict the subscriber
	publisher.Evict(sub)
	
	// Check that subscriber was removed from the map
	publisher.m.RLock()
	defer publisher.m.RUnlock()
	
	if len(publisher.subscribers) != 0 {
		t.Errorf("Expected 0 subscribers after eviction, got %d", len(publisher.subscribers))
	}
}

func TestPublisher_Close(t *testing.T) {
	publisher := NewPublisher(10)
	
	// Add multiple subscribers
	publisher.Subscribe()
	publisher.Subscribe()
	
	// Ensure subscribers are in the map
	publisher.m.RLock()
	if len(publisher.subscribers) != 2 {
		t.Errorf("Expected 2 subscribers before closing, got %d", len(publisher.subscribers))
	}
	publisher.m.RUnlock()
	
	// Close the publisher
	publisher.Close()
	
	// Check that all subscribers were removed from the map
	publisher.m.RLock()
	defer publisher.m.RUnlock()
	
	if len(publisher.subscribers) != 0 {
		t.Errorf("Expected 0 subscribers after closing, got %d", len(publisher.subscribers))
	}
}

func TestPublisher_Publish(t *testing.T) {
	publisher := NewPublisher(10)
	
	// Create a subscriber
	sub := publisher.Subscribe()
	
	// Create a message
	message := &Message{
		Event:     "test",
		Data:      "test data",
		Source:    "test source",
		TimeStamp: "2022-01-01T00:00:00Z",
		Expire:    300,
	}
	
	// Publish the message in a goroutine to avoid blocking
	go publisher.Publish(message)
	
	// Wait for the message to be received
	select {
	case receivedMsg := <-sub:
		if receivedMsg.Event != message.Event {
			t.Errorf("Expected event %s, got %s", message.Event, receivedMsg.Event)
		}
		if receivedMsg.Data != message.Data {
			t.Errorf("Expected data %s, got %s", message.Data, receivedMsg.Data)
		}
	case <-time.After(1 * time.Second):
		t.Error("Timeout waiting for message")
	}
}

func TestPublisher_SendTopic(t *testing.T) {
	publisher := NewPublisher(10)
	
	// Create a topic function
	topicFunc := func(v *Message) bool {
		return v.Event == "test"
	}
	
	// Subscribe with the topic function
	sub := publisher.SubscribeTopic(topicFunc)
	
	// Create a message that matches the topic
	message := &Message{
		Event:     "test",
		Data:      "test data",
		Source:    "test source",
		TimeStamp: "2022-01-01T00:00:00Z",
		Expire:    300,
	}
	
	// Send the message in a goroutine to avoid blocking
	var wg sync.WaitGroup
	wg.Add(1)
	go publisher.SendTopic(sub, topicFunc, message, &wg)
	
	// Wait for the message to be received
	select {
	case receivedMsg := <-sub:
		if receivedMsg.Event != message.Event {
			t.Errorf("Expected event %s, got %s", message.Event, receivedMsg.Event)
		}
	case <-time.After(1 * time.Second):
		t.Error("Timeout waiting for message")
	}
	
	// Create a message that doesn't match the topic
	message2 := &Message{
		Event:     "other",
		Data:      "other data",
		Source:    "other source",
		TimeStamp: "2022-01-01T00:00:00Z",
		Expire:    300,
	}
	
	// Send the message that doesn't match the topic
	wg.Add(1)
	go publisher.SendTopic(sub, topicFunc, message2, &wg)
	
	// Should not receive this message
	select {
	case <-sub:
		t.Error("Should not receive message that doesn't match topic")
	case <-time.After(100 * time.Millisecond):
		// Correctly did not receive message
	}
}

func TestPublisher_SendTopicWithTimeout(t *testing.T) {
	publisher := NewPublisher(0) // Buffer size 0 to force blocking
	
	// Subscribe with a topic function
	topicFunc := func(v *Message) bool {
		return v.Event == "test"
	}
	
	// Subscribe to the topic
	sub := publisher.SubscribeTopic(topicFunc)
	
	// Fill the subscriber channel to make it block
	message := &Message{
		Event:     "test",
		Data:      "test data",
		Source:    "test source",
		TimeStamp: "2022-01-01T00:00:00Z",
		Expire:    1, // 1 second timeout
	}
	
	// Send a message to fill the buffer
	var wg sync.WaitGroup
	wg.Add(1)
	go publisher.SendTopic(sub, topicFunc, message, &wg)
	
	// Give some time for the first message to be sent
	time.Sleep(100 * time.Millisecond)
	
	// Try to send another message which should timeout
	wg.Add(1)
	go publisher.SendTopic(sub, topicFunc, message, &wg)
	
	// Wait for the goroutines to finish
	done := make(chan struct{})
	go func() {
		wg.Wait()
		close(done)
	}()
	
	select {
	case <-done:
		// Success, both goroutines finished
	case <-time.After(2 * time.Second):
		t.Error("Timeout waiting for SendTopic to complete")
	}
}