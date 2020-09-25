学习笔记

1.括号生成：转换成左右括号匹配问题，然后递归解决。    
分析：  
    1）括号问题本身包含条件     
        a）左右括号先后问题，从时间角度就是左括号先于右括号，从个数角度就是左括号多余右括号。  
        b）两个类型 左右括号的处理不同于单一类型，存在多类型，即左右，类似的是二叉树，左右子树。因此在下沉由两个分支处理。  
    2) 终止条件：树的终止条件一般都是为空，排列放置的终止条件一般是满足个数，例如n个数放完。括号的是，左括号和右括号都为n个。
    3）下沉变换 考虑下沉之后的参数变换，通常是减一，多类型需要进行多类型状态变换分析。括号是左括号用了，左括号减一，右括号用了，右括号减一。
    4）回溯：核心在于每一层都会被回溯，且回溯的时候是栈回溯，先回溯底层状态，然后一层层向上回溯。因此每一层的回溯代码都是准确的回溯当前的状态。
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
分析：1）树的反转是同层反转，同层的处理采用队列进行层次遍历很合适。2）递归方式，树的题目都是处理左右子树，而非左右值。因此先递归得到左右子树，然后进行当前交换。该方式类似后续遍历，先得到左右子树，然后赋值当前root左右节点。
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
3.最小深度 1）递归 后续遍历，获取左右子树最小值，然后比较谁最小 2）广度搜索 
分析：1）递归 同后续遍历，先求左右子树最小值，然后比较最小 2）层次遍历 a) 初始化push根 2）每层需要删除，因此需要计数 3）pop数据 4）判断左右子树为空，为空则得到结果，否则继续插入数据。    
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