package mq

import "fmt"

type Options struct {
}

type Option func(*Options)

func AddBroker(server string) Option {
	return func(*Options) {
		_opts.AddBroker(fmt.Sprintf("tcp://%s", server))
	}
}

func ClientID(idStr string) Option {
	return func(*Options) {
		_opts.SetClientID(idStr)
	}
}
