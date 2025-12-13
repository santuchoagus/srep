package sqlite

import (
	"context"
	"database/sql"
	"time"
	"tui-app/app"
)

type SQLiteTopicStorage struct {
	db *sql.DB
}

func NewSQLiteTopicStorage(db *sql.DB) (*SQLiteTopicStorage) {
	return &SQLiteTopicStorage{db: db}
}


func (s *SQLiteTopicStorage) Create(ctx context.Context, t *app.Topic) error {
	return nil
}
func (s *SQLiteTopicStorage) Update(ctx context.Context, t *app.Topic) error {
	return nil
}
func (s *SQLiteTopicStorage) Delete(ctx context.Context, id string) error {
	return nil
}

func (s *SQLiteTopicStorage) ByID(ctx context.Context, id string) (*app.Topic, error) {
	return nil, nil
}

func (s *SQLiteTopicStorage) List(ctx context.Context) (*[]app.Topic, error) {
	ret := make([]app.Topic, 0, 16)

	var query string = `select * from topics;`
	rows, err := s.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	
	for rows.Next() {
		var t app.Topic
		var lastRecall int64

		err := rows.Scan(
			&t.Id,
			&t.Tag,
			&t.Skipped,
			&t.Completed,
			&t.Skippable,
			&lastRecall,
		)

		if err != nil {
        	return nil, err
   		}

		t.LastRecall = time.Unix(lastRecall, 0)
		ret = append(ret, t)
	}


	return &ret, nil
}