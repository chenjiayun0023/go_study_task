package main

import (
	"fmt"
	"sort"
)

//https://leetcode.cn/problems/merge-intervals/description/

/*
56. 合并区间：以数组 intervals 表示若干个区间的集合，
其中单个区间为 intervals[i] = [starti, endi] 。
请你合并所有重叠的区间，并返回一个不重叠的区间数组，
该数组需恰好覆盖输入中的所有区间。可以先对区间数组按照区间的起始位置进行排序，
然后使用一个切片来存储合并后的区间，遍历排序后的区间数组，
将当前区间与切片中最后一个区间进行比较，如果有重叠，则合并区间；
如果没有重叠，则将当前区间添加到切片中。

0 <= starti <= endi <= 10^4
*/

func merge(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return intervals
	}
	// 先对区间进行排序
	sort.Slice(intervals, func(i, j int) bool {
		// fmt.Println("排序前:", intervals[i][0], intervals[j][0])
		return intervals[i][0] < intervals[j][0]
	})
	fmt.Println("排序后的区间:", intervals)

	merged := [][]int{intervals[0]}
	fmt.Println("merged:", merged)
	for i := 1; i < len(intervals); i++ {
		fmt.Println("当前区间:", intervals[i], "已合并区间:", merged)
		// 如果有重叠，则合并区间
		fmt.Println("当前区间:", intervals[i][0], "和", merged[len(merged)-1][1], "比较")
		if intervals[i][0] <= merged[len(merged)-1][1] {

			merged[len(merged)-1][1] = max(merged[len(merged)-1][1], intervals[i][1])
		} else {
			// 如果没有重叠，则添加当前区间
			merged = append(merged, intervals[i])
		}
	}
	return merged
}

func main() {
	// intervals := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
	// intervals := [][]int{{2, 6}, {1, 4}, {8, 10}, {15, 18}}
	intervals := [][]int{{2, 3}, {1, 4}, {8, 10}, {15, 18}}
	fmt.Println(len(intervals))
	merged := merge(intervals)
	fmt.Println("合并后的区间:", merged)
}
