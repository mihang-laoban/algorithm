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
		c         = map[byte]bool{}
		n         = len(s)
	)
	for r < n {
		// 如果容器中不存在当前元素，则添加并移动右边界
		if !c[s[r]] {
			c[s[r]] = true
			r++
		} else {
			// 如果当前右边界所志向的元素已经存在于容器中，并且左边界还没有超过右边界
			for l < r && c[s[r]] {
				// 则移动左边界并移除左边界所指向的元素
				delete(c, s[l])
				l++
			}
		}
		// 更新结果
		if res < r-l {
			res = r - l
		}
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
	var (
		sSize, tSize = len(s), len(t)
		l, r, start  = 0, 0, 0
		count        = tSize
		winSize      = math.MaxInt32
		need         = make([]int, 128)
	)
	if sSize == 0 || tSize == 0 {
		return ""
	}
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
			var curSize = r - l + 1
			if curSize < winSize {
				winSize = curSize
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
	var (
		size = len(nums)
		q    = make([]int, 0, k)
		push = func(v int) {
			for len(q) > 0 && v > q[len(q)-1] {
				q = q[:len(q)-1]
			}
			q = append(q, v)
		}
		pop = func(v int) {
			if len(q) > 0 && q[0] == v {
				q = q[1:]
			}
		}
	)
	if size == 0 || k <= 0 {
		return
	}
	for i := 0; i < size; i++ {
		push(nums[i])
		var containerCur = i + 1 - k
		if containerCur < 0 {
			continue
		}
		res = append(res, q[0])
		pop(nums[containerCur])
	}
	return
}

func LongestOnes(nums []int, k int) (ans int) {
	var left, right, lSum, rSum int
	// 循环移动右边界
	for right < len(nums) {
		rSum += 1 - nums[right]
		// 什么时候开始移动左边界
		for lSum < rSum-k {
			lSum += 1 - nums[left]
			left++
		}
		ans = Max(ans, right-left+1) // 更新结果
		right++
	}
	return
}

func MinSubArrayLen(target int, nums []int) int {
	var (
		l, r, cur, curSize int
		size               = len(nums)
		winSize            = math.MaxInt32
	)
	for r < size {
		cur += nums[r]
		curSize++
		if cur >= target {
			for l < r && cur-nums[l] >= target {
				cur -= nums[l]
				curSize--
				l++
			}
			if curSize <= winSize {
				winSize = curSize
			}
		}
		r++
	}
	if winSize == math.MaxInt32 {
		return 0
	}
	return winSize
}

func ContainsNearbyDuplicate(nums []int, k int) bool {
	var w = map[int]int{}
	for key, val := range nums {
		if _, ok := w[val]; ok {
			if key-w[val] <= k {
				return true
			}
		}
		w[val] = key
	}
	return false
}

func FindAnagrams(s string, p string) []int {
	var (
		sSize = len(s)
		pSize = len(p)
		pWin  = [26]int{}
		sWin  = [26]int{}
		res   = []int{}
	)
	if sSize < pSize {
		return res
	}
	for i := 0; i < pSize; i++ {
		pWin[p[i]-'a']++
	}
	for l, r := 0, 0; r < sSize; r++ {
		sWin[s[r]-'a']++
		// 如果S串中当前元素的数量大于P串中的，则移动S串的左边界并在窗口中对应元素的数量减一
		for sWin[s[r]-'a'] > pWin[s[r]-'a'] {
			sWin[s[l]-'a']--
			l++
		}
		if r-l+1 == pSize {
			res = append(res, l)
		}
	}
	return res
}

func MaxSatisfied(customers []int, grumpy []int, minutes int) int {
	var (
		ans, tmp int
		n        = len(customers)
	)
	// 统计没有抑制情绪时满意的客户数量
	for i := 0; i < n; i++ {
		if grumpy[i] == 0 {
			ans += customers[i]
		}
	}
	// 滑动窗口
	tmp = ans
	for l, r := 0, 0; r < n; r++ {
		// 仅在生气的情况下，需要调整
		if grumpy[r] == 1 {
			// 窗口值大于minutes，缩小窗口
			for r-l+1 > minutes {
				if grumpy[l] == 1 {
					tmp -= customers[l]
				}
				l++
			}
			tmp += customers[r]
			if tmp > ans {
				ans = tmp
			}
		}
	}
	return ans
}
