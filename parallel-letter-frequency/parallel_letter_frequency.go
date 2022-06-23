package letter

import "sync"

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

// Frequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap.
func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

// ConcurrentFrequency counts the frequency of each rune in the given strings,
// by making use of concurrency.
func ConcurrentFrequency(l []string) FreqMap {
	m := FreqMap{}
	ch := make(chan FreqMap)
	wg := sync.WaitGroup{}
	for _, line := range l {
		wg.Add(1)
		go func(text string) {
			ch <- Frequency(text)
			wg.Done()
		}(line)
	}
	go func() {
		wg.Wait()
		close(ch)
	}()
	for results := range ch {
		for k, v := range results {
			m[k] += v
		}
	}

	return m
}
