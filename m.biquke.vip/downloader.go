package mbiqukevip

import (
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/djimenez/iconv-go"
)

const pageEncodeType = "utf-8"

type Downloader struct{}

func (d Downloader) ChapterURLList(entryURL string) ([]string, error) {
	chapterURLs := []string{}
	for len(entryURL) > 0 {
		doc, err := d.getPageContent(entryURL)
		if nil != err {
			return []string{}, err
		}
		chapters := doc.Find("li a").Map(func(i int, sel *goquery.Selection) string {
			str, _ := sel.Attr("href")
			return str
		})
		chapterURLs = append(chapterURLs, chapters...)
		nextPage, _ := doc.Find(".listpage .right a").First().Attr("href")
		entryURL = nextPage
	}

	return chapterURLs, nil
}

func (d Downloader) ChapterDetail(url string) (string, error) {
	doc, err := d.getPageContent(url)
	if nil != err {
		return "", err
	}

	lines := doc.Find("#chaptercontent p").Map(func(_ int, sel *goquery.Selection) string {
		return sel.Text()
	})
	return strings.Join(lines, "\n"), nil
}

func (d Downloader) getPageContent(url string) (*goquery.Document, error) {
	if !strings.HasPrefix(url, "https://m.biquke.vip/") {
		url = "https://m.biquke.vip" + url
	}
	log.Printf("load page %v", url)
	resp, err := http.Get(url)
	if nil != err {
		return nil, err
	}
	defer resp.Body.Close()

	var r io.Reader = resp.Body
	if pageEncodeType != "utf-8" {
		utfBody, err := iconv.NewReader(resp.Body, pageEncodeType, "utf-8")
		if err != nil {
			return nil, err
		}
		r = utfBody
	}

	return goquery.NewDocumentFromReader(r)
}
