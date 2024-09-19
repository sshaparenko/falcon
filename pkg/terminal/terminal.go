package terminal

import (
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"syscall"

	"github.com/creack/pty"
)

type LogReader struct {
	r     io.Reader
	w     io.Writer
	terms bool
}

func (l *LogReader) Read(p []byte) (n int, err error) {

	var buff []byte = p

	n, _ = l.r.Read(buff)

	switch l.terms {
	case true:
		if n > 0 {
			if n, err := l.w.Write(buff[:n]); err != nil {
				return n, err
			}
		}
		l.terms = false
		return 0, nil
	case false:
		if n != 0 && buff[n-1] == 32 && buff[n-2] == 36 {
			l.terms = true
		}
	}

	return
}

func CmdReader(r io.Reader, w io.Writer) io.Reader {
	return &LogReader{r, w, false}
}

func Run() {
	cmd, ptmx, err := createPTY()
	defer func() { _ = ptmx.Close() }()
	if err != nil {
		log.Fatalf("Failed to create PTY: %s\n", err.Error())
	}

	handleAbort(cmd)
	logFile, err := logInput(ptmx)
	if err != nil {
		log.Fatalf("Error occured while logging data: %s\n", err.Error())
	}
	defer func() { _ = logFile.Close() }()
}

func logInput(ptmx *os.File) (*os.File, error) {
	logFile, err := os.Create("pty_log")
	if err != nil {
		return nil, err
	}

	tee := CmdReader(ptmx, logFile)

	go func() {
		_, err := io.Copy(os.Stdout, tee)
		if err != nil {
			log.Fatalf("Failed to read from pty: %s\n", err.Error())
		}
	}()

	_, err = io.Copy(ptmx, os.Stdin)
	if err != nil {
		return nil, err
	}
	return logFile, nil
}

func createPTY() (*exec.Cmd, *os.File, error) {
	cmd := exec.Command("/bin/bash", "-i")
	ptmx, err := pty.Start(cmd)

	if err != nil {
		return nil, nil, err
	}
	return cmd, ptmx, err
}

func handleAbort(cmd *exec.Cmd) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		if err := cmd.Process.Kill(); err != nil {
			fmt.Println("Failed to kill a shell process")
		}
		os.Exit(1)
	}()
}
