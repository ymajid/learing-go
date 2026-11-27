package repl

import (
	"bufio"
	"fmt"
	"io"

	"monkey/lexer"
	"monkey/parser"
)

const (
	b      = "`"
	MONKEY = `
                 __,__
        .--.  .-"     "-.  .--.
       / .. \/  .-. .-.  \/ .. \
      | |  '|  /   Y   \  |'  | |
      | \   \  \ 0 | 0 /  /   / |
       \ '- ,\.-"` + b + b + ` ` + b + b + `"-./, -' /
        ` + b + `'-' /_   ^ ^   _\ '-'` + b + `
        .--'|  \._   _./  |'--. 
      /` + b + `    \   \ ` + b + `~` + b + ` /   /    ` + b + `\
     /       '._ '---' _.'       \
    /           '~---~'   |       \
   /        _.             \       \
  /   .'-./` + b + `/        .'~'-.|\       \
 /   /    ` + b + `\:       /      ` + b + `\'.      \
/   |       ;      |         '.` + b + `;    /
\   \       ;      \           \/   /
 '.  \      ;       \       \   ` + b + `  /
   '._'.     \       '.      |   ;/_
     /__>     '.       \_ _ _/   ,  '--.
   .'   '.   .-~~~~~-. /     |--'` + b + `~~-.  \
  // / .---'/  .-~~-._/ / / /---..__.'  /
 ((_(_/    /  /      (_(_(_(---.__    .'
           | |     _              ` + b + `~~` + b + `
           | |     \'.
            \ '....' |
             '.,___.'
`
)
const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Printf(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)
		p := parser.New(l)

		program := p.ParseProgram()
		if len(p.Errors()) != 0 {
			printParserErrors(out, p.Errors())
			continue
		}

		io.WriteString(out, program.String())
		io.WriteString(out, "\n")
	}
}

func printParserErrors(out io.Writer, errors []string) {
	io.WriteString(out, MONKEY)
	io.WriteString(out, "Whoops! We ran into some monkey business here!\n")
	io.WriteString(out, " parser errors:\n")
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}
