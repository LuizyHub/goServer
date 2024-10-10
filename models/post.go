// Package models
package models

// Post 블로그 포스트를 나타냅니다.
type Post struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

// CreatePost 새로운 포스트를 데이터베이스에 추가합니다.
func CreatePost(title, content string) (*Post, error) {
	result, err := DB.Exec("INSERT INTO posts (title, content) VALUES (?, ?)", title, content)
	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	return &Post{
		ID:      int(id),
		Title:   title,
		Content: content,
	}, nil
}

// GetAllPosts 모든 포스트를 불러옵니다.
func GetAllPosts() ([]Post, error) {
	rows, err := DB.Query("SELECT id, title, content FROM posts")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	posts := []Post{}
	for rows.Next() {
		var post Post
		err := rows.Scan(&post.ID, &post.Title, &post.Content)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	return posts, nil
}

// GetPostByID 특정 ID의 포스트를 불러옵니다.
func GetPostByID(id int) (*Post, error) {
	row := DB.QueryRow("SELECT id, title, content FROM posts WHERE id = ?", id)
	var post Post
	err := row.Scan(&post.ID, &post.Title, &post.Content)
	if err != nil {
		return nil, err
	}
	return &post, nil
}

// UpdatePost 기존 포스트를 수정합니다.
func UpdatePost(id int, title, content string) error {
	_, err := DB.Exec("UPDATE posts SET title = ?, content = ? WHERE id = ?", title, content, id)
	return err
}

// DeletePost 특정 ID의 포스트를 삭제합니다.
func DeletePost(id int) error {
	_, err := DB.Exec("DELETE FROM posts WHERE id = ?", id)
	return err
}
