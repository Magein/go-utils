package cli

import (
	"bufio"
	"os"
	"strings"
)

// ReadInput Get input values from the interactive command line
// Use "\r\n" to wrap lines in the window operating system
// Use "\n" to wrap lines in the linux operating system
func ReadInput() string {
	render := bufio.NewReader(os.Stdin)
	// 这里是当读取到指定的符号的时候结束输入  键盘表现为回车键  \n修改为1、a、3、aa的时候，在键盘上输入对应的值会结束
	input, _ := render.ReadString('\n')
	// 兼用win平台，且这两个顺序不能乱
	input = strings.Replace(input, "\r\n", "", 1)
	input = strings.Replace(input, "\n", "", 1)
	return input
}
