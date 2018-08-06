package main
import ("bufio"
"fmt"
"os")
func main(){
in := bufio.NewReader(os.Stdin)
b, err := in.ReadByte()
fmt.Println("Key code:", b, err)
}
