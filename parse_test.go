package expression

import (
	"testing"
)

func TestParse(t *testing.T) {
	exps := []string{
		"age+20",
		"age-10",
		"age*2",
		"age/2",

		"age+10%",

		"年龄+10",

		"(年龄+10)*2",

		"!@#+10",
		"age&10",
	}
	for _, exp := range exps {
		_, err := Parse(exp)
		if err != nil {
			t.Error(err)
		}
	}
}
