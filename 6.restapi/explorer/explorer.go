package explorer

import (
	"fmt"
	"log"
	"net/http"
	"text/template"

	"github.com/IMHYEWON/hyewoncoin/4.blockchain/blockchain"
)

const templateDir string = "explorer/templates/"

var templates *template.Template

type homeData struct {
	PageTitle string
	Blocks    []*blockchain.Block
}

// ResponseWriter : 응답을 작성하는 인터페이스
// Request : 클라이언트의 요청을 나타내는 구조체 (*포인터)
func home(rw http.ResponseWriter, r *http.Request) {
	blockchain := blockchain.GetBlockChain()
	data := homeData{PageTitle: "Blockchain Home Page", Blocks: blockchain.AllBlocks()}
	// Execute : 템플릿을 렌더링하여 출력
	templates.ExecuteTemplate(rw, "home", data)
}

func add(rw http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		templates.ExecuteTemplate(rw, "add", nil)
	case "POST":
		r.ParseForm()
		data := r.Form.Get("blockData")
		blockchain.GetBlockChain().AddBlock(data)
		http.Redirect(rw, r, "/", http.StatusPermanentRedirect)
	}
}

func Start(port int) {
	// http.NewServeMux : HTTP 요청을 처리하는 새로운 라우터 생성
	handler := http.NewServeMux()

	templates = template.Must(template.ParseGlob(templateDir + "pages/*.gohtml"))
	templates = template.Must(templates.ParseGlob(templateDir + "partials/*.gohtml"))

	// Node.js의 라우터와 비슷한 기능
	http.HandleFunc("/", home)
	http.HandleFunc("/add", add)
	fmt.Printf("Listening on http://localhost:%d\n", port)

	// log.Fatal : 프로그램 에러 발생시 프로그램 종료 후 메시지 출력
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), handler))
}
