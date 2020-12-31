package compile

import "github.com/urfave/cli"

func Cmd() cli.Command {
	return cli.Command{
		Name: "compile",
		ShortName: "c",
		Usage: "smarting compile [filePath]",
		Description: "compile .smrt source file",
		Action: func(c *cli.Context) error {
			return compile(c.Args().Get(0))
		},
	}
}

func compile(filePath string) error {
	return nil
}