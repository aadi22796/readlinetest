package main

import (
    "fmt"
    "os"
    "bufio"
    "sync"
	"github.com/boltdb/bolt"
)

var mutex sync.RWMutex

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
		fil.WriteString(inp)
		fwrite(fil)
	}
	return
	}

func fopen(filename string, mode int, perm os.FileMode){


	f, err := os.OpenFile(filename, mode, perm)
	if err!=nil{
		panic(err)
	}
	mutex.Lock()
	//lockerr:=syscall.Flock(int(f.Fd()), syscall.LOCK_EX)
	fmt.Println("File is now locked,")
	//if lockerr!=nil{
	//	panic(lockerr)
	//}
	defer f.Close()

	fwrite(f)
	mutex.Unlock()
	//lockerr2:=syscall.Flock(int(f.Fd()), syscall.LOCK_UN)
	fmt.Println("File is now unlocked,")
	//if lockerr2!=nil{
	//	panic(lockerr2)
	//}
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
	//fopen("data", os.O_APPEND|os.O_WRONLY, 0644)
	dbupdate()
}

