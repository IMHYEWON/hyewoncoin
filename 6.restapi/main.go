package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

const port string = ":4000"

type URLDescreption struct {
	URL         string `json:"url"` // json 태그를 사용하여 JSON 키 이름을 변경 (java의 @JsonProperty의 역할)
	Method      string `json:"method"`
	Description string `json:"description"`
	Payload     string `json:"payload,omitempty"` // omitempty : 값이 비어있으면 JSON에서 생략 (java의 @JsonInclude(Include.NON_NULL)의 역할)
	IgonreMe    string `json:"-"`                 // JSON으로 변환하지 않음 (java의 @JsonIgnore의 역할)
}

func documentation(rw http.ResponseWriter, r *http.Request) {
	data := []URLDescreption{
		{
			URL:         "/",
			Method:      "GET",
			Description: "See Documentation",
			IgonreMe:    "I'm not going to be in the JSON response",
		},
		{
			URL:         "/bloacks",
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

func main() {
	http.HandleFunc("/", documentation)
	fmt.Printf("Listening on http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
