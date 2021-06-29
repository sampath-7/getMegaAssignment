package main

import (
	"strconv"
)

func (ps *PubSub) Publish(topicID string, message string) {
	//Publishing the message to respective subscription channels

	ps.mu.Lock()
	defer ps.mu.Unlock()
	ps.messageCounter = ps.messageCounter + 1
	messageID := strconv.Itoa(ps.messageCounter)
	ps.messageIDMapsMessage[messageID] = message
	for _, subscriptionID := range ps.topicMapsSubs[topicID] {
		ps.SubscriberFunc(subscriptionID, messageID)
		// time.Sleep(1 * time.Millisecond)
	}

}
