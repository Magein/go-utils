package main

import (
	"fmt"
	"github.com/magein/utils/mb"
	oss "github.com/magein/utils/os"
	times "github.com/magein/utils/time"
	"github.com/magein/utils/tool"
)

func main() {
	fmt.Println("hello world")
	fmt.Println(times.Date("/"))
	fmt.Println(times.Unix())
	fmt.Println("is int", tools.IsInt(1))
	fmt.Println("index of", tools.IndexOf([]int{1, 2}, 2))
	fmt.Println("in array ", tools.InArray("1,2,3,png", "png"))
	fmt.Println("file extension", oss.FileExtension("a.jp.png"))
	fmt.Println("file extension tolower", oss.FileExtension("a.jp.PNG"))
	fmt.Println("汉字长度", len("1a小马哥"))
	fmt.Println("汉字长度", mb.Len("1a小马哥"))
	fmt.Println("汉字截取", mb.Sub("1a小马哥", 2, 2))
}
