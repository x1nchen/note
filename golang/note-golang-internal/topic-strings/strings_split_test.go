package topic_strings

import (
	"fmt"
	"strings"
)

// Split 和 SplitAfter 差别是 SplitAfter 分割出来的字符尾部会带上分隔符
func ExampleStringsSplit() {
	image := "dockerhub.com/project/service"

	fmt.Println(strings.SplitN(image, "/", 1))
	fmt.Println(strings.SplitN(image, "/", 2))
	fmt.Println(strings.Split(image, "/"))
	fmt.Println(strings.SplitAfterN(image, "/", 1))
	fmt.Println(strings.SplitAfterN(image, "/", 2))
	fmt.Println(strings.SplitAfter(image, "/"))
	// Output:
	// [dockerhub.com/project/service]
	// [dockerhub.com project/service]
	// [dockerhub.com project service]
	// [dockerhub.com/project/service]
	// [dockerhub.com/ project/service]
	// [dockerhub.com/ project/ service]
}
