package main

import (
	"sort"
)

type sortRunes []rune

func (s sortRunes) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortRunes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortRunes) Len() int {
	return len(s)
}

func sortString(s string) string {
	r := []rune(s)
	sort.Sort(sortRunes(r))
	return string(r)
}

func groupAnagrams(strs []string) [][]string {

	// make a map for the sorted words
	var count = make(map[string]int)
	var res = make([][]string, 0)
	var index int = 0

	for _, s := range strs {

		// sort before checking if the string exists
		x := sortString(s)
		if i, found := count[x]; found {
			// if exists just append it
			res[i] = append(res[i], s)

		} else {
			// else create a new one
			count[x] = index
			temp := make([]string, 0)
			temp = append(temp, s)
			res = append(res, temp)
			index++
		}
	}

	return res
}
