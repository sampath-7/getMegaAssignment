package main

import (
	"log"
	"testing"
)

func TestWrongSubscribe(t *testing.T) {
	//Negative case: trying to subscribe non existing subscription ID
	ps := PubSubFunc()
	ps.CreateTopic("tech")
	ps.AddSubscription("tech", "sub1")
	ch1, ch1_status := ps.Subscribe("sub2")
	t.Log(ch1_status)
	if ch1_status {
		ps.Publish("tech", "test_channel")
		go ps.Ack("sub2", ch1)
		messageID, ok := <-ch1
		if ok {
			t.Log("Successfully retrieved the message from channel", ps.messageIDMapsMessage[messageID])
		}
	} else {
		t.Log("As subscription ID not exist")
	}
}

func TestSubscribe(t *testing.T) {
	//Testing the subscribe function and sending message and receiving the message
	ps := PubSubFunc()
	ps.CreateTopic("tech")
	ps.AddSubscription("tech", "sub1")
	message := "test_channel"
	is_message_found := false
	ch1, _ := ps.Subscribe("sub1")

	go ps.Ack("sub1", ch1)
	t.Log("Publising the 'test_channel' as message")
	ps.Publish("tech", message)
	for _, v := range ps.messageIDMapsMessage {
		if v == message {
			is_message_found = true
		}
	}
	if is_message_found {
		t.Log("Successfully message sent to Subscriber")
	} else {
		t.Error("Successfully message sent to Subscriber")
	}
}

func TestAddSubscriptions(t *testing.T) {
	ps := PubSubFunc()
	topicID := "topic_1"
	subscriptionID := "sub_1"
	ps.CreateTopic(topicID)
	ps.AddSubscription(topicID, subscriptionID)
	if subscriptionList, ok := ps.topicMapsSubs[topicID]; ok {
		for _, client := range subscriptionList {
			if client == subscriptionID {
				t.Log("Added Subccessfully")
			}
		}
	}
	//negative case: adding same subscriptionID
	err := ps.AddSubscription(topicID, subscriptionID)
	if err == nil {
		t.Log("Trying to add same subscriptionID")
	}
}

func TestDeleteSubscriptions(t *testing.T) {
	ps := PubSubFunc()
	subscriptionID := "sub_1"
	//Delete the existing subscriptionID
	ps.CreateTopic("tech")
	ps.AddSubscription("tech", subscriptionID)
	ps.DeleteSubscription(subscriptionID)
	log.Println("Succefully deleted a subscription ")
	//Deleting the non existing subscriptionID
	wrongSubscriptionID := "wrong_sub1"
	ps.DeleteSubscription(wrongSubscriptionID)
	log.Println(" a wrong subscription not present in list")
}

func TestDeleteTopic(t *testing.T) {
	ps := PubSubFunc()
	topicID := "topic_1"
	ps.CreateTopic(topicID)
	ps.DeleteTopic(topicID)
	//checking for positive case
	if _, isExist := ps.topicMapsSubs[topicID]; isExist {
		t.Error(" Error found while deleting topic ")
	} else {
		t.Log(" Succefully deleted a topic ")
	}
	// checking for negative case
	ps.DeleteTopic(topicID)
	t.Log(" No error deleting for same topicId")
}

func TestCreateTopic(t *testing.T) {
	ps := PubSubFunc()
	topicID := "topic_1"
	ps.CreateTopic(topicID)
	//Checking for Positive case
	if _, isExist := ps.topicMapsSubs[topicID]; isExist {
		t.Log("Added topicID successfully")
	} else {
		t.Error(" Error found while creating a topic ")
	}
	//Checking for Negative case, adding same topicID
	ps.CreateTopic(topicID)
	t.Log("No error raised while adding same topicID")
}
