package sqlite

import (
	"context"
	"database/sql"
	"time"

	"github.com/santuchoagus/srep/app"
)

type SQLiteTopicStorage struct {
	db *sql.DB
}

func NewSQLiteTopicStorage(db *sql.DB) *SQLiteTopicStorage {
	return &SQLiteTopicStorage{db: db}
}

func (s *SQLiteTopicStorage) Create(ctx context.Context, t *app.Topic) error {
	var query string = `
		insert into topics
		(id, tag, skipped, completed, skippable, last_recall)
		values (?, ?, ?, ?, ?, ?);`

	_, err := s.db.Exec(query, t.Id, t.Tag, t.Skipped, t.Completed, t.Skippable, t.LastRecall.Unix())
	if err != nil {
		return err
	}
	return nil
}
func (s *SQLiteTopicStorage) Update(ctx context.Context, t *app.Topic) error {
	var query string = `update topics set flag=? where id=?;`
	_, err := s.db.Exec(query, t.Tag, t.Id)
	return err
}
func (s *SQLiteTopicStorage) Delete(ctx context.Context, id string) error {
	var query string = `delete from topics where id=?;`
	_, err := s.db.Exec(query, id)
	return err
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
	defer rows.Close()

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
