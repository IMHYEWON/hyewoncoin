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