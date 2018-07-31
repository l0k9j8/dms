// +build ignore

package main

import (
	"bufio"
	"flag"
	"io"
	"log"
	"os"
	"time"

	"../../misc"
)

func main() {
	ss := flag.String("ss", "", "")
	t := flag.String("t", "", "")
	flag.Parse()
	if flag.NArg() != 1 {
		log.Fatalln("wrong argument count")
	}
	go func() {
		buf := bufio.NewWriterSize(os.Stdout, 1234)
		n, err := io.Copy(buf, r)
		log.Println("copied", n, "bytes")
		if err != nil {
			log.Println(err)
		}
	}()
	time.Sleep(time.Second)
	go r.Close()
	time.Sleep(time.Second)
}
