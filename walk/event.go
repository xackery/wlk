// Copyright 2011 The Walk Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build windows
// +build windows

package walk

type eventHandlerInfo struct {
	handler EventHandler
	once    bool
}

// EventHandler is a generic function
type EventHandler func()

// Event is used to handle events
type Event struct {
	handlers []eventHandlerInfo
}

// Attach allows you to subscribe to an event
func (e *Event) Attach(handler EventHandler) int {
	handlerInfo := eventHandlerInfo{handler, false}

	for i, h := range e.handlers {
		if h.handler == nil {
			e.handlers[i] = handlerInfo
			return i
		}
	}

	e.handlers = append(e.handlers, handlerInfo)

	return len(e.handlers) - 1
}

// Detach removes a handle by id
func (e *Event) Detach(handle int) {
	e.handlers[handle].handler = nil
}

// Once is fired once, then removed
func (e *Event) Once(handler EventHandler) {
	i := e.Attach(handler)
	e.handlers[i].once = true
}

type EventPublisher struct {
	event Event
}

// Event returns the event
func (p *EventPublisher) Event() *Event {
	return &p.event
}

// Publish emits an event to all handlers
func (p *EventPublisher) Publish() {
	// This is a kludge to find the form that the event publisher is
	// affiliated with. It's only necessary because the event publisher
	// doesn't keep a pointer to the form on its own, and the call
	// to Publish isn't providing it either.
	if form := App().ActiveForm(); form != nil {
		fb := form.AsFormBase()
		fb.inProgressEventCount++
		defer func() {
			fb.inProgressEventCount--
			if fb.inProgressEventCount == 0 && fb.layoutScheduled {
				fb.layoutScheduled = false
				fb.startLayout()
			}
		}()
	}

	for i, h := range p.event.handlers {
		if h.handler != nil {
			h.handler()

			if h.once {
				p.event.Detach(i)
			}
		}
	}
}

type genericEventHandlerInfo[T any] struct {
	handler GenericEventHandler[T]
	once    bool
}

// GenericEventHandler allows a generic function to be used
type GenericEventHandler[T any] func(param T)

// GenericEvent is used to handle events
type GenericEvent[T any] struct {
	handlers []genericEventHandlerInfo[T]
}

// Attach allows you to subscribe to an event
func (e *GenericEvent[T]) Attach(handler GenericEventHandler[T]) int {
	handlerInfo := genericEventHandlerInfo[T]{handler, false}

	for i, h := range e.handlers {
		if h.handler == nil {
			e.handlers[i] = handlerInfo
			return i
		}
	}

	e.handlers = append(e.handlers, handlerInfo)

	return len(e.handlers) - 1
}

// Detach removes a handle by id
func (e *GenericEvent[T]) Detach(handle int) {
	e.handlers[handle].handler = nil
}

// Once is fired once, then removed
func (e *GenericEvent[T]) Once(handler GenericEventHandler[T]) {
	i := e.Attach(handler)
	e.handlers[i].once = true
}

// GenericEventPublisher is used to publish events
type GenericEventPublisher[T any] struct {
	event GenericEvent[T]
}

// Event returns the event
func (p *GenericEventPublisher[T]) Event() *GenericEvent[T] {
	return &p.event
}

// Publish emits an event to all handlers
func (p *GenericEventPublisher[T]) Publish(param T) {
	for i, h := range p.event.handlers {
		if h.handler != nil {
			h.handler(param)

			if h.once {
				p.event.Detach(i)
			}
		}
	}
}
