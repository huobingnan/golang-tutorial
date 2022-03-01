package main

import (
	"bufio"
	"fmt"
	"os"
)

func ReadFromStdin() {
	reader := bufio.NewReader(os.Stdin)
	line, b, err := reader.ReadLine()
	if err != nil {
		return
	}
	fmt.Printf("line : %s, isPrefix : %v\n", string(line), b)
}

func ReadViaScanf() {
	// 使用Scanf读取输入
	var name string
	fmt.Printf("Please input your name: ")
	flag, err := fmt.Scanf("%s %s", &name, &name)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("name = %v, flag = %v\n", name, flag)

}

func main() {
	//ReadFromStdin()
	ReadViaScanf()
}
