package injector

import (
	"github.com/labstack/echo"
	"github.com/s-kmmr/viewtrivia/domain/repository"
	"github.com/s-kmmr/viewtrivia/handler"
	"github.com/s-kmmr/viewtrivia/infrastructure/api/trivia"
	"github.com/s-kmmr/viewtrivia/usecase"
)

// Injector Injector
type Injector struct {
	trivia handler.Trivia
}

// TriviaHandler Getter(handler.Trivia)
func (i *Injector) TriviaHandler() handler.Trivia {
	return i.trivia
}

// NewInjector Injectorコンストラクタ
func NewInjector(e *echo.Echo) {
	i := Injector{}
	i.trivia = i.newTriviaHandler()

	InitRouting(e, i)
	e.Validator = NewValidator()

}

func (i *Injector) newTriviaHandler() handler.Trivia {
	return handler.NewTrivia(i.newTriviaUseCase())
}

func (i *Injector) newTriviaUseCase() usecase.Trivia {
	return usecase.NewTrivia(i.newTriviaRepository())
}

func (i *Injector) newTriviaRepository() repository.Trivia {
	return trivia.NewTrivia()
}
