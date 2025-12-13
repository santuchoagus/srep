package app

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"
)


type TopicService struct {
	store TopicStore
}

func NewTopicService(store TopicStore) *TopicService {
	return &TopicService{store: store}
}

func (s *TopicService) ListTopics(w io.Writer, ctx context.Context) {
	data, err := s.store.List(ctx)
	if err != nil {
		log.Fatal("Couldn't list topics, err: ", err)
	}

	if data == nil {
		log.Fatal("Nill data for some reason");
	}

	for _, topic := range *data {
		fmt.Fprintf(w, "- %s\n", topic.Id)
	}
}

func (s *TopicService) ListTopicsVerbose(w io.Writer, ctx context.Context) {
	data, err := s.store.List(ctx)
	if err != nil {
		log.Fatal("Couldn't list topics, err: ", err)
	}

	if data == nil {
		log.Fatal("Nill data for some reason");
	}

	for _, topic := range *data {
		fmt.Fprintf(w, "- %s %s\n", topic.Id, topic.LastRecall.Format(time.DateOnly))
	}
}