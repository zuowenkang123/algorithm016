学习笔记
1. 最小基因：从基因库出发，对比差异为1的，深度遍历。
2. 每个树行最大值
3. 单词接龙：深度遍历模板
````
【广度优先模板】
if indexOf(endWord, wordList) == -1 {
    return 0
}
step := 0
visited := make([]bool, len(wordList))
queue := []string{beginWord}
for len(queue) > 0 {
    l := len(queue)
    step++
    for i := 0; i < l; i++ {
        if queue[i] == endWord {
            return step
        }
        for j := 0; j < len(wordList); j++ {
            if !visited[j] && oneDiff(queue[i], wordList[j]) {
                queue = append(queue, wordList[j])
                visited[j] = true
            }
        }
    }
    queue = queue[l:]
}
return 0
````
4. 单词接龙二：困难
5. 岛屿个数：有一个岛计算后，周围的下沉，防止重复计算。
6. 扫雷游戏：踩雷直接爆炸，否则开始搜索。搜索八方，如果大于0，则填入，如果为0，则继续递归搜索。
7. 柠檬水找零：贪心算法，枚举各种情况，先凑大的，再凑小的。
8. 股票系列二：贪心法，后一个买进前一个卖出 
9. 开平方：二分查找
```
【二分模板】
if x == 0 {
    return 0
}
l, r := 1, x/2
for l < r {
    mid := (r+l)/2 + 1
    if mid*mid > x {
        r = mid - 1

    } else {
        l = mid
    }
}
return l
```

