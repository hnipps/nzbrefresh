package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/Tensai75/nzbrefresh/internal/arguments"
	"github.com/Tensai75/nzbrefresh/pkg/refresh"
)

type (
	Provider struct {
		Name                  string
		Host                  string
		Port                  uint32
		SSL                   bool
		SkipSslCheck          bool
		Username              string
		Password              string
		MaxConns              uint32
		ConnWaitTime          time.Duration
		IdleTimeout           time.Duration
		HealthCheck           bool
		MaxTooManyConnsErrors uint32
		MaxConnErrors         uint32
	}
)

var (
	args *arguments.Args
)

func init() {
	arguments.ParseArguments()
	args = arguments.Arguments
	fmt.Println(args.Version())

	if args.Debug {
		logFileName := strings.TrimSuffix(filepath.Base(args.NZBFile), filepath.Ext(filepath.Base(args.NZBFile))) + ".log"
		f, err := os.Create(logFileName)
		if err != nil {
			exit(fmt.Errorf("unable to open debug log file: %v", err))
		}
		log.SetOutput(f)
	} else {
		log.SetOutput(io.Discard)
	}

	refresh.Prepare(
		refresh.WithNZBFile(args.NZBFile),
		refresh.WithCheckOnly(args.CheckOnly),
		refresh.WithProvider(args.Provider),
		refresh.WithDebug(args.Debug),
		refresh.WithCsv(args.Csv),
	)
}

func main() {
	refresh.Run()
}

func exit(err error) {
	if err != nil {
		fmt.Printf("Fatal error: %v\n", err)
		log.Fatal(err)
	} else {
		os.Exit(0)
	}
}
