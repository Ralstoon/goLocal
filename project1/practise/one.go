package main

import (
	"fmt"
	"sync/atomic"
)

var (
	x int    = 1
	y string = "wuyuhang"
)

func Fprint(x interface{}) {
	fmt.Printf("%T:\t%v\n", x, x)
}

type matrix struct {
	X, Y int
}

func funA(a, b int) string {
	return "wuyuhang"
}

func (m matrix) funcC() {
	Fprint(m.X)
}

func funcB(head string, fn func(int, int) string) {
	Fprint(head + fn(1, 2))
}

// 闭包
func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

type F float32

type A interface {
	funcC()
}

func func1() (v int) {
	v = 1
	return
}

func main() {
	// readBytes := make([]byte, 1)
	// var buffer bytes.Buffer
	// reader := bufio.NewReader(os.Stdin)

	// for {
	// 	// _, err := os.Stdin.Read(readBytes)
	// 	s, err := reader.ReadString('@')
	// 	if err != nil {
	// 		fmt.Println(err)
	// 		break
	// 	}
	// 	fmt.Println(s)

	// end := readBytes[0]
	// if end == '\n' {
	// 	os.Stdout.Write(buffer.Bytes())
	// 	buffer.Reset()
	// } else {
	// 	buffer.Write(readBytes)
	// }
	// }

	var countVal atomic.Value
	countVal.Store([]int{1, 3, 5, 7})
	func(one atomic.Value) {
		one.Store([]int{2, 4, 6, 8})
	}(countVal)
	fmt.Println(countVal.Load())

}
