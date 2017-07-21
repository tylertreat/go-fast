package iface

import (
	"fmt"
	"testing"
)

func BenchmarkArgIface(b *testing.B) {
	print := false
	if *cpuProfile != "" {
		defer setupProfiling(b, "iface-arg")()
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		argIface("foo", print)
	}
}

func BenchmarkArgString(b *testing.B) {
	print := false
	if *cpuProfile != "" {
		defer setupProfiling(b, "string-arg")()
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		argString("foo", print)
	}
}

func argIface(arg interface{}, print bool) {
	if print {
		fmt.Println(arg)
	}
}

func argString(arg string, print bool) {
	if print {
		fmt.Println(arg)
	}
}
