package hash

import (
	"fmt"
	"testing"
)

/*给定两个字符串 s 和 t ，编写一个函数来判断 t 是否是 s 的字母异位词。

示例 1:

输入: s = "anagram", t = "nagaram"
输出: true
示例 2:

输入: s = "rat", t = "car"
输出: false
说明:
你可以假设字符串只包含小写字母。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/valid-anagram
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func Test(t *testing.T) {
	fmt.Println(isAnagram("anagram", "nagaram"))
	fmt.Println(IsAnagram2("anagram", "nagaram"))
}

func isAnagram(s string, t string) bool {
	text1 := map[byte]int{}
	for i := 0; i < len(s); i++ {
		text1[s[i]]++
	}
	for i := 0; i < len(t); i++ {
		if _, ok := text1[t[i]]; ok {
			text1[t[i]]--
			if text1[t[i]] == 0 {
				delete(text1, t[i])
			}
		} else {
			return false
		}
	}

	return len(text1) == 0
}

/*不使用任何内建的哈希表库设计一个哈希集合（HashSet）。

实现 MyHashSet 类：

void add(key) 向哈希集合中插入值 key 。
bool contains(key) 返回哈希集合中是否存在这个值 key 。
void remove(key) 将给定值 key 从哈希集合中删除。如果哈希集合中没有这个值，什么也不做。

示例：

输入：
["MyHashSet", "add", "add", "contains", "contains", "add", "contains", "remove", "contains"]
[[], [1], [2], [1], [3], [2], [2], [2], [2]]
输出：
[null, null, null, true, false, null, true, null, false]

解释：
MyHashSet myHashSet = new MyHashSet();
myHashSet.add(1);      // set = [1]
myHashSet.add(2);      // set = [1, 2]
myHashSet.contains(1); // 返回 True
myHashSet.contains(3); // 返回 False ，（未找到）
myHashSet.add(2);      // set = [1, 2]
myHashSet.contains(2); // 返回 True
myHashSet.remove(2);   // set = [1]
myHashSet.contains(2); // 返回 False ，（已移除）

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/design-hashset
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

func TestHashSet(t *testing.T) {
	hashSet := HashSetConstructor()
	hashSet.Add(1)
	hashSet.Add(2)
	fmt.Println(hashSet.Contains(1))
	fmt.Println(hashSet.Contains(3))
	hashSet.Add(2)
	fmt.Println(hashSet.Contains(2))
	hashSet.Remove(2)
	fmt.Println(hashSet.Contains(2))
}

func TestH(t *testing.T) {
	var a, b uint
	a = 2
	b = 1
	fmt.Println(b - a)
}
