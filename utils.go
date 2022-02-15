package cli

import (
	"fmt"
	"os"
)

func fatal(v ...interface{}) {
	fmt.Println(v...)
	os.Exit(1)
}
