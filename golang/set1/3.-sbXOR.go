package main
import(
  "fmt"
  "encoding/hex"
)
//1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736
var freq = map[rune]float64{
  'a': 0.0651738,
  'b': 0.0124248,
  'c': 0.0217339,
  'd': 0.0349835,
  'e': 0.1041442,
  'f': 0.0197881,
  'g': 0.0158610,
  'h': 0.0492888,
  'i': 0.0558094,
  'j': 0.0009033,
  'k': 0.0050529,
  'l': 0.0331490,
  'm': 0.0202124,
  'n': 0.0564513,
  'o': 0.0596302,
  'p': 0.0137645,
  'q': 0.0008606,
  'r': 0.0497563,
  's': 0.0515760,
  't': 0.0729357,
  'u': 0.0225134,
  'v': 0.0082903,
  'w': 0.0171272,
  'x': 0.0013692,
  'y': 0.0145984,
  'z': 0.0007836,
  ' ': 0.1918182,
}
func scoreText(text string)float64 {
  var res float64 = 0
  for _,val:=range(text){
    res+=freq[val]
  }
  return res
}
func Decode(text1b []byte)(string, float64, byte) {
  result := make([]byte, len(text1b))
  var score float64 = 0
  var decoded string
  var chr byte
  for i:=33; i<=122; i++ {
    for j,v := range(text1b){
      result[j] = v^byte(i)
    }
    score2 := scoreText(string(result))
    if score2 > score {
      score = score2
      decoded = string(result)
      chr = byte(i)
    }
  }
  return decoded, score, chr
}

func caca() {
  var textEnc string
  fmt.Scanln(&textEnc)
  text1b, _:= hex.DecodeString(textEnc)
  result, _, chr := Decode(text1b)
  fmt.Println(result+"\n"+string(chr))
}
