package fbs_test

import (
	"encoding/json"
	"testing"
	"time"

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
	// jsonArticle := byt
	var articleD ClientArticleObj
	err := json.Unmarshal([]byte(jsonArticle), &articleD)
	assert.NoError(t, err)
	assert.Equal(t, mockArticle.Categories, articleD.Categories)
	assert.Equal(t, mockArticle.Images, articleD.Images)
}
func TestFbsMarshall(t *testing.T) {
	builder := flatbuffers.NewBuilder(0)
	handler := &fbs.ArticleFbsHandler{}

	buf := handler.MakeArticle(builder, mockArticle)

	a := handler.ReadArticle(buf)
	assert.Equal(t, mockArticle.Categories, a.Categories)
	assert.Equal(t, mockArticle.Images, a.Images)
	assert.Equal(t, mockArticle.Content, a.Content)

}

func BenchmarkCreateArticleJSON(b *testing.B) {
	b.N = BenchCall
	for i := 0; i < b.N; i++ {

		byt, _ := json.Marshal(mockArticle)
		jsonArticle := string(byt)
		// jsonArticle := byt
		var articleD ClientArticleObj
		err := json.Unmarshal([]byte(jsonArticle), &articleD)
		if !assert.NoError(b, err) {
			break
		}
		assert.Equal(b, mockArticle.Categories, articleD.Categories)
		assert.Equal(b, mockArticle.Images, articleD.Images)
	}
}

func BenchmarkCreateArticleFBS(b *testing.B) {
	b.N = BenchCall

	for i := 0; i < b.N; i++ {
		builder := flatbuffers.NewBuilder(0)
		handler := &fbs.ArticleFbsHandler{}

		buf := handler.MakeArticle(builder, mockArticle)
		a := handler.ReadArticle(buf)
		assert.Equal(b, mockArticle.Categories, a.Categories)
		assert.Equal(b, mockArticle.Images, a.Images)
		if !assert.Equal(b, mockArticle.Content, a.Content) {
			break
		}

	}

}
