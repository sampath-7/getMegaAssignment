package main

import (
	"fmt"
	"sync"
	"time"
)

var w sync.WaitGroup

type topicMapsSubsList map[string][]string
type subsMapsTopicList map[string][]string
type messageQueue map[string][]string
type messageIDMapsMessage map[string]string
type subsMapsChannels map[string]chan string

type PubSub struct {
	mu                   sync.RWMutex
	topicMapsSubs        topicMapsSubsList
	subsMapsChannels     subsMapsChannels
	subsMapsTopic        subsMapsTopicList
	messageIDMapsMessage messageIDMapsMessage
	messageQueue         messageQueue
	messageCounter       int
}

func (ps *PubSub) Close() {
	ps.mu.Lock()
	defer ps.mu.Unlock()
	for subscriptionID, ch := range ps.subsMapsChannels {
		if !ps.IsaClosed(ch) {
			time.Sleep(1000 * time.Millisecond)
			close(ch)
			fmt.Println(subscriptionID, " is closed")

		}
	}
}

func (ps *PubSub) IsaClosed(ch chan string) bool {
	select {
	case <-ch:
		return true
	default:
	}

	return false
}

func main() {
	fmt.Println("PubSub Service Initiated")
	initTopicMapsSubs := make(topicMapsSubsList)
	initSubsMapsTopic := make(subsMapsTopicList)
	initSubsMapsChannels := make(subsMapsChannels)
	initMessageQueue := make(messageQueue)
	initMessageIDMapsMessage := make(messageIDMapsMessage)
	initMessageCounter := 0

	ps := PubSub{
		topicMapsSubs:        initTopicMapsSubs,
		subsMapsTopic:        initSubsMapsTopic,
		subsMapsChannels:     initSubsMapsChannels,
		messageQueue:         initMessageQueue,
		messageIDMapsMessage: initMessageIDMapsMessage,
		messageCounter:       initMessageCounter,
	}
	ps.CreateTopic("tech")
	ps.AddSubscription("tech", "sub1")
	ps.AddSubscription("tech", "sub2")
	ch1 := ps.Subscribe("sub1")
	ch2 := ps.Subscribe("sub2")

	listener := func(subscriptionID string, ch <-chan string) {
		for messageID := range ch {
			ps.Ack(subscriptionID, messageID)
		}
		fmt.Printf("[%s] is sent all topic messages to its subscriptions :) \n", subscriptionID)
	}
	go listener("sub1", ch1)
	go listener("sub2", ch2)
	ps.Publish("tech", "Tablets Main")
	ps.Publish("tech", "robots")
	ps.Publish("tech", "drones")
	ps.Publish("tech", "tablets1")
	ps.Publish("tech", "robots123")
	ps.Publish("tech", "drones1244")
	time.Sleep(1 * time.Millisecond)
	w.Wait()
	ps.Close()
}
