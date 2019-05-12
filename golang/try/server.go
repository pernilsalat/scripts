package main
import(
  "fmt"
  "net"
  "os"
)
func Usage() {
  fmt.Println("needs 1 arguments:\n port")
  os.Exit(1)
}
func check(e error, s string) {
  if e!=nil {
    panic(s)
  }
}
func handle(conn net.Conn) {
  buff := make([]byte, 1024)
  _, e := conn.Read(buff)
  check(e, "canot read a message")
  _, e1 := conn.Write([]byte("de vaca"))
  check(e1, "canot send a message")

  fmt.Println("message recived: ", string(buff))
  fmt.Println("message send: de vaca  ")
  conn.Close()
}
func main() {
  arguments := os.Args
  if len(arguments) != 2 {
    Usage()
  }
  ln, e := net.Listen("tcp", ":"+arguments[1])
  defer ln.Close()
  check(e, "canot listen any connection")

  for {
    conn, e := ln.Accept()
    check(e, "cannot accept a new connection")
    go handle(conn)
  }
}
