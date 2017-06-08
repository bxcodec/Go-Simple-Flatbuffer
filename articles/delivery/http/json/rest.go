package json

import (
	"net/http"
	"time"

	"github.com/bxcodec/Go-Simple-Flatbuffer/articles"
	httpdlv "github.com/bxcodec/Go-Simple-Flatbuffer/articles/delivery/http"
	"github.com/labstack/echo"
)

type ArticleHandler struct {
}

var (
	contentObjs = []articles.ContentObj{
		&articles.ParagraphObj{"Kurio will stand over the world", "text"},
		&articles.ImageContentObj{
			articles.ImageObj{
				URL:    "http://kurio.com/images/gambar.png",
				Mime:   "png",
				Width:  200,
				Height: 200,
			},
			"image",
		},
	}
	images = []*articles.ImageObj{
		&articles.ImageObj{
			URL:    "http://kurio.com/images/winner.png",
			Mime:   "png",
			Width:  200,
			Height: 200,
		},
	}
	now         = time.Now()
	mockArticle = &articles.ArticleObj{
		ID:         int64(63),
		Title:      "Kurio On Fire",
		Excerpt:    "We will won, whatever happen for us",
		Html:       "<h3> Kurio On Fire </h3> <br> <p> We will Won, Whatever Happer for us </p>",
		Content:    contentObjs,
		Url:        "http://kurio.com/article/23",
		Images:     images,
		PubishTime: now,
		Categories: []int{2, 3, 4},
		SourceID:   23,
	}
)

func (u *ArticleHandler) Get(c echo.Context) error {

	return c.JSON(http.StatusOK, mockArticle)
}

func (u *ArticleHandler) GetListArticle(c echo.Context) error {
	list := make([]*articles.ArticleObj, 0)
	for i := 0; i < httpdlv.DATA_SIZE; i++ {
		list = append(list, mockArticle)
	}
	return c.JSON(http.StatusOK, list)
}
func NewArticleJSONDelivery(e *echo.Echo) {

	handler := &ArticleHandler{}

	e.GET("/article", handler.Get)
	// e.GET("/Articlelist", handler.GetListArticle)

}
