package main

import "fmt"

func main() {

	st := stack{}

	for i := 0; i < 10; i++ {
		st.push(i)

	}
	fmt.Println(st)
	fmt.Println(st.pop())
	fmt.Println(st)

}

type stack struct {
	values []int
}

func (st *stack) push(value int) {
	st.values = append(st.values, value)
}

func (st *stack) pop() int {
	var x int

	if len(st.values) == 0 {
		x = 0

	} else {
		x = st.values[len(st.values)-1]
		st.values = st.values[:len(st.values)-1]

	}

	return x
}

func (st *stack) peek() int {
	x := int(st.values[len(st.values)-1])
	return x
}
