# 6. REST API
## 6.1 Marshal and Field Tags
- Marshal : sturct를 JSON으로 변환
  - `b, err = json.Marshal(data)`
- json.NewEncoder : JSON 인코딩을 위한 인코더 생성
  - `json.NewEncoder(rw).Encode(data)`
- sturct 의 json property (field tags)
  - `json:"url"` 
    - json 태그를 사용하여 JSON 키 이름을 변경 (java의 @JsonProperty의 역할)
  - `json:"payload,omitempty"` 
    - omitempty : 값이 비어있으면 JSON에서 생략 (java의 @JsonInclude(Include.NON_NULL)의 역할)
  - `json:"-"`
    - JSON으로 변환하지 않음 (java의 @JsonIgnore의 역할)

## 6.2 Marshal Text
- Stringers Interface
  - fmt package의 인터페이스로 이 인터페이스를 구현하면 struct의 형태를 fmt Print로 조절가능
  - Go에는 java, python과 같이 상속의 개념이 없기 때문에
  - Go에게 Stringer interface라고 말해줄 필요가 없음, 단지 **Signature가 동일한 인터페이스의 메소드를 오버라이딩 하면 됨**
- MarshalText 
  - Field가 Json String으로써 어떻게 보일지 결정하는 메소드

## 6.3 JSON DECODE
- POST method에서 Request Bo용dy를 받기 위해 새로운 리퀘스트 struct 생성 및 decode
  - `json.NewDecoder(r.Body).Decode(&addBlockBody)` 

## 6.4 NewServeMux
- Mux : Multiplexer
  - handling request by url
  - 같은 multiplexer 를 사용해서 rest.go, explore.go 핸들링
  - 새로운 handler를 생성해서 이를 사용

## 6.5 Gorila Mux
- Gorila Mux
  - Gorila : go의 toolkit 으로 context, session, mux, rpc ...사용 가능
  - Gorila Mux는 path variable 등 기능 지원 
  - `go get -u github.com/gorilla/mux`
  - mux.Vars : URL에서 변수를 추출하여 map으로 반환

## 6.6 Atoi
- Height를 리퀘스트로 받아 특정 블록하나만 보여주는 함수
  - mux.Vars의 Map Value는 String 이기 때문에 int로 사용하려면 변환해야함
- Go에는 형변환만을 위한 패키지가 있음
  - `package strconv` 

## 6.7 Error Handling
- method에서 error 리턴
  -  API 에서는 error response를 만들어서 이를 리턴

## 6.8 Middlewares
- API Response를 json으로 설정하는 미들웨어 함수 생성
  - 내부적으로 NextServeHTTP 메서드를 호출하여 다음 핸들러로 요청을 전달
  - 라우터에서 미들웨어 함수 사용
- Adapter Pattern :
  - 호환되지 않는 인터페이스를 가진 두 클래스가 함께 동작할 수 있도록 연결하는 역할을 합니다. 마치 "어댑터"가 전압이 다른 전자제품을 연결하는 것처럼, 특정 객체의 인터페이스를 클라이언트가 기대하는 인터페이스로 변환
  - Handler 타입객체를 직접 생성 & ServeHTTP 인터페이스 함수를 구현하는 것이 아니고, HandlerFunc으로 정의된 타입의 객체를 생성하면 (적절한 Signature를 전했을 때) HandlerFunc이 대신 구현해줌
  - Adapter Pattern을 사용하는 이유? 
    - 다른 인터페이스를 사용하는 기존 코드와 통합해야 할 때 유용
    - 유연한 미들웨어 구성: 여러 미들웨어를 조합하여 요청 처리 로직을 동적으로 구성할 때.
    - 인터페이스와 함수의 결합: 특정 인터페이스를 만족시키기 위해 함수를 래핑할 때.
  - Middleware 함수의 argument, return value 모두 `Handler` Type
    - Go의 HTTP 서버는 모든 요청 처리를 위해 http.Handler 인터페이스를 사용
    ``` Go
    type Handler interface {
      ServeHTTP(ResponseWriter, *Request)
    }
    ```
  - HandlerFunc (Adapter)
    ``` Go
    // The HandlerFunc type is an adapter to allow the use of
    // ordinary functions as HTTP handlers. If f is a function
    // with the appropriate signature, HandlerFunc(f) is a
    // [Handler] that calls f.
    // HandlerFunc는 ServeHTTP 메서드를 구현함으로써 Handler 인터페이스를 충족
    type HandlerFunc func(ResponseWriter, *Request)

    // ServeHTTP calls f(w, r).
    func (f HandlerFunc) ServeHTTP(w ResponseWriter, r *Request) {
      f(w, r) // // 실제로는 정의된 함수 f를 호출
    }
    ```