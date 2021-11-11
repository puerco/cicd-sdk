package secret

import (
	"github.com/sirupsen/logrus"
)

type Value map[string]interface{}

func (val *Value) GetString(key string) string {
	entry, ok := (*val)[key]
	if !ok {
		return ""
	}
	return entry.(string)
}

func (val *Value) Contains(key string) bool {
	_, ok := (*val)[key]
	return ok
}

func NewManager() *Manager {
	return NewManagerWithOptions(&Options{})
}

func NewManagerWithOptions(opts *Options) *Manager {
	m := &Manager{
		opts: opts,
	}

	impl, err := NewVaultManager()
	if err == nil {
		m.impl = impl
	} else {
		logrus.Fatal(err)
	}

	return m
}

type Manager struct {
	impl ManagerImplementation
	opts *Options
}

func (m *Manager) Get(key string) (*Secret, error) {
	return m.impl.Get(key)
}

func (m *Manager) Write(key string, val Value) error {
	return m.impl.Write(key, val)
}

type Options struct {
	SystemOptions interface{}
}

type Secret struct {
	Key   string
	Value Value
}

type ManagerImplementation interface {
	Get(string) (*Secret, error)
	Write(string, Value) error
	List() ([]string, error)
}
