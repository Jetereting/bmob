package bmob

import (
	"fmt"
	"testing"
)

func TestIsPay(t *testing.T) {
	fmt.Println(IsPay("测试未支付"))
	fmt.Println(IsPay("测试过期"))
	fmt.Println(IsPay("测试过期"))
}
