package main

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"
)

func TestAll(t *testing.T) {
	// path, err := os.Getwd()
	// if err != nil {
	// 	log.Panicln(err)
	// }
	//	testdata := filepath.Join(path, "testdata")
	analysistest.Run(t, analysistest.TestData(), Analyzer)
}
