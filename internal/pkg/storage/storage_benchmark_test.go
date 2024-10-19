package storage

import (
	"testing"

	"go.uber.org/zap"
)

func BenchmarkSet(b *testing.B) {
	s, err := NewStorage()
	if err != nil {
		b.Fatalf("error! achtung: %v", err)
	}
	s.logger = zap.NewNop()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s.Set("AAAA", "AAAA")
	}
}

func BenchmarkGet(b *testing.B) {
	s, err := NewStorage()
	if err != nil {
		b.Fatalf("error! achtung: %v", err)
	}
	s.Set("AAAA", "AAAA")
	s.logger = zap.NewNop()
	var curvalue *Value
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		curvalue = s.Get("AAAA")
	}
	if curvalue.StringField == "1" {
		b.Fatalf("error! achtung: %v", err)
	}
}
func BenchmarkSetGet(b *testing.B) {
	s, err := NewStorage()
	if err != nil {
		b.Fatalf("error! achtung: %v", err)
	}
	s.logger = zap.NewNop()
	var curvalue *Value
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s.Set("AAAA", "AAAA")
		curvalue = s.Get("AAAA")
	}
	if curvalue.StringField == "1" {
		b.Fatalf("error! achtung: %v", err)
	}
}
