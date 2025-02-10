package cli

import (
	"fmt"
	"runtime"

	"github.com/Kshitiz-Mhto/mock-my-commit/utility"
	"github.com/spf13/cobra"
)

const logo = `
███╗   ███╗ ██████╗  ██████╗██╗  ██╗     ███╗   ███╗██╗   ██╗      ██████╗ ██████╗ ███╗   ███╗███╗   ███╗██╗████████╗
████╗ ████║██╔═══██╗██╔════╝██║ ██╔╝     ████╗ ████║╚██╗ ██╔╝     ██╔════╝██╔═══██╗████╗ ████║████╗ ████║██║╚══██╔══╝
██╔████╔██║██║   ██║██║     █████╔╝█████╗██╔████╔██║ ╚████╔╝█████╗██║     ██║   ██║██╔████╔██║██╔████╔██║██║   ██║   
██║╚██╔╝██║██║   ██║██║     ██╔═██╗╚════╝██║╚██╔╝██║  ╚██╔╝ ╚════╝██║     ██║   ██║██║╚██╔╝██║██║╚██╔╝██║██║   ██║   
██║ ╚═╝ ██║╚██████╔╝╚██████╗██║  ██╗     ██║ ╚═╝ ██║   ██║        ╚██████╗╚██████╔╝██║ ╚═╝ ██║██║ ╚═╝ ██║██║   ██║   
╚═╝     ╚═╝ ╚═════╝  ╚═════╝╚═╝  ╚═╝     ╚═╝     ╚═╝   ╚═╝         ╚═════╝ ╚═════╝ ╚═╝     ╚═╝╚═╝     ╚═╝╚═╝   ╚═╝   
                                                                                                                     
`
const (
	CLI_NAME    = "mock-my-comit"
	CLI_VERSION = "1.0.0-stable"
)

var (
	quiet      bool
	verbose    bool
	versionCMD = &cobra.Command{
		Use:   "version",
		Short: "Version will output the current build information",
		Run: func(cmd *cobra.Command, args []string) {
			buildDate := utility.GetBuildDate()
			switch {
			case verbose:
				fmt.Print(logo)
				fmt.Printf("CLI version: v%s\n", CLI_VERSION)
				fmt.Printf("Go version (client): %s\n", runtime.Version())
				if buildDate != "" {
					fmt.Printf("Build date (client): %s\n", buildDate)
				}
				fmt.Printf("OS/Arch (client): %s/%s\n", runtime.GOOS, runtime.GOARCH)
			case quiet:
				fmt.Printf("v%s\n", CLI_VERSION)
			default:
				fmt.Printf("mock-my-commit CLI v%s\n", CLI_VERSION)
			}
		},
	}
)

func init() {
	versionCMD.Flags().BoolVarP(&quiet, "quiet", "q", false, "Use quiet output for simple output")
	versionCMD.Flags().BoolVarP(&verbose, "verbose", "v", false, "Use verbose output to see full information")
}
