// Package handlers handlers/postsHandler.go
package handlers

import (
	"encoding/json"
	"goServer/models"
	"log"
	"net/http"
)

// PostsHandler /posts 경로의 요청을 처리합니다.
type PostsHandler struct{}

// NewPostsHandler PostsHandler의 인스턴스를 생성합니다.
func NewPostsHandler() *PostsHandler {
	return &PostsHandler{}
}

// ServeHTTP HTTP 요청을 처리합니다.
func (h *PostsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.GetAllPosts(w, r)
	case http.MethodPost:
		h.CreatePost(w, r)
	default:
		http.Error(w, "허용되지 않는 메서드입니다.", http.StatusMethodNotAllowed)
	}
}

// GetAllPosts 모든 포스트를 반환합니다.
func (h *PostsHandler) GetAllPosts(w http.ResponseWriter, r *http.Request) {
	log.Println("GetAllPosts 호출")

	posts, err := models.GetAllPosts()
	if err != nil {
		http.Error(w, "포스트를 불러오는 데 실패했습니다.", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(posts)
}

// CreatePost 새로운 포스트를 생성합니다.
func (h *PostsHandler) CreatePost(w http.ResponseWriter, r *http.Request) {
	log.Println("CreatePost 호출")

	var post models.Post
	err := json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		http.Error(w, "JSON 파싱에 실패했습니다.", http.StatusBadRequest)
		return
	}

	newPost, err := models.CreatePost(post.Title, post.Content)
	if err != nil {
		http.Error(w, "포스트를 생성하는 데 실패했습니다.", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newPost)
}
