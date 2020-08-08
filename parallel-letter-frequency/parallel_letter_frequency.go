package letter

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

// ConcurrentFrequency uses Frequency function concurrently to count
// the frequency of each rune in each string of array of strings.
func ConcurrentFrequency(ss []string) FreqMap {
	mres := FreqMap{}
	channel := make(chan FreqMap, len(ss))
	for _, s := range ss {
		go func(str string) {
			channel <- Frequency(str)
		}(s)
	}

	for range ss {
		for i, l := range <-channel {
			mres[i] += l
		}
	}

	return mres
}
