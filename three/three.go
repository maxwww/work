package two

import (
	"github.com/maxwww/work/one"
	"log"
)

func Three() string {
	log.Println("In Three")
	return one.One()
}
