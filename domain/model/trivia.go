package model

// TriviaResponse Trivia知識
type TriviaResponse struct {
	ResponseCode int                    `json:"response_code"`
	Result       []TriviaResponseResult `json:"results"`
}

// TriviaResponseResult Triviaの説明
type TriviaResponseResult struct {
	Category        string   `json:"category"`
	Type            string   `json:"type"`
	Difficulty      string   `json:"difficulty"`
	Question        string   `json:"question"`
	CorrectAnswer   string   `json:"correct_answer"`
	InCorrectAnswer []string `json:"incorrect_answers"`
}

// TriviaRequest Triviaリクエスト用モデル
type TriviaRequest struct {
	Amount int
}
