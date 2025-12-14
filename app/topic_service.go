package app

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"regexp"
	"time"
)

type TopicService struct {
	store TopicStore
}

var ErrInvalidIdFormatting error = errors.New("invalid ID, only lowercase letters separated by \"-\"")

func NewTopicService(store TopicStore) *TopicService {
	return &TopicService{store: store}
}

func (s *TopicService) Add(ctx context.Context, t *Topic) error {
	if !isValidId(t.Id) {
		return ErrInvalidIdFormatting
	}
	err := s.store.Create(ctx, t)
	fmt.Println(err)
	if err != nil {
		return err
	}
	return nil
}

func (s *TopicService) RemoveTopic(ctx context.Context, id string) error {
	if !isValidId(id) {
		return ErrInvalidIdFormatting
	}

	s.store.Delete(ctx, id)
	return nil
}

func (s *TopicService) GetCurrentTopic(ctx context.Context) (*Topic, error) {
	return s.store.GetCurrentTopic(ctx)
}

func (s *TopicService) GetTopics(ctx context.Context) (*[]Topic, error) {
	return s.store.List(ctx)
}

func (s *TopicService) ListTopicsVerbose(w io.Writer, ctx context.Context) {
	data, err := s.store.List(ctx)
	if err != nil {
		log.Fatal("Couldn't list topics, ", err)
	}

	if data == nil {
		fmt.Print("No topics in the database")
		return
	}

	for _, topic := range *data {
		fmt.Fprintf(w, "- %s %s\n", topic.Id, topic.LastRecall.Format(time.DateOnly))
	}

}

func (s *TopicService) SetCurrentTopic(ctx context.Context, id string) {
	err := s.store.SetCurrentTopic(ctx, id)
	if err != nil {
		log.Println(err)
	}
}

func isValidId(id string) bool {
	re := regexp.MustCompile(`^[a-z]+(-[a-z]+)*$`)
	return re.MatchString(id)
}
