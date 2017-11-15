package main

import (
	"bufio"
	"fmt"
	"strings"
)

func main() {
	input := "foo   bar    baz blockchain"
	// 创建scanner 扫描仪
	scanner := bufio.NewScanner(strings.NewReader(input))
	// 设置split方法，支持自定义
	scanner.Split(bufio.ScanWords)
	// 遍历循环
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
