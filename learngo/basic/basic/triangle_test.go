package main

import "testing"

func TestTriangle(t *testing.T) {
	tests := []struct{ a, b, c int }{
		{3, 4, 5},
		{5, 12, 13},
		{8, 15, 17},
		{30000, 40000, 50000},
	}

	for _, tt := range tests {
		if actual := calTriangle(tt.a, tt.b); actual != tt.c {
			t.Errorf("calTriangle(%d, %d) got %d; expected %d",
				tt.a, tt.b, actual, tt.c)
		}
	}
}

func BenchmarkTriangle(b *testing.B) {
	tests := []struct{ a, b, c int }{
		{3, 4, 5},
		{5, 12, 13},
		{8, 15, 17},
		{30000, 40000, 50000},
	}

	for i := 0; i < b.N; i++ {
		for _, tt := range tests {
			if actual := calTriangle(tt.a, tt.b); actual != tt.c {
				b.Errorf("calTriangle(%d, %d) got %d; expected %d",
					tt.a, tt.b, actual, tt.c)
			}
		}
	}
}
