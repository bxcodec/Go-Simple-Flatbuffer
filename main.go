package main

import (
	articleFbs "github.com/bxcodec/Go-Simple-Flatbuffer/articles/delivery/http/fbs"
	userFbs "github.com/bxcodec/Go-Simple-Flatbuffer/users/delivery/http/fbs"
	userJson "github.com/bxcodec/Go-Simple-Flatbuffer/users/delivery/http/jsondlv"
	"github.com/labstack/echo"
)

func main() {

	e := echo.New()
	articleFbs.NewFBSDelivery(e)
	userJson.NewJSONDelivery(e)
	userFbs.NewFBSDelivery(e)
	e.Start(":3000")

}
