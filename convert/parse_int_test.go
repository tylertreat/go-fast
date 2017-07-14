package parsing

import (
	"strconv"
	"testing"
)

var (
	n0                  = []byte("0")
	n8                  = []byte("8")
	n65                 = []byte("65")
	n1024               = []byte("1024")
	n1000000            = []byte("1000000")
	n1212121212         = []byte("1212121212")
	n998877665544332211 = []byte("998877665544332211")
)

func BenchmarkParseInt(b *testing.B) {

	b.Run("Value = 0", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_, _ = ParseInt64(n0)
		}
	})

	b.Run("Value = 8", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_, _ = ParseInt64(n8)
		}
	})

	b.Run("Value = 65", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_, _ = ParseInt64(n65)
		}
	})

	b.Run("Value = 1024", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_, _ = ParseInt64(n1024)
		}
	})

	b.Run("Value = 1000000", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_, _ = ParseInt64(n1000000)
		}
	})

	b.Run("Value = 1212121212", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_, _ = ParseInt64(n1212121212)
		}
	})

	b.Run("Value = 998877665544332211", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_, _ = ParseInt64(n998877665544332211)
		}
	})
}

func BenchmarkParseIntStrconv(b *testing.B) {

	b.Run("Value = 0", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_, _ = strconv.ParseInt(string(n0), 10, 64)
		}
	})

	b.Run("Value = 8", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_, _ = strconv.ParseInt(string(n8), 10, 64)
		}
	})

	b.Run("Value = 65", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_, _ = strconv.ParseInt(string(n65), 10, 64)
		}
	})

	b.Run("Value = 1024", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_, _ = strconv.ParseInt(string(n1024), 10, 64)
		}
	})

	b.Run("Value = 1000000", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_, _ = strconv.ParseInt(string(n1000000), 10, 64)
		}
	})

	b.Run("Value = 1212121212", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_, _ = strconv.ParseInt(string(n1212121212), 10, 64)
		}
	})

	b.Run("Value = 998877665544332211", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_, _ = strconv.ParseInt(string(n998877665544332211), 10, 64)
		}
	})
}
