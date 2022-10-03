package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {

	host, _ := os.Hostname()

	c := make(chan string)
	for {
		go getMessage(c, host)
		fmt.Println(<-c)
	}
}

func getMessage(c chan string, host string) {
	wait := getRandomWaitInSeconds()
	time.Sleep(wait)
	c <- fmt.Sprintf("%s | %s | hello! | %s", host, time.Now(), wait)
}

func getRandomWaitInSeconds() time.Duration {
	min := 1
	max := 5
	rand.Seed(time.Now().UnixNano())
	sec := rand.Intn(max-min) + min
	return time.Duration(sec) * time.Second
}
