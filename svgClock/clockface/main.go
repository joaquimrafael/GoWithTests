package main

import (
	"os"
	"time"

	clockface "github.com/joaquimprieto/gowithtests/svgClock"
)

func main() {
	t := time.Now()
	clockface.SVGWriter(os.Stdout, t)
}
