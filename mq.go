package mqttx

import (
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/wl955/log"
)

var _opts = mqtt.NewClientOptions()

var client mqtt.Client

func Init(opts ...Option) (mqtt.Client, error) {
	custom := Options{}

	for _, opt := range opts {
		opt(&custom)
	}

	_opts.SetDefaultPublishHandler(pubHandler)
	_opts.SetOnConnectHandler(connectHandler)
	_opts.SetConnectionLostHandler(connectLostHandler)

	client = mqtt.NewClient(
		_opts,
	)

	if token := client.Connect(); token.Wait() && token.Error() != nil {
		return client, token.Error()
	}

	for _, route := range routes {
		if token := client.Subscribe(route.topic, route.qos, route.callback); token.Wait() && token.Error() != nil {
			return client, token.Error()
		}
	}

	return client, nil
}

func Disconnect() {
	if client != nil {
		client.Disconnect(250)
	}
}

func Sub(topic string, qos byte, callback mqtt.MessageHandler) {
	routes = append(routes, &Route{
		topic:    topic,
		qos:      qos,
		callback: callback,
	})
}

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
