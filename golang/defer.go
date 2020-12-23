package main

import "fmt"

func main() {
	defer func() {
		panic("触发异常1")
	}()

	defer func() { // 必须要先声明defer，否则recover()不能捕获到panic异常
		if err := recover(); err != nil {
			fmt.Println(err) //err 就是panic传入的参数
		}
		fmt.Println("打印中")
	}()

	panic("触发异常")
}
