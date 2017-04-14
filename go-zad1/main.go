package main

import "fmt"
import "time"

func worker(done chan bool) {
	fmt.Print("working..")
	time.Sleep(time.Second * 3)
	fmt.Println("done")

	done <- true
}

func main() {
	done := make(chan bool, 1)
	go worker(done)

	<-done
}

//time in windows: Measure-Command { ./main.exe | Out-Host }
