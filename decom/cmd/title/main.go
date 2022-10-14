// Reformats text following English title capitalization rules. Takes input
// text through arguments or STDIN and prints titleized version.
//
// Implementation details and caveats can be found in library source code.
//
// Lowercase:
//
//	a, an, and, as, at, but, by, en, etc, for, from,
//	if, in, of, on, or, the, to, via, von, vs, with
//
// Example:
//
//	$ title "tears for fears @ rule the world: the greatest hits"
//	Tears for Fears @ Rule the World: The Greatest Hits
//
// Usage:
//
//	title <text>
//	title < <file>
//	cat <file> | title
//
// Author: Göran Gustafsson <gustafsson.g@gmail.com>
//
// License: BSD 3-Clause
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/ggustafsson/godis/pkg/title"
)

func main() {
	if len(os.Args) > 1 {
		// When executed as: title <text>
		text := strings.Join(os.Args[1:], " ")
		fmt.Println(title.Titleize(text))
	} else {
		// When executed as: title < <file>
		stdin := bufio.NewScanner(os.Stdin)
		for stdin.Scan() {
			fmt.Println(title.Titleize(stdin.Text()))
		}
		if err := stdin.Err(); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}
}
