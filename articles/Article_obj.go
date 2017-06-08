package articles

import "time"

type ArticleObj struct {
	ID         int64
	Title      string
	Excerpt    string
	Html       string
	Content    []ContentObj
	Url        string
	Images     []*ImageObj
	PubishTime time.Time
	Categories []int
	SourceID   int
}

type ContentObj interface {
	Type() string
}
type ImageObj struct {
	URL    string
	Mime   string
	Width  int
	Height int
}
type ImageContentObj struct {
	ImageObj
	TypeContent string
}

func (i ImageContentObj) Type() string {
	return i.TypeContent
}

type ParagraphObj struct {
	Text        string
	TypeContent string
}

func (p ParagraphObj) Type() string {
	return p.TypeContent
}
