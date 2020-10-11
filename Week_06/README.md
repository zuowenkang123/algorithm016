学习笔记    
1.模拟行走机器人    
```
坐标系如下图

                   北
                   0(0,1)
                   |
(左)西(-1,0) 3 ——  —— 1(1,0) 东（右）
                   |
                   2(0,-1)
                   北
转换为二维数组：
int[][] directions = {{0, 1}, {1, 0}, {0, -1}, {-1, 0}};

（0,1）-> 右->(1,0)  (数组索引为2)
（0,1）-> 左->(-1,0) (数组索引为3)

```
2.最大子序和 dp 递推 自底向上    
3.最长公共自学列，dp存在最有子结构  
````
【dp 模板】
// 定义
dp := make([][]int, len(text1)+1)
for i := 0; i <= len(text1); i++ {
    dp[i] = make([]int, len(text2)+1)
}
// 递推
for i := 1; i <= len(text1); i++ {
    // 情景区分
    if text1[i-1] == text2[j-1] {
        dp[i][j] = dp[i-1][j-1] + 1 
    } else {
        dp[i][j] = max(dp[i-1][j], dp[i][j-1])å
    }
}
// 返回结果
return dp[len(text1)][len(text2)]
````
4.不同路径，核心在于两边为固定，递推中间   
5.有障碍的不同路径  
6.搜索二维矩阵：二分法  
7.零钱兑换  
8.三角形最短路径：dp  
9.乘积最大子数组   
10.打家劫舍：dp 先得到0和1的情况，然后循环2到len-1 .团灭链接：https://labuladong.gitbook.io/algo/
````
【dp】
if len(nums) == 0 {
    return 0
}
if len(nums) == 1 {
    return nums[0]
}
dp := make([]int, len(nums))
dp[0] = nums[0]
dp[1] = max(nums[0], nums[1])

for i := 2; i < len(nums); i++ {
    dp[i] = max(dp[i-2]+nums[i], dp[i-1])
}
return dp[len(nums)-1]
````
11.打家劫舍二：拆解为包含头或包含尾。采用空间优化方式。
````
【空间优化，适用于比较最优】
cur, pre := 0, 0
	for _, v := range nums {
		cur, pre = max(pre+v, cur), cur
	}
	return cur
````