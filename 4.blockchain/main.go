package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/IMHYEWON/hyewoncoin/4.blockchain/blockchain"
)

const port string = ":4000"

type homeData struct {
	PageTitle string
	Blocks    []*blockchain.Block
}

// ResponseWriter : 응답을 작성하는 인터페이스
// Request : 클라이언트의 요청을 나타내는 구조체 (*포인터)
func home(rw http.ResponseWriter, r *http.Request) {
	// Fprint : 문자열을 콘솔이 아닌 writer에 출력
	// fmt.Fprint(rw, "Hello from home")
	//tmpl, err := template.ParseFiles("templates/home.html")

	// Go에는 Exception이나 Error가 없으므로, 직접 처리해야함
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// Must : 템플릿을 파싱하고 에러가 있으면 프로그램을 종료 (에러핸들링을 하지 않아도 됨핸
	tmpl := template.Must(template.ParseFiles("templates/home.html"))

	blockchain := blockchain.GetBlockChain()
	data := homeData{PageTitle: "Blockchain Home Page", Blocks: blockchain.AllBlocks()}
	// Execute : 템플릿을 렌더링하여 출력
	tmpl.Execute(rw, data)
}

func main() {
	// Node.js의 라우터와 비슷한 기능
	http.HandleFunc("/", home)
	fmt.Printf("Listening on http://localhost%s\n", port)

	// log.Fatal : 프로그램 에러 발생시 프로그램 종료 후 메시지 출력
	log.Fatal(http.ListenAndServe(port, nil))
}
