package storage

import (
	"strconv"

	"go.uber.org/zap"
)

type Value struct {
	StringField string
	IntField    int
	ValueType   string
}

type Storage struct {
	inner  map[string]Value
	logger *zap.Logger
}

func NewStorage() (*Storage, error) {
	logger, err := zap.NewProduction()
	if err != nil {
		return nil, err
	}

	defer logger.Sync()
	logger.Info("created new storage")

	return &Storage{
		inner:  make(map[string]Value),
		logger: logger,
	}, nil
}

func (r Storage) Set(key string, value string) {
	newValue, err := strconv.Atoi(value)
	var x Value
	if err == nil {
		x = Value{
			IntField:  newValue,
			ValueType: "D",
		}
	} else {
		x = Value{
			StringField: value,
			ValueType:   "S",
		}
	}
	r.inner[key] = x

	r.logger.Info("key set")
	r.logger.Sync()
}

func (r *Storage) Get(key string) *Value {
	res, ok := r.inner[key]
	if !ok {
		return nil
	}
	return &res
}

func (r *Storage) GetKind(key string) string {
	var current = r.Get(key)
	if current == nil {
		r.logger.Info("no such key in the storage")
		r.logger.Sync()
		return "no such key in storage"
	}
	return (*current).ValueType
}

func (r *Storage) GetType(key string) string {
	res, ok := r.inner[key]
	if !ok {
		r.logger.Info("no such key in the storage")
		r.logger.Sync()
	}
	return res.ValueType
}
