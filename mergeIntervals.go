package main

import (
	"fmt"
	"sort"
)

// 定义一个区间结构体
type Interval struct {
	Start int
	End   int
}

// 定义max方法 返回两个整数的最大值
func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

// 定义Merge函数 用于合并区间
func Merge(intervals []Interval) []Interval {
	// 先对intervals进行排序,先根据起点，再根据终点
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i].Start != intervals[j].Start {
			return intervals[i].Start < intervals[j].Start
		}
		return intervals[i].End < intervals[j].End
	})

	merged := []Interval{}

	for _, interval := range intervals {
		if len(merged) == 0 || interval.Start > merged[len(merged)-1].End {
			merged = append(merged, interval)
		} else {
			merged[len(merged)-1].End = max(merged[len(merged)-1].End, interval.End)
		}
	}

	return merged

}

func main() {
	intervals := []Interval{{3, 6}, {1, 3}, {9, 10}, {9, 15}, {8, 10}, {12, 16}}
	fmt.Println("before merge:", intervals)
	mergedIntervals := Merge(intervals)

	fmt.Println("after merge:", mergedIntervals)
}
