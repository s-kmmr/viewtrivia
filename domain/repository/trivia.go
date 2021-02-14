package repository

import "github.com/s-kmmr/viewtrivia/domain/model"

// Trivia Triviaリポジトリインターフェース定義
type Trivia interface {
	Show(req *model.TriviaRequest) (*model.TriviaResponse, error)
}
