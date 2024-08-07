package errorx

import (
	"fmt"
	"testing"
)

func TestAddErrorx(t *testing.T) {
	var (
		CODE1, _ = AddErrorx(0, "code1 text")
		CODE2, _ = AddErrorx(0, "code2 text")
		CODE3, _ = AddErrorx(2, "code3 text")
		CODE4, _ = AddErrorx(3, "code4 text")
		CODE5, _ = AddErrorx(4, "code5 text")
		CODE6, _ = AddErrorx(5, "code6 text")
		CODE7, _ = AddErrorx(6, "code7 text")
	)

	fmt.Printf("%+v\n", GetAllRegCodes())

	fmt.Println(CODE1.Code())
	fmt.Println(CODE2.WithError(fmt.Errorf("code2 with error")).Error())
	fmt.Println(CODE3.WithMessage("code3 with message").Error())
	fmt.Println(CODE4.WithMessagef("code%d with messagef", 4).Error())
	fmt.Println(CODE5.WithMessagef("code%d with messagef", 5).WithError(fmt.Errorf("code5 with error")).Error())
	fmt.Println(CODE6.WithMessagef("code%d with messagef", 6).WithMessage("code6 with message").Error())
	fmt.Println(CODE7.WithMessagef("code%d with messagef", 6).WithMessagef("code%d with messagef", 6).Error())
}
