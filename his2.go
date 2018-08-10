package main

import (
	"io"
	"bytes"
	"os"
	"os/exec"
	"bufio"
	"fmt"
)

func arrows() string {
	// disable input buffering
	exec.Command("stty", "-F", "/dev/tty", "cbreak", "min", "1").Run()
	// do not display entered characters on the screen
	exec.Command("stty", "-F", "/dev/tty", "-echo").Run()

	//var b = make([]byte, 100)
	var b string
	for {
		fmt.Scan(&b)

		if b == "^[[A" {
			exec.Command("stty", "-F", "/dev/tty", "echo").Run()
			exec.Command("stty", "-F", "/dev/tty", "-cbreak", "min", "1").Run()
			return "Up"
		} else if b == "^[[B" {
			exec.Command("stty", "-F", "/dev/tty", "echo").Run()
			exec.Command("stty", "-F", "/dev/tty", "-cbreak", "min", "1").Run()
			return "Down"
		} else if b == "^[[C" {
			exec.Command("stty", "-F", "/dev/tty", "echo").Run()
			exec.Command("stty", "-F", "/dev/tty", "-cbreak", "min", "1").Run()
			return "esc"
		} else {
			exec.Command("stty", "-F", "/dev/tty", "echo").Run()
			exec.Command("stty", "-F", "/dev/tty", "-cbreak", "min", "1").Run()
			return "Other value caught"
		}

	}
}
func lineCounter(r io.Reader) (int) {
	buf := make([]byte, 32*1024)
	count := 0
	lineSep := []byte{'\n'}

	for {
		c, err := r.Read(buf)
		count += bytes.Count(buf[:c], lineSep)

		switch {
		case err == io.EOF:
			return count
		}
	}
}

func ReadLine(inp string, n int) (string, error) {
	f, err := os.Open(inp)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	bf := bufio.NewReader(f)
	var line string
	for lnum := 0; lnum < n; lnum++ {
		line, err = bf.ReadString('\n')
	}
	if err != nil {
		return "Error", err
	}

	return line, err

}

func updown(lines int, maxline int) {
	a, _ := ReadLine("data", lines)
	fmt.Print(a)
	//fmt.Println("Enter + to go down, - to go up")
	if arrows() == "Down" {
		if lines == maxline {
			fmt.Print("You're already on the last line: ")
			updown(lines, maxline)
		} else {
			lines = lines + 1
			updown(lines, maxline)
		}
	} else if arrows() == "Up" {
		if lines == 1 {
			fmt.Print("You're already on the first line: ")
			updown(lines, maxline)
		} else {
			lines = lines - 1
			updown(lines, maxline)
		}
	} else if arrows() == "esc" {
		fmt.Println("Exited Program")
	} else {
		fmt.Println("Invalid Syntax")
		updown(lines, maxline)
	}
}
func main() {
	fil, err := os.Open("data")
	if err != nil {
		panic(err)
	}
	var lncnt = lineCounter(fil)
	updown(lncnt, lncnt)
}
