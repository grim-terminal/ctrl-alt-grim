// ctrl-alt-grim filters a log stream on stdin: tail -f app.log | cag -level warn
package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/grim-terminal/ctrl-alt-grim/internal/logfilter"
)

func main() {
	level := flag.String("level", "info", "minimum level: debug|info|warn|error")
	grep := flag.String("grep", "", "only lines containing this text")
	since := flag.Duration("since", 0, "only lines newer than this age, e.g. 10m (lines start with an RFC3339 timestamp)")
	flag.Parse()

	min, ok := logfilter.ParseLevel(*level)
	if !ok {
		fmt.Fprintf(os.Stderr, "unknown level %q\n", *level)
		os.Exit(2)
	}

	var cutoff time.Time
	if *since > 0 {
		cutoff = time.Now().Add(-*since)
	}

	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		line := sc.Text()
		if logfilter.Keep(line, min, *grep) && logfilter.Since(line, cutoff) {
			fmt.Println(line)
		}
	}
	if err := sc.Err(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
