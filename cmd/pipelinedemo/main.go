package main

import (
	"fmt"
	"imooc/pipeline"
	"os"
	"bufio"
)

func main() {
	const filename  = "small.in"
	const n  = 64

	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	write := bufio.NewWriter(file)
	pipeline.WriterSink(write,pipeline.RandomResource(n))
	write.Flush()

	/*file, err := os.Open("small.in")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	p := pipeline.ReaderSource(file,-1)
	for v := range p {
		fmt.Println(v)
	}*/
}

func MergeDemo(){
	p := pipeline.Merge(pipeline.InMemSort(pipeline.ArraySource(1,3,2,6,9,7,8)),pipeline.InMemSort(pipeline.ArraySource(1,3,2,6,9,7,8,12)))
	for val := range p {
		fmt.Println(val)
	}
}