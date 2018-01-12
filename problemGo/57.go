package main

import "fmt"

//Given a set of non-overlapping intervals, insert a new interval into the intervals (merge if necessary).
//
//You may assume that the intervals were initially sorted according to their start times.
//
//Example 1:
//Given intervals [1,3],[6,9], insert and merge [2,5] in as [1,5],[6,9].
//
//Example 2:
//Given [1,2],[3,5],[6,7],[8,10],[12,16], insert and merge [4,9] in as [1,2],[3,10],[12,16].
//
//This is because the new interval [4,9] overlaps with [3,5],[6,7],[8,10].

/**
 * Definition for an interval.
 * type Interval struct {
 *	   Start int
 *	   End   int
 * }
 */
type Interval struct {
	Start int
	End   int
}

func insert(intervals []Interval, newInterval Interval) []Interval {
	its := make([]Interval, 0, len(intervals)+1)
	hasIns := false
	for i := 0; i < len(intervals); i++ {
		if !hasIns && intervals[i].Start > newInterval.Start {
			its = append(its, newInterval)
			hasIns = true
		}

		its = append(its, intervals[i])

	}
	if !hasIns {
		its = append(its, newInterval)
		hasIns = true
	}
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
	//a := []Interval{Interval{1,2},Interval{3,5},Interval{6,7}, Interval{8,10}, Interval{12, 16}}
	a := []Interval{}
	x := Interval{4, 9}
	fmt.Println(insert(a, x))
}
