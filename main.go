// ctrl-alt-grim filters a log stream on stdin: tail -f app.log | cag -level warn
package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"

	"github.com/grim-terminal/ctrl-alt-grim/internal/logfilter"
	"github.com/grim-terminal/ctrl-alt-grim/internal/notify"
)

func main() {
	level := flag.String("level", "info", "minimum level: debug|info|warn|error")
	grep := flag.String("grep", "", "only lines containing this text")
	hook := flag.String("notify", "", "command to run for every error line, e.g. notify-send")
	flag.Parse()

	min, ok := logfilter.ParseLevel(*level)
	if !ok {
		fmt.Fprintf(os.Stderr, "unknown level %q\n", *level)
		os.Exit(2)
	}

	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		line := sc.Text()
		if !logfilter.Keep(line, min, *grep) {
			continue
		}
		fmt.Println(line)
		if *hook != "" && logfilter.Keep(line, logfilter.Error, "") {
			if err := notify.Run(*hook, line); err != nil {
				fmt.Fprintln(os.Stderr, "notify:", err)
			}
		}
	}
	if err := sc.Err(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
