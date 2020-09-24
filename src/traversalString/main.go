package main

import "fmt"

func main() {
	traversalString()

	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
}

// 字符串底层是一个byte数组，所以可以和[]byte类型相互转换。字符串是不能修改的 字符串是由byte字节组成，所以字符串的长度是byte字节的长度。
// rune类型用来表示utf8字符，一个rune字符由一个或多个byte组成。
func traversalString() {
	s := "hello沙河"
	for i := 0; i < len(s); i++ { //byte
		fmt.Printf("%v(%c) ", s[i], s[i])
	}
	fmt.Println()
	for _, r := range s { //rune
		fmt.Printf("%v(%c) ", r, r)
	}
	fmt.Println()
}
