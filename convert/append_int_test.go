package parsing

import (
	"strconv"
	"testing"
)

const (
	_PUB_P_ = "PUB "
)

var scratch [512]byte

func BenchmarkOldAppendInt(b *testing.B) {

	b.Run("Value = 0", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			msg := scratch[:len(_PUB_P_)]
			msg = OldAppendInt(msg, 0)
		}
	})

	b.Run("Value = 8", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			msg := scratch[:len(_PUB_P_)]
			msg = OldAppendInt(msg, 8)
		}
	})

	b.Run("Value = 65", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			msg := scratch[:len(_PUB_P_)]
			msg = OldAppendInt(msg, 65)
		}
	})

	b.Run("Value = 1024", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			msg := scratch[:len(_PUB_P_)]
			msg = OldAppendInt(msg, 1024)
		}
	})

	b.Run("Value = 1000000", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			msg := scratch[:len(_PUB_P_)]
			msg = OldAppendInt(msg, 1000000)
		}
	})

	b.Run("Value = 1212121212", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			msg := scratch[:len(_PUB_P_)]
			msg = OldAppendInt(msg, 1212121212)
		}
	})

	b.Run("Value = 998877665544332211", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			msg := scratch[:len(_PUB_P_)]
			msg = OldAppendInt(msg, 998877665544332211)
		}
	})
}

func BenchmarkAppendInt(b *testing.B) {

	b.Run("Value = 0", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			msg := scratch[:len(_PUB_P_)]
			msg = AppendInt(msg, 0)
		}
	})

	b.Run("Value = 8", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			msg := scratch[:len(_PUB_P_)]
			msg = AppendInt(msg, 8)
		}
	})

	b.Run("Value = 65", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			msg := scratch[:len(_PUB_P_)]
			msg = AppendInt(msg, 65)
		}
	})

	b.Run("Value = 1024", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			msg := scratch[:len(_PUB_P_)]
			msg = AppendInt(msg, 1024)
		}
	})

	b.Run("Value = 1000000", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			msg := scratch[:len(_PUB_P_)]
			msg = AppendInt(msg, 1000000)
		}
	})

	b.Run("Value = 1212121212", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			msg := scratch[:len(_PUB_P_)]
			msg = AppendInt(msg, 1212121212)
		}
	})

	b.Run("Value = 998877665544332211", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			msg := scratch[:len(_PUB_P_)]
			msg = AppendInt(msg, 998877665544332211)
		}
	})
}

func BenchmarkStrconvAppendInt(b *testing.B) {

	b.Run("Value = 0", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			msg := scratch[:len(_PUB_P_)]
			msg = strconv.AppendInt(msg, 0, 10)
		}
	})

	b.Run("Value = 8", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			msg := scratch[:len(_PUB_P_)]
			msg = strconv.AppendInt(msg, 8, 10)
		}
	})

	b.Run("Value = 65", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			msg := scratch[:len(_PUB_P_)]
			msg = strconv.AppendInt(msg, 65, 10)
		}
	})

	b.Run("Value = 1024", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			msg := scratch[:len(_PUB_P_)]
			msg = strconv.AppendInt(msg, 1024, 10)
		}
	})

	b.Run("Value = 1000000", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			msg := scratch[:len(_PUB_P_)]
			msg = strconv.AppendInt(msg, 1000000, 10)
		}
	})

	b.Run("Value = 1212121212", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			msg := scratch[:len(_PUB_P_)]
			msg = strconv.AppendInt(msg, 1212121212, 10)
		}
	})

	b.Run("Value = 998877665544332211", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			msg := scratch[:len(_PUB_P_)]
			msg = strconv.AppendInt(msg, 998877665544332211, 10)
		}
	})
}
