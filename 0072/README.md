### 72. 编辑距离


### 题目
给你两个单词 word1 和 word2，请你计算出将 word1 转换成 word2 所使用的最少操作数 。

你可以对一个单词进行如下三种操作：

- 插入一个字符
- 删除一个字符
- 替换一个字符
 

示例 1：

```
输入：word1 = "horse", word2 = "ros"
输出：3
解释：
horse -> rorse (将 'h' 替换为 'r')
rorse -> rose (删除 'r')
rose -> ros (删除 'e')
```
示例 2：

```
输入：word1 = "intention", word2 = "execution"
输出：5
解释：
intention -> inention (删除 't')
inention -> enention (将 'i' 替换为 'e')
enention -> exention (将 'n' 替换为 'x')
exention -> exection (将 'n' 替换为 'c')
exection -> execution (插入 'u')

```
 

提示：

0 <= word1.length, word2.length <= 500
word1 和 word2 由小写英文字母组成


### 思路

#### 动态规划：
dp[i][j] 代表word1第i个位置的字符转换到word2第j个字符需要的操作最小步数。
分以下两种情况：
- word1[i] == word2[j], 相等说明无需操作，故dp[i][j]=dp[i-1][j-1]
- word1[i] != word2[j], 那么相应的有三种操作来处理这种情况
	- 插入：在word1[i]后面插入word2[j], 相等抵消，故dp[i][j] = dp[i][j-1]
	- 删除：dp[i][j] = dp[i][j-1]
	- 替换: dp[i][j] = dp[i-1][j-1]

进而我们可以推导出状态迁移表达式：
- 当word1[i] == word2[j]:
	- $D[i][j]=min(D[i][j−1]+1,D[i−1][j]+1,D[i−1][j−1])=1+min(D[i][j−1],D[i−1][j],D[i−1][j−1]−1)$
- 不等时:
    - $D[i][j]=1+min(D[i][j−1],D[i−1][j],D[i−1][j−1])$

### 推荐
[https://alchemist-al.com/algorithms/edit-distance](https://alchemist-al.com/algorithms/edit-distance)
