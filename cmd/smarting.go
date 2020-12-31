package main

import (
	"github.com/fatih/color"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli"
	"os"
	"smart/cmd/compile"
	"smart/cmd/execute"
	"smart/cmd/lex"
	"smart/cmd/parse"
	"time"
)

const PROMPT = ">> "
const EXIT = "exit()"

const smartlang = `
 _______ _______ _______  ______ _______        _______ __   _  ______
 |______ |  |  | |_____| |_____/    |    |      |_____| | \  | |  ____
 ______| |  |  | |     | |    \_    |    |_____ |     | |  \_| |_____|
`

func prompt() {
	color.Cyan(smartlang)
	bold := color.New(color.Bold)
	_, _ = bold.Println("Use exit() or Ctrl-C to exit")
}

func repl() {
	prompt()
	ps1 := color.New(color.Bold)
	_, _ = ps1.Print(PROMPT)
}

func main()  {
	app := cli.NewApp()
	app.Name = "smarting"
	app.Usage = "smart lang compiler"
	app.Description = "implementation compiler for smartlang - an experimental smart language"
	app.Version = "0.1"
	app.Compiled = time.Now()

	app.Commands = []cli.Command{}
	app.Commands = append(app.Commands, lex.Cmd())
	app.Commands = append(app.Commands, parse.Cmd())
	app.Commands = append(app.Commands, compile.Cmd())
	app.Commands = append(app.Commands, execute.Cmd())

	app.Action = func(c *cli.Context) error {
		repl()
		return nil
	}
	if err := app.Run(os.Args); err != nil {
		log.Error(err)
	}
}
