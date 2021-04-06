package main

import (
	"encoding/binary"
	"fmt"
	"io"
	"os"
	"path"
	"strconv"
	"strings"
)

func main() {

	args := os.Args[1:]

	if len(args) != 1 {
		panic("Invalid args")
	}

	fiPath := args[0]

	fmt.Println(fiPath)

	fi, err := os.Open(fiPath)
	if err != nil {
		panic(err)
	}
	defer fi.Close()

	fo, err := os.Create(path.Join(path.Dir(fiPath), "out.csv"))
	if err != nil {
		panic(err)
	}
	defer fo.Close()

	valuesPerColumn := 2

	h := 0.0
	var b int8
	fileEnded := false

	for !fileEnded {
		cols := []string{fmt.Sprintf("%.2f", h)}
		for i := 1; i <= valuesPerColumn; i++ {
			err = binary.Read(fi, binary.LittleEndian, &b)
			if err != nil {
				if err == io.EOF {
					fileEnded = true
					break
				} else {
					panic(err)
				}
			}
			cols = append(cols, strconv.Itoa(int(b)))
		}
		if fileEnded {
			break
		}
		fmt.Println(strings.Join(cols, ","))
		_, err := fo.WriteString(strings.Join(cols, ",") + "\n")
		if err != nil {
			panic(err)
		}
		h += 0.2
	}
}
