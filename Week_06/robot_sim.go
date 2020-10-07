package Week_06

// https://leetcode-cn.com/problems/walking-robot-simulation/description/
// 2020-10-07

type pos struct {
	x int
	y int
}

/*
 *要点1：如何表示东南西北
 *要点2：如何表示左右转向
 *要点3：如何判断是否遇到障碍
 */
func robotSim(commands []int, obstacles [][]int) int {
	res := 0
	var x, y int

	dir_x := 0
	dir_y := 1 // 初始面向北方

	var obs = make(map[pos]struct{}, 0)
	for _, ls := range obstacles {
		obs[pos{ls[0], ls[1]}] = struct{}{}
	}
	// 执行命令集
	for _, c := range commands {
		// 大于0，直走，遇到阻碍停下来
		if c > 0 {
			// 移动位置
			for ; c > 0; c-- {
				if _, ok := obs[pos{x + dir_x, y + dir_y}]; ok {
					break
				}
				x += dir_x
				y += dir_y
			}
			res = max(res, x*x+y*y)
		} else if c == -1 { // 右转向
			dir_x, dir_y = dir_y, -dir_x
		} else { // 左转向
			dir_x, dir_y = -dir_y, dir_x
		}
	}
	return res
}

func robotSim1(commands []int, obstacles [][]int) int {
	// 方向控制
	directions := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	// 阻碍map
	obs := make(map[pos]struct{}, 0)
	for _, ls := range obstacles {
		obs[pos{ls[0], ls[1]}] = struct{}{}
	}
	// 初始值
	res, x, y, direction := 0, 0, 0, 0

	// 执行命令集
	for _, c := range commands {
		if c == -1 { // Turns left
			direction = (direction + 1) % 4
		} else if c == -2 { // Turns right
			direction = (direction + 3) % 4
		} else { // Moves forward commands[i] steps
			step := 0
			for step < c {
				if _, ok := obs[pos{x + directions[direction][0], y + directions[direction][1]}]; ok {
					break
				}
				x += directions[direction][0]
				y += directions[direction][1]
				step++
			}
		}
		res = max(res, x*x+y*y)

	}
	return res
}
