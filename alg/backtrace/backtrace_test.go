package backtrace

import (
	"dp/alg/dp"
	. "dp/ds"
	. "dp/tools"
	"fmt"
	"testing"
)

func init() {
	Max(1, 2)
}

/*给你一个整数数组 nums ，返回该数组所有可能的子集（幂集）。解集不能包含重复的子集。

示例 1：
输入：nums = [1,2,3]
输出：[[],[1],[2],[1,2],[3],[1,3],[2,3],[1,2,3]]*/
func TestSubset(t *testing.T) {
	nums := []int{1, 2, 3}
	fmt.Println(Subset(nums))
	fmt.Println(subset(nums))
}

func subset(nums []int) interface{} {
	res := [][]int{}
	trace := []int{}
	var get func([]int, int, []int)
	get = func(nums []int, start int, trace []int) {
		res = append(res, trace)
		for i := start; i < len(nums); i++ {
			trace = append(trace, nums[i])
			get(nums, i+1, trace)
			trace = trace[:len(trace)-1]
		}
	}
	get(nums, 0, trace)
	return res
}

func TestN(t *testing.T) {
	// 当元素数量超过容量
	// 切片会在底层申请新的数组
	slice := make([]int, 5, 5)
	slice1 := slice
	slice = append(slice, 1)
	slice[0] = 1
	fmt.Println(slice)  //[1 0 0 0 0 1]
	fmt.Println(slice1) //[0 0 0 0 0]
	// copy 函数提供深拷贝功能
	// 但需要在拷贝前申请空间
	slice2 := make([]int, 4, 4)
	slice3 := make([]int, 5, 5)
	fmt.Println(copy(slice2, slice)) //4
	fmt.Println(copy(slice3, slice)) //5
	slice2[1] = 2
	slice3[1] = 3
	fmt.Println(slice)  //[1 0 0 0 0 1]
	fmt.Println(slice2) //[1 2 0 0]
	fmt.Println(slice3) //[1 3 0 0 0]
}

/*
给定一个仅包含数字 2-9 的字符串，返回所有它能表示的字母组合。
给出数字到字母的映射如下（与电话按键相同）。注意 1 不对应任何字母。

示例:

输入："23"
输出：["ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"].
说明:
尽管上面的答案是按字典序排列的，但是你可以任意选择答案输出的顺序。

phoneMap := map[string]string{
	"2": "abc",
	"3": "def",
	"4": "ghi",
	"5": "jkl",
	"6": "mno",
	"7": "pqrs",
	"8": "tuv",
	"9": "wxyz",
}

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/letter-combinations-of-a-phone-number
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

func TestPhoneNum(t *testing.T) {
	fmt.Println(PhoneNum("23"))
	fmt.Println(Phone("23"))
}

func Phone(str string) interface{} {
	res := []string{}
	return res
}

/*n 皇后问题 研究的是如何将 n 个皇后放置在 n×n 的棋盘上，并且使皇后彼此之间不能相互攻击。
给你一个整数 n ，返回所有不同的 n 皇后问题 的解决方案。
每一种解法包含一个不同的 n 皇后问题 的棋子放置方案，该方案中 'Q' 和 '.' 分别代表了皇后和空位。

示例 1：


输入：n = 4
输出：[[".Q..","...Q","Q...","..Q."],["..Q.","Q...","...Q",".Q.."]]
解释：如上图所示，4 皇后问题存在两个不同的解法。
示例 2：

输入：n = 1
输出：[["Q"]]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/n-queens
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

func TestNQueue(t *testing.T) {
	//res := SolveNQueens(4)
	res := SolveNQueen(4)
	for _, v := range res {
		for _, s := range v {
			fmt.Println(s)
		}
		fmt.Println("====")
	}
}

func SolveNQueen(queenNum int) [][]string {
	res := [][]string{}
	return res
}

/*给定一个非负整数数组，你最初位于数组的第一个位置。
数组中的每个元素代表你在该位置可以跳跃的最大长度。
判断你是否能够到达最后一个位置。

示例 1:
输入: [2,3,1,1,4]
输出: true
解释: 我们可以先跳 1 步，从位置 0 到达 位置 1, 然后再从位置 1 跳 3 步到达最后一个位置。

示例 2:
输入: [3,2,1,0,4]
输出: false
解释: 无论怎样，你总会到达索引为 3 的位置。但该位置的最大跳跃长度是 0 ， 所以你永远不可能到达最后一个位置。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/jump-game
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

func TestJumpGame(t *testing.T) {
	//nums := []int{2, 3, 1, 1, 4}
	nums := []int{3, 2, 1, 0, 4}
	fmt.Println(JumpG(nums))
	fmt.Println(dp.JumpGame(nums))
}

func JumpG(nums []int) bool {
	pos := 0
	for i := 0; i < len(nums); i++ {
		if i > pos {
			return false
		}
		pos = Max(pos, i+nums[i])
	}
	return true
}

/*给你一个由 '1'（陆地）和 '0'（水）组成的的二维网格，请你计算网格中岛屿的数量。

岛屿总是被水包围，并且每座岛屿只能由水平方向和/或竖直方向上相邻的陆地连接形成。

此外，你可以假设该网格的四条边均被水包围。



示例 1：

输入：grid = [
["1","1","1","1","0"],
["1","1","0","1","0"],
["1","1","0","0","0"],
["0","0","0","0","0"]
]
输出：1
示例 2：

输入：grid = [
["1","1","0","0","0"],
["1","1","0","0","0"],
["0","0","1","0","0"],
["0","0","0","1","1"]
]
输出：3

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/number-of-islands
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

func TestIsland(t *testing.T) {
	grid := [][]byte{
		[]byte{'1', '1', '0', '0', '0'},
		[]byte{'1', '1', '0', '0', '0'},
		[]byte{'0', '0', '1', '0', '0'},
		[]byte{'0', '0', '0', '1', '1'},
	}
	fmt.Println(numIslands(grid))
}

func numIslands(grid [][]byte) int {
	rows := len(grid)
	if rows == 0 {
		return 0
	}
	cols := len(grid[0])

	// 空地的数量
	spaces := 0
	unionFind := UnionFind{}
	unionFind.UnionFind(rows * cols)
	directions := [][]int{
		[]int{1, 0},
		[]int{0, 1},
	}
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == '0' {
				spaces++
			} else {
				// 此时 grid[i][j] == '1'
				for _, direction := range directions {
					newX := i + direction[0]
					newY := j + direction[1]
					// 先判断坐标合法，再检查右边一格和下边一格是否是陆地
					if newX < rows && newY < cols && grid[newX][newY] == '1' {
						unionFind.Union(i*cols+j, newX*cols+newY)
					}
				}
			}
		}
	}

	return unionFind.Count - spaces
}
