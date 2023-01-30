package main

// 结构体里嵌套接口的目的：当前结构体实例可以用所有实现了该接口的其他结构体来初始化（即使他们的属性不完全一致）

import (
	"fmt"
	"strconv"
	"time"
)

// 接口：一组方法的集合
// OpenCloser 接口定义两个方法 返回 error
type OpenCloseFlag interface {
	Open() error
	Close() error
}

type Door struct {
	open bool
	lock bool
}

func (d *Door) Open() error {
	fmt.Println("door open...")
	d.open = true
	return nil
}

func (d *Door) Close() error {
	fmt.Println("door close...")
	d.open = false
	return nil
}

type AutoDoor struct {
	OpenCloseFlag        // 匿名接口
	delay         int    // 延迟多长时间开启
	msg           string // 自动开启时的警报
}

func (a *AutoDoor) OpenDoor() error {
	fmt.Println("Open after " + strconv.Itoa(a.delay) + " seconds")
	time.Sleep(time.Duration(a.delay) * time.Second)
	fmt.Println("Door is opening: ", a.msg)
	return nil
}

func main() {
	autoDoor := &AutoDoor{
		OpenCloseFlag: &Door{
			open: false,
			lock: false,
		},
		delay: 1,
		msg:   "Hello, it's opening...",
	}
	autoDoor.OpenDoor()
	if v, ok := autoDoor.OpenCloseFlag.(*Door); ok {
		v.Open()
		fmt.Printf("door properties status: %v\n", v)
	}
	if v, ok := autoDoor.OpenCloseFlag.(*Door); ok {
		v.Close()
		fmt.Printf("door properties status: %v\n", v)
	}
}
