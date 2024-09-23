package main

import (
	"fmt"
	"github.com/rs/zerolog"
)

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	var path = startup()
	k := Korwarder{make(map[string]PortForward), path}
	fmt.Printf("%v/n", k)
}
