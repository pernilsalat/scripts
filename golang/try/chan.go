package main
import "fmt"
import "os"
import "os/exec"
import "syscall"


func main() {
  message := make(chan string)
  signals := make(chan bool)

  select {
  case msg := <-message:
    fmt.Println("recived message", msg)
  default:
    fmt.Println("no message recived")
  }
  msg := "hi"
  select {
  case message<-msg:
    fmt.Println("message sent", msg)
  default:
    fmt.Println("no message sent")
  }

  select {
  case msg := <-message:
    fmt.Println("message recived", msg)
  case sig := <-signals:
    fmt.Println("recived signal", sig)
  default:
    fmt.Println("no activity")
  }

  go func() {
    binary, err := exec.LookPath("ls")
    if err != nil{
      panic(err)
    }
    args := []string{"ls", "-a", "-l", "-h"}
    env := os.Environ()
    execErr := syscall.Exec(binary,args,env)
    if execErr != nil{
      panic(execErr)
    }
  }()

  
  fmt.Println("Cya!!")
}
