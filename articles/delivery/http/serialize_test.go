package http_test

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"testing"
	"time"

	msgpack "gopkg.in/vmihailenco/msgpack.v2"

	"github.com/bxcodec/Go-Simple-Flatbuffer/articles"
	"github.com/bxcodec/Go-Simple-Flatbuffer/articles/delivery/http/fbs"
	"github.com/stretchr/testify/assert"

	flatbuffers "github.com/google/flatbuffers/go"
)

const BenchCall = 100000

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

type ClientArticleObj struct {
	ID         int64
	Title      string
	Excerpt    string
	Html       string
	Content    []interface{}
	Url        string
	Images     []*articles.ImageObj
	PubishTime time.Time
	Categories []int
	SourceID   int
}

func TestJsonMarshall(t *testing.T) {

	byt, _ := json.Marshal(mockArticle)
	jsonArticle := string(byt)
	// fmt.Println(jsonArticle)
	// jsonArticle := byt
	var articleD ClientArticleObj
	err := json.Unmarshal([]byte(jsonArticle), &articleD)
	size := binary.Size(byt)
	fmt.Println("JSON: ", size)
	assert.NoError(t, err)
	assert.Equal(t, mockArticle.Categories, articleD.Categories)
	assert.Equal(t, mockArticle.Images, articleD.Images)
}
func TestFbsMarshall(t *testing.T) {
	builder := flatbuffers.NewBuilder(0)
	handler := &fbs.ArticleFbsHandler{}

	buf := handler.MakeArticle(builder, mockArticle)

	size := binary.Size(buf)
	fmt.Println("FBS: ", size)
	a := handler.ReadArticle(buf)
	assert.Equal(t, mockArticle.Categories, a.Categories)
	assert.Equal(t, mockArticle.Images, a.Images)
	assert.Equal(t, mockArticle.Content, a.Content)

}

func TestMsgPackMarshall(t *testing.T) {
	client := &ClientArticleObj{
		ID:         mockArticle.ID,
		Title:      mockArticle.Title,
		Excerpt:    mockArticle.Excerpt,
		Html:       mockArticle.Html,
		Url:        mockArticle.Url,
		Images:     mockArticle.Images,
		PubishTime: mockArticle.PubishTime,
		Categories: mockArticle.Categories,
		SourceID:   mockArticle.SourceID,
	}

	for _, v := range mockArticle.Content {
		client.Content = append(client.Content, v)
	}

	bench, err := msgpack.Marshal(mockArticle)
	assert.NoError(t, err)

	size := binary.Size(bench)
	fmt.Println("MsgPack: ", size)
	var a ClientArticleObj
	err = msgpack.Unmarshal(bench, &a)
	assert.NoError(t, err)
	assert.Equal(t, mockArticle.Categories, a.Categories)

}
func BenchmarkSerializeWithFBS(bench *testing.B) {
	bench.N = BenchCall

	for i := 0; i < bench.N; i++ {
		builder := flatbuffers.NewBuilder(0)
		handler := &fbs.ArticleFbsHandler{}

		buf := handler.MakeArticle(builder, mockArticle)
		a := handler.ReadArticle(buf)
		assert.Equal(bench, mockArticle.Categories, a.Categories)
		assert.Equal(bench, mockArticle.Images, a.Images)
		if !assert.Equal(bench, mockArticle.Content, a.Content) {
			break
		}

	}

}

func BenchmarkSerializeWithMsgPakc(bench *testing.B) {
	bench.N = benchCall
	for i := 0; i < bench.N; i++ {
		b, err := msgpack.Marshal(mockArticle)
		assert.NoError(bench, err)

		var a ClientArticleObj
		err = msgpack.Unmarshal(b, &a)
		assert.NoError(bench, err)
		assert.Equal(bench, mockArticle.Categories, a.Categories)

	}

}

func BenchmarkSerializeWithJSON(bench *testing.B) {
	bench.N = BenchCall
	for i := 0; i < bench.N; i++ {

		byt, _ := json.Marshal(mockArticle)
		jsonArticle := string(byt)
		// jsonArticle := byt
		var articleD ClientArticleObj
		err := json.Unmarshal([]byte(jsonArticle), &articleD)
		if !assert.NoError(bench, err) {
			break
		}
		assert.Equal(bench, mockArticle.Categories, articleD.Categories)
		assert.Equal(bench, mockArticle.Images, articleD.Images)
	}
}
