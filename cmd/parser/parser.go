package parser

import "github.com/urfave/cli"

func Cmd() cli.Command {
	return cli.Command{
		Name: "parse",
		ShortName: "p",
		Usage: "smarting parse [filePath]",
		Description: "parse .smrt file to ast",
		Action: func(c *cli.Context) error {
			return parse(c.Args().Get(0))
		},
	}
}

func parse(filePath string) error {
	return nil
}