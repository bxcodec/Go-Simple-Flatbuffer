package main

import (
	"github.com/bxcodec/Go-Simple-Flatbuffer/users/delivery/http/fbs"
	"github.com/bxcodec/Go-Simple-Flatbuffer/users/delivery/http/jsondlv"
	"github.com/labstack/echo"
)

func main() {

	e := echo.New()

	jsondlv.NewJSONDelivery(e)
	fbs.NewFBSDelivery(e)
	e.Start(":3000")

}
