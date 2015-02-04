package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"
)

var runes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

var (
	count = flag.Int("count", 1000, "")
	min   = flag.Int("min", 50, "")
	max   = flag.Int("max", 100, "")
	delay = flag.Int("delay", 0, "")
)

func randomString(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = runes[rand.Intn(len(runes))]
	}
	return string(b)
}

func main() {
	flag.Parse()

	for i := 0; i < *count; i++ {
		n := int(rand.Float64()*(float64(*max)-float64(*min)) + float64(*min))
		s := randomString(n)
		fmt.Println(s)
		time.Sleep(time.Millisecond * time.Duration(*delay))
	}
}
