package parsing

const digits = "0123456789"

// OldAppendInt performs the original behavior from go-nats,
// and is included here for benchmarking
func OldAppendInt(dst []byte, i int64) (d []byte) {
	var b [32]byte
	var itr = len(b)
	if i > 0 {
		for ; i > 0; i /= 10 {
			itr--
			b[itr] = digits[i%10]
		}
	} else {
		itr--
		b[itr] = digits[0]
	}

	d = append(dst, b[itr:]...)
	return
}
