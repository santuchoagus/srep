package app

import "context"
import "errors"

var ErrTopicNotFound = errors.New("topic not found")

type TopicStore interface {
	Create(ctx context.Context, t *Topic) error
	Update(ctx context.Context, t *Topic) error
	Delete(ctx context.Context, id string) error
	ByID(ctx context.Context, id string) (*Topic, error)
	List(ctx context.Context) (*[]Topic, error)
}
