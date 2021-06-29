package main

import "fmt"

func (ps *PubSub) CreateTopic(topicID string) {
	//create the topic in map
	var subscriptionList []string
	if _, isExist := ps.topicMapsSubs[topicID]; isExist {
		fmt.Println("Sorry The Topic Id is already exist")
	} else {
		ps.topicMapsSubs[topicID] = subscriptionList
	}
}

func (ps *PubSub) DeleteTopic(topicID string) {
	// delete the topic from map
	if _, isExist := ps.topicMapsSubs[topicID]; isExist {
		delete(ps.topicMapsSubs, topicID)
	} else {
		fmt.Println("Sorry, Topic ID doesnot exist ")
	}
}
