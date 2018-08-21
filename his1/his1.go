package his1

import (
	"os"
	"strconv"
	"strings"
	"io/ioutil"
	"syscall"
)

func circbuf(fil *os.File, inp string) {
	
	//creating lastline file if not exists

	/*const s1 = 5
	//reading lastline file and storing in int
	f, err := os.OpenFile("/home/aadi22796/go/src/github.com/readlinetest/lastline", os.O_RDWR, 0600)
	if err != nil {
		panic(err)
	}
	bf := bufio.NewReader(f)
	l1, _ := bf.ReadString('\n')
	lastline, _ := strconv.Atoi(strings.Trim(l1, "\n"))
	*/

	data1, _ := ioutil.ReadFile("/home/aadi22796/go/src/github.com/readlinetest/data")
	slice1:= strings.Split(string(data1), "\n")
	s1,_:=strconv.Atoi(slice1[0])
	lastline,_:=strconv.Atoi(slice1[1])

	//conditions for circular insertion
	switch con := lastline / s1; con {
	case 0:
		{
			/*if lastline%s1 == s1-1 {
				fil.WriteString(strings.Trim(inp, "\n"))
				lastline++
				ioutil.WriteFile("/home/aadi22796/go/src/github.com/readlinetest/lastline", []byte(strconv.Itoa(lastline)+"\n"), 0600)

			} else {
				fil.WriteString(inp)
				lastline++
				ioutil.WriteFile("/home/aadi22796/go/src/github.com/readlinetest/lastline", []byte(strconv.Itoa(lastline)+"\n"), 0600)
			}*/
			slice1=append(slice1,strings.Trim(inp, "\n"))
			lastline++
			slice1[1]=strconv.Itoa(lastline)
			ioutil.WriteFile(fil.Name(), []byte(strings.Join(slice1, "\n")), 0600)
		}
	default:
		{
			d, _ := ioutil.ReadFile("/home/aadi22796/go/src/github.com/readlinetest/data")
			lines := strings.Split(string(d), "\n")
			lines1:=lines[:2]
			lines2 := lines[3:]
			lines2=append(lines2, strings.Trim(inp, "\n"))
			lastline++
			lines1[1]=strconv.Itoa(lastline)
			lines=append(lines1,lines2...)
			output := strings.Join(lines, "\n")
			ioutil.WriteFile("/home/aadi22796/go/src/github.com/readlinetest/data", []byte(output), 0600)
			ioutil.WriteFile("/home/aadi22796/go/src/github.com/readlinetest/lastline", []byte(strconv.Itoa(lastline)+"\n"), 0600)
		}
	}
}

//openfile
func Fopen(str string) {

	if _, err := os.Stat("/home/aadi22796/go/src/github.com/readlinetest/data"); os.IsNotExist(err) {
		f, _ := os.OpenFile("/home/aadi22796/go/src/github.com/readlinetest/data", os.O_RDWR|os.O_CREATE, 0600)
		f.WriteString("5\n0")}

	f, err := os.OpenFile("/home/aadi22796/go/src/github.com/readlinetest/data", os.O_APPEND|os.O_RDWR, 0600)
	if err!=nil{
		panic(err)
	}

	lockerr := syscall.Flock(int(f.Fd()), syscall.LOCK_EX)
	if lockerr != nil {
		panic(lockerr)
	}
	circbuf(f, str)
	unlockerr := syscall.Flock(int(f.Fd()), syscall.LOCK_UN)
	if unlockerr != nil {
		panic(lockerr)
	}
	f.Close()
	}
