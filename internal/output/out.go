package output

import (
	"fmt"
	"os"
	"strings"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/kmarkela/duffman/internal/pcollection"
	"github.com/kmarkela/duffman/pkg/jsonparser"
)

func Header(col *pcollection.Collection, wl, rt int, blacklist []int) {

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
		r += len(v.Parameters.Get) + len(v.Parameters.Post) + len(v.Parameters.Path)
	}

	line = fmt.Sprintf("# [*] Amount of parameters: %d", r)
	fmt.Printf("%s%s#\n", line, strings.Repeat(" ", length-len(line)))

	line = fmt.Sprintf("# [*] Total to fuzz: %d", r*wl)
	fmt.Printf("%s%s#\n", line, strings.Repeat(" ", length-len(line)))

	if len(blacklist) > 0 {
		line = fmt.Sprintf("# [*] Status Code Blacklist: %s", strings.Trim(strings.Replace(fmt.Sprint(blacklist), " ", ",", -1), "[]"))
		fmt.Printf("%s%s#\n", line, strings.Repeat(" ", length-len(line)))
	}

	if rt > 0 {
		line = fmt.Sprintf("# [*] Hide Response Time less than (ms): %d", rt)
		fmt.Printf("%s%s#\n", line, strings.Repeat(" ", length-len(line)))
	}

	fmt.Printf("%s#\n", strings.Repeat("#", length))
	fmt.Println()

}

// moveCursorUp moves the cursor up by n lines.
func moveCursorUp(n int) {
	fmt.Printf("\033[%dA", n)
}

// clearLine clears the entire line.
func clearLine() {
	// fmt.Print("\033[2K")
}

func RenderTable(rl []Results) {

	t := table.NewWriter()

	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"Enpoint", "Method", "Parameter", "FUZZ", "Code", "Length", "Time"})

	for _, v := range rl {
		t.AppendRows([]table.Row{{v.Endpoint, v.Method, jsonparser.Param2Str(v.Param), v.Word, v.Code, v.Length, v.Time}})
	}

	if len(rl) > 1 {
		moveCursorUp(len(rl) + 3)
		clearLine()
	}

	t.Render()

}

func RenderErrors(rl []Results) {

	fmt.Println()
	if len(rl) == 1 {
		fmt.Printf("[-] %d Error occurs during Fuzz:\n", 1)
	} else {
		fmt.Printf("[-] %d Errors occur during Fuzz:\n", len(rl))
	}

	for _, v := range rl {
		fmt.Printf("  - Endpoint %s: \n", v.Endpoint)
		fmt.Printf("    * Param: %s\n", v.Param)
		fmt.Printf("    * Error: %s\n", v.Err)
	}

}
