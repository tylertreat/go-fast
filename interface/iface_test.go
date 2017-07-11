package iface

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"runtime/pprof"
	"sort"
	"testing"

	"github.com/Workiva/go-datastructures/sort"

	"github.com/tylertreat/go-fast/serialization"
)

var cpuProfile = flag.String("prof", "", "Write CPU profile")

func init() {
	flag.Parse()
}

func BenchmarkMethodCall(b *testing.B) {
	s := serialization.MakeStruct()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s.Foo()
	}
}

func BenchmarkMethodCallIface(b *testing.B) {
	var s serialization.Iface = serialization.MakeStruct()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s.Foo()
	}
}

type SortableIface interface {
	Number() int
}

type Sortable struct {
	number int
}

func (s Sortable) Number() int {
	return s.number
}

func (s Sortable) Compare(other merge.Comparator) int {
	otherSortable := other.(Sortable)
	if s.number == otherSortable.number {
		return 0
	}
	if s.number > otherSortable.number {
		return 1
	}
	return -1
}

type SortableIfaceByNumber []SortableIface

func (a SortableIfaceByNumber) Len() int           { return len(a) }
func (a SortableIfaceByNumber) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a SortableIfaceByNumber) Less(i, j int) bool { return a[i].Number() < a[j].Number() }

type SortableByNumber []Sortable

func (a SortableByNumber) Len() int           { return len(a) }
func (a SortableByNumber) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a SortableByNumber) Less(i, j int) bool { return a[i].Number() < a[j].Number() }

func shuffle(a SortableByNumber) {
	for i := range a {
		j := rand.Intn(i + 1)
		a[i], a[j] = a[j], a[i]
	}
}

func shuffleIfaces(a SortableIfaceByNumber) {
	for i := range a {
		j := rand.Intn(i + 1)
		a[i], a[j] = a[j], a[i]
	}
}

func shuffleComparators(a []merge.Comparator) {
	for i := range a {
		j := rand.Intn(i + 1)
		a[i], a[j] = a[j], a[i]
	}
}

func BenchmarkSortStruct(b *testing.B) {
	s := make([]SortableByNumber, b.N)
	for i := 0; i < b.N; i++ {
		s[i] = make(SortableByNumber, 10000000)
		for j := 0; j < 10000000; j++ {
			s[i][j] = Sortable{j}
		}
		shuffle(s[i])
	}
	if *cpuProfile != "" {
		defer setupProfiling(b, "struct")()
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		sort.Sort(s[i])
	}
}

func BenchmarkSortIface(b *testing.B) {
	s := make([]SortableIfaceByNumber, b.N)
	for i := 0; i < b.N; i++ {
		s[i] = make(SortableIfaceByNumber, 10000000)
		for j := 0; j < 10000000; j++ {
			s[i][j] = Sortable{j}
		}
		shuffleIfaces(s[i])
	}
	if *cpuProfile != "" {
		defer setupProfiling(b, "iface")()
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		sort.Sort(s[i])
	}
}

func BenchmarkSortParallelSymMerge(b *testing.B) {
	s := make([]merge.Comparator, 10000000)
	for i := 0; i < 10000000; i++ {
		s[i] = Sortable{i}
	}
	shuffleComparators(s)
	if *cpuProfile != "" {
		defer setupProfiling(b, "symmerge")()
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = merge.MultithreadedSortComparators(s)
	}
}

func setupProfiling(b *testing.B, prefix string) func() {
	file := fmt.Sprintf("%s-%s", prefix, *cpuProfile)
	f, err := os.Create(file)
	if err != nil {
		b.Fatal(err)
	}
	if err := pprof.StartCPUProfile(f); err != nil {
		b.Fatal(err)
	}
	return pprof.StopCPUProfile
}
