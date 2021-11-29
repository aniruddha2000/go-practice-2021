package main

import (
	"fmt"
	"os"
	"os/exec"
	// "strings"
	// "syscall"
)

func main() {
	pid := os.Getpid()
	ppid := os.Getppid()
	fmt.Print("Process Id of current process: ", pid)
	fmt.Print("\nProcess Id of parent process: ", ppid)
	fmt.Println("\n")

	// check if ls command exists in the PATH
	// binary, lookErr := exec.LookPath("ls")
	// if lookErr != nil {
	// 	panic(lookErr)
	// }

	// // the program we want to execute
	// args := []string{"ls", "-a", "-l", "-h"}

	// execErr := syscall.Exec(binary, args, os.Environ())

	// syscall.Kill(os.Getpid(), syscall.SIGTERM)

	// // lspid := os.Getpid()
	// // lsppid := os.Getppid()
	// // fmt.Print("Process Id of current process: ", lspid)
	// // fmt.Print("\nProcess Id of parent process: ", lsppid)

	// // catch error if any
	// if execErr != nil {
	// 	panic(execErr)
	// }

	// args := []string{"ls", "-a", "-l", "-h"}
	cmd := exec.Command("ls")
	cmd.Run()
	lspid := cmd.Process.Pid
	fmt.Print(lspid)
}
