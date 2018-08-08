package main

import (
    "fmt"
    "os"
    "bufio"
	"github.com/boltdb/bolt"
	"github.com/theckman/go-flock"
	"strconv"
	"strings"
	"io/ioutil"

)

//var mutex sync.Mutex

func fwrite(fil *os.File){

	fmt.Println("Enter your string")
	rd:=bufio.NewReader(os.Stdin)
	inp, err:=rd.ReadString('\n')
	if err!=nil{
		panic(err)
	}
	if inp == "exit\n"{
		fmt.Println("Exited program")
	} else {
		//fil.WriteString(inp)
		circbuf(fil,inp)
		fwrite(fil)
	}
	return
	}

func circbuf(fil *os.File, inp string) {


	//creating lastline file if not exists
	if _, err := os.Stat("lastline"); os.IsNotExist(err) {
		f, _ := os.OpenFile("lastline", os.O_WRONLY|os.O_CREATE, 0600)
		f.WriteString("0\n")
	}
	const size =5
	//reading lastline file and storing in int
		f, err := os.OpenFile("lastline", os.O_RDWR, 0600)
		if err != nil {
			panic(err)
		}
		bf := bufio.NewReader(f)
		ll, _ := bf.ReadString('\n')
		lastline, _ := strconv.Atoi(strings.Trim(ll, "\n"))


	//conditions for circular insertion
	switch con:=lastline/size;con{
	case 0:{
		fil.WriteString(inp)
		lastline++
		fmt.Println("In case 0, lastline status:",lastline)
		ioutil.WriteFile("lastline",[]byte(strconv.Itoa(lastline)+"\n"),0600)
		fwrite(fil)
	}
	default:{
		var n=(lastline%size)-1
		d,_:=ioutil.ReadFile("data")
		fmt.Println(string(d))
		lines := strings.Split(string(d), "\n")
		if n==-1{
			lines[len(lines)-1]=inp
		}else{
			lines[n]=inp
			}
		output := strings.Join(lines, "\n")
		ioutil.WriteFile("data",[]byte(output),0600)
		lastline++
		fmt.Println("In case default, lastline status:",lastline)
		ioutil.WriteFile("lastline",[]byte(strconv.Itoa(lastline)+"\n"),0600)
		fwrite(fil)

	}

	}


	}




func fopen(filename string, mode int, perm os.FileMode){

	//mutex.Lock()
	//defer mutex.Unlock()
	fileLock := flock.NewFlock("/home/aadi22796/readlinetest/his1.go")
	_,errf:=fileLock.TryLock()
	if errf!=nil{
		panic(errf)
	}
	f, err := os.OpenFile(filename, mode, perm)
	if err!=nil{
		panic(err)
	}
	defer f.Close()

	//lockerr:=syscall.Flock(int(f.Fd()), syscall.LOCK_EX)
	//fmt.Println("File is now locked,")
	//if lockerr!=nil{
	//	panic(lockerr)
	//}

	fwrite(f)
	//mutex.Unlock()
	//lockerr2:=syscall.Flock(int(f.Fd()), syscall.LOCK_UN)
	//fmt.Println("File is now unlocked,")
	//if lockerr2!=nil{
	//	panic(lockerr2)
	//}
	fileLock.Unlock()
	}

//func dbopen(dbname string){
	//db,_:=bolt.Open(dbname,0644, nil)
	//defer db.Close()
	//dbupdate(db)

//}

func dbupdate(){
	db,_:=bolt.Open("mydb.db",0644, nil)
	defer db.Close()
	fmt.Println("Enter your string")
	rd:=bufio.NewReader(os.Stdin)
	inp, err:=rd.ReadString('\n')
	if err!=nil{
		panic(err)
	}
	if inp == "exit\n"{
		fmt.Println("Exited program")
	} else {
		tx,_:=db.Begin(true)
		defer tx.Rollback()
		b,_:=tx.CreateBucketIfNotExists([]byte("MyBucket5"))
		err:=b.Put([]byte("answer1"), []byte("^[[A"))
		if err!=nil{panic(err)}
		fmt.Println(b.Get([]byte("answer1")))
		//tx.DeleteBucket([]byte("MyBucket5"))
		tx.Commit()
	}
	return
}

func main() {
	fopen("data", os.O_APPEND|os.O_WRONLY, os.ModeExclusive)
	//dbupdate()
}

