package greedy

import "sort"

func FindContentChildren(g, s []int) (count int) {
	sort.Ints(g)
	sort.Ints(s)
	// 查找饼干
	for j := 0; count < len(g) && j < len(s); j++ {
		// 如果饼干能满足当前孩子，就看下一个孩子，如果不能就看下一个饼干
		if g[count] <= s[j] {
			count++
		}
	}
	return count
}
