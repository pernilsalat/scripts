package main
import(
  "fmt"
  "os"
  "bufio"
  "encoding/hex"
)
func readfile(name string, p chan<- []byte) {
  f, _ := os.Open(name)
  defer f.Close()
  scanner := bufio.NewScanner(f)
  for scanner.Scan() {
    text1b, _:= hex.DecodeString(scanner.Text())
    p <- text1b
  }
  p <- []byte("caca725")
}

func main() {
  p := make(chan []byte)
  go readfile("ins.txt", p)
  var score float64 = 0
  var result string

  for {
    message := <-p
    if string(message)=="caca725" {
      break
    }
    res, sc := Decode(message)
    if sc > score{
      score = sc
      result = res
    }
  }
  fmt.Println(result)

}
