package bmob

import (
	"fmt"
	"testing"
)

func TestIsPay(t *testing.T) {
	fmt.Println(IsPay("测试2"))
	fmt.Println(IsPay("测试"))
	fmt.Println(IsPay("测试"))
	fmt.Println(IsPay("测试"))
	fmt.Println(IsPay("测试"))
}
