package captcha

import (
	"fmt"
)

func Captcha(typeCaptcha, first, oparate, second int) string {
	ops := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	oprs := []string{"+", "-", "*"}

	var captchaString string
	if typeCaptcha == 1 {
		captchaString = fmt.Sprintf("%v %s %v", first, oprs[oparate-1], ops[second-1])
	} else {
		captchaString = fmt.Sprintf("%v %s %v", ops[first-1], oprs[oparate-1], second)

	}

	return captchaString
}
