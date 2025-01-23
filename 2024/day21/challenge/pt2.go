package challenge

import (
	"fmt"
	"time"

	"github.com/atotto/clipboard"
)

func Pt2() {
	result := 0

	defer func(t time.Time) {
		fmt.Printf("\nPart 2 result -> %v, runned in %s\n", result, time.Since(t))
		_ = clipboard.WriteAll(fmt.Sprintf("%v", result))
	}(time.Now())
}
