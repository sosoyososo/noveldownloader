package xianqihaotianmiorg

import (
	"errors"
	"io"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/djimenez/iconv-go"
)

const pageEncodeType = "utf-8"

var trimStringList = []string{}

type Downloader struct{}

func (d Downloader) ChapterURLList(entryURL string) ([]string, error) {
	chapterURLs := []string{}
	for len(entryURL) > 0 {
		doc, err := d.getPageContent(entryURL)
		if nil != err {
			return []string{}, err
		}
		chapters := doc.Find(".list-charts li a").Map(func(i int, sel *goquery.Selection) string {
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

		lines := doc.Find(".content-body").Map(func(_ int, sel *goquery.Selection) string {
			return sel.Text()
		})
		if len(lines) <= 0 {
			log.Panic("no content")
		}

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

	if len(chapterContent) <= 0 {
		log.Panicf("download chapter content empty")
	}

	return chapterContent, nil
}

func (d Downloader) fullUrl(url string) string {
	if !strings.HasPrefix(url, "http://www.xianqihaotianmi.org") {
		url = "http://www.xianqihaotianmi.org" + url
	}
	return url
}

func (d Downloader) getPageContent(url string) (*goquery.Document, error) {
	sendReq := func() (*http.Response, error) {
		url = d.fullUrl(url)
		log.Printf("load page %v", url)

		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			log.Fatalln(err)
		}
		req.Header.Set("User-Agent", "Mozilla/5.0")
		client := &http.Client{}
		return client.Do(req)
	}

	resp, err := sendReq()
	if nil != err {
		return nil, err
	}
	if resp.StatusCode > 299 {
		time.Sleep(2 * time.Second)
		resp, err = sendReq()
		if nil != err {
			return nil, err
		}
		if resp.StatusCode > 299 {
			return nil, errors.New("http status code is not 200")
		}
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
