package main

import (

	"os"

	"strconv"
)

func main(){
//	lastline,err:=strconv.Atoi(strings.Trim("0\n","\n"))
//	if err!=nil{panic(err)}
//	fmt.Println(reflect.TypeOf(lastline))

	//if _, err := os.Stat("lastline"); os.IsNotExist(err) {
		f,_:=os.OpenFile("lastline",os.O_RDWR|os.O_CREATE,0600)
	f.WriteString(strconv.Itoa(3)+"\n")
	//}
}
