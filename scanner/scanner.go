package scanner

import (
	"fmt"
	"os"
	"runtime/debug"

	"github.com/jessevdk/go-flags"
)

type Options struct {
	Version bool `short:"V" long:"version" required:"false" description:"Display the current version of binary"`
}

var opts Options

var Version = ""

func LoadResources() {
	opts := InitializeOptions()
	DisplayVersion(Version)
	fmt.Println("Options: ", opts)
}

func InitializeOptions() *Options {
	_, err := flags.Parse(&opts)
	if err != nil {
		os.Exit(1)
	}
	return &opts
}

func DisplayVersion(version string) {
	if !IsVersion() {
		return
	}
	if version != "" {
		fmt.Printf("%s\n", version)
	} else {
		buildInfo, ok := debug.ReadBuildInfo()
		if !ok {
			fmt.Println("Version information not available.")
		}
		if buildInfo.Main.Version != "" {
			fmt.Printf("%s\n", buildInfo.Main.Version)
		} else {
			fmt.Println("unknown")
		}
	}
	os.Exit(0)
}

func IsVersion() bool {
	return opts.Version
}
