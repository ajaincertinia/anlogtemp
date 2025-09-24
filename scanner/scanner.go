package scanner
import (
	"fmt"
	"github.com/jessevdk/go-flags"
	"os"
)

type Options struct {
	Version      bool   `short:"V" long:"version" required:"false" description:"Display the current version of binary"`
}

var opts Options

var Version = "1.0.4"

func LoadResources() {
	opts := InitializeOptions()
	DisplayVersion(Version)
	fmt.Println( "Options: ", opts)
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
	fmt.Printf("%s\n", version)
	os.Exit(0)
}

func IsVersion() bool {
	return opts.Version
}