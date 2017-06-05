package http_test

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/bxcodec/Go-Simple-Flatbuffer/users/delivery/http/fbs"
	"github.com/bxcodec/Go-Simple-Flatbuffer/users/delivery/http/jsondlv"
	"github.com/labstack/echo"
)

func BenchmarkJSONSimple(b *testing.B) {

	e := echo.New()
	b.N = 100000

	for i := 0; i < b.N; i++ {

		req, _ := http.NewRequest(echo.GET, `/user`, strings.NewReader(``))

		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		handler := jsondlv.UserHandler{}

		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		handler.Get(c)

		if rec.Code != http.StatusOK {
			os.Exit(0)
		}
		body := rec.Body
		dta, err := ioutil.ReadAll(body)
		if err != nil {
			os.Exit(0)
		}

		if i == 1 {
			fmt.Println(string(dta))
			fmt.Print("")
		}

	}

}

func BenchmarkFbsSimple(b *testing.B) {

	e := echo.New()
	b.N = 100000

	for i := 0; i < b.N; i++ {

		req, _ := http.NewRequest(echo.GET, `/userfbs`, strings.NewReader(``))

		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		handler := fbs.HttpHandlerFbs{}

		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		handler.Get(c)

		if rec.Code != http.StatusOK {
			os.Exit(0)
		}

		body := rec.Body
		dta, err := ioutil.ReadAll(body)

		if err != nil {
			fmt.Println("Something happen")

			os.Exit(0)
		}

		if i == 1 {
			data := handler.ReadUser(dta)
			fmt.Println(data)
			fmt.Print("")
		}

	}

}

func BenchmarkJSONSimpleList(b *testing.B) {

	e := echo.New()
	b.N = 100000

	for i := 0; i < b.N; i++ {

		req, _ := http.NewRequest(echo.GET, `/userlist`, strings.NewReader(``))

		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		handler := jsondlv.UserHandler{}

		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		handler.Get(c)

		if rec.Code != http.StatusOK {
			os.Exit(0)
		}
		body := rec.Body
		_, err := ioutil.ReadAll(body)
		if err != nil {
			os.Exit(0)
		}

		if i == 1 {
			// fmt.Println(string(dta))
			fmt.Print("")
		}

	}

}

func BenchmarkFbsSimpleList(b *testing.B) {

	e := echo.New()
	b.N = 100000

	for i := 0; i < b.N; i++ {

		req, _ := http.NewRequest(echo.GET, `/userfbslist`, strings.NewReader(``))

		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		handler := fbs.HttpHandlerFbs{}

		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		handler.Get(c)

		if rec.Code != http.StatusOK {
			os.Exit(0)
		}

		body := rec.Body
		dta, err := ioutil.ReadAll(body)

		if err != nil {
			fmt.Println("Something happen")

			os.Exit(0)
		}

		if i == 1 {
			handler.ReadUser(dta)
			// fmt.Printf("Name : %s , id : %d ", name, id)
			fmt.Print("")
		}

	}

}
