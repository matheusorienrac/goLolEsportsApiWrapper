package main

import (
	"fmt"
	"time"

	"github.com/matheusorienrac/goLolEsportsApiWrapper/wrapper"
)

func main() {
	now := time.Now()
	fmt.Println(wrapper.GetWindow(110413246204026169, now))
}
