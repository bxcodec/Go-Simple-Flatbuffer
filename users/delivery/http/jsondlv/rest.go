package jsondlv

import (
	"net/http"

	"github.com/bxcodec/Go-Simple-Flatbuffer/users"
	httpdlv "github.com/bxcodec/Go-Simple-Flatbuffer/users/delivery/http"
	"github.com/labstack/echo"
)

type UserHandler struct {
}

func (u *UserHandler) Get(c echo.Context) error {

	user := &users.UserObj{int64(2), "Iman"}

	return c.JSON(http.StatusOK, user)
}

func (u *UserHandler) GetListUser(c echo.Context) error {

	user := &users.UserObj{int64(2), "Iman"}
	list := make([]*users.UserObj, 0)
	for i := 0; i < httpdlv.DATA_SIZE; i++ {
		list = append(list, user)
	}
	return c.JSON(http.StatusOK, list)
}
func NewJSONDelivery(e *echo.Echo) {

	handler := &UserHandler{}

	e.GET("/user", handler.Get)
	e.GET("/userlist", handler.GetListUser)

}
