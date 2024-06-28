package output

import (
	"fmt"
	"os"
	"strings"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/kmarkela/duffman/internal/pcollection"
)

func Header(col *pcollection.Collection, wl int, blacklist []int) int {

	var length int = 55

	fmt.Printf("%s#\n", strings.Repeat("#", length))

	f := (length - 7) / 2
	line := fmt.Sprintf("%sDuffMan%s", strings.Repeat(" ", f), strings.Repeat(" ", length-f-8))
	fmt.Printf("#%s#\n", line)

	line = fmt.Sprintf("# [*] Wordlist count: %d", wl)
	fmt.Printf("%s%s#\n", line, strings.Repeat(" ", length-len(line)))

	line = fmt.Sprintf("# [*] Amount of request: %d", len(col.Requests))
	fmt.Printf("%s%s#\n", line, strings.Repeat(" ", length-len(line)))

	var r int
	for _, v := range col.Requests {
		r += len(v.Parameters.Get) + len(v.Parameters.Post)
	}

	line = fmt.Sprintf("# [*] Amount of parameters: %d", r)
	fmt.Printf("%s%s#\n", line, strings.Repeat(" ", length-len(line)))

	line = fmt.Sprintf("# [*] Total to fuzz: %d", r*wl)
	fmt.Printf("%s%s#\n", line, strings.Repeat(" ", length-len(line)))

	if len(blacklist) > 0 {
		line = fmt.Sprintf("# [*] Status Code Blacklist: %s", strings.Trim(strings.Replace(fmt.Sprint(blacklist), " ", ",", -1), "[]"))
		fmt.Printf("%s%s#\n", line, strings.Repeat(" ", length-len(line)))
	}

	fmt.Printf("%s#\n", strings.Repeat("#", length))
	fmt.Println()
	return r * wl

}

// moveCursorUp moves the cursor up by n lines.
func moveCursorUp(n int) {
	fmt.Printf("\033[%dA", n)
}

// clearLine clears the entire line.
func clearLine() {
	fmt.Print("\033[2K")
}

func RenderTable(rl []Results, i int) {

	t := table.NewWriter()

	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"Enpoint", "Method", "Parameter", "FUZZ", "Code", "Length", "Time"})

	for _, v := range rl {
		t.AppendRows([]table.Row{{v.Endpoint, v.Method, v.Param, v.Word, v.Code, v.Length, v.Time}})
	}

	if len(rl) > 1 {
		moveCursorUp(len(rl) + 3)
		clearLine()
	}

	t.Render()

}
