package main

import (
	"container/list"
	"fmt"
)

func main() {
	var l *list.List
	l = list.New()
	l.PushFront(3)
	e := l.Front()
	fmt.Println(e.Value)
}
