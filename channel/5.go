package main

import "fmt"

// 解决前面的问题的方法还有一种就是使用有缓冲区的通道。
// 只要通道的容量大于零，那么该通道就是有缓冲的通道，通道的容量表示通道中能存放元素的数量。
// 就像你小区的快递柜只有那么个多格子，格子满了就装不下了，就阻塞了，等到别人取走一个快递员就能往里面放一个。
// 我们可以使用内置的 len 函数获取通道内元素的数量，使用 cap 函数获取通道的容量，虽然我们很少会这么做。

func main() {
	ch := make(chan int, 1) // 创建一个容量为 1 的有缓冲区通道
	ch <- 10
	fmt.Println("发送成功")
}
