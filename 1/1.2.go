package main

import (
	"fmt"
	"os"
	"strconv"
)
//
//func main() {
//	var s, sep string
//	for i := 1; i < len(os.Args); i++ {
//		s += sep + os.Args[i]
//		sep = " "
//	}
//	fmt.Println(s)
//}

func main(){
	//var s string
	//for _, arg := range os.Args[1:]{
	//	s += sep + arg
	//	sep = " "
	//}
	//s = os.Args[0]

	for i, arg := range os.Args[1:]{
		fmt.Println(strconv.Itoa(i) +" "+ arg)
	}
	//fmt.Println(s)
}