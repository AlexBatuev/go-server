package sqlite

import (
	"database/sql"
	"errors"
	"go-server/pkg/models"
	"time"
)

// SnippetModel - Определяем тип который обертывает пул подключения sql.DB
type SnippetModel struct {
	DB *sql.DB
}

func (m *SnippetModel) Insert(title, content string, expires int64) (int, error) {
	query := "INSERT INTO snippets (title,content,created,expires) VALUES (?,?,?,?);"
	now := time.Now()
	result, err := m.DB.Exec(query, title, content, now.Unix(), now.Unix()+expires)
	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(id), nil
}

// Get - Метод для возвращения данных заметки по её идентификатору ID.
func (m *SnippetModel) Get(id int) (*models.Snippet, error) {
	query := "SELECT ROWID, title, content, created, expires FROM snippets WHERE expires > ? AND ROWID = ?;"
	now := time.Now().Unix()
	row := m.DB.QueryRow(query, now, id)
	s := &models.Snippet{}
	err := row.Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, models.ErrNoRecord
		} else {
			return nil, err
		}
	}
	return s, nil
}

// Latest - Метод возвращает 10 наиболее часто используемые заметки.
func (m *SnippetModel) Latest() ([]*models.Snippet, error) {
	return nil, nil
}
