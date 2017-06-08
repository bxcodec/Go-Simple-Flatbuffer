package http_test

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	msgpack "gopkg.in/vmihailenco/msgpack.v2"

	// "github.com/bxcodec/Go-Simple-Flatbuffer/users"
	// "github.com/bxcodec/Go-Simple-Flatbuffer/users/delivery/http/jsondlv"
	fbsdlv "github.com/bxcodec/Go-Simple-Flatbuffer/articles/delivery/http/fbs"
	jsondlv "github.com/bxcodec/Go-Simple-Flatbuffer/articles/delivery/http/json"
	msgpackdlv "github.com/bxcodec/Go-Simple-Flatbuffer/articles/delivery/http/msgpack"
	"github.com/bxcodec/Go-Simple-Flatbuffer/users"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

const benchCall = 100000

func TestRESTwithFBS(bench *testing.T) {
	e := echo.New()

	req, _ := http.NewRequest(echo.GET, `/articlefbs`, strings.NewReader(``))

	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	handler := fbsdlv.HttpHandlerFbs{}

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	handler.Get(c)

	if rec.Code != http.StatusOK {
		os.Exit(0)
	}
	body := rec.Body
	dta, err := ioutil.ReadAll(body)
	assert.NoError(bench, err)

	data := handler.ReadArticle(dta)

	assert.Equal(bench, mockArticle.Categories, data.Categories)

}
func BenchmarkRESTwithFBS(bench *testing.B) {
	bench.N = benchCall
	for i := 0; i < bench.N; i++ {
		e := echo.New()

		req, _ := http.NewRequest(echo.GET, `/articlefbs`, strings.NewReader(``))

		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		handler := fbsdlv.HttpHandlerFbs{}

		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		handler.Get(c)

		if rec.Code != http.StatusOK {
			os.Exit(0)
		}
		body := rec.Body
		dta, err := ioutil.ReadAll(body)
		assert.NoError(bench, err)

		data := handler.ReadArticle(dta)

		assert.Equal(bench, mockArticle.Categories, data.Categories)
	}
}

func TestRESTwithMsgPack(bench *testing.T) {
	e := echo.New()

	req, _ := http.NewRequest(echo.GET, `/articlemsgpack`, strings.NewReader(``))

	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	handler := msgpackdlv.HttpHandlerMsgPack{}

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	handler.Get(c)

	if rec.Code != http.StatusOK {
		os.Exit(0)
	}
	body := rec.Body
	dta, err := ioutil.ReadAll(body)
	assert.NoError(bench, err)

	var a ClientArticleObj
	err = msgpack.Unmarshal(dta, &a)
	assert.NoError(bench, err)
	assert.Equal(bench, mockArticle.Categories, a.Categories)

}
func BenchmarkRESTwithMsgPack(bench *testing.B) {
	bench.N = benchCall
	e := echo.New()
	for i := 0; i < bench.N; i++ {

		req, _ := http.NewRequest(echo.GET, `/articlemsgpack`, strings.NewReader(``))

		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		handler := msgpackdlv.HttpHandlerMsgPack{}

		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		handler.Get(c)

		if rec.Code != http.StatusOK {
			os.Exit(0)
		}
		body := rec.Body
		dta, err := ioutil.ReadAll(body)
		assert.NoError(bench, err)

		var a ClientArticleObj
		err = msgpack.Unmarshal(dta, &a)
		assert.NoError(bench, err)
		assert.Equal(bench, mockArticle.Categories, a.Categories)

	}
}

func TestRESTwithJSON(bench *testing.T) {
	e := echo.New()

	req, _ := http.NewRequest(echo.GET, `/article`, strings.NewReader(``))

	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	handler := jsondlv.ArticleHandler{}

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	handler.Get(c)

	if rec.Code != http.StatusOK {
		os.Exit(0)
	}
	body := rec.Body
	dta, err := ioutil.ReadAll(body)
	assert.NoError(bench, err)
	var user users.UserObj
	err = json.Unmarshal(dta, &user)
	assert.NoError(bench, err)

}
func BenchmarkRESTwithJSON(bench *testing.B) {
	bench.N = benchCall
	e := echo.New()
	for i := 0; i < bench.N; i++ {

		req, _ := http.NewRequest(echo.GET, `/article`, strings.NewReader(``))

		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		handler := jsondlv.ArticleHandler{}

		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		handler.Get(c)

		if rec.Code != http.StatusOK {
			os.Exit(0)
		}
		body := rec.Body
		dta, err := ioutil.ReadAll(body)
		assert.NoError(bench, err)

		var user users.UserObj
		err = json.Unmarshal(dta, &user)
		assert.NoError(bench, err)

	}

}
