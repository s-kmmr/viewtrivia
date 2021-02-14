package usecase

import (
	"github.com/s-kmmr/viewtrivia/domain/model"
	"github.com/s-kmmr/viewtrivia/domain/repository"
	"golang.org/x/xerrors"
)

type trivia struct {
	tr repository.Trivia
}

// Trivia Triviaユースケースインターフェース定義
type Trivia interface {
	Show(req *model.TriviaRequest) (*model.TriviaResponse, error)
}

// NewTrivia Triviaユースケースコンストラクタ
func NewTrivia(tr repository.Trivia) Trivia {
	return &trivia{
		tr: tr,
	}
}

// Show Triviaを見る
func (t *trivia) Show(req *model.TriviaRequest) (*model.TriviaResponse, error) {
	response, err := t.tr.Show(req)
	if err != nil {
		return nil, xerrors.Errorf("repository.trivia.Show(): %w", err)
	}
	return response, nil
}
