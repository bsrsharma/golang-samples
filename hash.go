package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"os"
	"log"
	"math/rand"
)

func mymd5(s string) []byte {
	h := md5.New()
	io.WriteString(h, s)
        return h.Sum(nil)
}

func makeNiceStr(AoB  []byte )  string  {
  var AoC = make([]byte,len(AoB))
  var myInt uint = 0
  var c byte = ' '
  var niceAoC = "ABCDEFGHJKLMNPQRSTUVWXYZ23456789"
   
   for i := 0; i < len(AoB); i++  {
      myInt = uint(AoB[i])
      myInt = myInt & 31
      c = niceAoC[myInt]
      AoC[i] = c
   }
   return(string(AoC))
}

func main() {

  fmt.Println("Hash for 'Hello'\n")

  var hashAoB = mymd5("Hello")

  var myStr string

  io.WriteString(os.Stdout,"")

  myStr = makeNiceStr(hashAoB)
  
  fmt.Println(myStr)
   
  // hash a bunch of names

  fmt.Println("Delaware Pennsylvania California Texas...")
  fmt.Println(makeNiceStr(mymd5("Delaware Pennsylvania California Texas...")))

  fmt.Println("Ecuador Japan Bhutan Paraguay")
  fmt.Println(makeNiceStr(mymd5("Ecuador Japan Bhutan Paraguay")))

  fmt.Println("Russia Romania")
  fmt.Println(makeNiceStr(mymd5("Russia Romania")))

  fmt.Println("Turkmenistan Denmark")
  fmt.Println(makeNiceStr(mymd5("Turkmenistan Denmark")))

  fmt.Println("Tajikistan Jamaica")
  fmt.Println(makeNiceStr(mymd5("Tajikistan Jamaica")))

  fmt.Println("Panama Djibouti Nepal")
  fmt.Println(makeNiceStr(mymd5("Panama Djibouti Nepal")))

  fmt.Println("Bolivia Srilanka Singapore")
  fmt.Println(makeNiceStr(mymd5("Bolivia Srilanka Singapore")))  


// Generate an array of random Bytes


  file, err := os.Create("arrs.dat")

  if err != nil {
	log.Fatal(err)
  }


  var rb [30]byte
  var rblen = 0  // length of the random string

  for i := 0; i < 10000; i++ {
     for j := 0; j < 2; j++  {
       for k := 0; k < 30; k++ {
         rb[k] = byte(rand.Intn(256))
       }
       rblen = rand.Intn(20)
       myStr = makeNiceStr(rb[0:30])
       file.WriteString(myStr[0:10+rblen])
       file.WriteString(" , ")
     }
     for k := 0; k < 30; k++ {
       rb[k] = byte(rand.Intn(256))
     }
     rblen = rand.Intn(20)
     myStr = makeNiceStr(rb[0:30])
     file.WriteString(myStr[0:10+rblen])
     file.WriteString("\n")     
  }

  file.Close()

  fmt.Println(""); fmt.Println("wrote arrs.dat"); fmt.Println("")
}  
