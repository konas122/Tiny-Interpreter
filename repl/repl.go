package repl

import (
	"Interpreter/compiler"
	"Interpreter/lexer"
	"Interpreter/parser"
	"Interpreter/vm"
	"bufio"
	"fmt"
	"io"
)

const PROMPT = ">> "
const MONKEY_FACE = `
            __,__
   .--.  .-"     "-.  .--.
  / .. \/  .-. .-.  \/ .. \
 | |  '|  /   Y   \  |'  | |
 | \   \  \ 0 | 0 /  /   / |
  \ '- ,\.-"""""""-./, -' /
   ''-' /_   ^ ^   _\ '-''
       |  \._   _./  |
       \   \ '~' /   /
        '._ '-=-' _.'
           '-----'

`

func Start(in io.Reader, out io.Writer) {
	io.WriteString(out, MONKEY_FACE)
	scanner := bufio.NewScanner(in)

	// env := object.NewEnvironment()
	// macroEnv := object.NewEnvironment()

	for {
		fmt.Fprint(out, PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)
		p := parser.New(l)

		program := p.ParseProgram()
		if len(p.Errors()) != 0 {
			printParseErrors(out, p.Errors())
			continue
		}

		// evaluator.DefineMacros(program, macroEnv)
		// expanded := evaluator.ExpandMacros(program, macroEnv)

		// evaluated := evaluator.Eval(expanded, env)
		// if evaluated != nil {
		// 	io.WriteString(out, evaluated.Inspect())
		// 	io.WriteString(out, "\n")
		// }

		comp := compiler.New()
		if err := comp.Compile(program); err != nil {
			fmt.Fprintf(out, "Woops! Compilation failed:\n\t%s\n", err)
			continue
		}

		machine := vm.New(comp.Bytecode())
		if err := machine.Run(); err != nil {
			fmt.Fprintf(out, "Woops! Executing bytecode failed:\n\t%s\n", err)
			continue
		}

		lastPopped := machine.LastPoppedStackElem()
		io.WriteString(out, lastPopped.Inspect())
		io.WriteString(out, "\n")
	}
}

func printParseErrors(out io.Writer, errors []string) {
	io.WriteString(out, "\n========================================================\n")
	io.WriteString(out, "Woops! We ran into some monkey business here!\n")
	io.WriteString(out, " parser errors:\n")
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
	io.WriteString(out, "========================================================\n\n")
}
