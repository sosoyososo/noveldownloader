package bi24cc

import (
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/djimenez/iconv-go"
)

const pageEncodeType = "utf-8"

var trimStringList = []string{
	"『点此报错』『加入书签』",
	"请收藏本站：https://www.bi24.cc。笔趣阁手机版：https://m.bi24.cc",
}

type Downloader struct{}

func (d Downloader) ChapterURLList(entryURL string) ([]string, error) {
	doc, err := d.getPageContent(entryURL)
	if nil != err {
		return []string{}, err
	}
	chapterURLs := doc.Find(".listmain dd a").Map(func(i int, sel *goquery.Selection) string {
		if rel, _ := sel.Attr("rel"); rel == "nofollow" {
			return ""
		}
		str, _ := sel.Attr("href")
		return str
	})

	retUrls := []string{}
	for _, url := range chapterURLs {
		if len(url) <= 0 {
			continue
		}
		retUrls = append(retUrls, url)
	}

	return retUrls, nil
}

func (d Downloader) ChapterDetail(url string) (string, error) {
	doc, err := d.getPageContent(url)
	if nil != err {
		return "", err
	}

	pageLines := doc.Find("#chaptercontent").Map(func(_ int, sel *goquery.Selection) string {
		return sel.Text()
	})

	// 删除需要过滤的信息
	chapterContent := strings.Join(pageLines, "\n")
	for _, item := range trimStringList {
		chapterContent = strings.ReplaceAll(chapterContent, item, "")
	}

	return chapterContent, nil
}

func (d Downloader) fullUrl(url string) string {
	if !strings.HasPrefix(url, "https://www.bi24.cc/") {
		url = "https://www.bi24.cc" + url
	}
	return url
}

func (d Downloader) getPageContent(url string) (*goquery.Document, error) {
	url = d.fullUrl(url)
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
