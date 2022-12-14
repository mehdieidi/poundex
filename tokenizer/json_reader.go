package tokenizer

import "encoding/json"

type Tweet struct {
	Text string `json:"text"`
}

func ReadJSON(text string) (Tweet, error) {
	var tweet Tweet
	if err := json.Unmarshal([]byte(text), &tweet); err != nil {
		return Tweet{}, err
	}
	return tweet, nil
}
