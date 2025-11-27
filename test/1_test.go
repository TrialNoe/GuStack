package test

import (
	"GuStack"
	"fmt"
	"testing"
)

func Test1(t *testing.T) {
	stack := GuStack.New()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())

}

func Benchmark1(b *testing.B) {
	stack := GuStack.New()
	for i := 0; i < b.N; i++ {
		stack.Push(i)
		stack.Pop()
	}
	b.ReportAllocs()

}
