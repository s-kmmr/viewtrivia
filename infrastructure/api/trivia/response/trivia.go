package response

import "github.com/s-kmmr/viewtrivia/domain/model"

// Trivia Trivia知識
type Trivia struct {
	ResponseCode int            `json:"response_code"`
	Result       []TriviaResult `json:"results"`
}

// TriviaResult Triviaの説明
type TriviaResult struct {
	Category        string   `json:"category"`
	Type            string   `json:"type"`
	Difficulty      string   `json:"difficulty"`
	Question        string   `json:"question"`
	CorrectAnswer   string   `json:"correct_answer"`
	InCorrectAnswer []string `json:"incorrect_answers"`
}

// MakeTrivia domain用Triviaを作る
func (t *Trivia) MakeTrivia() *model.TriviaResponse {
	if t == nil {
		return nil
	}
	response := &model.TriviaResponse{}
	result := model.TriviaResponseResult{}
	results := make([]model.TriviaResponseResult, len(t.Result))

	for i, v := range t.Result {
		result.Category = v.Category
		result.Type = v.Type
		result.Difficulty = v.Difficulty
		result.Question = v.Question
		result.CorrectAnswer = v.CorrectAnswer
		result.InCorrectAnswer = v.InCorrectAnswer
		results[i] = result
	}
	response.ResponseCode = t.ResponseCode
	response.Result = results

	return response
}
