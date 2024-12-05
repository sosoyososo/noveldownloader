package xianqihaotianmiorg

import (
	"log"
	"testing"
)

func TestDownloadChapters(t *testing.T) {
	urls, err := Downloader{}.ChapterURLList("http://www.xianqihaotianmi.org/book/124728.html")
	if nil != err {
		t.Error(err)
	}
	t.Log(urls)
}

func TestDownloadChapterContent(t *testing.T) {
	content, err := Downloader{}.ChapterDetail("http://www.xianqihaotianmi.org/read/124728_52062430.html")
	if nil != err {
		t.Error(err)
	}
	log.Println(content)
}
