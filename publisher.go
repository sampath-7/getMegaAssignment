package main

import (
	"fmt"
	"strconv"
)

func (ps *PubSub) Publish(topicID string, message string) {
	// fmt.Println("Publish Method Entered", topicID, ps.topicMapsSubs
	ps.mu.Lock()
	defer ps.mu.Unlock()
	ps.messageCounter = ps.messageCounter + 1
	messageID := strconv.Itoa(ps.messageCounter)
	ps.messageIDMapsMessage[messageID] = message
	for _, subscriptionID := range ps.topicMapsSubs[topicID] {
		fmt.Println(subscriptionID)
		ps.SubscriberFunc(subscriptionID, messageID)
		fmt.Println(subscriptionID, "DONEEEE")
		// time.Sleep(1 * time.Millisecond)
	}

}
