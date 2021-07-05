package main

import (
	"fmt"
	"sync"
	"time"
)

var w sync.WaitGroup

type topicMapsSubsList map[string][]string
type subsMapsTopicList map[string]string
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

func PubSubFunc() *PubSub {
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
	return &ps
}

func main() {
	fmt.Println("PubSub Service Initiated")
	ps := PubSubFunc()
	ps.CreateTopic("tech")
	ps.AddSubscription("tech", "sub1")
	ps.AddSubscription("tech", "sub2")
	ch1, ch1_status := ps.Subscribe("sub1")
	ch2, ch2_status := ps.Subscribe("sub2")
	if ch1_status {
		go ps.Ack("sub1", ch1)
	}
	if ch2_status {
		go ps.Ack("sub2", ch2)
	}
	ps.Publish("tech", "Tablets Main")
	ps.Publish("tech", "robots")
	ps.Publish("tech", "drones")
	ps.Publish("tech", "tablets1")
	ps.Publish("tech", "robots123")
	ps.Publish("tech", "drones1244")
	time.Sleep(100 * time.Millisecond)
	w.Wait()
	ps.Close()

}
