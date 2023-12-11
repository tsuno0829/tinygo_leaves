package main

import (
	"bufio"
	"bytes"
	"embed"
	"syscall/js"
	"github.com/dmitryikh/leaves"
)

//go:embed model.txt
var fs embed.FS

var model *leaves.Ensemble

func modelInit() (error) {
	useTransformation := true
	data, err := fs.ReadFile("model.txt")
	if err != nil {
		return err
	}

	reader := bytes.NewReader(data)
	bufferedReader := bufio.NewReader(reader)
	model, err = leaves.LGEnsembleFromReader(bufferedReader, useTransformation)
	if err != nil {
		return err
	}

	return nil
}

func inference(this js.Value, args []js.Value) interface{} {
	// JSから渡される配列を取得
	inputJSArray := args[0]
	rows := args[1].Int()
	cols := args[2].Int() 

	// 配列の長さを取得
	length := inputJSArray.Length()

	// Goのfloat64スライスに変換
	input := make([]float64, length)
	for i := 0; i < length; i++ {
		input[i] = float64(inputJSArray.Index(i).Float())
	}

	// 推論処理
	predictions, err := inference_impl(input, rows, cols)
	if err != nil {
		panic(err)
	}

	// 予測結果をJavaScriptの配列に変換
	result := make([]interface{}, rows)
	for i, v := range predictions {
		result[i] = v
	}

	return js.ValueOf(result)
}

func inference_impl(v []float64, rows int, cols int) ([]float64, error) {
	predictions := make([]float64, rows)
	err := model.PredictDense(v, rows, cols, predictions, 0, 1)
	if err != nil {
		return predictions, err
	}
	return predictions, nil
}

func main() {
	err := modelInit()
	if err != nil {
		panic(err)
	}

	wait := make(chan struct{}, 0)
	js.Global().Set("inference", js.FuncOf(inference))
	<-wait
}
