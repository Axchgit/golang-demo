package main

import (
	"fmt"
	"github.com/Axchgit/golang-demo/pkg/stringutils"
	"github.com/Axchgit/golang-demo/internal/config"
)

func main() {
	fmt.Println("=== Hello Go! ===")
	
	// 使用 pkg 下的工具
	name := "abcdefg"
	fmt.Printf("Hello, %s!\n", stringutils.Reverse(name))
	
	// 使用 internal 包
	cfg := config.Load()
	fmt.Printf("Config loaded: %+v\n", cfg)
}
