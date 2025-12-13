package app

import (
	"context"
	"github.com/urfave/cli/v3"
    "fmt"
    "os"
    "log"
)

func Srep(ctx context.Context) {

}

func StartCli() {
    addCmd := &cli.Command{
        Name: "add",
        Usage: "Adds the topic with the tag if specificied",
        Flags: []cli.Flag{
            &cli.StringFlag{
                Name: "tag",
                Aliases: []string{"t"},
                Usage: "Set the specific tag for the topic",
            },
        },
        Action: actionAdd,
    }

    bucketCmd := &cli.Command{
        Name: "bucket",
        Aliases: []string{"b"},
        Usage: "Takes a topic from the bucket, reusing this command skip it and increment the topic's skipped count",
        Action: actionBucket,
    }

    listCmd := &cli.Command{
        Name: "list",
        Aliases: []string{"ls"},
        Usage: "Lists every topic",
        Action: actionList,
    }

	cmd := &cli.Command{
        Commands: []*cli.Command{
            addCmd,
            bucketCmd,
            listCmd,
        },
		Name:  "srep",
		Usage: "Spaced repetition tool",
        Arguments: []cli.Argument{
            &cli.StringArg{
                Name: "topic",
            },
        },

		Action: func(ctx context.Context, cmd *cli.Command) error {
            fmt.Println("Spaced repetition tool: \"srep help\" to show help\n<!> Current topic: \"Topic\"")
            return nil
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}

func actionBucket(ctx context.Context, cmd *cli.Command) error {
    if !cmd.Bool("bucket") {
        return nil
    }
    if cmd.NArg() > 0 {
        fmt.Println("Bucket command take no arguments")
        return nil
    }
    return nil
}

func actionAdd(ctx context.Context, cmd *cli.Command) error {
    tag := cmd.String("tag")
    args := cmd.Args().Slice()

    if len(args) == 0 {
        fmt.Println("Unespecified topics to add")
        return nil
    }
    
    fmt.Println(tag, args)
    return nil
}

func actionList(ctx context.Context, cmd *cli.Command) error {
    if cmd.NArg() > 0 {
        fmt.Println("List command take no arguments")
        return nil
    }

    return nil
}