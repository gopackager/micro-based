package client

import (
	"errors"
	"fmt"
	"time"

	"github.com/gopackager/micro-based/helper"
	"github.com/hibiken/asynq"
)

type AsynQueueClient interface {
	NewTask(name string, params []byte) error
	Priority(name string)
	Delay(t time.Duration, name string, params []byte) error
	Flush() (*asynq.TaskInfo, error)
	Close() error
}

type engine struct {
	Client  *asynq.Client
	Task    *asynq.Task
	Options asynq.Option
}

func New(host, port, password string) AsynQueueClient {
	client := asynq.NewClient(asynq.RedisClientOpt{
		Addr:     fmt.Sprintf("%v:%v", host, port),
		Password: password,
		DB:       0,
	})
	return &engine{
		Client:  client,
		Task:    &asynq.Task{},
		Options: nil,
	}
}

const defaultTaskName = "queue:asynq"
const defaultPrior = "default"
const defaultDelay = 5 * time.Minute
const defaultUniqueTime = 30 * time.Second

func (e engine) Close() error {
	return e.Client.Close()
}

func (e *engine) NewTask(name string, params []byte) error {
	if name == "" {
		name = defaultTaskName
	}
	if params == nil {
		return errors.New("params is empty")
	}
	e.Task = asynq.NewTask(name, params)
	return nil
}

func (e *engine) Priority(prior string) {
	if prior == "" {
		prior = defaultPrior
	}

	if !helper.InArray(prior, []string{"default", "low", "high"}) {
		prior = defaultPrior
	}

	e.Options = asynq.Queue(prior)
}

func (e *engine) Delay(t time.Duration, name string, params []byte) error {
	if t == 0 {
		t = defaultDelay
	}
	if err := e.NewTask(name, params); err != nil {
		return err
	}
	e.Options = asynq.ProcessIn(t)
	return nil
}

func (e *engine) TaskUnique(ttl time.Duration, name string, params []byte) error {
	var err error
	er := e.NewTask(name, params)
	switch {
	case errors.Is(err, asynq.ErrDuplicateTask):
		err = er
	default:
		err = er
	}
	if err != nil {
		return err
	}
	if ttl == 0 {
		ttl = defaultUniqueTime
	}
	e.Options = asynq.Unique(ttl)
	return nil
}

func (e engine) Flush() (*asynq.TaskInfo, error) {
	return e.Client.Enqueue(e.Task, e.Options)
}
