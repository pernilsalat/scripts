package main

import (
	"bufio"
	b64 "encoding/base64"
	"encoding/hex"
	"fmt"
	"os"
	"sort"
)

func readfile(name string, p chan<- []byte) {
	f, _ := os.Open(name)
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		text1b, _ := hex.DecodeString(scanner.Text())
		p <- text1b
	}
	p <- []byte("caca725")
}
func stringToBin(s string) (binString string) {
	for _, c := range s {
		binString = fmt.Sprintf("%s%08b", binString, byte(c))
	}
	return
}
func edit(t1, t2 string) int {
	res := 0
	bin1 := stringToBin(t1)
	bin2 := stringToBin(t2)
	for i, v := range bin1 {
		if v != rune(bin2[i]) {
			res++
		}
	}
	return res
}
func sizeKey(line string) int {
	var size []float64
	for i := 3; i < len(line)/2; i++ {
		//fmt.Println(line[0]==line[i], i)
		size = append(size, float64(edit(line[:i], line[i:2*i]))/float64(i))
	}
	sort.Float64s(size)
	//fmt.Println(size, len(size))
	s := (size[0] + size[1] + size[2]) / 3
	return int(s + 0.5)
}
func findKey(slice [][]byte, size int) string {
	key := ""
	for i := 0; i < size; i++ {
		_, _, chr := Decode(slice[i])
		key += string(chr)
	}
	return key
}
func decodeText(f *os.File, key string) {
	_, _ = f.Seek(0, 0)
	scanner3 := bufio.NewScanner(f)
	for scanner3.Scan() {
		line3, _ := b64.StdEncoding.DecodeString(scanner3.Text())
		res := make([]byte, len(line3))
		for i, v := range line3 {
			res[i] = byte(v) ^ key[i%len(key)]
		}

		fmt.Println(string(res))
	}
}
func main() {
	//p := make(chan []byte)
	//go readfile("6.txt", p)
	f, _ := os.Open("6.txt")
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Scan()
	line, _ := b64.StdEncoding.DecodeString(scanner.Text()) //////////////////
	scanner.Scan()
	linesec, _ := b64.StdEncoding.DecodeString(scanner.Text()) //////////////////
	lineTotal := string(line) + string(linesec)
	size := sizeKey(lineTotal)
	fmt.Println(size,  edit("this is a test", "wokka wokka!!!"))

	_, _ = f.Seek(0, 0)
	scanner2 := bufio.NewScanner(f)
	//i:=0
	slice := make([][]byte, size)
	for scanner2.Scan() {
		line2, _ := b64.StdEncoding.DecodeString(scanner2.Text())
		for i, v := range line2 { ///////
			slice[i%size] = append(slice[i%size], byte(v))
		}
	}
	key := findKey(slice, size)
	decodeText(f, key) ////////////////
}
