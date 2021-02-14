package handler

import (
	"log"
	"net/http"

	"github.com/labstack/echo"
	"github.com/s-kmmr/viewtrivia/domain/model"
	"github.com/s-kmmr/viewtrivia/usecase"
)

type trivia struct {
	tu usecase.Trivia
}

// TriviaRequest リクエストバインド用
type TriviaRequest struct {
	Amount int `json:"amount" query:"amount" validate:"required"`
}

// Trivia Triviaハンドラーインターフェース定義
type Trivia interface {
	ShowTrivia(c echo.Context) error
}

// NewTrivia Triviaハンドラーコンストラクタ
func NewTrivia(u usecase.Trivia) Trivia {
	return &trivia{
		tu: u,
	}
}

// ShowTrivia Trivia知識を返却する
func (t *trivia) ShowTrivia(c echo.Context) error {
	r := &TriviaRequest{}
	if err := c.Bind(r); err != nil {
		log.Printf("%+v", err)
		return err
	}

	if err := c.Validate(r); err != nil {
		log.Printf("%+v", err)
		return err
	}

	req := &model.TriviaRequest{
		Amount: r.Amount,
	}
	res, err := t.tu.Show(req)
	if err != nil {
		log.Printf("%+v", err)
		return c.JSON(http.StatusBadRequest, res)
	}
	return c.JSON(http.StatusOK, res)

}
