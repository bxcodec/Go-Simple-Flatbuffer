package fbs

import (
	"reflect"
	"time"

	"github.com/bxcodec/Go-Simple-Flatbuffer/articles"
	flatbuffers "github.com/google/flatbuffers/go"
)

type ArticleFbsHandler struct {
}

func (u *ArticleFbsHandler) imageList(b *flatbuffers.Builder, imglist []*articles.ImageObj) []flatbuffers.UOffsetT {
	ptrs := make([]flatbuffers.UOffsetT, len(imglist))
	for i := 0; i < len(imglist); i++ {
		url := b.CreateString(imglist[i].URL)
		mime := b.CreateString(imglist[i].Mime)
		articles.ImageStart(b)
		articles.ImageAddUrl(b, url)
		articles.ImageAddMime(b, mime)
		articles.ImageAddHeight(b, int32(imglist[i].Height))
		articles.ImageAddWidth(b, int32(imglist[i].Width))
		image := articles.ImageEnd(b)
		ptrs[i] = image
	}
	return ptrs
}

func (u *ArticleFbsHandler) contentList(b *flatbuffers.Builder, contents []articles.ContentObj) []flatbuffers.UOffsetT {

	// articles.ArticleStartContentVector(b, len(contents))
	ptrs := make([]flatbuffers.UOffsetT, len(contents))

	for i := len(contents) - 1; i >= 0; i-- {

		contentType := b.CreateString(contents[i].Type())
		articles.ContentStart(b)
		articles.ContentAddType(b, contentType)
		baseContent := articles.ContentEnd(b)

		switch reflect.TypeOf(contents[i]).String() {
		case "articles.ImageContentObj", "*articles.ImageContentObj":

			imgContentObj := contents[i].(*articles.ImageContentObj)

			url := b.CreateString(imgContentObj.URL)
			mime := b.CreateString(imgContentObj.Mime)
			// tipe := b.CreateString(imgContentObj.Type())

			articles.ImageContentStart(b)

			articles.ImageContentAddUrl(b, url)
			articles.ImageContentAddMime(b, mime)
			articles.ImageContentAddHeight(b, int32(imgContentObj.Height))
			articles.ImageContentAddWidth(b, int32(imgContentObj.Width))
			articles.ImageContentAddBase(b, baseContent)
			image := articles.ImageContentEnd(b)
			articles.ContentContainerStart(b)
			articles.ContentContainerAddContentsType(b, articles.ContentUnionImageContent)
			articles.ContentContainerAddContents(b, image)
			imageContent := articles.ContentContainerEnd(b)
			ptrs[i] = imageContent
			break
		case "articles.ParagraphObj", "*articles.ParagraphObj":
			// articles.ContentAddType(b, articles.ContentUnionParagraph)
			prgraphContentObj := contents[i].(*articles.ParagraphObj)

			text := b.CreateString(prgraphContentObj.Text)

			// tipe := b.CreateString(prgraphContentObj.Type())

			articles.ParagraphStart(b)
			articles.ParagraphAddText(b, text)
			articles.ParagraphAddBase(b, baseContent)

			paragraph := articles.ParagraphEnd(b)

			articles.ContentContainerStart(b)
			articles.ContentContainerAddContentsType(b, articles.ContentUnionParagraph)
			articles.ContentContainerAddContents(b, paragraph)
			paragraphContent := articles.ContentContainerEnd(b)
			ptrs[i] = paragraphContent
			break
		default:
			panic("Unknown Type of Reflection")
			break
		}

	}
	return ptrs
}

