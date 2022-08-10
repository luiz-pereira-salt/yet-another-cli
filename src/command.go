package main

// https://blog.kowalczyk.info/article/wOYk/advanced-command-execution-in-go-with-osexec.html

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"runtime"
)

func main() {
	cmd := exec.Command("ls", "-lah")
	if runtime.GOOS == "windows" {
		cmd = exec.Command("tasklist")
	}

	var stdoutBuf, stderrBuf bytes.Buffer
	cmd.Stdout = io.MultiWriter(os.Stdout, &stdoutBuf)
	cmd.Stderr = io.MultiWriter(os.Stderr, &stderrBuf)

	err := cmd.Run()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
	outStr, errStr := string(stdoutBuf.Bytes()), string(stderrBuf.Bytes())
	fmt.Printf("\nout:\n%s\nerr:\n%s\n", outStr, errStr)
}