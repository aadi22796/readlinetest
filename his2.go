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
	defer exec.Command("stty", "-F", "/dev/tty", "echo").Run()
	defer exec.Command("stty", "-F", "/dev/tty", "-cbreak", "min", "1").Run()


	var b = make([]byte, 100)
	for {
		os.Stdin.Read(b)
		if b[0] == 27 {
			if b[2] == 65 {
			return "Up"
		} else if b[2] == 66 {
			return "Down"
		} else  {
			return "esc"
		} }else {
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
	arval:=arrows()
	if arval == "Down" {
		if lines == maxline {
			fmt.Print("You're already on the last line: ")
			updown(lines, maxline)
		} else {
			lines = lines + 1
			updown(lines, maxline)
		}
	} else if arval == "Up" {
		if lines == 1 {
			fmt.Print("You're already on the first line: ")
			updown(lines, maxline)
		} else {
			lines = lines - 1
			updown(lines, maxline)
		}
	} else if arval == "esc" {
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
