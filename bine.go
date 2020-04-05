package main

import (
	"bytes"
	"flag"
	"fmt"
	"os"
	"strconv"
)

var (
	offset int64
	length int64
	file   string
)

func init() {
	flag.Int64Var(&offset, "s", 0, "offset")
	flag.Int64Var(&length, "n", 0, "length")
	flag.StringVar(&file, "f", "", "file name")
	flag.Parse()
}

func main() {
	args := flag.Args()
	if file == "" || len(args) == 0 {
		usage()
		return
	}

	buf := new(bytes.Buffer)
	for _, arg := range args {
		v, err := strconv.ParseUint(arg, 0, 8)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		buf.WriteByte(byte(v))
	}

	if err := write(file, offset, length, buf.Bytes()...); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func usage() {
	fmt.Println("Usage: bine -f FILENAME [-s OFFSET] [-n LENGTH] VALUE [VALUE...]")
}

func write(filename string, offset, count int64, values ...byte) (err error) {
	f, err := os.OpenFile(file, os.O_WRONLY, 0400)
	if err != nil {
		return
	}
	defer f.Close()

	if count > 0 {
		if int64(len(values)) >= count {
			values = values[:count]
		} else {
			b := make([]byte, count-int64(len(values)))
			values = append(values, b...)
		}
	}

	_, err = f.WriteAt(values, offset)
	return
}
