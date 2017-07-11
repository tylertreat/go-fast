package serialization

import (
	"bytes"
	"encoding/json"
	"testing"

	"git.apache.org/thrift.git/lib/go/thrift"
	capn "github.com/glycerine/go-capnproto"

	"github.com/tylertreat/go-fast/serialization/easyjson"
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

func makeEasyJSONStruct() *easyjson.Struct {
	return &easyjson.Struct{
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

func makeStructAvro() *StructAvro {
	return &StructAvro{
		Field1: "foo",
		Field2: 42,
		Field3: make([]string, 10),
		Field4: 100,
		Field5: "bar",
		Field6: "baz",
		Field7: make([]byte, 10),
	}
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

func BenchmarkJSONReflectionMarshalIface(b *testing.B) {
	var s Iface = makeStruct()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, err := json.Marshal(s)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkJSONReflectionUnmarshal(b *testing.B) {
	s := makeStruct()
	buf, err := json.Marshal(s)
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		if err := json.Unmarshal(buf, s); err != nil {
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

func BenchmarkFFJSONUnmarshal(b *testing.B) {
	s := makeStruct()
	buf, err := s.MarshalJSON()
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		err := s.UnmarshalJSON(buf)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkEasyJSONMarshal(b *testing.B) {
	s := makeEasyJSONStruct()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, err := s.MarshalJSON()
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkEasyJSONUnmarshal(b *testing.B) {
	s := makeEasyJSONStruct()
	buf, err := s.MarshalJSON()
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		err := s.UnmarshalJSON(buf)
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

func BenchmarkMsgpackUnmarshal(b *testing.B) {
	s := makeStruct()
	var data []byte
	buf, err := s.MarshalMsg(data)
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, err := s.UnmarshalMsg(buf)
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

func BenchmarkProtobufUnmarshal(b *testing.B) {
	s := makeStructPB()
	buf, err := s.Marshal()
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		err := s.Unmarshal(buf)
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

func BenchmarkThriftUnmarshal(b *testing.B) {
	s := makeStructThrift()
	serializer := thrift.NewTSerializer()
	deserializer := thrift.NewTDeserializer()
	buf, err := serializer.Write(s)
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		err := deserializer.Read(s, buf)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkCapnpMarshal(b *testing.B) {
	s := makeStructCapnp()
	var buf bytes.Buffer

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

func BenchmarkCapnpUnmarshal(b *testing.B) {
	s := makeStructCapnp()
	var buf bytes.Buffer

	_, err := s.WriteTo(&buf)
	if err != nil {
		b.Fatal(err)
	}
	segBuf := bytes.NewBuffer(make([]byte, 0, 1<<20))

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		r := bytes.NewBuffer(buf.Bytes())
		seg, err := capn.ReadFromStream(r, segBuf)
		if err != nil {
			b.Fatalf("WriteTo: %v", err)
		}
		_ = ReadRootStructCapnp(seg)
	}
}

func BenchmarkAvroMarshal(b *testing.B) {
	s := makeStructAvro()
	var buf bytes.Buffer
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		buf.Reset()
		err := s.Serialize(&buf)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkAvroUnmarshal(b *testing.B) {
	s := makeStructAvro()
	var buf bytes.Buffer
	err := s.Serialize(&buf)
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		r := bytes.NewBuffer(buf.Bytes())
		_, err := DeserializeStructAvro(r)
		if err != nil {
			b.Fatal(err)
		}
	}
}
