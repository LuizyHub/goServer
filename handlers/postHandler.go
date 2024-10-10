// Package handlers handlers/postHandler.go
package handlers

import (
	"database/sql"
	"encoding/json"
	"errors"
	"goServer/models"
	"goServer/utils"
	"log"
	"net/http"
	"strconv"
	"strings"
)

// PostHandler PostHandler는 /posts/{id} 경로의 요청을 처리합니다.
type PostHandler struct{}

// NewPostHandler PostHandler의 인스턴스를 생성합니다.
func NewPostHandler() *PostHandler {
	return &PostHandler{}
}

// ServeHTTP HTTP 요청을 처리합니다.
func (h *PostHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, PostsPathPrefix+"/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "유효하지 않은 ID입니다.", http.StatusBadRequest)
		return
	}

	switch r.Method {
	case http.MethodGet:
		h.GetPost(w, r, id)
	case http.MethodPut:
		h.UpdatePost(w, r, id)
	case http.MethodDelete:
		h.DeletePost(w, r, id)
	default:
		http.Error(w, "허용되지 않는 메서드입니다.", http.StatusMethodNotAllowed)
	}
}

// GetPost 특정 포스트를 반환합니다.
func (h *PostHandler) GetPost(w http.ResponseWriter, r *http.Request, id int) {
	log.Println("GetPost 호출")

	post, err := models.GetPostByID(id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			utils.WriteErrorResponse(w, http.StatusNotFound, "포스트를 찾을 수 없습니다.")
			return
		}
		utils.WriteErrorResponse(w, http.StatusInternalServerError, "서버 오류가 발생했습니다.")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(post)
}

// UpdatePost 포스트를 수정합니다.
func (h *PostHandler) UpdatePost(w http.ResponseWriter, r *http.Request, id int) {
	log.Println("UpdatePost 호출")

	var post models.Post
	err := json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		http.Error(w, "JSON 파싱에 실패했습니다.", http.StatusBadRequest)
		return
	}

	err = models.UpdatePost(id, post.Title, post.Content)
	if err != nil {
		http.Error(w, "포스트를 수정하는 데 실패했습니다.", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// DeletePost 포스트를 삭제합니다.
func (h *PostHandler) DeletePost(w http.ResponseWriter, r *http.Request, id int) {
	log.Println("DeletePost 호출")

	err := models.DeletePost(id)
	if err != nil {
		http.Error(w, "포스트를 삭제하는 데 실패했습니다.", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
