package serialization

// AUTO GENERATED - DO NOT EDIT

import (
	"bufio"
	"bytes"
	"encoding/json"
	C "github.com/glycerine/go-capnproto"
	"io"
)

type StructCapnp C.Struct

func NewStructCapnp(s *C.Segment) StructCapnp      { return StructCapnp(s.NewStruct(16, 5)) }
func NewRootStructCapnp(s *C.Segment) StructCapnp  { return StructCapnp(s.NewRootStruct(16, 5)) }
func AutoNewStructCapnp(s *C.Segment) StructCapnp  { return StructCapnp(s.NewStructAR(16, 5)) }
func ReadRootStructCapnp(s *C.Segment) StructCapnp { return StructCapnp(s.Root(0).ToStruct()) }
func (s StructCapnp) Field1() string               { return C.Struct(s).GetObject(0).ToText() }
func (s StructCapnp) Field1Bytes() []byte          { return C.Struct(s).GetObject(0).ToDataTrimLastByte() }
func (s StructCapnp) SetField1(v string)           { C.Struct(s).SetObject(0, s.Segment.NewText(v)) }
func (s StructCapnp) Field2() int64                { return int64(C.Struct(s).Get64(0)) }
func (s StructCapnp) SetField2(v int64)            { C.Struct(s).Set64(0, uint64(v)) }
func (s StructCapnp) Field3() C.TextList           { return C.TextList(C.Struct(s).GetObject(1)) }
func (s StructCapnp) SetField3(v C.TextList)       { C.Struct(s).SetObject(1, C.Object(v)) }
func (s StructCapnp) Field4() uint64               { return C.Struct(s).Get64(8) }
func (s StructCapnp) SetField4(v uint64)           { C.Struct(s).Set64(8, v) }
func (s StructCapnp) Field5() string               { return C.Struct(s).GetObject(2).ToText() }
func (s StructCapnp) Field5Bytes() []byte          { return C.Struct(s).GetObject(2).ToDataTrimLastByte() }
func (s StructCapnp) SetField5(v string)           { C.Struct(s).SetObject(2, s.Segment.NewText(v)) }
func (s StructCapnp) Field6() string               { return C.Struct(s).GetObject(3).ToText() }
func (s StructCapnp) Field6Bytes() []byte          { return C.Struct(s).GetObject(3).ToDataTrimLastByte() }
func (s StructCapnp) SetField6(v string)           { C.Struct(s).SetObject(3, s.Segment.NewText(v)) }
func (s StructCapnp) Field7() []byte               { return C.Struct(s).GetObject(4).ToData() }
func (s StructCapnp) SetField7(v []byte)           { C.Struct(s).SetObject(4, s.Segment.NewData(v)) }
func (s StructCapnp) WriteJSON(w io.Writer) error {
	b := bufio.NewWriter(w)
	var err error
	var buf []byte
	_ = buf
	err = b.WriteByte('{')
	if err != nil {
		return err
	}
	_, err = b.WriteString("\"field1\":")
	if err != nil {
		return err
	}
	{
		s := s.Field1()
		buf, err = json.Marshal(s)
		if err != nil {
			return err
		}
		_, err = b.Write(buf)
		if err != nil {
			return err
		}
	}
	err = b.WriteByte(',')
	if err != nil {
		return err
	}
	_, err = b.WriteString("\"field2\":")
	if err != nil {
		return err
	}
	{
		s := s.Field2()
		buf, err = json.Marshal(s)
		if err != nil {
			return err
		}
		_, err = b.Write(buf)
		if err != nil {
			return err
		}
	}
	err = b.WriteByte(',')
	if err != nil {
		return err
	}
	_, err = b.WriteString("\"field3\":")
	if err != nil {
		return err
	}
	{
		s := s.Field3()
		{
			err = b.WriteByte('[')
			if err != nil {
				return err
			}
			for i, s := range s.ToArray() {
				if i != 0 {
					_, err = b.WriteString(", ")
				}
				if err != nil {
					return err
				}
				buf, err = json.Marshal(s)
				if err != nil {
					return err
				}
				_, err = b.Write(buf)
				if err != nil {
					return err
				}
			}
			err = b.WriteByte(']')
		}
		if err != nil {
			return err
		}
	}
	err = b.WriteByte(',')
	if err != nil {
		return err
	}
	_, err = b.WriteString("\"field4\":")
	if err != nil {
		return err
	}
	{
		s := s.Field4()
		buf, err = json.Marshal(s)
		if err != nil {
			return err
		}
		_, err = b.Write(buf)
		if err != nil {
			return err
		}
	}
	err = b.WriteByte(',')
	if err != nil {
		return err
	}
	_, err = b.WriteString("\"field5\":")
	if err != nil {
		return err
	}
	{
		s := s.Field5()
		buf, err = json.Marshal(s)
		if err != nil {
			return err
		}
		_, err = b.Write(buf)
		if err != nil {
			return err
		}
	}
	err = b.WriteByte(',')
	if err != nil {
		return err
	}
	_, err = b.WriteString("\"field6\":")
	if err != nil {
		return err
	}
	{
		s := s.Field6()
		buf, err = json.Marshal(s)
		if err != nil {
			return err
		}
		_, err = b.Write(buf)
		if err != nil {
			return err
		}
	}
	err = b.WriteByte(',')
	if err != nil {
		return err
	}
	_, err = b.WriteString("\"field7\":")
	if err != nil {
		return err
	}
	{
		s := s.Field7()
		buf, err = json.Marshal(s)
		if err != nil {
			return err
		}
		_, err = b.Write(buf)
		if err != nil {
			return err
		}
	}
	err = b.WriteByte('}')
	if err != nil {
		return err
	}
	err = b.Flush()
	return err
}
func (s StructCapnp) MarshalJSON() ([]byte, error) {
	b := bytes.Buffer{}
	err := s.WriteJSON(&b)
	return b.Bytes(), err
}
func (s StructCapnp) WriteCapLit(w io.Writer) error {
	b := bufio.NewWriter(w)
	var err error
	var buf []byte
	_ = buf
	err = b.WriteByte('(')
	if err != nil {
		return err
	}
	_, err = b.WriteString("field1 = ")
	if err != nil {
		return err
	}
	{
		s := s.Field1()
		buf, err = json.Marshal(s)
		if err != nil {
			return err
		}
		_, err = b.Write(buf)
		if err != nil {
			return err
		}
	}
	_, err = b.WriteString(", ")
	if err != nil {
		return err
	}
	_, err = b.WriteString("field2 = ")
	if err != nil {
		return err
	}
	{
		s := s.Field2()
		buf, err = json.Marshal(s)
		if err != nil {
			return err
		}
		_, err = b.Write(buf)
		if err != nil {
			return err
		}
	}
	_, err = b.WriteString(", ")
	if err != nil {
		return err
	}
	_, err = b.WriteString("field3 = ")
	if err != nil {
		return err
	}
	{
		s := s.Field3()
		{
			err = b.WriteByte('[')
			if err != nil {
				return err
			}
			for i, s := range s.ToArray() {
				if i != 0 {
					_, err = b.WriteString(", ")
				}
				if err != nil {
					return err
				}
				buf, err = json.Marshal(s)
				if err != nil {
					return err
				}
				_, err = b.Write(buf)
				if err != nil {
					return err
				}
			}
			err = b.WriteByte(']')
		}
		if err != nil {
			return err
		}
	}
	_, err = b.WriteString(", ")
	if err != nil {
		return err
	}
	_, err = b.WriteString("field4 = ")
	if err != nil {
		return err
	}
	{
		s := s.Field4()
		buf, err = json.Marshal(s)
		if err != nil {
			return err
		}
		_, err = b.Write(buf)
		if err != nil {
			return err
		}
	}
	_, err = b.WriteString(", ")
	if err != nil {
		return err
	}
	_, err = b.WriteString("field5 = ")
	if err != nil {
		return err
	}
	{
		s := s.Field5()
		buf, err = json.Marshal(s)
		if err != nil {
			return err
		}
		_, err = b.Write(buf)
		if err != nil {
			return err
		}
	}
	_, err = b.WriteString(", ")
	if err != nil {
		return err
	}
	_, err = b.WriteString("field6 = ")
	if err != nil {
		return err
	}
	{
		s := s.Field6()
		buf, err = json.Marshal(s)
		if err != nil {
			return err
		}
		_, err = b.Write(buf)
		if err != nil {
			return err
		}
	}
	_, err = b.WriteString(", ")
	if err != nil {
		return err
	}
	_, err = b.WriteString("field7 = ")
	if err != nil {
		return err
	}
	{
		s := s.Field7()
		buf, err = json.Marshal(s)
		if err != nil {
			return err
		}
		_, err = b.Write(buf)
		if err != nil {
			return err
		}
	}
	err = b.WriteByte(')')
	if err != nil {
		return err
	}
	err = b.Flush()
	return err
}
func (s StructCapnp) MarshalCapLit() ([]byte, error) {
	b := bytes.Buffer{}
	err := s.WriteCapLit(&b)
	return b.Bytes(), err
}

type StructCapnp_List C.PointerList

func NewStructCapnpList(s *C.Segment, sz int) StructCapnp_List {
	return StructCapnp_List(s.NewCompositeList(16, 5, sz))
}
func (s StructCapnp_List) Len() int             { return C.PointerList(s).Len() }
func (s StructCapnp_List) At(i int) StructCapnp { return StructCapnp(C.PointerList(s).At(i).ToStruct()) }
func (s StructCapnp_List) ToArray() []StructCapnp {
	n := s.Len()
	a := make([]StructCapnp, n)
	for i := 0; i < n; i++ {
		a[i] = s.At(i)
	}
	return a
}
func (s StructCapnp_List) Set(i int, item StructCapnp) { C.PointerList(s).Set(i, C.Object(item)) }
