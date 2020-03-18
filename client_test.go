package main

import (
  "testing"
  "os"
)

  func TestCommandLine(t *testing.T) {
    t.Run("should set clientID and Broker URL attributes based on command line arguments", func(t *testing.T) {
      URL  := "tcp://localhost:61616"
      clientID := "sample_client_id"

      oldArgs := os.Args
      defer func() { os.Args = oldArgs }()

      arguments := []string{"cmd",URL, clientID}
      os.Args = arguments
      newClient := MQTTClient{}
      newClient.mqttObj = newClient.setupClient()
      got := []string {newClient.brokerURL, newClient.clientID}
      want := []string {URL, clientID} 

      if reflect.DeepCompare(got,want) {
        t.Errorf("got %v want %v", got, want)
      }
    })
  }

//should use those command-line arguments to subscribe to MQTT Topic
//should throw error if command-line arguments are incorrect
//should be able to send messages to subscribed Topic
//should receive messages to subscribed topic
//should be able to successfully disconnect

