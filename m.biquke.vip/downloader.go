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

var trimStringList = []string{
	"『加入书签，方便阅读』",
	"天才一秒记住本站地址：[",
	"笔趣阁]",
	"https://m.biquke.vip最快更新！无广告！",
}

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

	// 获取页面内容和下一页地址
	pageContent := func(url string) ([]string, string, error) {
		doc, err := d.getPageContent(url)
		if nil != err {
			return []string{}, "", err
		}

		lines := doc.Find("#chaptercontent p").Map(func(_ int, sel *goquery.Selection) string {
			return sel.Text()
		})

		nextURL, _ := doc.Find("#pb_next").First().Attr("href")
		return lines, nextURL, nil
	}

	// 判断下一页是否是同一章，如果是就合并为一章内容进行返回
	nextUrl := url
	pageLines := []string{}
	for strings.HasPrefix(strings.TrimRight(d.fullUrl(nextUrl), ".html"), strings.TrimRight(d.fullUrl(url), ".html")) {
		lines, pageNextUrl, err := pageContent(nextUrl)
		if nil != err {
			return "", err
		}
		pageLines = append(pageLines, lines...)
		nextUrl = pageNextUrl
	}

	// 删除需要过滤的信息
	chapterContent := strings.Join(pageLines, "\n")
	for _, item := range trimStringList {
		chapterContent = strings.ReplaceAll(chapterContent, item, "")
	}

	return chapterContent, nil
}

func (d Downloader) fullUrl(url string) string {
	if !strings.HasPrefix(url, "https://m.biquke.vip/") {
		url = "https://m.biquke.vip" + url
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
