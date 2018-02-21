package main

import (
	"fmt"
	"imooc/pipeline"
	"os"
	"bufio"
)

func main()  {
	file, err := os.Create("small.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	out := pipeline.RandomResource(10000000)
	pipeline.WriterSink(bufio.NewWriter(file),out)

	fh, err := os.Open("small.txt")
	if err != nil {
		panic(err)
	}
	put := pipeline.ReaderSource(bufio.NewReader(fh))
	for v := range put {
		fmt.Println(v)
	}
}

func MergeDemo(){
	p := pipeline.Merge(pipeline.InMemSort(pipeline.ArraySource(1,3,2,4,6,7,5,8)),pipeline.InMemSort(pipeline.ArraySource(1,3,2,4,6,7,5,8)))
	for v := range p {
		fmt.Println(v)
	}
}