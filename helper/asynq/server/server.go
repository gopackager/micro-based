package server

import (
	"context"
	"fmt"
	"sync"

	"github.com/hibiken/asynq"
)

type AsynQueueServer interface {
	Handler(fn []HandlerAsynq)
	Listen() error
}

var (
	Concurent = 10
)

type engine struct {
	Server   *asynq.Server
	FNHandle *asynq.ServeMux
}

func New(host, port, password string) AsynQueueServer {
	option := asynq.RedisClientOpt{
		Addr:     fmt.Sprintf("%v:%v", host, port),
		Password: password,
		DB:       0,
	}

	config := asynq.Config{
		Concurrency:    Concurent,
		StrictPriority: true,
	}

	return &engine{
		Server:   asynq.NewServer(option, config),
		FNHandle: nil,
	}
}

type HandlerAsynq struct {
	Name string
	Fn   func(context.Context, *asynq.Task) error
}

func (e *engine) Handler(fn []HandlerAsynq) {
	var m sync.Mutex
	mux := asynq.NewServeMux()
	for _, f := range fn {
		m.Lock()
		mux.HandleFunc(f.Name, f.Fn)
		m.Unlock()
		e.FNHandle = mux
	}
}

func (e engine) Listen() error {
	return e.Server.Run(e.FNHandle)
}
