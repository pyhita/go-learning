package dependency_inject

import (
	"fmt"
	"io"
)

// 测试 go 依赖注入

func Greet(writer io.Writer, name string) {

	fmt.Fprintf(writer, "Hello, %s!", name)
}
