package cli

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"worldboxing/lib/utils"
)

func Start() {
	for {
		var input string
		print("> ")
		_, e := fmt.Scanln(&input)
		utils.Unwrap(e)
		if len(input) > 0 {
			print(input, "\n")
		}
	}
}

func readInp() ([]string, errs.Err) {
	inp := bufio.NewScanner(os.Stdin)
	fmt.Print("> ")
	ok := inp.Scan()
	if !ok {
		fmt.Print("\n")
		return nil, quit(nil)
	}

	text := inp.Text()
	lastInp = text
	return strings.Fields(text), nil
}
