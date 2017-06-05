package http_test

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/bxcodec/Go-Simple-Flatbuffer/users"
	httpdlv "github.com/bxcodec/Go-Simple-Flatbuffer/users/delivery/http"
	"github.com/bxcodec/Go-Simple-Flatbuffer/users/delivery/http/fbs"
	"github.com/bxcodec/Go-Simple-Flatbuffer/users/delivery/http/jsondlv"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

const benchCall = 100000

func BenchmarkJSONSimple(b *testing.B) {

	e := echo.New()
	b.N = benchCall

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
		assert.NoError(b, err)

		var user users.UserObj
		err = json.Unmarshal(dta, &user)
		assert.NoError(b, err)
		assert.Equal(b, "Iman", user.Name)

	}

}

func BenchmarkFbsSimple(b *testing.B) {

	e := echo.New()
	b.N = benchCall

	for i := 0; i < b.N; i++ {

		req, _ := http.NewRequest(echo.GET, `/userfbs`, strings.NewReader(``))

		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		handler := fbs.HttpHandlerFbs{}

		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		handler.Get(c)

		assert.Equal(b, http.StatusOK, rec.Code)

		body := rec.Body
		dta, err := ioutil.ReadAll(body)
		assert.NoError(b, err)

		data := handler.ReadUser(dta)
		assert.Equal(b, int64(64), data.ID)

	}

}

func BenchmarkJSONSimpleList(b *testing.B) {

	e := echo.New()
	b.N = benchCall

	for i := 0; i < b.N; i++ {

		req, _ := http.NewRequest(echo.GET, `/userlist`, strings.NewReader(``))

		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		handler := jsondlv.UserHandler{}

		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		handler.GetListUser(c)
		assert.Equal(b, http.StatusOK, rec.Code)
		body := rec.Body
		dta, err := ioutil.ReadAll(body)
		assert.NoError(b, err)
		var listUser []*users.UserObj
		err = json.Unmarshal(dta, &listUser)
		assert.NoError(b, err)

		assert.Len(b, listUser, httpdlv.DATA_SIZE)
		// assert.Equal(b, "Iman", listUser[0].Name)

	}

}

func BenchmarkFbsSimpleList(b *testing.B) {

	e := echo.New()
	b.N = benchCall

	for i := 0; i < b.N; i++ {

		req, _ := http.NewRequest(echo.GET, `/userfbslist`, strings.NewReader(``))

		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		handler := fbs.HttpHandlerFbs{}

		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		handler.GetListUser(c)

		assert.Equal(b, http.StatusOK, rec.Code)
		body := rec.Body
		dta, err := ioutil.ReadAll(body)
		assert.NoError(b, err)

		listUser := handler.ReadUserList(dta)

		assert.Len(b, listUser, httpdlv.DATA_SIZE)
		assert.Equal(b, "Iman", listUser[0].Name)

	}

}
