package di

import (
	"context"
	"fmt"
	"sync"
)

type Scope int

const (
	Singleton Scope = iota
	Scoped
)

type contextKey int

const containerKey contextKey = 1

type DepFactoryFunc func(c Container) (any, error)

type tempValue = chan struct{}

// Container is the port to external services
type Container interface {
	AddSingleton(key string, fn DepFactoryFunc)
	AddScoped(key string, fn DepFactoryFunc)
	Scoped(ctx context.Context) context.Context
	Get(key string) any
}

type depInfo struct {
	key     string
	scope   Scope
	factory DepFactoryFunc
}

var _ Container = (*container)(nil)

type container struct {
	parent  *container
	deps    map[string]depInfo
	vals    map[string]any
	tracked tracked
	mu      sync.Mutex
}

func New() Container {
	return &container{
		deps: make(map[string]depInfo),
		vals: make(map[string]any),
	}
}

func (c *container) AddSingleton(key string, fn DepFactoryFunc) {
	c.deps[key] = depInfo{
		key:     key,
		scope:   Singleton,
		factory: fn,
	}
}

func (c *container) AddScoped(key string, fn DepFactoryFunc) {
	c.deps[key] = depInfo{
		key:     key,
		scope:   Scoped,
		factory: fn,
	}
}

func (c *container) Scoped(ctx context.Context) context.Context {
	return context.WithValue(ctx, containerKey, c.scoped())
}

func (c *container) Get(key string) any {
	info, exits := c.deps[key]
	if !exits {
		panic(fmt.Sprintf("there is no dependency registered with `%s`", key))
	}

	if _, exists := c.tracked[info.key]; exists {
		panic(fmt.Sprintf("cyclic dependency encountered while building `%s`, tracked: %s", info.key, c.tracked))
	}

	if info.scope == Singleton {
		return c.getFromParent(info)
	}
	return c.get(info)
}

func (c *container) getFromParent(info depInfo) any {
	if c.parent != nil {
		return c.parent.getFromParent(info)
	}
	return c.get(info)
}

func (c *container) get(info depInfo) any {
	c.mu.Lock()

	v, exists := c.vals[info.key]
	if !exists {
		tv := make(tempValue)
		c.vals[info.key] = tv
		c.mu.Unlock()
		return c.build(info, tv)
	}

	c.mu.Unlock()
	tv, isTemp := v.(tempValue)
	if !isTemp {
		return v
	}
	<-tv

	return c.get(info)
}

func (c *container) build(info depInfo, tv tempValue) any {
	v, err := info.factory(c.builder(info))

	c.mu.Lock()

	if err != nil {
		delete(c.vals, info.key)
		c.mu.Unlock()
		close(tv)
		panic(fmt.Sprintf("error building dependency `%s`: %s", info.key, err))
	}
	c.vals[info.key] = v
	c.mu.Unlock()
	close(tv)
	return v
}

func (c *container) builder(info depInfo) *container {
	return &container{
		parent:  c.parent,
		deps:    c.deps,
		vals:    c.vals,
		tracked: c.tracked.add(info),
	}
}
func (c *container) scoped() *container {
	return &container{
		parent: c,
		deps:   c.deps,
		vals:   make(map[string]any),
	}
}
