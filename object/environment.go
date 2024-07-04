package object

import "sync"

type Environment struct {
	mu sync.RWMutex
	store map[string]Object
	outer *Environment
}

func NewEnvironment() *Environment {
	s := make(map[string]Object)
	return &Environment{store: s, outer: nil}
}

func NewEnclosedEnvironment(outer *Environment) *Environment {
	env := NewEnvironment()
	env.outer = outer
	return env
}

func (e *Environment) Get(name string) (Object, bool) {
	e.mu.RLock()
	defer e.mu.RUnlock()
	obj, ok := e.store[name]
	if !ok && e.outer != nil {
		obj, ok = e.outer.Get(name)
	}
	return obj, ok
}

func (e *Environment) Set(name string, val Object) bool {
	e.mu.Lock()
	defer e.mu.Unlock()
	e.store[name] = val
	return true
}