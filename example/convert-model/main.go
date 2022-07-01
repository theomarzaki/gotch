package main

import (
	"fmt"
	"log"

	"github.com/theomarzaki/gotch/ts"
)

func main() {
	// NOTE. Python script to save model to .npz can be found at https://github.com/theomarzaki/pytorch-pretrained/bert/bert-base-uncased-to-npz.py
	filepath := "../../data/convert-model/bert/model.npz"

	namedTensors, err := ts.ReadNpz(filepath)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Num of named tensor: %v\n", len(namedTensors))
	outputFile := "/home/theomarzaki/projects/transformer/data/bert/model.gt"
	err = ts.SaveMultiNew(namedTensors, outputFile)
	if err != nil {
		log.Fatal(err)
	}
}
