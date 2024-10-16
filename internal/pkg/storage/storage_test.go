package storage

import (
	"strconv"
	"testing"
)

type testCase1 struct {
	name  string
	key   string
	value string
}

type testCase2 struct {
	name  string
	key   string
	value string
	Type  string
}

var testcases1 = []testCase1{
	{"hello world", "hello", "world"},
	{"Welcome back", "Welcome", "back"},
	{"228 1488", "228", "1488"},
	{"AAAA AAAA", "AAAA", "AAAA"},
}

var testcases2 = []testCase2{
	{"test1", "hello", "world", "S"},
	{"test2", "Welcome", "back", "S"},
	{"test3", "228", "1488", "D"},
}

func TestGet(t *testing.T) {
	s, err := NewStorage()
	if err != nil {
		t.Errorf("new storage: %v", err)
	}

	for _, c := range testcases1 {
		t.Run(c.name, func(t *testing.T) {
			s.Set(c.key, c.value)

			sValue := *s.Get(c.key)
			switch s.GetKind(c.key) {
			case "D":
				{
					var new_c, _ = strconv.Atoi(c.value)
					if sValue.d != new_c {
						t.Errorf("values not equal")
					}
				}
			case "S":
				{
					if sValue.s != c.value {
						t.Errorf("values not equal")
					}
				}
			}
		})
	}
}
func TestGetKind(t *testing.T) {
	s, err := NewStorage()
	if err != nil {
		t.Errorf("new storage: %v", err)
	}

	for _, c := range testcases2 {
		t.Run(c.name, func(t *testing.T) {
			s.Set(c.key, c.value)

			sValue := *s.Get(c.key)
			if sValue.ValueType != c.Type {
				t.Errorf("value types not equal")
			}
		})
	}
}
