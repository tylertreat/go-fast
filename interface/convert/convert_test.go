package convert

import "testing"

type myint int64

type Inccer interface {
	inc()
}

func (i *myint) inc() {
	*i = *i + 1
}

func BenchmarkIncrInt(b *testing.B) {
	i := new(myint)
	incrementInt(i, b.N)
}

func BenchmarkIncrIface(b *testing.B) {
	i := new(myint)
	incrementIface(i, b.N)
}

func BenchmarkIncrTypeSwitch(b *testing.B) {
	i := new(myint)
	incrementTypeSwitch(i, b.N)
}

func BenchmarkIncrTypeAsser(b *testing.B) {
	i := new(myint)
	incrementTypeAssert(i, b.N)
}

func incrementInt(i *myint, n int) {
	for k := 0; k < n; k++ {
		i.inc()
	}
}

func incrementIface(any Inccer, n int) {
	for k := 0; k < n; k++ {
		any.inc()
	}
}

func incrementTypeSwitch(any Inccer, n int) {
	for k := 0; k < n; k++ {
		switch v := any.(type) {
		case *myint:
			v.inc()
		}
	}
}

func incrementTypeAssert(any Inccer, n int) {
	for k := 0; k < n; k++ {
		if newint, ok := any.(*myint); ok {
			newint.inc()
		}
	}
}
