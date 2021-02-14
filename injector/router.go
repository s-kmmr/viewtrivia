package injector

import (
	"github.com/labstack/echo"
)

// InitRouting リクエストルーティング設定
func InitRouting(e *echo.Echo, i Injector) {
	e.GET("/", i.TriviaHandler().ShowTrivia)
}
