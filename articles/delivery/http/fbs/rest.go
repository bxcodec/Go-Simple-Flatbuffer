package fbs

import (
	"encoding/binary"
	"net/http"
	"strconv"
	"time"

	"github.com/bxcodec/Go-Simple-Flatbuffer/articles"
	flatbuffers "github.com/google/flatbuffers/go"
	"github.com/labstack/echo"
)

type HttpHandlerFbs struct {
	*ArticleFbsHandler
}

var (
	contentObjs = []articles.ContentObj{
		&articles.ParagraphObj{"Kurio will stand over the world", "text"},
		&articles.ImageContentObj{
			articles.ImageObj{
				URL:    "http://kurio.com/images/winner.png",
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

	mockArticle = &articles.ArticleObj{
		ID:         int64(63),
		Title:      "Kurio On Fire",
		Excerpt:    "We will won, whatever happen for us",
		Html:       "<h3> Kurio On Fire </h3> <br> <p> We will Won, Whatever Happer for us </p>",
		Content:    contentObjs,
		Url:        "http://kurio.com/article/23",
		Images:     images,
		PubishTime: time.Now(),
		Categories: []int{2, 3, 4},
		SourceID:   23,
	}
)

func (u *HttpHandlerFbs) Get(c echo.Context) error {

	b := flatbuffers.NewBuilder(0)
	buf := u.MakeArticle(b, mockArticle)

	c.Response().Header().Set(echo.HeaderContentType, echo.MIMEOctetStream)
	c.Response().Header().Set(echo.HeaderContentLength, strconv.Itoa(binary.Size(buf)))
	c.Response().WriteHeader(http.StatusOK)

	c.Response().Write(buf)

	c.Response().Flush()

	return nil
}

// func (u *HttpHandlerFbs) GetListUser(c echo.Context) error {
//
// 	b := flatbuffers.NewBuilder(0)
// 	user := &users.UserObj{int64(42), "Iman"}
// 	list := make([]*users.UserObj, 0)
// 	for i := 0; i < httpdlv.DATA_SIZE; i++ {
// 		list = append(list, user)
// 	}
//
// 	buf := u.MakeListUser(b, list)
// 	size := binary.Size(buf)
// 	// fmt.Println(size)
// 	c.Response().Header().Set(echo.HeaderContentType, echo.MIMEOctetStream)
// 	c.Response().Header().Set(echo.HeaderContentLength, strconv.Itoa(size))
// 	c.Response().WriteHeader(http.StatusOK)
// 	c.Response().Write(buf)
//
// 	c.Response().Flush()
//
// 	return nil
// }

func NewFBSDelivery(e *echo.Echo) {
	handler := &HttpHandlerFbs{}
	e.GET("/articlefbs", handler.Get)

	// e.GET("/userfbslist", handler.GetListUser)
}
