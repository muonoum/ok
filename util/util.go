package util

import (
	"flag"
	"fmt"
	"os"
	"time"
)

var debug = flag.Bool("debug", false, "")

func Debug(args ...interface{}) {
	if *debug {
		fmt.Fprintln(os.Stderr, args...)
	}
}

func Debugf(s string, args ...interface{}) {
	if *debug {
		fmt.Fprintf(os.Stderr, s, args...)
	}
}

func MaybeAfter(timeout time.Duration, maybe bool) <-chan time.Time {
	if maybe {
		return time.After(timeout)
	}

	return nil
}

func Limit(source, min, max int) int {
	if source > max {
		return max
	} else if source < min {
		return min
	} else {
		return source
	}
}
