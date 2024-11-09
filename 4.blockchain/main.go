package main

import (
	"fmt"
	"log"
	"net/http"
)

const port string = ":4000"

// ResponseWriter : 응답을 작성하는 인터페이스
// Request : 클라이언트의 요청을 나타내는 구조체 (*포인터)
func home(rw http.ResponseWriter, r *http.Request) {
	// Fprint : 문자열을 콘솔이 아닌 writer에 출력
	fmt.Fprint(rw, "Hello from home")
}

func main() {
	// Node.js의 라우터와 비슷한 기능
	http.HandleFunc("/", home)
	fmt.Printf("Listening on http://localhost%s\n", port)

	// log.Fatal : 프로그램 에러 발생시 프로그램 종료 후 메시지 출력
	log.Fatal(http.ListenAndServe(port, nil))
}
