/*
 * CODE GENERATED AUTOMATICALLY WITH github.com/alanctgardner/gogen-avro
 * THIS FILE SHOULD NOT BE EDITED BY HAND
 *
 * SOURCE:
 *     struct.avro
 */

package serialization

import (
	"io"
)

type StructAvro struct {
	Field1 string
	Field2 int64
	Field3 []string
	Field4 int64
	Field5 string
	Field6 string
	Field7 []byte
}

func DeserializeStructAvro(r io.Reader) (*StructAvro, error) {
	return readStructAvro(r)
}

func NewStructAvro() *StructAvro {
	v := &StructAvro{
		Field3: make([]string, 0),
	}

	return v
}

func (r *StructAvro) Schema() string {
	return "{\"fields\":[{\"name\":\"Field1\",\"type\":\"string\"},{\"name\":\"Field2\",\"type\":\"long\"},{\"name\":\"Field3\",\"type\":{\"items\":\"string\",\"type\":\"array\"}},{\"name\":\"Field4\",\"type\":\"long\"},{\"name\":\"Field5\",\"type\":\"string\"},{\"name\":\"Field6\",\"type\":\"string\"},{\"name\":\"Field7\",\"type\":\"bytes\"}],\"name\":\"StructAvro\",\"type\":\"record\"}"
}

func (r *StructAvro) Serialize(w io.Writer) error {
	return writeStructAvro(r, w)
}
