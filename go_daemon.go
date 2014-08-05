package main

import (
	"fmt"
	"log"
	"os/exec"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go reportLoop(&wg)
	// go otherLoop(&wg)
	wg.Wait()
}

func reportLoop(wg *sync.WaitGroup) {
	c := time.Tick(1 * time.Minute)
	for now := range c {
		var _ = now
		// fmt.Printf("report loop %v\n", now)
		cmd := exec.Command("/home/matt/.scout/scout_cron.sh")
		err := cmd.Run()
		if err != nil {
			log.Fatal(err)
		}
	}
	wg.Done()
}

func otherLoop(wg *sync.WaitGroup) {
	c := time.Tick(10 * time.Second)
	for now := range c {
		fmt.Printf("other loop %v\n", now)
	}
	wg.Done()
}
