package execute

import "github.com/urfave/cli"

func Cmd() cli.Command {
	return cli.Command{
		Name: "execute",
		ShortName: "e",
		Usage: "smarting execute [filePath]",
		Description: "execute smart lang bytecodes",
		Action: func(c *cli.Context) error {
			return compile(c.Args().Get(0))
		},
	}
}

func compile(filePath string) error {
	return nil
}