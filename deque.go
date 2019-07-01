package main

import "fmt"

func main() {

	de := deque{}

	for i := 0; i < 10; i++ {
		de.push_front(i)

	}
	fmt.Println(de)
	for i := 0; i < 10; i++ {
		de.push_back(i)

	}
	fmt.Println(de)
	fmt.Println(de.pop_back())
	fmt.Println(de)
	fmt.Println(de.pop_front())
	fmt.Println(de)
	fmt.Println(de.peekdeque('f'))
	fmt.Println(de.peekdeque('b'))
}

type deque struct {
	values [] int
}

func (de *deque) push_front(value int) {
	if len(de.values) == 0 {
		de.values = append(de.values, value)
	} else {
		de.values = append(de.values, 0)
		copy(de.values[1:], de.values[0:])
		de.values[0] = value
	}
}

func (de *deque) push_back(value int) {
	de.values = append(de.values, value)
}

func (de *deque) pop_back() int {
	x := de.values[len(de.values)-1]
	de.values = de.values[:len(de.values)-1]
	return x
}

func (de *deque) pop_front() int {
	x := de.values[0]
	de.values = append(de.values[:0], de.values[1:]...)
	return x
}

func (de *deque) peekdeque(side rune) int {
	var x int
	if side == 'b' {
		x = de.values[len(de.values)-1]
	} else if side == 'f' {
		x = de.values[0]
	}
	return x
}
