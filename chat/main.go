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
	"github.com/stretchr/objx"
	"os"
	"trace"
)

var avatars Avatar=TryAvatars{
	UseFileSystemAvatar,
	UseAuthAvatar,
	UseGravatar,
}

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
	data :=map[string]interface{}{
		"Host": r.Host,
	}
	if authCookie, err := r.Cookie("auth"); err == nil{
		data["UserData"] = objx.MustFromBase64(authCookie.Value)
	}
	
	t.templ.Execute(w, data)
}

func main() {
	var addr = flag.String("addr",":8080","アプリケーションのアドレス")
	flag.Parse();

	//Gomniauthのセットアップ
	gomniauth.SetSecurityKey("hfioeferuielfjeho137udje73d83thgd62")
	gomniauth.WithProviders(
		facebook.New("クライアントID","秘密鍵","http://localhost:8080/auth/callback/facebook"),
		github.New("クライアントID","秘密鍵","http://localhost:8080/auth/callback/github"),
		google.New("191903503672-2ne2dr00ucfcr4l81g5q7opmfbbfnhss.apps.googleusercontent.com","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","http://localhost:8080/auth/callback/google"),
	)

	r := newRoom(UseFileSystemAvatar)

	//取得したログをコンソールに表示
	r.tracer = trace.New(os.Stdout)

	//ルート
	http.Handle("/chat", MustAuth(&templateHandler{filename: "chat.html"}))
	http.Handle("/login",&templateHandler{filename: "login.html"})
	http.HandleFunc("/auth/",loginHandler)
	http.Handle("/room", r)
	http.HandleFunc("/logout",func(w http.ResponseWriter, r *http.Request){
		http.SetCookie(w, &http.Cookie{
			Name: "auth",
			Value: "",
			Path: "/",
			MaxAge: -1,
		})
		w.Header()["Location"] = []string{"/chat"}
		w.WriteHeader(http.StatusTemporaryRedirect)
	})
	http.Handle("/upload",&templateHandler{filename: "upload.html"})
	http.HandleFunc("/uploader",uploadHandler)
	http.Handle("/avatars/",
		http.StripPrefix("/avatars/",
			http.FileServer(http.Dir("./avatars"))))

	//チャットルームを開始
	go r.run()

	//webサーバの開始
	log.Println("Webサーバーを開始します　ポート: ",*addr)
	if err := http.ListenAndServe(*addr, nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
