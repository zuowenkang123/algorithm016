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
5. 岛屿个数：