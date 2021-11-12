package main

import (
	"io"
	"log"
	"net"
	"net/http"
	"sync"
	"time"

	"github.com/garyburcd/go-oauth/oauth"
	"github.com/joeshaw/envdecode"
)

var conn net.Conn
func deal(netw, addr string) (net.Conn, error){
	if conn != nil{
		conn.Close()
		conn = nil
	}
	netc, err := net.DialTimeout(netw, addr, 5*time.Second)
	if err != nil{
		return nil,err
	}
	conn = netc
	return conn,err
}

var reader io.ReadCloser
func closeConn(){
	if conn != nil{
		conn.Close()
	}
	if reader != nil{
		reader.Close()
	}
}

var (
	authClient *oauth.Client
	creds *oauth.Credentials
)

func setupTwitterAuth(){
	var ts struct{
		ConsumarKey string `env:"SP_TWITTER_KEY,required"`
		ConsumarSecret string `env:"SP_TWITTER_SECRET,required"`
		AccessToken string `env:"SP_TWITTER_ACCESSTOKEN,required"`
		AccessSectet string `env:"SP_TWITTER_ACCESSSECRET,required"`
 	}
	if err := envdecode.Decode(&ts); err!= nil{
		log.Fatalln(err)
	}
	creds = &oauth.Credentials{
		Token: ts.AccessToken,
		Secret: ts.AccessSectet,
	}
	authClient = &oauth.Client{
		Credentials: oauth.Credentials{
			Token: ts.ConsumarKey,
			Secret: ts.ConsumarSecret,
		},
	}

	var (
		authSetupOnce sync.Once
		httpClient *http.Client
	)
	func makeRequest(req *http.Request, params url.Values) (*http.Respose, error){
		authSetupOnce.Do(sunc(){
			setupTwitterAuth()
			httpClient = &http.Client{
				Transport: &http.Transport{
					Dial: dial,
				},
			}
		})
		formEnc := param.Encode()
		req.Header.Set("Content-Type","application/x-www-form-urlencoded")
		req.Header.Set("Content-Length",strconv.Itoa(len(formEnc)))
		req.Header.Set("Authorization",authClient.AuthorizationHeader(creds,"POAT",req.URL, params))
		return httpClient.Do(req)
	}
}