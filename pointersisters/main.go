package main

import (
	f "fmt"
	"math/rand"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func main() {
   s := make([]*string, 20)
   for i:=0; i < 20; i++ {
     rand.Seed(int64(i))
     aa := RandStringBytesRmndr(11)
     s[i] = &aa
   }
   Print(s)
}

func Print(s []*string) {
	for _, v := range s {
		f.Println(*v)
	}
}

func RandStringBytesRmndr(n int) string {
    b := make([]byte, n)
    for i := range b {
        b[i] = letterBytes[rand.Int63() % int64(len(letterBytes))]
    }
    return string(b)
}
