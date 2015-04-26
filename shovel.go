package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"
)

//we'll use this to check for errors
func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	//in flight futures default
	maxFutures := 100
	currentFutures := 0
	var futuresChannel chan int = make(chan int)

	verbose := false

	file, err := os.Open("../data/urls.txt")
	check(err)

	defer file.Close()

	reader := bufio.NewReader(file)
	scanner := bufio.NewScanner(reader)

	thisUrl := ""
	for scanner.Scan() {

		thisUrl = scanner.Text()
		fmt.Println(thisUrl)

		currentFutures++
		fmt.Println(currentFutures)
		go callWget(thisUrl, futuresChannel, verbose)

		for currentFutures > maxFutures {
			currentFutures = currentFutures - <-futuresChannel
			fmt.Println(currentFutures)
			time.Sleep(time.Second * 1)
		}

	}
}

//call wget
func callWget(url string, futuresChannel chan int, verbose bool) /* *exec.Cmd */ {
	url = "www." + url
	cmd := exec.Command("wget", "--random-wait", "-r", "-t", "10", "-A", "-R", "-nc", "--tries=5", "--wait=10", url)
	//if verbose drop wget output into stdout
	if verbose {
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
	}

	//Kick off the job and increment futures
	//	currentFutures++
	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}

	//	return cmd
	log.Printf("Waiting for wget to finish...")
	err = cmd.Wait()
	log.Printf("Command wget with error: %v", err)
	fmt.Println(cmd)

	//I'm Done!
	futuresChannel <- 1
}
