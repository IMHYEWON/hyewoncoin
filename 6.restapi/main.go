package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/IMHYEWON/hyewoncoin/6.restapi/utils"
)

const port string = ":4000"

type URLDescreption struct {
	URL         string
	Method      string
	Description string
}

func documentation(rw http.ResponseWriter, r *http.Request) {
	data := []URLDescreption{
		{
			URL:         "/",
			Method:      "GET",
			Description: "See Documentation",
		},
	}

	rw.Header().Add("Content-Type", "application/json")
	// json.Marshal : 구조체를 JSON으로 변환
	b, err := json.Marshal(data)
	utils.HandleErr(err)
	fmt.Printf("%s", b)
	fmt.Fprintf(rw, "%s", b)
}

func main() {
	http.HandleFunc("/", documentation)
	fmt.Printf("Listening on http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
