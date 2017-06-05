package fbs

import (
	"net/http"

	"github.com/bxcodec/Go-Simple-Flatbuffer/users"
	httpdlv "github.com/bxcodec/Go-Simple-Flatbuffer/users/delivery/http"
	flatbuffers "github.com/google/flatbuffers/go"
	"github.com/labstack/echo"
)

type HttpHandlerFbs struct {
	*UserFbsHandler
}

func (u *HttpHandlerFbs) Get(c echo.Context) error {

	b := flatbuffers.NewBuilder(0)
	buf := u.MakeUser(b, &users.UserObj{int64(64), "Iman"})

	c.Response().Header().Set(echo.HeaderContentType, echo.MIMEOctetStream)
	c.Response().WriteHeader(http.StatusOK)
	c.Response().Write(buf)

	c.Response().Flush()

	return nil
}

func (u *HttpHandlerFbs) GetListUser(c echo.Context) error {

	b := flatbuffers.NewBuilder(0)
	user := &users.UserObj{int64(42), "Arthur Dent"}
	list := make([]*users.UserObj, 0)
	for i := 0; i < httpdlv.DATA_SIZE; i++ {
		list = append(list, user)
	}

	buf := u.MakeListUser(b, list)

	c.Response().Header().Set(echo.HeaderContentType, echo.MIMEOctetStream)
	c.Response().WriteHeader(http.StatusOK)
	c.Response().Write(buf)

	c.Response().Flush()

	return nil
}

func NewFBSDelivery(e *echo.Echo) {
	handler := &HttpHandlerFbs{}
	e.GET("/userfbs", handler.Get)

	e.GET("/userfbslist", handler.GetListUser)
}
