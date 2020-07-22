### 回溯算法套路
解决回溯问题，实际上就是一个决策树的遍历过程。
#### 关键字
- 选择列表
- 路径
- 结束条件

#### 模版
```
for chose in choses:
    # do chose
    remove chose from choses
    path.add(chose)
    backtrack(path, choses)
    # rollback chose
    path.remove(chose)
```
