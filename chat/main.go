package main

import (
	"flag"
	"log"
	"net/http"
	"path/filepath"
	"sync"
	"text/template"

	"github.com/stretchr/gomniauth"
	"github.com/stretchr/gomniauth/providers/facebook"
	"github.com/stretchr/gomniauth/providers/github"
	"github.com/stretchr/gomniauth/providers/google"
	//"os"
	//"trace"
)

//templは一つのテンプレートを表す
type templateHandler struct {
	once     sync.Once
	filename string
	templ    *template.Template
}

//ServeHTTPはHTTPリクエストを処理する
func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t.once.Do(func() {
		t.templ = template.Must(template.ParseFiles(filepath.Join("templates", t.filename)))
	})
	t.templ.Execute(w, r)
}

func main() {
	var addr = flag.String("addr",":8080","アプリケーションのアドレス")
	flag.Parse();
	//Gomniauthのセットアップ
	gomniauth.SetSecurityKey("シークレットキー")
	gomniauth.WithProviders(
		facebook.New("クライアントID","秘密鍵","http://localhost:8080/auth/callback/facebook"),
		github.New("クライアントID","秘密鍵","http://localhost:8080/auth/callback/github"),
		google.New("クライアントID","秘密鍵","http://localhost:8080/auth/callback/google"),
	)
	r := newRoom()
	//取得したログをコンソールに表示
	//r.tracer = trace.New(os.Stdout)
	//ルート
	http.Handle("/chat", MustAuth(&templateHandler{filename: "chat.html"}))
	http.Handle("/login",&templateHandler{filename: "login.html"})
	http.HandleFunc("/auth/",loginHandler)
	http.Handle("/room", r)
	//チャットルームを開始
	go r.run()
	//webサーバの開始
	log.Println("Webサーバーを開始します　ポート: ",*addr)
	if err := http.ListenAndServe(*addr, nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
