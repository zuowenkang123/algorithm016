学习笔记

1.括号生成：转换成左右括号匹配问题，然后递归解决   
```
【分治模板】
func divide_conquer(Problem problem, ) {

if (problem == nil) {
    res = process_last_result();
    return res
}
subProblems = split_problem(problem)

res0 := divide_conquer(subProblems[0])
res1 := divide_conquer(subProblems[1])

result = process_result(res0, res1)

return result
}
``` 
```
【分治】
// termination
if left == n && right == n {
    res = append(res, str)
    return
}

// process
// drill down
if left < n {
    generate_parenthesis(left+1, right, n, str+"(")
}
if right < left {
    generate_parenthesis(left, right+1, n, str+")")
}
// reverse

return
```
	  
2.反转二叉树：1）递归 终止条件 处理+下沉 2）迭代 起始装入根，非空循环，弹出处理，非空入列，空结束。 
```
【bfs】
if root == nil {
    return nil
}
queue := TreeNodeQueue{}
// 加入初始值，开始非空循环
queue.push(root)
for queue.size() != 0 {
    // 弹出处理
    node := queue.pop()
    // 交换
    node.Left, node.Right = node.Right, node.Left
    // 非空入列
    if node.Left != nil {
        queue.push(node.Left)
    }
    if node.Right != nil {
        queue.push(node.Right)
    }
} 
```  
3.最小深度 1）递归 4中情况的处理 2）深度搜索     
```
【dfs】
if k == 0 {
    // 需要copy
    comb := make([]int, len(arr))
    copy(comb, arr)
    resInt = append(resInt, comb)
    return
}
for i := start; i <= n-k+1; i++ {
    arr = append(arr, i)
    // 装入当前后，下一层k-1,i+1
    dfs1(n, k-1, i+1, arr)
    // 尝试装完后，回溯处理上一层 https://leetcode-cn.com/problems/permutations/solution/hui-su-suan-fa-python-dai-ma-java-dai-ma-by-liweiw/
    arr = arr[:len(arr)-1] // 当前回溯，所有都在回溯
}
```
4.替换空格 循环字符串，替换    
5.生成树：inorder决定左右子树位置，preorder决定root位置。  
6.验证二叉树：预留树中最大和最小，然后递归左右子树    
7.倒序打印链表：栈、递归、正序反转   
8.合并二叉树：深度和广度优先  
9.八皇后：1.依次放入n个 2.终止条件 3.设置状态下沉 4.恢复状态继续。状态是根据递归栈先恢复最底层，最后恢复最上层。       
10.全排列 
```
【回溯模板】
result = []
func backtrack(路径，选择列表) {
   if 满足结束条件 {
       result.add(路径)
        return
   }
   for 选择 in 选择列表 {
       做选择
       backtrack(路径，选择列表)
       撤销选择
   }
}
```
11.重复全排列，排序，剪枝