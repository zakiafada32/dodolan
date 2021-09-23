package main

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

func main() {
	fmt.Println(time.Now())
	fmt.Println(uuid.NewString())
}
