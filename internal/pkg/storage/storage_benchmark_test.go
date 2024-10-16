package storage

import "testing"

func BenchmarkSet(b *testing.B) {
	s, err := NewStorage()
	if err != nil {
		b.Fatalf("error! achtung: %v", err)
	}
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
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s.Get("AAAA")
	}
}
func BenchmarkSetGet(b *testing.B) {
	s, err := NewStorage()
	if err != nil {
		b.Fatalf("error! achtung: %v", err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s.Set("AAAA", "AAAA")
		s.Get("AAAA")
	}
}
