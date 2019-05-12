package main
import(
  hex "encoding/hex"
  "fmt"
)

func main() {
  var text1, text2 string
  fmt.Scanln(&text1)
  fmt.Scanln(&text2)
  text1b, _:= hex.DecodeString(text1)
  text2b, _:= hex.DecodeString(text2)

  result := make([]byte, len(text1b))

  for i, v := range(text1b) {
    result[i] = v^text2b[i]
  }
  got := hex.EncodeToString(result)
  fmt.Println("result: ", got)
}
