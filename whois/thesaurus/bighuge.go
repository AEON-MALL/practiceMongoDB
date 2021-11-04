package thesaurus

import (
	"encoding/json"
	"net/http"
	"fmt"
)

type BigHugh struct{
	APIkey string
}

type synonyms struct{
	Noun *words `json:"noun"`
	Verb *words `json:"verb"`
}

type words struct {
	Syn [] string `json:"syn"` 
}

func (b *BigHugh) Synonyms(term string) ([]string ,error){
	var syns []string
	response, err := http.Get("http://words.bighugelabs.co/api/2/" + b.APIkey + "/ " + term + "/json")
	if err != nil {
		return syns , fmt.Errorf("bighuge : %qの類語検索に失敗しました: %v", term, err)
	}
	var data synonyms
	defer response.Body.Close()
	if err := json.NewDecoder(response.Body).Decode(&data); err != nil{
		return syns, err 
	}
	syns = append(syns, data.Noun.Syn...)
	syns = append(syns, data.Verb.Syn...)
	return syns , nil
}