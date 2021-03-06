package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"os"
	"runtime"
)

var rootCmd = &cobra.Command{
	Use:   "gvm",
	Short: "GVM is a go version manager",
	Long: `GVM is a go version manager. You can use commands of install, list, switch 
and remove to manage local installation of multiple Go versions.

GVM is just support for Unix-like system yet, and the working mechanism of it 
is very simple. GVM will create a gvm directory in /usr/local to host multiple 
versions of GOROOT, and create a symbol link named go in /usr/local referring 
to the specific version in gvm directory. So, you just need to add /usr/local/go/bin 
to PATH environment variable to run go command, and use gvm to switch the 
reference of the symbol link.

Multiple versions of Go installed by GVM can share the same GOPATH compatibly, 
and this is also advocated by GVM.`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		if runtime.GOOS == "windows" {
			log.Fatal("Sorry, GVM does not support Windows platform at the moment.")
		}
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func isRootUser(cmd *cobra.Command, args []string) {
	if os.Getuid() != 0 {
		log.Fatal("Permission denied, please execute this command as the root user.")
	}
}
