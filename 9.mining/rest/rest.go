package rest

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/IMHYEWON/hyewoncoin/9.mining/blockchain"
	"github.com/IMHYEWON/hyewoncoin/9.mining/utils"
	"github.com/gorilla/mux"
)

var port string

type url string

// TestMarshaler : URL 타입을 JSON 형식으로 변환하는 인터페이스
// MarshalText() : URL 타입을 특정 형식으로 변환하는 TextMarshaler 인터페이스의 메서드
// Go에서는 인터페이스를 명시적으로 구현하지 않음 (signature만 맞으면 자동으로 구현)
func (u url) MarshalText() ([]byte, error) {
	return []byte(fmt.Sprintf("http://localhost%s%s", port, u)), nil
}

type urlDescreption struct {
	URL         url    `json:"url"` // json 태그를 사용하여 JSON 키 이름을 변경 (java의 @JsonProperty의 역할)
	Method      string `json:"method"`
	Description string `json:"description"`
	Payload     string `json:"payload,omitempty"` // omitempty : 값이 비어있으면 JSON에서 생략 (java의 @JsonInclude(Include.NON_NULL)의 역할)
	IgonreMe    string `json:"-"`                 // JSON으로 변환하지 않음 (java의 @JsonIgnore의 역할)
}

type addBlockBody struct {
	Message string
}

type errorResponse struct {
	ErrorMessage string `json:"errorMessage"`
}

// String() : fmt.Stringer 인터페이스를 구현하여 String() 메서드를 오버라이딩
// URLDescreption 타입을 fmt.Stringer 인터페이스로 사용할 수 있음
func (u urlDescreption) String() string {
	return "Hello I'm the URL Description"
}

func documentation(rw http.ResponseWriter, r *http.Request) {
	data := []urlDescreption{
		{
			URL:         url("/"),
			Method:      "GET",
			Description: "See Documentation",
			IgonreMe:    "I'm not going to be in the JSON response",
		},
		{
			URL:         url("/status"),
			Method:      "GET",
			Description: "See the Status of the Blockchain",
		},
		{
			URL:         url("/blocks"),
			Method:      "POST",
			Description: "Add a block",
			Payload:     "data:string",
		},
		{
			URL:         url("/blocks/{height}"),
			Method:      "GET",
			Description: "See a block",
		},
	}

	// json.NewEncoder : JSON 인코딩을 위한 인코더 생성
	json.NewEncoder(rw).Encode(data)
}

func blocks(rw http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		rw.Header().Add("Content-Type", "application/json")
		json.NewEncoder(rw).Encode(blockchain.BlockChain().Blocks())
	case "POST":
		// request body to block struct
		var addBlockBody addBlockBody

		// json.NewDecoder : JSON 디코딩을 위한 디코더 생성
		utils.HandleErr(json.NewDecoder(r.Body).Decode(&addBlockBody))
		fmt.Println(addBlockBody)

		blockchain.BlockChain().AddBlock(addBlockBody.Message)
		rw.WriteHeader(http.StatusCreated)
	}

}

func block(rw http.ResponseWriter, r *http.Request) {
	// mux.Vars : URL에서 변수를 추출하여 map으로 반환
	vars := mux.Vars(r)
	// strconv.Atoi : 문자열을 정수로 변환
	hash := vars["hash"]

	block, err := blockchain.FindBlock(hash)
	encoder := json.NewEncoder(rw)
	if err == blockchain.ErrNotFound {
		encoder.Encode(errorResponse{fmt.Sprint(err)})
	} else {
		encoder.Encode(block)
	}
}

func status(rw http.ResponseWriter, r *http.Request) {
	json.NewEncoder(rw).Encode(blockchain.BlockChain())
}

func jsonContentTypeMiddleWare(next http.Handler) http.Handler {
	// 내부적으로 NextServeHTTP 메서드를 호출하여 다음 핸들러로 요청을 전달
	// API 요청이 들어올 때 getBlock같은 메소드 전에 실행

	// HandlerFunc : function이 아닌 type이므로 실제로 http.HandleFunc ~ 처럼 사용시에는 type 변수가 생성됨,
	// 그럼 어떻게 HandlerFunc type이 HandlerFunc Interface로 인식하고 구현되는걸까?
	// -> adapter pattern : adapter에게 적절한 argument를 넘겨주면 adapter가 해당 argument를 받아서 적절한 type으로 변환하여 사용
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(rw, r)
	})
}

func Start(aPort int) {
	// http.NewServeMux : HTTP 요청을 처리하는 새로운 라우터 생성
	router := mux.NewRouter()

	// port 전역 변수에 포트 번호 저전
	port = fmt.Sprintf(":%d", aPort)

	router.Use(jsonContentTypeMiddleWare)
	router.HandleFunc("/", documentation).Methods("GET")
	router.HandleFunc("/status", status).Methods("GET")
	router.HandleFunc("/blocks", blocks).Methods("GET", "POST")
	router.HandleFunc("/blocks/{hash:[a-f0-9]+}", block).Methods("GET") // {id:[0-9]+} : 정규표현식으로 숫자만 받음

	fmt.Printf("Listening on http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, router))
}
