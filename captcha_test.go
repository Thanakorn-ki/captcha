package captcha_test

import (
	"testing"

	"github.com/salapao2136/captcha"
)

func TestCaptchaTwoPlusFive(t *testing.T) {
	expected := "2 * fives"
	result := captcha.Captcha(1, 2, 3, 5)
	if expected != result {
		t.Errorf("it should say %q but get %q", expected, result)
	}

}

func TestCaptchaNine(t *testing.T) {
	expected := "nine + 3"
	result := captcha.Captcha(2, 9, 1, 3)
	if expected != result {
		t.Errorf("it should say %q but get %q", expected, result)
	}
}

func TestCaptchaSevenMinusThree(t *testing.T) {
	expected := "seven - 3"
	result := captcha.Captcha(2, 7, 2, 3)
	if expected != result {
		t.Errorf("it should say %q but get %q", expected, result)
	}
}
