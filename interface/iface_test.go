package serialization

import (
	"math/rand"
	"sort"
	"testing"

	"github.com/Workiva/go-datastructures/sort"

	"github.com/tylertreat/go-fast/serialization"
)

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
		s[i] = make(SortableByNumber, 100000000)
		for j := 0; j < 100000000; j++ {
			s[i][j] = Sortable{j}
		}
		shuffle(s[i])
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		sort.Sort(s[i])
	}
}

func BenchmarkSortIface(b *testing.B) {
	s := make([]SortableIfaceByNumber, b.N)
	for i := 0; i < b.N; i++ {
		s[i] = make(SortableIfaceByNumber, 100000000)
		for j := 0; j < 100000000; j++ {
			s[i][j] = Sortable{j}
		}
		shuffleIfaces(s[i])
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		sort.Sort(s[i])
	}
}

func BenchmarkSortParallelSymMerge(b *testing.B) {
	s := make([]merge.Comparator, 100000000)
	for i := 0; i < 100000000; i++ {
		s[i] = Sortable{i}
	}
	shuffleComparators(s)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = merge.MultithreadedSortComparators(s)
	}
}
