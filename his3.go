package main
import(
"os"
"fmt"
"os/exec"
)

func arrows2() string{
	// disable input buffering
exec.Command("stty", "-F", "/dev/tty", "cbreak", "min", "1").Run()
// do not display entered characters on the screen
exec.Command("stty", "-F", "/dev/tty", "-echo").Run()
	
	var b []byte = make([]byte, 100)
for {
os.Stdin.Read(b)
if b[0]==27{
	if b[2]==65{
			exec.Command("stty", "-F", "/dev/tty", "echo").Run()
	exec.Command("stty", "-F", "/dev/tty", "-cbreak", "min", "1").Run()
	return "Up Caught"
}else if b[2]==66{
				exec.Command("stty", "-F", "/dev/tty", "echo").Run()
	exec.Command("stty", "-F", "/dev/tty", "-cbreak", "min", "1").Run()
	return "Down Caught"
}else{
				exec.Command("stty", "-F", "/dev/tty", "echo").Run()
	exec.Command("stty", "-F", "/dev/tty", "-cbreak", "min", "1").Run()
return "esc caught"
}
	}else{
					exec.Command("stty", "-F", "/dev/tty", "echo").Run()
	exec.Command("stty", "-F", "/dev/tty", "-cbreak", "min", "1").Run()
	return "Other value caught"
}
}
}
func main(){
	fmt.Println(arrows2())
}
