package main

import (
	"fmt"
	"math/rand"
	"time"
)

var all = []string{"A", "B", "C", "D", "E", "F", "G"}

func main() {
	var s []string
	var j int64
	for j = 0; j < 5; j++ {
		tmp := funcName(j)
		for k := 0; k < len(tmp); k++ {
			s = append(s, tmp[k])
		}
	}
	fmt.Println()
	fmt.Println()
	//flag := 0
	for i := 0; i < len(s); i++ {
		//if flag == 5 {
		fmt.Println()
		fmt.Println()
		fmt.Println()
		//flag = 0
		//}
		fmt.Print(s[i])
		//fmt.Print("\t\t")
		//flag++
	}

}

func funcName(i int64) []string {
	var tmp []string
	r := rand.New(rand.NewSource(time.Now().Unix() * (i + 2)))
	for _, i := range r.Perm(len(all)) {
		val := all[i]
		tmp = append(tmp, val)
		fmt.Print(val)
		fmt.Print("\t")
	}
	fmt.Println()

	return tmp
}
