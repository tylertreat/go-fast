package serialization

func MakeStruct() *Struct {
	return &Struct{
		Field1: "foo",
		Field2: 42,
		Field3: make([]string, 10),
		Field4: 100,
		Field5: "bar",
		Field6: "baz",
		Field7: make([]byte, 10),
	}
}

type Iface interface {
	Foo()
}

type Struct struct {
	Field1 string
	Field2 int
	Field3 []string
	Field4 uint64
	Field5 string
	Field6 string
	Field7 []byte
}

func (s *Struct) Foo() {}
