package mq

import (
	"fmt"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/wl955/log"
)

var client mqtt.Client

//func init() {
//	opts := mqtt.NewClientOptions()
//
//	opts.AddBroker("tcp://broker.emqx.io:1883")
//
//	opts.SetClientID("go_mqtt_client1")
//
//	opts.SetDefaultPublishHandler(pubHandler)
//	opts.SetOnConnectHandler(connectHandler)
//	opts.SetConnectionLostHandler(connectLostHandler)
//
//	client = mqtt.NewClient(opts)
//
//	if token := client.Connect(); token.Wait() && token.Error() != nil {
//		panic(token.Error())
//	}
//
//	return
//}

func Setup(broker string) (mqtt.Client, error) {
	opts := mqtt.NewClientOptions()

	opts.AddBroker(fmt.Sprintf("tcp://%s", broker))

	opts.SetClientID("go_mqtt_client1")

	opts.SetDefaultPublishHandler(pubHandler)
	opts.SetOnConnectHandler(connectHandler)
	opts.SetConnectionLostHandler(connectLostHandler)

	client = mqtt.NewClient(opts)

	if token := client.Connect(); token.Wait() && token.Error() != nil {
		return client, token.Error()
	}

	return client, nil
}

//func Disconnect() {
//	if client != nil {
//		client.Disconnect(250)
//	}
//}

func Serve() {
	for _, route := range routes {
		client.Subscribe(route.topic, route.qos, route.callback)
	}
}

//func Sub(topic string, qos byte, callback mqtt.MessageHandler) mqtt.Token {
//	return client.Subscribe(topic, qos, callback)
//}

func Pub(topic string, qos byte, retained bool, payload interface{}) mqtt.Token {
	return client.Publish(topic, qos, retained, payload)
}

var pubHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	log.Infof("Received message: %s from topic: %s\n", msg.Payload(), msg.Topic())
}

var connectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
	log.Info("Connected")
}

var connectLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
	log.Infof("Connect lost: %v", err)
}
