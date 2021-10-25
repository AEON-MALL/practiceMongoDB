package main
import "testing"
func TestAuthAvatar(t *testing.T){
	var authAvatar AuthAvatar
	client :=new(client)
	url , err := authAvatar.AvatarURL(client)
	if err != ErrNoAvatarURL{
		t.Error("値が存在しない場合、AuthAvatar.AvatarURLは"+"ErrNoAvatarURLを返すべきです")
	}
	//値をセット
	testUrl := "http://url-to-avatar"
	client.userData = map[string]interface{}{"avatar_url":testUrl}
	url, err = authAvatar.AvaterURL(client)
	if err != nil{
		t.Error("値が存在する場合、AuthAvatar.AvatarURLは"+"エラーを返すべきではありません")
	}else{
		if url != testUrl{
			t.Error("AuthAvatar.AvatarURLは正しいURLを返すべきです")
		}
	}
}