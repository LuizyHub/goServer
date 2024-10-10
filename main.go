// main.go
package main

import (
	"goServer/handlers"
	"goServer/models"
	"log"
	"net/http"
)

func main() {
	// 데이터베이스 초기화
	models.InitDB("data/blog.db")

	// 로그 설정
	log.Println("서버를 시작합니다...")

	// 핸들러 인스턴스 생성
	postsHandler := handlers.NewPostsHandler()
	postHandler := handlers.NewPostHandler()

	// 라우트 설정
	http.Handle(handlers.PostsPathPrefix+"/", postHandler) // 먼저 등록
	http.Handle(handlers.PostsPathPrefix, postsHandler)

	// 정적 파일 제공
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)

	// 서버 시작
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("서버 시작 실패:", err)
	}
}
