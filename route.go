package mq

import mqtt "github.com/eclipse/paho.mqtt.golang"

type Route struct {
	topic    string
	qos      byte
	callback mqtt.MessageHandler
}

var routes []*Route

func Sub(topic string, qos byte, callback mqtt.MessageHandler) {
	routes = append(routes, &Route{
		topic:    topic,
		qos:      qos,
		callback: callback,
	})
}
