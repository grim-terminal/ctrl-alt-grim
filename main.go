// ctrl-alt-grim filters a log stream on stdin: tail -f app.log | cag -level warn
package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"

	"github.com/grim-terminal/ctrl-alt-grim/internal/logfilter"
)

func main() {
	level := flag.String("level", "info", "minimum level: debug|info|warn|error")
	grep := flag.String("grep", "", "only lines containing this text")
	invert := flag.Bool("v", false, "with -grep, keep only lines that do not contain the text")
	flag.Parse()

	min, ok := logfilter.ParseLevel(*level)
	if !ok {
		fmt.Fprintf(os.Stderr, "unknown level %q\n", *level)
		os.Exit(2)
	}

	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		line := sc.Text()
		if logfilter.Keep(line, min, *grep, *invert) {
			fmt.Println(line)
		}
	}
	if err := sc.Err(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