func (u *ArticleFbsHandler) MakeArticle(b *flatbuffers.Builder, articleObj *articles.ArticleObj) []byte {
	b.Reset()
	title := b.CreateString(articleObj.Title)
	excerpt := b.CreateString(articleObj.Excerpt)
	html := b.CreateString(articleObj.Html)
	url := b.CreateString(articleObj.Url)
	publishTime := b.CreateString(articleObj.PubishTime.Format(time.RFC3339))

	imglist := articleObj.Images
	ptrs := u.imageList(b, imglist)
	arrContent := u.contentList(b, articleObj.Content)

	// FROM HERE NO MORE CREATION MEMBER

	articles.ArticleStartImagesVector(b, len(imglist))
	for i := len(articleObj.Images) - 1; i >= 0; i-- {
		b.PrependUOffsetT(ptrs[i])
	}

	listImage := b.EndVector(len(imglist))
	// b.Finish(listImage)
	articles.ArticleStartCategoriesVector(b, len(articleObj.Categories))
	for i := len(articleObj.Categories) - 1; i >= 0; i-- {
		b.PrependInt64(int64(articleObj.Categories[i]))
		// b.PrependByte(byte())
	}
	catList := b.EndVector(len(articleObj.Categories))

	articles.ArticleStartContentVector(b, len(arrContent))
	for i := len(arrContent) - 1; i >= 0; i-- {
		b.PrependUOffsetT(arrContent[i])
	}
	contents := b.EndVector(len(arrContent))
	articles.ArticleStart(b)
	articles.ArticleAddId(b, articleObj.ID)
	articles.ArticleAddTitle(b, title)
	articles.ArticleAddExcerpt(b, excerpt)
	articles.ArticleAddHtml(b, html)

	articles.ArticleAddContent(b, contents)

	articles.ArticleAddUrl(b, url)

	articles.ArticleAddImages(b, listImage)

	articles.ArticleAddPublishTime(b, publishTime)

	articles.ArticleAddCategories(b, catList)

	articles.ArticleAddSourceId(b, int64(articleObj.SourceID))
	article := articles.ArticleEnd(b)
	b.Finish(article)

	return b.FinishedBytes()
}

func (u *ArticleFbsHandler) readImages(a *articles.Article) []*articles.ImageObj {
	listImg := make([]*articles.ImageObj, a.ImagesLength())
	img := new(articles.Image)
	for i := 0; i < a.ImagesLength(); i++ {
		if a.Images(img, i) {
			tempImgObj := &articles.ImageObj{
				URL:    string(img.Url()),
				Mime:   string(img.Mime()),
				Height: int(img.Height()),
				Width:  int(img.Width()),
			}
			listImg[i] = tempImgObj
		}
	}
	return listImg
}

func (u *ArticleFbsHandler) readCategories(a *articles.Article) []int {
	categoriesList := make([]int, a.CategoriesLength())
	for i := 0; i < a.CategoriesLength(); i++ {
		categoriesList[i] = int(a.Categories(i))

	}
	return categoriesList

}

func (u *ArticleFbsHandler) readContent(a *articles.Article) []articles.ContentObj {
	res := make([]articles.ContentObj, a.ContentLength())
	contentContainer := &articles.ContentContainer{}
	for i := 0; i < a.ContentLength(); i++ {
		if a.Content(contentContainer, i) {
			unionTable := new(flatbuffers.Table)
			if contentContainer.Contents(unionTable) {
				switch int(contentContainer.ContentsType()) {
				case articles.ContentUnionImageContent:
					img := &articles.ImageContent{}
					img.Init(unionTable.Bytes, unionTable.Pos)

					contentBase := &articles.Content{}
					contentBase = img.Base(contentBase)
					imgContentObj := &articles.ImageContentObj{
						articles.ImageObj{
							URL:    string(img.Url()),
							Mime:   string(img.Mime()),
							Height: int(img.Height()),
							Width:  int(img.Width()),
						},
						string(contentBase.Type()),
					}
					res[i] = imgContentObj

					break
				case articles.ContentUnionParagraph:
					paragraph := &articles.Paragraph{}
					paragraph.Init(unionTable.Bytes, unionTable.Pos)

					contentBase := &articles.Content{}
					contentBase = paragraph.Base(contentBase)
					p := &articles.ParagraphObj{
						Text:        string(paragraph.Text()),
						TypeContent: string(contentBase.Type()),
					}
					res[i] = p
					break
				default:
					panic("Unknown Content Type")

				}
			}
		}

	}

	return res
}

func (u *ArticleFbsHandler) ReadArticle(buf []byte) *articles.ArticleObj {
	articleRes := &articles.ArticleObj{}
	article := articles.GetRootAsArticle(buf, 0)
	articleRes.ID = article.Id()
	articleRes.Title = string(article.Title())
	articleRes.Excerpt = string(article.Excerpt())
	articleRes.Html = string(article.Html())
	articleRes.Url = string(article.Url())
	articleRes.SourceID = int(article.SourceId())

	images := u.readImages(article)
	articleRes.Images = images

	pubTime := string(article.PublishTime())

	var err error
	articleRes.PubishTime, err = time.Parse(time.RFC3339, pubTime)
	if err != nil {
		panic(err)

	}
	articleRes.Content = u.readContent(article)
	articleRes.Categories = u.readCategories(article)
	return articleRes
}
