package main
import (
  "fmt"
  hex "encoding/hex"
  b64 "encoding/base64"
)
func main() {
  fmt.Print("Enter text: ")
  var input string
  fmt.Scanln(&input)
  textb, e:= hex.DecodeString(input)
  if e!=nil {
    panic("Error in decoding the hex value: ", e)
  }

  sEnc := b64.StdEncoding.EncodeToString(textb)
  fmt.Println(sEnc)
}
