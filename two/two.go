package two

import (
	"github.com/maxwww/work/one"
	"log"
)

func Two() string {
	log.Println("in Two")
	return one.One()
}
