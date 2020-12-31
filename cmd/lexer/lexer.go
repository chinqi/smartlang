package lexer

import "github.com/urfave/cli"

func Cmd() cli.Command {
	return cli.Command{
		Name: "lexer",
		ShortName: "l",
		Usage: "smarting lexer [filePath]",
		Description: "lexer source file",
		Action: func(c *cli.Context) error {
			return lex(c.Args().Get(0))
		},
	}
}

func lex(filePath string) error {
	return nil
}