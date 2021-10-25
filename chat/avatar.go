package main
import(
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