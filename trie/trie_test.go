package trie

import (
	"math/rand"
	"strconv"
	"sync"
	"testing"

	"github.com/tylertreat/fast-topic-matching"
)

func BenchmarkMultithreaded16Thread5050Trie(b *testing.B) {
	numItems := 10000
	numThreads := 16
	benchmark5050(b, numItems, numThreads, func(items [][]string) matching.Matcher {
		return matching.NewTrieMatcher()
	})
}

func BenchmarkMultithreaded16Thread5050CSTrie(b *testing.B) {
	numItems := 10000
	numThreads := 16
	benchmark5050(b, numItems, numThreads, func(items [][]string) matching.Matcher {
		return matching.NewCSTrieMatcher()
	})
}

func populateMatcher(m matching.Matcher, num, topicSize int) {
	for i := 0; i < num; i++ {
		prefix := ""
		topic := ""
		for j := 0; j < topicSize; j++ {
			topic += prefix + strconv.Itoa(rand.Int())
			prefix = "."
		}
		m.Subscribe(topic, matching.Subscriber(topic))
	}
}

func benchmark5050(b *testing.B, numItems, numThreads int, factory func([][]string) matching.Matcher) {
	itemsToInsert := make([][]string, 0, numThreads)
	for i := 0; i < numThreads; i++ {
		items := make([]string, 0, numItems)
		for j := 0; j < numItems; j++ {
			topic := strconv.Itoa(j%10) + "." + strconv.Itoa(j%50) + "." + strconv.Itoa(j)
			items = append(items, topic)
		}
		itemsToInsert = append(itemsToInsert, items)
	}

	var wg sync.WaitGroup
	sub := matching.Subscriber("abc")
	m := factory(itemsToInsert)
	populateMatcher(m, 1000, 5)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		wg.Add(numThreads)
		for j := 0; j < numThreads; j++ {
			go func(j int) {
				if j%2 != 0 {
					for _, key := range itemsToInsert[j] {
						m.Subscribe(key, sub)
					}
				} else {
					for _, key := range itemsToInsert[j] {
						m.Lookup(key)
					}
				}
				wg.Done()
			}(j)
		}
		wg.Wait()
	}
}
