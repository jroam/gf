package testdata

import "github.com/gogf/gf/os/gres"

func init() {
	if err := gres.Add("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"); err != nil {
		panic("add binary content to resource manager failed: " + err.Error())
	}
}
