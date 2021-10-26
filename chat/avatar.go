package main

import (
	"errors"
)

//ErrNoAvatarはAvatarインスタンスがアバターのURLを返すことができない場合に発生するエラー
var ErrNoAvatarURL = errors.New("chat: アバターのURLを取得できません。")

//Avatarはユーザーのプロフィール画像を表す型です
type Avatar interface {
	//AvatarURLは指定されたクライアントのアバターのURLを返す。
	//問題が発生した場合エラーを返す。特にURLを取得できなかった場合にはErrNoAvataURLを返す
	AvatarURL(c *client) (string, error)
}

type AuthAvatar struct {}
var UseAuthAvatar AuthAvatar
func (AuthAvatar) AvatarURL(c *client)(string,error){
	if url, ok := c.userData["avatar_url"]; ok{
		if urlStr, ok := url.(string); ok{
			return urlStr,nil
		}
	}
	return "",ErrNoAvatarURL
}

type GravatarAvatar struct{}
var UseGravatar GravatarAvatar
func (GravatarAvatar) AvatarURL(c *client) (string,error){
	if userid, ok := c.userData["userid"]; ok{
		if useridStr ,ok := userid.(string); ok{
			return "//www.gravatar.com/avatar/"+ useridStr, nil
		}
	}
	return "",ErrNoAvatarURL
}

