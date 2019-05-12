package main
import(
  "fmt"
  "encoding/hex"
  "bufio"
  "os"
)
func main() {
  var key string
  fmt.Scanln(&key)

  scanner := bufio.NewScanner(os.Stdin)
  for scanner.Scan() {
    res := make([]byte, len(scanner.Text()))
    for i,v:=range(scanner.Text()) {
      res[i] = byte(v)^key[i%len(key)]
    }

    fmt.Println(scanner.Text()+"\n"+ hex.EncodeToString(res))
  }
}
