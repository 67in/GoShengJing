package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

//func main() {
//	counts := make(map[string]int)
//	input := bufio.NewScanner(os.Stdin)
//	for input.Scan() {
//		counts[input.Text()]++
//	}
//	//用input.Scan()，即读入下一行，并移除行末的换行符；
//	//读取的内容可以调用input.Text()得到。
//	//Scan函数在读到一行时返回true，不再有输入时返回false。
//
//	for line, n := range counts {
//		if n > 1 {
//			fmt.Printf("%d\t%s\n", n, line)
//		}
//	}
//}

//func main(){
//	counts := make(map[string]int)
//	files := os.Args[1:]
//	if len(files) == 0{
//		countLines(os.Stdin, counts)
//	}else {
//		for _, arg := range files {
//			f, err := os.Open(arg)
//			if err != nil{
//				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
//				continue
//			}
//			countLines(f, counts)
//			f.Close()
//		}
//	}
//	for line , n := range counts{
//		if n>1{
//			fmt.Printf("%d\t%s\n", n, line)
//		}
//	}
//}
//
//func countLines(f *os.File, counts map[string]int){
//	input := bufio.NewScanner(f)
//	for input.Scan(){
//		counts[input.Text()]++
//	}
//}

func main(){
	counts := make(map[string]int)
	for _, filename := range os.Args[1:]{
		data, err := ioutil.ReadFile(filename)
		//ReadFile函数返回一个字节切片（byte slice），必须把它转换为string，才能用strings.Split分割
		if err != nil{
			fmt.Fprintf(os.Stderr, "dup3: %v\n",err)
			continue
		}
		for _, line := range strings.Split(string(data), "\n"){
			counts[line]++
		}
	}
}