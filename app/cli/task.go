package cli

import "fmt"

func CliRunTask(args []string) {
	fmt.Println(args, len(args))
}
