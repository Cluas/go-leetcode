//205. 同构字符串
//给定两个字符串 t 和 t，判断它们是否是同构的。
//
//如果 t 中的字符可以被替换得到 t ，那么这两个字符串是同构的。
//
//所有出现的字符都必须用另一个字符替换，同时保留字符的顺序。两个字符不能映射到同一个字符上，但字符可以映射自己本身。
//
//示例 1:
//
//输入: t = "egg", t = "add"
//输出: true
//示例 2:
//
//输入: t = "foo", t = "bar"
//输出: false
//示例 3:
//
//输入: t = "paper", t = "title"
//输出: true
//说明:
//你可以假设 t 和 t 具有相同的长度。
package leetcode

func isIsomorphic(s string, t string) bool {
	s2t := make(map[byte]byte)
	t2s := make(map[byte]byte)

	for i := 0; i < len(s); i++ {
		if v, ok := s2t[s[i]]; ok && v != t[i] {
			return false
		}
		if v, ok := t2s[t[i]]; ok && v != s[i] {
			return false
		}
		s2t[s[i]] = t[i]
		t2s[t[i]] = s[i]
	}

	return true
}
