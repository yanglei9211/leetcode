package main

import (
	"sort"
	"fmt"
)
//Given a collection of intervals, merge all overlapping intervals.
//
//For example,
//Given [1,3],[2,6],[8,10],[15,18],
//return [1,6],[8,10],[15,18].

/**
 * Definition for an interval.
 * type Interval struct {
 *	   Start int
 *	   End   int
 * }
 */

type Interval struct {
	Start int
	End int
}

type Intervals []Interval

func (s Intervals) Less(i, j int) bool {
	if s[i].Start < s[j].Start {
		return true
	} else if s[i].Start == s[j].Start {
		return s[i].End < s[j].End
	}
	return false
}

func (s Intervals) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s Intervals) Len() int {
	return len(s)
}

func merge(intervals []Interval) []Interval {
	if len(intervals) == 0 {
		return intervals
	}

    its := make(Intervals, 0, len(intervals))
	for _, i := range(intervals) {
		its = append(its, i)
	}
	sort.Sort(its)
	res := []Interval{}
	st := its[0].Start
	ed := its[0].End
	//maxx := its[0].End
	for i := 1; i < len(its); i++ {
		if its[i].Start <= ed {
			if its[i].End > ed {
				ed = its[i].End
			}
		} else {
			res = append(res, Interval{st, ed})
			st = its[i].Start
			ed = its[i].End
		}
	}
	res = append(res, Interval{st, ed})
	return res
}

func main() {
	a := []Interval{Interval{1,3},Interval{2,6},Interval{8,10}, Interval{15,18}}
	fmt.Println(merge(a))
}