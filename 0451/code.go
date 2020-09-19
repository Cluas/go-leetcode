//451. 根据字符出现频率排序
//给定一个字符串，请将字符串里的字符按照出现的频率降序排列。
//
//示例 1:
//
//输入:
//"tree"
//
//输出:
//"eert"
//
//解释:
//'e'出现两次，'r'和't'都只出现一次。
//因此'e'必须出现在'r'和't'之前。此外，"eetr"也是一个有效的答案。
//示例 2:
//
//输入:
//"cccaaa"
//
//输出:
//"cccaaa"
//
//解释:
//'c'和'a'都出现三次。此外，"aaaccc"也是有效的答案。
//注意"cacaca"是不正确的，因为相同的字母必须放在一起。
//示例 3:
//
//输入:
//"Aabb"
//
//输出:
//"bbAa"
//
//解释:
//此外，"bbaA"也是一个有效的答案，但"Aabb"是不正确的。
//注意'A'和'a'被认为是两种不同的字符。
package leetcode

import (
	"bytes"
	"sort"
)

func frequencySort(s string) string {
	freqC := make(map[byte]int)
	for _, v := range []byte(s) {
		freqC[v]++
	}
	cm := make(map[int][]byte)

	for k, v := range freqC {
		cm[v] = append(cm[v], k)
	}

	var keys []int
	for k := range cm {
		keys = append(keys, k)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(keys)))
	var buf bytes.Buffer
	for _, v := range keys {
		for _, b := range cm[v] {
			for i := 0; i < v; i++ {
				buf.WriteByte(b)
			}
		}
	}
	return buf.String()

}
