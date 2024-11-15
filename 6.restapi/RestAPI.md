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