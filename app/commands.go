package app

import (
	"context"
	"errors"
	"fmt"
	"log"
	"math/rand"
	"os"
	"regexp"
	"time"

	"github.com/urfave/cli/v3"
)

var ErrInvalidTagFormatting error = errors.New("tags can only be of the form: foo, foo-bar, foo-bar-baz")

func Srep(ctx context.Context) {

}

func StartCli(service *TopicService) {
	addCmd := &cli.Command{
		Name:  "add",
		Usage: "Adds a topic with the specified tag",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "tag",
				Aliases: []string{"t"},
				Usage:   "Sets the specific tag for the topic",
			},
		},
		Action: actionAdd(service),
	}

	removeCmd := &cli.Command{
		Name:    "remove",
		Aliases: []string{"rm"},
		Usage:   "Removes a topic from the database along with its associated data",
		Action:  actionRemove(service),
	}

	bucketCmd := &cli.Command{
		Name:    "bucket",
		Aliases: []string{"b"},
		Usage:   "Selects a topic from the bucket, reusing this command skips it and increments the topic's skipped count",
		Action:  actionBucket(service),
	}

	completeCmd := &cli.Command{
		Name:    "complete",
		Aliases: []string{"c"},
		Usage:   "Clears the current topic and significantly reduces its skipped count",
		Action:  actionComplete(service),
	}

	listCmd := &cli.Command{
		Name:    "list",
		Aliases: []string{"ls"},
		Usage:   "Lists all topics",
		Action:  actionList(service),
	}

	cmd := &cli.Command{
		Commands: []*cli.Command{
			addCmd,
			removeCmd,
			bucketCmd,
			completeCmd,
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
			fmt.Fprintf(os.Stdout, "Spaced repetition tool.\nUse \"srep help\" to show available commands.\n")

			t, err := service.GetCurrentTopic(ctx)

			if err != nil {
				fmt.Fprintf(os.Stdout, "No current topic set - Use \"srep bucket\" to roll a topic\n")
				return nil
			}

			fmt.Fprintf(os.Stdout, "<!> Current topic: \"%s\"\n", t.Id)
			return nil
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}

func actionBucket(service *TopicService) cli.ActionFunc {
	return func(ctx context.Context, cmd *cli.Command) error {
		if cmd.NArg() > 0 {
			fmt.Println("Bucket command take no arguments")
			return nil
		}

		topics, err := service.GetTopics(ctx)
		if err != nil {
			return nil
		}

		if len(*topics) == 0 {
			fmt.Println("No topics to select from")
			return nil
		}

		var selectedTopic Topic

		total := 0
		for _, topic := range *topics {
			total += topic.Skipped
		}

		if total == 0 {
			r := rand.Intn(len(*topics))
			selectedTopic = (*topics)[r]
		} else {
			r := rand.Intn(total)
			fmt.Println("r ", r)
			for _, topic := range *topics {
				if r < topic.Skipped {
					selectedTopic = topic
					break
				}
				r -= topic.Skipped
			}
		}

		oldTopic, err := service.store.GetCurrentTopic(ctx)
		if err == nil {
			oldTopic.Skipped++
			service.store.Update(ctx, oldTopic)
		}

		fmt.Fprintf(os.Stdout, "New topic: \"%s\"\n", selectedTopic.Id)

		service.store.SetCurrentTopic(ctx, selectedTopic.Id)
		return nil
	}
}

func actionComplete(service *TopicService) cli.ActionFunc {
	return func(ctx context.Context, cmd *cli.Command) error {
		if cmd.NArg() > 0 {
			fmt.Println("Complete command take no arguments")
			return nil
		}

		topic, err := service.GetCurrentTopic(ctx)
		if err == nil && topic.Id != "" {
			topic.Skipped = 1
			service.store.Update(ctx, topic)
		} else {
			fmt.Fprintf(os.Stderr, "No current topic to complete\n")
			return nil
		}
		service.SetCurrentTopic(ctx, "")
		fmt.Fprintf(os.Stdout, "Completed topic \"%s\"\n", topic.Id)
		return nil
	}
}

func actionRemove(service *TopicService) cli.ActionFunc {
	return func(ctx context.Context, cmd *cli.Command) error {
		args := cmd.Args().Slice()

		for i, arg := range args {
			err := service.RemoveTopic(ctx, arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "%d. skipped topic \"%s\" is not a valid topic\n", i+1, arg)
				continue
			}
		}
		fmt.Fprintf(os.Stdout, "ok\n")
		return nil
	}
}

func actionAdd(service *TopicService) cli.ActionFunc {
	return func(ctx context.Context, cmd *cli.Command) error {
		tag := cmd.String("tag")
		args := cmd.Args().Slice()

		re := regexp.MustCompile(`^[a-z]+(-[a-z]+)*$`)
		if tag != "" && !re.MatchString(tag) {
			return ErrInvalidTagFormatting
		}

		if len(args) == 0 {
			fmt.Println("Unespecified topics to add")
			return nil
		}

		for _, arg := range args {
			err := service.Add(ctx, &Topic{
				Id:         arg,
				Tag:        tag,
				Skipped:    1,
				Skippable:  true,
				LastRecall: time.Now(),
			})

			if err != nil {
				fmt.Fprintf(os.Stderr, "Topic \"%s\" cannot be added: %s\n", arg, err.Error())
				continue
			}
			fmt.Fprintf(os.Stdout, "Topic \"%s\" added\n", arg)
		}
		return nil
	}
}

func actionList(service *TopicService) cli.ActionFunc {
	return func(ctx context.Context, cmd *cli.Command) error {
		if cmd.NArg() > 0 {
			fmt.Println("List command take no arguments")
			return nil
		}
		data, err := service.GetTopics(ctx)

		if err != nil || data == nil {
			log.Fatal("Couldn't list topics, err: ", err)
		}

		if len(*data) == 0 {
			fmt.Fprintf(os.Stdout, "No topics in the database\n")
			return nil
		}

		for _, topic := range *data {
			fmt.Fprintf(os.Stdout, "- %s\n", topic.Id)
		}
		return nil
	}
}
