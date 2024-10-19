package storage

import (
	"strconv"
	"testing"
)

type testCase1 struct {
	key   string
	value string
}

type testCase2 struct {
	key   string
	value string
	Type  string
}

var testcases1 = []testCase1{
	{"hello", "world"},
	{"Welcome", "back"},
	{"228", "1488"},
	{"AAAA", "AAAA"},
}

var testcases2 = []testCase2{
	{"hello", "world", "S"},
	{"Welcome", "back", "S"},
	{"228", "1488", "D"},
}

func TestGet(t *testing.T) {
	s, err := NewStorage()
	if err != nil {
		t.Errorf("new storage: %v", err)
	}

	for _, c := range testcases1 {
		t.Run(c.key, func(t *testing.T) {
			s.Set(c.key, c.value)

			sValue := *s.Get(c.key)
			switch s.GetKind(c.key) {
			case "D":
				{
					var new_c, _ = strconv.Atoi(c.value)
					if sValue.IntField != new_c {
						t.Errorf("values not equal")
					}
				}
			case "S":
				{
					if sValue.StringField != c.value {
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
		t.Run(c.key, func(t *testing.T) {
			s.Set(c.key, c.value)

			sValue := *s.Get(c.key)
			if sValue.ValueType != c.Type {
				t.Errorf("value types not equal")
			}
		})
	}
}
