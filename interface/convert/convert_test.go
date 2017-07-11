package convert

import "testing"

type myStruct struct{}

type myIface interface {
	Foo()
}

func (*myStruct) Foo() {}

func BenchmarkMethodConcrete(b *testing.B) {
	m := new(myStruct)
	invokeConcrete(m, b.N)
}

func BenchmarkMethodIface(b *testing.B) {
	m := new(myStruct)
	invokeIface(m, b.N)
}

func BenchmarkMethodTypeSwitch(b *testing.B) {
	m := new(myStruct)
	invokeTypeSwitch(m, b.N)
}

func BenchmarkMethodTypeAssert(b *testing.B) {
	m := new(myStruct)
	invokeTypeAssert(m, b.N)
}

func invokeConcrete(m *myStruct, n int) {
	for k := 0; k < n; k++ {
		m.Foo()
	}
}

func invokeIface(m myIface, n int) {
	for k := 0; k < n; k++ {
		m.Foo()
	}
}

func invokeTypeSwitch(m myIface, n int) {
	for k := 0; k < n; k++ {
		switch v := m.(type) {
		case *myStruct:
			v.Foo()
		}
	}
}

func invokeTypeAssert(m myIface, n int) {
	for k := 0; k < n; k++ {
		if n, ok := m.(*myStruct); ok {
			n.Foo()
		}
	}
}
