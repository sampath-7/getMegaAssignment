package main

import (
	"fmt"
	"strconv"
)

func (ps *PubSub) AddSubscription(topicID string, subscriptionID string) (err error) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
		}
	}()
	if _, isExist := ps.topicMapsSubs[topicID]; isExist {
		subscriptionList := ps.topicMapsSubs[topicID]
		if subscriptionList, ok := ps.topicMapsSubs[topicID]; ok {
			for _, client := range subscriptionList {
				if client == subscriptionID {
					panic("SubscriptionID already exist")
				}
			}
		}
		subscriptionList = append(subscriptionList, subscriptionID)
		ps.topicMapsSubs[topicID] = subscriptionList
		ps.subsMapsTopic[subscriptionID] = topicID
	} else {
		fmt.Println("Topic ID is not exist")
	}
	return
}

func (ps *PubSub) DeleteSubscription(subscriptionID string) {
	if _, isExist := ps.subsMapsTopic[subscriptionID]; isExist {
		topicID := ps.subsMapsTopic[subscriptionID]
		if subscriptionList, ok := ps.topicMapsSubs[topicID]; ok {
			var updateSubscriptionList []string
			for _, client := range subscriptionList {
				if client != subscriptionID {
					updateSubscriptionList = append(updateSubscriptionList, client)
				}
			}
			ps.topicMapsSubs[topicID] = updateSubscriptionList
		}

		delete(ps.subsMapsTopic, subscriptionID)
	} else {
		fmt.Println("Subscription ID doesnot exist", subscriptionID)
	}
}

func (ps *PubSub) Subscribe(subscriptionID string) (chan string, bool) {
	ps.mu.Lock()
	defer ps.mu.Unlock()
	ch := make(chan string, 1)
	state := true
	if _, isExist := ps.subsMapsTopic[subscriptionID]; !isExist {
		fmt.Println("Sorry The SubscriptionId is not exist")
		state = false
	} else {
		ps.subsMapsChannels[subscriptionID] = ch
		ps.messageCounter += 1
		messageID := strconv.Itoa(ps.messageCounter)
		ps.messageIDMapsMessage[messageID] = subscriptionID + "is successfully Subscribed"
		ps.SubscriberFunc(subscriptionID, messageID)
	}
	return ch, state
}

func (ps *PubSub) SubscriberFunc(subscriptionID string, messageID string) {
	w.Add(1)
	fmt.Println("Subscriber FUn started")
	ch := ps.subsMapsChannels[subscriptionID]
	ch <- messageID
	fmt.Println("Subscriber FUn doe", messageID)
}

func (ps *PubSub) UnSubscribe(subscriptionID string) {
	ps.mu.Lock()
	defer ps.mu.Unlock()
	ch := ps.subsMapsChannels[subscriptionID]
	if !ps.IsaClosed(ch) {
		close(ch)
	}
}
