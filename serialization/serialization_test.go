package serialization

import (
	"bytes"
	"encoding/json"
	"testing"

	"git.apache.org/thrift.git/lib/go/thrift"
	capn "github.com/glycerine/go-capnproto"
)

func makeStruct() *Struct {
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

func makeStructPB() *StructPB {
	return &StructPB{
		Field1: "foo",
		Field2: 42,
		Field3: make([]string, 10),
		Field4: 100,
		Field5: "bar",
		Field6: "baz",
		Field7: make([]byte, 10),
	}
}

func makeStructThrift() *StructThrift {
	return &StructThrift{
		Field1: "foo",
		Field2: 42,
		Field3: make([]string, 10),
		Field4: 100,
		Field5: "bar",
		Field6: "baz",
		Field7: make([]byte, 10),
	}
}

func makeStructCapnp() *capn.Segment {
	seg := capn.NewBuffer(make([]byte, 0, 1<<20))
	s := NewRootStructCapnp(seg)
	s.SetField1("foo")
	s.SetField2(42)
	s.SetField3(seg.NewTextList(10))
	s.SetField4(100)
	s.SetField5("bar")
	s.SetField6("baz")
	s.SetField7(make([]byte, 10))
	return seg
}

func BenchmarkJSONReflectionMarshal(b *testing.B) {
	s := makeStruct()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, err := json.Marshal(s)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkFFJSONMarshal(b *testing.B) {
	s := makeStruct()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, err := s.MarshalJSON()
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkMsgpackMarshal(b *testing.B) {
	s := makeStruct()
	var data []byte
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, err := s.MarshalMsg(data)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkProtobufMarshal(b *testing.B) {
	s := makeStructPB()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, err := s.Marshal()
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkThriftMarshal(b *testing.B) {
	s := makeStructThrift()
	serializer := thrift.NewTSerializer()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, err := serializer.Write(s)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkCapnpMarshal(b *testing.B) {
	s := makeStructCapnp()
	var buf bytes.Buffer
	b.ResetTimer()

	_, err := s.WriteTo(&buf)
	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		buf.Reset()
		_, err := s.WriteTo(&buf)
		if err != nil {
			b.Fatalf("WriteTo: %v", err)
		}
	}
}
