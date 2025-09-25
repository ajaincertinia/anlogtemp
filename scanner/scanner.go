package scanner
import (
	"fmt"
	"github.com/jessevdk/go-flags"
	"os"
	 "runtime/debug"
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
	buildInfo, ok := debug.ReadBuildInfo()
	if !ok {
		fmt.Println("Unable to determine version information.")
		return
	}
	if buildInfo.Main.Version != "" {
		fmt.Printf("Version: %s\n", buildInfo.Main.Version)
	} else {
		fmt.Println("Version: unknown")
	}
	os.Exit(0)
}

func IsVersion() bool {
	return opts.Version
}