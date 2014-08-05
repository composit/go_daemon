package main

import (
	"log"
	"os/exec"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	go listenForRealtime(&wg)
	go reportLoop(&wg)
	wg.Wait()
}

func reportLoop(wg *sync.WaitGroup) {
	c := time.Tick(1 * time.Second)
	for _ = range c {
		// fmt.Printf("report loop\n")
		cmd := exec.Command("/home/matt/.scout/scout_cron.sh")
		err := cmd.Run()
		if err != nil {
			log.Fatal(err)
		}
	}
	wg.Done()
}

func listenForRealtime(wg *sync.WaitGroup) {
	wg.Done()
}
