package main

import (
	"fmt"
	"sort"
)

func main() {
	s := []string{"eat", "tea", "tan", "tan", "ate", "nat", "bat", "eat"}
	fmt.Println(groupAnagrams(s))
}

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

func groupAnagrams(strs []string) [][]string {
	ans := [][]string{}
	mp := make(map[string]int)
	cnt := 0
	for i := 0; i < len(strs); i++ {
		s := adjust(strs[i])
		if mp[s] == 0 {
			cnt++
			mp[s] = cnt
			ans = append(ans, []string{strs[i]})
		} else {
			j := mp[s] - 1
			ans[j] = append(ans[j], strs[i])
		}
	}
	return ans
}

func adjust(s string) string {
	val := []rune(s)
	sort.Sort(sortRunes(val))
	//fmt.Println(val)
	return string(val)
}
