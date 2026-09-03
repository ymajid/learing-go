package main

import (
	"fmt"
	"strings"
)

func main() {
	days := []string{
		"first", "second", "third", "fourth",
		"fifth", "sixth", "seventh", "eighth",
		"ninth", "tenth", "eleventh", "twelfth",
	}
	gifts := []string{
		"partridge in a pear tree",
		"Two turtle doves",
		"Three french hens",
		"Four calling birds",
		"Five golden rings",
		"Six geese a-laying",
		"Seven swans a-swimming",
		"Eight maids a-milking",
		"Nine ladies dancing",
		"Ten lords a-leaping",
		"Eleven pipers piping",
		"Twelve drummers drumming",
	}

	/* Initial attempt
	for i, day := range days {
		line := make([]string, i+1)
		copy(line, gifts[:i+1])
		if i == 0 {
			line[0] = "A " + line[0] + "."
		} else {
			line[0] = "And a" + line[0]
		}
		slices.Reverse(line)
		jg := strings.Join(line, "\n")
		fmt.Printf("On the %v day of Christmas\nMy true love gave to me:\n%v\n\n", day, jg)
	}
	*/

	/* Second attempt
	for i, day := range days {
		var lines []string // reversed list
		for k := i; k > 0; k-- {
			lines = append(lines, gifts[k])
		}

		if i == 0 {
			lines = append(lines, "A "+gifts[0]+".")
		} else {
			lines = append(lines, "And a "+gifts[0]+".")
		}

		fmt.Printf("On the %v day of Christmas\nMy true love gave to me:\n%v\n\n", day, strings.Join(lines, "\n"))
	}
	*/

	var rev []string
	for i, day := range days {
		if i >= 1 {
			rev = append([]string{gifts[i]}, rev...)
		}

		var lines []string
		lines = append(lines, rev...)
		if i == 0 {
			lines = append(lines, "A "+gifts[0]+".")
		} else {
			lines = append(lines, "And a "+gifts[0]+".")
		}
		fmt.Printf("On the %v day of Christmas\nMy true love gave to me:\n%v\n\n", day, strings.Join(lines, "\n"))
	}
}
