package main

import (
	"fmt"
)

func (ps *PubSub) Ack(subscriptionID string, ch <-chan string) {
	listener := func(subscriptionID string, ch <-chan string) {
		for messageID := range ch {
			messageData := ps.messageIDMapsMessage[messageID]
			fmt.Println("For Subscription ID:", subscriptionID, "got", messageID, " message :", messageData)
			w.Done()
		}

		fmt.Printf("[%s] is sent all topic messages to its subscriptions :) \n", subscriptionID)
	}
	listener(subscriptionID, ch)
}
