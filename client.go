package main

import (
  MQTT "github.com/eclipse/paho.mqtt.golang"
  "os"
  "fmt"
)
//for when a message received doesn't match any known subscriptions
var f MQTT.MessageHandler = func(client MQTT.Client, msg MQTT.Message) {
  fmt.Printf("TOPIC: %s\n", msg.Topic())
  fmt.Printf("MSG: %s\n", msg.Payload())
}

type MQTTClient struct{
  mqttObj MQTT.Client
  brokerURL string
  clientID string
}
func (m *MQTTClient) ParseCommandLineArguments() {
  m.brokerURL = os.Args[1]
  m.clientID = os.Args[2]
}

func ConnectToBroker(clientWithOpts MQTT.Client) bool {
  if token := clientWithOpts.Connect(); token.Wait() && token.Error() != nil {
    panic(token.Error())
  }
  return clientWithOpts.IsConnected()
}

func (m MQTTClient) setupClient() MQTT.Client {
  m.ParseCommandLineArguments()
  fmt.Println("BROKER: ", m.brokerURL)
  fmt.Println("CLIENTID: ", m.clientID)
  opts := MQTT.NewClientOptions().AddBroker(m.brokerURL).SetClientID(m.clientID).SetDefaultPublishHandler(f)
  client := MQTT.NewClient(opts)
  fmt.Println("Is Connected? : " , ConnectToBroker(client))
  client.Disconnect(250)
  fmt.Println("Is Connected? : " , client.IsConnected())
  return MQTT.NewClient(opts)
}

func main(){
  newClient := MQTTClient{}
  newClient.setupClient()
}

  
