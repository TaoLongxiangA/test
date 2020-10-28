package interfaceTest

import (
	"fmt"
	"log"
)

type dog interface {
	dogVoice() string
}

type cat interface {
	catVoice() string
}

type dogMIT string

func (m dogMIT) catVoice() string {
	return "mia"
}

func (d dogMIT) dogVoice() string {
	return "wan"
}

func main() {
	var d dog
	if c, ok := d.(cat); ok {
		fmt.Println(c.catVoice())
	} else {
		log.Println(c)
		log.Println(ok)
		fmt.Println("bad")
	}
}
