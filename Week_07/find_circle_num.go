package Week_07

// https://leetcode-cn.com/problems/friend-circles/solution/
// 2020-11-01

func findCircleNum(M [][]int) int {
	var uf UnionFind
	uf.Init(len(M))

	for i, row := range M {
		for j, v := range row {
			if v == 1 {
				uf.Union(i, j)
			}
		}
	}
	return uf.count
}

// 并查集
type UnionFind struct {
	count  int
	parent []int
}

// 初始化都设置默认-1
func (u *UnionFind) Init(count int) {
	u.count = count
	u.parent = make([]int, count)
	for i, _ := range u.parent {
		u.parent[i] = -1
	}
}

// 查找元素，递归查找父元素
func (u *UnionFind) Find(node int) int {
	if u.parent[node] < 0 {
		return node
	}
	result := u.Find(u.parent[node])
	u.parent[node] = result
	return result
}

// 合并集合 todo
func (u *UnionFind) Union(node1 int, node2 int) {
	root1 := u.Find(node1)
	root2 := u.Find(node2)
	// 父节点相同则是合并
	if root1 == root2 {
		return
	}
	//
	if u.parent[root1] < u.parent[root2] {
		u.parent[root2] = root1
	} else {
		if u.parent[root1] == u.parent[root2] {
			u.parent[root2]--
		}
		u.parent[root1] = root2
	}
	u.count--
}

// dfs
func findCircleNumDfs(M [][]int) int {
	visited := make(map[int]int)
	count := 0
	for i := 0; i < len(M); i++ {
		if _, ok := visited[i]; !ok {
			dfsCircleNum(M, visited, i)
			count++
		}

	}
	return count
}
func dfsCircleNum(M [][]int, visited map[int]int, i int) {
	for j := 0; j < len(M); j++ {
		if _, ok := visited[j]; !ok && M[i][j] == 1 {
			visited[j] = 1
			dfsCircleNum(M, visited, j)
		}
	}
}
