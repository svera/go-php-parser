package main

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"time"

	"github.com/z7zmey/php-parser/pkg/conf"
	"github.com/z7zmey/php-parser/pkg/errors"
	"github.com/z7zmey/php-parser/pkg/parser"
	"github.com/z7zmey/php-parser/pkg/version"
	"github.com/z7zmey/php-parser/pkg/visitor/dumper"
)

func main() {
	var m runtime.MemStats
	start := time.Now()
	src := loadSource()

	// Error handler

	var parserErrors []*errors.Error
	errorHandler := func(e *errors.Error) {
		parserErrors = append(parserErrors, e)
	}

	// Parse

	rootNode, err := parser.Parse(src, conf.Config{
		Version:          &version.Version{Major: 7, Minor: 0},
		ErrorHandlerFunc: errorHandler,
	})

	if err != nil {
		log.Fatal("Error:" + err.Error())
	}

	// Dump

	goDumper := dumper.NewDumper(os.Stdout)

	rootNode.Accept(goDumper)
	end := time.Now()
	runtime.ReadMemStats(&m)
	fmt.Printf("\nExecution time: %s\n", end.Sub(start).String())
	fmt.Printf("Memory usage: %v Mb\n", m.Alloc/1024/1024)
}

func loadSource() []byte {
	dat, err := os.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	return dat
}
