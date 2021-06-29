package main

import (
	"fmt"
	"strconv"
)

func (ps *PubSub) AddSubscription(topicID string, subscriptionID string) {
	//Add subscription to the topicID if not exist in map
	if _, isExist := ps.topicMapsSubs[topicID]; isExist {
		subscriptionList := ps.topicMapsSubs[topicID]
		subscriptionList = append(subscriptionList, subscriptionID)
		ps.topicMapsSubs[topicID] = subscriptionList
	} else {
		fmt.Println("Topic ID is not exist")
	}

}

func (ps *PubSub) DeleteSubscription(subscriptionID string) {
	//Delete subscription to the topicID List and subscription Key from subscription map
	if _, isExist := ps.subsMapsTopic[subscriptionID]; isExist {
		for _, topicID := range ps.subsMapsTopic[subscriptionID] {
			if subscriptionList, ok := ps.topicMapsSubs[topicID]; ok {
				//remove one client chan in chan List
				var updateSubscriptionList []string
				for _, client := range subscriptionList {
					if client != subscriptionID {
						updateSubscriptionList = append(updateSubscriptionList, client)
					}
				}
				ps.topicMapsSubs[topicID] = updateSubscriptionList
			}
		}

		delete(ps.subsMapsTopic, subscriptionID)
	} else {
		fmt.Println("Subscription ID doesnot exist", subscriptionID)
	}
}

func (ps *PubSub) Subscribe(subscriptionID string) chan string {
	//create a channel to receive the topic messages
	ps.mu.Lock()
	defer ps.mu.Unlock()
	ch := make(chan string, 1)
	ps.subsMapsChannels[subscriptionID] = ch
	ps.messageCounter += 1
	messageID := strconv.Itoa(ps.messageCounter)
	ps.messageIDMapsMessage[messageID] = subscriptionID + "is successfully Subscribed"
	ps.SubscriberFunc(subscriptionID, messageID)
	return ch
}

func (ps *PubSub) SubscriberFunc(subscriptionID string, messageID string) {
	//add message to the channel
	w.Add(1)
	ch := ps.subsMapsChannels[subscriptionID]
	ch <- messageID
}

func (ps *PubSub) UnSubscribe(subscriptionID string) {
	//closing the channel incase it is opened
	ps.mu.Lock()
	defer ps.mu.Unlock()
	ch := ps.subsMapsChannels[subscriptionID]
	if !ps.IsaClosed(ch) {
		close(ch)
	}
}
