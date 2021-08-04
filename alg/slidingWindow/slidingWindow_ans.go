package slidingWindow

import (
	. "dp/tools"
	"math"
)

func LengthOfLongestSubstring(s string) int {
	// 哈希集合，记录每个字符是否出现过
	set := map[byte]bool{}
	right, ans, size := 0, 0, len(s)
	for left := 0; left < size; left++ {
		if left != 0 {
			// 左指针向右移动一格，移除一个字符
			delete(set, s[left-1])
		}
		// 记录右指针走过的节点，直到当前字母已经出现过或者到达最后一个字母
		for right < size && !set[s[right]] {
			set[s[right]] = true
			right++
		}
		// 第 left 到 right 个字符是一个极长的无重复字符子串
		ans = Max(ans, right-left)
	}
	return ans
}

func LengthOfLongestSubstring2(s string) int {
	// l，r记录左右，c记录容器长度（初始条件）
	var (
		l, r, res int
		c         = map[string]bool{}
		n         = len(s)
	)
	for r < n {
		// 如果容器中不存在当前元素，则添加并移动右边界
		if !c[Ele(s, r)] {
			c[Ele(s, r)] = true
			r++
		} else {
			// 如果当前右边界所志向的元素已经存在于容器中，并且左边界还没有超过右边界
			for l < r && c[Ele(s, r)] {
				// 则移动左边界并移除左边界所指向的元素
				delete(c, Ele(s, l))
				l++
			}
		}
		// 更新结果
		res = Max(res, r-l)
	}
	return res
}

func CharacterReplacement1(s string, k int) int {
	var memo [26]int
	size := len(s)
	max := math.MinInt64
	l := 0
	for r, v := range s {
		memo[v-'A']++
		max = Max(max, memo[v-'A'])
		if r-l+1 > max+k {
			memo[s[l]-'A']--
			l++
		}
	}
	return size - l
}

func CheckInclusion(s1, s2 string) bool {
	n, m := len(s1), len(s2)
	if n > m {
		return false
	}
	cnt := [26]int{}
	for i, ch := range s1 {
		cnt[ch-'a']--
		cnt[s2[i]-'a']++
	}
	diff := 0
	for _, c := range cnt {
		if c != 0 {
			diff++
		}
	}
	if diff == 0 {
		return true
	}
	for i := n; i < m; i++ {
		l, r := s2[i]-'a', s2[i-n]-'a'
		if l == r {
			continue
		}

		if cnt[l] == 0 {
			diff++
		}
		cnt[l]++
		if cnt[l] == 0 {
			diff--
		}

		if cnt[r] == 0 {
			diff++
		}
		cnt[r]--
		if cnt[r] == 0 {
			diff--
		}

		if diff == 0 {
			return true
		}
	}
	return false
}

func CheckInclusion1(s1 string, s2 string) bool {
	size1, size2 := len(s1), len(s2)
	if size1 > size2 {
		return false
	}
	var n, m [26]int
	for key, value := range s1 {
		n[value-'a']++
		m[s2[key]-'a']++
	}
	if n == m {
		return true
	}
	for i := size1; i < size2; i++ {
		m[s2[i]-'a']++
		m[s2[i-size1]-'a']--
		if n == m {
			return true
		}
	}
	return false
}

func MinWindow(s string, t string) string {
	sSize, tSize := len(s), len(t)
	if sSize == 0 || tSize == 0 {
		return ""
	}
	l, r, start, count, winSize, need := 0, 0, 0, tSize, math.MaxInt32, make([]int, 128)
	for i := 0; i < count; i++ {
		need[t[i]]++
	}
	// 右边界到达S字符串的末尾为止
	for r < sSize {
		// 如果S字符串中的字符在T字符串中存在，则计数器标记找到了一个字符
		if need[s[r]] > 0 {
			count--
		}
		// 标记S字符串中存在的字符，负数为不需要的字符
		need[s[r]]--
		// 如果所有的字符全部找到
		if count == 0 {
			// 右指针与左指针之间还存在空隙，并且当前左指针指向的字符串已经被标记为多余字符，则移动左边界
			for l < r && need[s[l]] < 0 {
				need[s[l]]++
				l++
			}
			// 如果当前窗口小于现有记录则更新窗口大小和左边界
			if r-l+1 < winSize {
				winSize = r - l + 1
				start = l
			}
			// 移动左边界，需要重新寻找丢失的字符
			need[s[l]]++
			l++
			count++
		}
		// 每次都要移动右边界
		r++
	}
	if winSize == math.MaxInt32 {
		return ""
	} else {
		return s[start : start+winSize]
	}
}

func MaxSlidingWindow(nums []int, k int) (res []int) {
	queue := make([]int, 0, k)
	push := func(v int) {
		for len(queue) > 0 && v > queue[len(queue)-1] {
			queue = queue[:len(queue)-1]
		}
		queue = append(queue, v)
	}

	pop := func(v int) {
		if len(queue) > 0 && queue[0] == v {
			queue = queue[1:]
		}
	}
	for i := 0; i < len(nums); i++ {
		push(nums[i])
		if i >= i-k+1 {
			res = append(res, queue[0])
			pop(nums[i-k+1])
		}
	}
	return
}
