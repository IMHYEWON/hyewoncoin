package rest

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/IMHYEWON/hyewoncoin/6.restapi/blockchain"
	"github.com/IMHYEWON/hyewoncoin/6.restapi/utils"
)

var port string

type URL string

// TestMarshaler : URL 타입을 JSON 형식으로 변환하는 인터페이스
// MarshalText() : URL 타입을 특정 형식으로 변환하는 TextMarshaler 인터페이스의 메서드
// Go에서는 인터페이스를 명시적으로 구현하지 않음 (signature만 맞으면 자동으로 구현)
func (u URL) MarshalText() ([]byte, error) {
	return []byte(fmt.Sprintf("http://localhost%s%s", port, u)), nil
}

type urlDescreption struct {
	URL         URL    `json:"url"` // json 태그를 사용하여 JSON 키 이름을 변경 (java의 @JsonProperty의 역할)
	Method      string `json:"method"`
	Description string `json:"description"`
	Payload     string `json:"payload,omitempty"` // omitempty : 값이 비어있으면 JSON에서 생략 (java의 @JsonInclude(Include.NON_NULL)의 역할)
	IgonreMe    string `json:"-"`                 // JSON으로 변환하지 않음 (java의 @JsonIgnore의 역할)
}

type addBlockBody struct {
	Message string
}

// String() : fmt.Stringer 인터페이스를 구현하여 String() 메서드를 오버라이딩
// URLDescreption 타입을 fmt.Stringer 인터페이스로 사용할 수 있음
func (u urlDescreption) String() string {
	return "Hello I'm the URL Description"
}

func documentation(rw http.ResponseWriter, r *http.Request) {
	data := []urlDescreption{
		{
			URL:         URL("/"),
			Method:      "GET",
			Description: "See Documentation",
			IgonreMe:    "I'm not going to be in the JSON response",
		},
		{
			URL:         URL("/blocks"),
			Method:      "POST",
			Description: "Add a block",
			Payload:     "data:string",
		},
	}

	// 응답을 application/json으로 설정
	rw.Header().Add("Content-Type", "application/json")

	// json.NewEncoder : JSON 인코딩을 위한 인코더 생성
	json.NewEncoder(rw).Encode(data)
}

func blocks(rw http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		rw.Header().Add("Content-Type", "application/json")
		json.NewEncoder(rw).Encode(blockchain.GetBlockChain().AllBlocks())
	case "POST":
		// request body to block struct
		var addBlockBody addBlockBody

		// json.NewDecoder : JSON 디코딩을 위한 디코더 생성
		utils.HandleErr(json.NewDecoder(r.Body).Decode(&addBlockBody))
		fmt.Println(addBlockBody)

		blockchain.GetBlockChain().AddBlock(addBlockBody.Message)
		rw.WriteHeader(http.StatusCreated)
	}

}

func Start(aPort int) {
	// http.NewServeMux : HTTP 요청을 처리하는 새로운 라우터 생성
	handler := http.NewServeMux()

	// port 전역 변수에 포트 번호 저전
	port = fmt.Sprintf(":%d", aPort)

	handler.HandleFunc("/", documentation)
	handler.HandleFunc("/blocks", blocks)

	fmt.Printf("Listening on http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, handler))
}