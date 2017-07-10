/*
 * CODE GENERATED AUTOMATICALLY WITH github.com/alanctgardner/gogen-avro
 * THIS FILE SHOULD NOT BE EDITED BY HAND
 *
 * SOURCE:
 *     struct.avro
 */

package serialization

import (
	"fmt"
	"io"
	"math"
)

type ByteWriter interface {
	Grow(int)
	WriteByte(byte) error
}

type StringWriter interface {
	WriteString(string) (int, error)
}

func encodeInt(w io.Writer, byteCount int, encoded uint64) error {
	var err error
	var bb []byte
	bw, ok := w.(ByteWriter)
	// To avoid reallocations, grow capacity to the largest possible size
	// for this integer
	if ok {
		bw.Grow(byteCount)
	} else {
		bb = make([]byte, 0, byteCount)
	}

	if encoded == 0 {
		if bw != nil {
			err = bw.WriteByte(0)
			if err != nil {
				return err
			}
		} else {
			bb = append(bb, byte(0))
		}
	} else {
		for encoded > 0 {
			b := byte(encoded & 127)
			encoded = encoded >> 7
			if !(encoded == 0) {
				b |= 128
			}
			if bw != nil {
				err = bw.WriteByte(b)
				if err != nil {
					return err
				}
			} else {
				bb = append(bb, b)
			}
		}
	}
	if bw == nil {
		_, err := w.Write(bb)
		return err
	}
	return nil

}

func readArrayString(r io.Reader) ([]string, error) {
	var err error
	var blkSize int64
	var arr = make([]string, 0)
	for {
		blkSize, err = readLong(r)
		if err != nil {
			return nil, err
		}
		if blkSize == 0 {
			break
		}
		if blkSize < 0 {
			blkSize = -blkSize
			_, err = readLong(r)
			if err != nil {
				return nil, err
			}
		}
		for i := int64(0); i < blkSize; i++ {
			elem, err := readString(r)
			if err != nil {
				return nil, err
			}
			arr = append(arr, elem)
		}
	}
	return arr, nil
}

func readBytes(r io.Reader) ([]byte, error) {
	size, err := readLong(r)
	if err != nil {
		return nil, err
	}
	bb := make([]byte, size)
	_, err = io.ReadFull(r, bb)
	return bb, err
}

func readLong(r io.Reader) (int64, error) {
	var v uint64
	buf := make([]byte, 1)
	for shift := uint(0); ; shift += 7 {
		if _, err := io.ReadFull(r, buf); err != nil {
			return 0, err
		}
		b := buf[0]
		v |= uint64(b&127) << shift
		if b&128 == 0 {
			break
		}
	}
	datum := (int64(v>>1) ^ -int64(v&1))
	return datum, nil
}

func readString(r io.Reader) (string, error) {
	len, err := readLong(r)
	if err != nil {
		return "", err
	}

	// makeslice can fail depending on available memory.
	// We arbitrarily limit string size to sane default (~2.2GB).
	if len < 0 || len > math.MaxInt32 {
		return "", fmt.Errorf("string length out of range: %d", len)
	}

	bb := make([]byte, len)
	_, err = io.ReadFull(r, bb)
	if err != nil {
		return "", err
	}
	return string(bb), nil
}

func readStructAvro(r io.Reader) (*StructAvro, error) {
	var str = &StructAvro{}
	var err error
	str.Field1, err = readString(r)
	if err != nil {
		return nil, err
	}
	str.Field2, err = readLong(r)
	if err != nil {
		return nil, err
	}
	str.Field3, err = readArrayString(r)
	if err != nil {
		return nil, err
	}
	str.Field4, err = readLong(r)
	if err != nil {
		return nil, err
	}
	str.Field5, err = readString(r)
	if err != nil {
		return nil, err
	}
	str.Field6, err = readString(r)
	if err != nil {
		return nil, err
	}
	str.Field7, err = readBytes(r)
	if err != nil {
		return nil, err
	}

	return str, nil
}

func writeArrayString(r []string, w io.Writer) error {
	err := writeLong(int64(len(r)), w)
	if err != nil || len(r) == 0 {
		return err
	}
	for _, e := range r {
		err = writeString(e, w)
		if err != nil {
			return err
		}
	}
	return writeLong(0, w)
}

func writeBytes(r []byte, w io.Writer) error {
	err := writeLong(int64(len(r)), w)
	if err != nil {
		return err
	}
	_, err = w.Write(r)
	return err
}

func writeLong(r int64, w io.Writer) error {
	downShift := uint64(63)
	encoded := uint64((r << 1) ^ (r >> downShift))
	const maxByteSize = 10
	return encodeInt(w, maxByteSize, encoded)
}

func writeString(r string, w io.Writer) error {
	err := writeLong(int64(len(r)), w)
	if err != nil {
		return err
	}
	if sw, ok := w.(StringWriter); ok {
		_, err = sw.WriteString(r)
	} else {
		_, err = w.Write([]byte(r))
	}
	return err
}

func writeStructAvro(r *StructAvro, w io.Writer) error {
	var err error
	err = writeString(r.Field1, w)
	if err != nil {
		return err
	}
	err = writeLong(r.Field2, w)
	if err != nil {
		return err
	}
	err = writeArrayString(r.Field3, w)
	if err != nil {
		return err
	}
	err = writeLong(r.Field4, w)
	if err != nil {
		return err
	}
	err = writeString(r.Field5, w)
	if err != nil {
		return err
	}
	err = writeString(r.Field6, w)
	if err != nil {
		return err
	}
	err = writeBytes(r.Field7, w)
	if err != nil {
		return err
	}

	return nil
}
