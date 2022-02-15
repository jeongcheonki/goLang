package myapp

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// json 담을 스트럭쳐 만들기
type User struct {
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	createdAt time.Time `json:"createdAT"`
}

// fooHandler 인스턴스를 만듬
type fooHandler struct{}

// fooHandler 인터페이스를 구현함.
func (f *fooHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprint(w, "Hello Foo")
	user := new(User)
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Bad request: ", err)
		return
	}
	user.createdAt = time.Now()

	data, _ := json.Marshal(user)
	w.Header().Add("content-type", "application/json")
	// 문서 헤더에 문서 형태가 json이라 알려줌.
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, string(data))

}

func barhandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		name = "Worlddd"
	}
	fmt.Fprintf(w, "Hello %s!", name)
	// 브라우져에 http://localhost:3000/bar?name=bbbbbb 주소를 넣어본다.
}

func NewHttpHandler() http.Handler {
	// 라오터 클래스를 만들어 등록 함.
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// HandleFunc 핸들러를 등록한다. 어떤 경로의 요청이 들어왔을때 어떤 작업을 한것인지 핸들.
		// "/" 인덱스 절대 경로
		// w, r 인자가 있다.
		fmt.Fprint(w, "Hellow 월드")
		// Fprint는 writer에 프린트해라.
		// w -> http.ResponseWriter
		// Hellow 월드 라고 리스폰스를 줘라.
	})

	/*	http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "hello bar")
	}) */
	mux.HandleFunc("/bar", barhandler)
	// 핸들러를 인스턴스 형태로 등록했을 경우
	mux.Handle("/foo", &fooHandler{})
	return mux

	// ListenAndServe 웹서버를 실행 해라. 구동~~~~
	// 3000 번 포트 사용
	// http.ListenAndServe(":3000", mux)
	// mux로 라우터 등록

}
