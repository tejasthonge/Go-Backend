package sports

import "fmt"

type Cricket struct {
	PCount int
}

func (c *Cricket) Paly(gName string) {
	fmt.Println("Palying the ", gName)
}
