package sports

import "fmt"

type Football struct {
	PCount int
}

func (f *Football) Pay(gName string) {
	fmt.Println("Playing the ", gName)
}
