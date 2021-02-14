package trivia

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"

	"golang.org/x/xerrors"

	"github.com/s-kmmr/viewtrivia/domain/model"
	"github.com/s-kmmr/viewtrivia/domain/repository"
	"github.com/s-kmmr/viewtrivia/infrastructure/api/trivia/request"
	"github.com/s-kmmr/viewtrivia/infrastructure/api/trivia/response"
)

type trivia struct {
	h *http.Client
	u string
}

// NewTrivia Triviaリポジトリコンストラクタ
func NewTrivia() repository.Trivia {
	return &trivia{
		h: new(http.Client),
		u: "https://opentdb.com/api.php",
	}
}

const (
	queryParamAmount = "amount"
)

// Show TriviaAPIにリクエストして、質問と解答を受け取る
func (t *trivia) Show(req *model.TriviaRequest) (*model.TriviaResponse, error) {

	u, err := url.Parse(t.u)
	if err != nil {
		return nil, xerrors.Errorf("url.Parse(): %w", err)
	}

	r := request.NewTrivia(req)
	q := u.Query()

	if r.Amount > 0 {
		q.Add(queryParamAmount, strconv.Itoa(r.Amount))
	}
	u.RawQuery = q.Encode()

	httpreq, err := http.NewRequest(http.MethodGet, u.String(), nil)
	httpres, err := t.h.Do(httpreq)
	defer func(httpres *http.Response) {
		if httpres != nil && httpres.Body != nil {
			httpres.Body.Close()
		}
	}(httpres)

	if err != nil {
		return nil, xerrors.Errorf("http Do(): %w", err)
	}

	byteArray, err := ioutil.ReadAll(httpres.Body)
	if err != nil {
		return nil, xerrors.Errorf("ioutil.ReadAll(): %w", err)
	}

	trivia := &response.Trivia{}
	if err := json.Unmarshal(byteArray, &trivia); err != nil {
		return nil, xerrors.Errorf("json.Unmarshal(): %w", err)
	}

	if trivia.ResponseCode != 0 {
		return nil, xerrors.Errorf("trivia return response_code(%v)", trivia.ResponseCode)
	}

	return trivia.MakeTrivia(), nil
}
