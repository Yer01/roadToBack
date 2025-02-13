package main

func getNameCounts(names []string) map[rune]map[string]int {
	counts := map[rune]map[string]int{}

	for _, name := range names {
		c := []rune(name)
		if _, ok := counts[c[0]]; !ok {
			counts[c[0]] = make(map[string]int)
		}
		counts[c[0]][name]++
	}
	return counts
}
