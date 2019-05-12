package main
import (
  "fmt"
  "net"
  "os"
  "strconv"
  "sync"
)
func Usage() {
  fmt.Println("needs 3 arguments:\nip, port, numberMessages")
  os.Exit(1)
}
func check(e error, s string) {
  if e!=nil {
    panic(s)
  }
}
func jobToDo(address string, wg sync.WaitGroup) {
  con, e := net.Dial("tcp", address)
  check(e, "cannot establish connection")
  defer con.Close()
  defer wg.Done()

  _, e1 := con.Write([]byte("caca"))
  check(e1, "cannot send message")
  buff := make([]byte,1024)
  _, e2:= con.Read(buff)
  check(e2, "cannot recive message")

  fmt.Println("message send: caca")
  fmt.Println("message recived: ", string(buff))
}

func main() {
  arguments := os.Args
  if len(arguments) !=4 {
    Usage()
  }
  var wg sync.WaitGroup
  address := arguments[1]+":"+arguments[2]
  rep, _ := strconv.Atoi(arguments[3])
  for i:=0; i<rep;i++ {
    wg.Add(1)
    go jobToDo(address, wg)
  }
  wg.Wait()
  //fmt.Println(address)
}
