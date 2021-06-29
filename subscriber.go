package main

import (
	"fmt"
)

func (ps *PubSub) Ack(subscriptionID string, messageID string) bool {
	//Its acknowlegdement sent back to Pubsub
	w.Done()
	messageData := ps.messageIDMapsMessage[messageID]
	fmt.Println("For Subscription ID:", subscriptionID, "got", messageID, " message :", messageData)
	return true
}
