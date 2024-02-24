package main

import (
	"fmt"

	"github.com/stevenwilkin/deribit-funding/deribit"
)

func main() {
	d := &deribit.Deribit{}

	for f := range d.Funding() {
		fmt.Println(f)
	}
}
