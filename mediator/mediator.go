package mediator

import (
	"context"
	"fmt"
	"reflect"
	"sync"
)

var lock = &sync.Mutex{}
var singletonMediator = &Mediator{}

func GetInstance() *Mediator {
	if singletonMediator == nil {
		lock.Lock()
		defer lock.Unlock()
		if singletonMediator == nil {
			singletonMediator = &Mediator{}
		}
	}
	return singletonMediator
}

type Mediator struct {
	handlers map[string][]NotificationHandler
}

func (m *Mediator) Register(notification interface{}, handler NotificationHandler) {
	if m.handlers == nil {
		m.handlers = make(map[string][]NotificationHandler)
	}
	typ := reflect.TypeOf(notification)
	m.handlers[typ.Name()] = append(m.handlers[typ.Name()], handler)
}

func (m *Mediator) Notify(ctx context.Context, notification interface{}) error {
	typ := reflect.TypeOf(notification)
	typeName := typ.Name()
	for _, h := range m.handlers[typeName] {
		fmt.Printf("(MEDIATOR NOTIFICATION) %v: %v\n", typeName, notification)
		h(ctx, notification)
	}
	return nil
}

type NotificationHandler func(context.Context, interface{}) error
