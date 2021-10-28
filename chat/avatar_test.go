package main

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

import gomniauthtest "github.com/stretchr/gomniauth/test"
func TestAuthAvatar(t *testing.T){
	var authAvatar AuthAvatar
	testUser := &gomniauthtest.TestUser{}
	testUser.On("AvatarURL").Return("",ErrNoAvatarURL)
	testChatUser := &chatUser{User: testUser}
	_, err := authAvatar.GetAvatarURL(testChatUser)
	if err != ErrNoAvatarURL{
		t.Error("値が存在しない場合、AuthAvatar.AvatarURLはErrNoAvatarURLを返すべきです")
	}
	//値をセット
	testUrl := "http://url-to-avatar"
	testChatUser.User = testUser
	testUser.On("AvatarURL").Return(testUrl,nil)
	url, err := authAvatar.GetAvatarURL(testChatUser)
	if err != nil{
		t.Error("値が存在する場合、AuthAvatar.AvatarURLはエラーを返すべきではありません")
	}else{
		if url != testUrl{
			t.Error("AuthAvatar.AvatarURLは正しいURLを返すべきです")
		}
	}
}

func TestGravatarAvatar (t *testing.T){
	var gravatarAvatar GravatarAvatar
	user := &chatUser{uniqueID: "abc"}
	url,err := gravatarAvatar.GetAvatarURL(user)
	if err != nil{
		t.Error("GravatarAvatar.AvatarURLはエラーを返すべきではありません")
	}
	if url != "//www.gravatar.com/avatar/abc"{
		t.Errorf("GravatarAvatar.AvatarURLが%sという誤った値を返しました", url)
	}
}

func TestFileSystemAvatar(t *testing.T){
	//テスト用のアバターのファイルを生成する
	filename := filepath.Join("avatars","abc.jpg")
	ioutil.WriteFile(filename, []byte{},0777)
	defer func ()  { os.Remove(filename) } ()

	var fileSystemAvatar FileSystemAvatar
	user := &chatUser{uniqueID: "abc"}
	url, err := fileSystemAvatar.GetAvatarURL(user)
	if err != nil{
		t.Error("FileSystemAvatar.AvatarURLはエラーを返すべきではありません")
	}
	if url != "/avatars/abc.jpg" {
		t.Errorf("FileSystemAvatar.AvatarURLが%sという誤った値を返しました、",url)
	}
}