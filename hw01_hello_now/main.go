package main

import (
	"fmt"
	"log"
	"time"

	"github.com/beevik/ntp"
)

func main() {
	// Place your code here
	t, err := ntp.Time("0.beevik-ntp.pool.ntp.org")
	if err != nil {
		log.Fatalf("error from calling ntp.time function %s", err)
	}
	fmt.Printf(`current time: %s
exact time: %s
`, time.Now().Round(0), t.Round(0))
}
