package request

import "github.com/s-kmmr/viewtrivia/domain/model"

// Trivia Triviaリクエスト用モデル
type Trivia struct {
	Amount int
}

// NewTrivia Triviaリクエストコンストラクタ
func NewTrivia(trivia *model.TriviaRequest) *Trivia {
	return &Trivia{
		Amount: trivia.Amount,
	}
}
