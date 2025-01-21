package challenge

import (
	"fmt"
	"time"
)

func Pt2() {
	result := ""

	defer func(t time.Time) {
		fmt.Printf("\nPart 2 result -> %v, runned in %s\n", result, time.Since(t))
	}(time.Now())
}
