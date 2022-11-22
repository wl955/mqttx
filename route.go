package mq

import mqtt "github.com/eclipse/paho.mqtt.golang"

type Route struct {
	topic    string
	qos      byte
	callback mqtt.MessageHandler
}

var routes []*Route

func Sub(route *Route) {
	routes = append(routes, route)
}
