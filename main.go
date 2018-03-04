package main

import (
	"github.com/sony/sonyflake"
)

// for task id
var sf *sonyflake.Sonyflake

func init() {
	var st sonyflake.Settings
	sf = sonyflake.NewSonyflake(st)
	if sf == nil {
		panic("sonyflake not created")
	}
}

func main() {

	//	r := engine()

	// r.Run(":3000")

	log()
}
