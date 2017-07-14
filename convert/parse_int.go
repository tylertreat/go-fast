package parsing

import "errors"

// Ascii numbers 0-9
const (
	asciiZero = 48
	asciiNine = 57
)

// ParseInt64 expects decimal positive numbers. We
// return -1 to signal error
func ParseInt64(d []byte) (n int64, e error) {
	if len(d) == 0 {
		return 0, errors.New("error")
	}
	for _, dec := range d {
		if dec < asciiZero || dec > asciiNine {
			return 0, errors.New("error")
		}
		n = n*10 + (int64(dec) - asciiZero)
	}
	return n, nil
}
