package storage

import (
	"strconv"

	"go.uber.org/zap"
)

type Value struct {
	s         string
	d         int
	ValueType string
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
	TypeOfvalue := r.GetKind(value)
	var x Value
	switch TypeOfvalue {
	case "D":
		{
			newValue, _ := strconv.Atoi(value)
			x = Value{
				d:         newValue,
				ValueType: "int",
			}
		}
	case "S":
		{
			x = Value{
				s:         value,
				ValueType: "string",
			}
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
	if _, err := strconv.Atoi(key); err == nil {
		return "D"
	}
	return "S"
}

func (r *Storage) GetType(key string) string {
	res, ok := r.inner[key]
	if !ok {
		r.logger.Info("no such key in the storage")
		r.logger.Sync()
	}
	return res.ValueType
}
