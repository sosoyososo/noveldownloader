package bi24cc

import (
	"log"
	"testing"
)

func TestDownloadChapters(t *testing.T) {
	urls, err := Downloader{}.ChapterURLList("https://www.bi24.cc/book/163233/")
	if nil != err {
		t.Error(err)
	}
	log.Println(urls)
}

func TestDownloadChapterContent(t *testing.T) {
	content, err := Downloader{}.ChapterDetail("https://www.bi24.cc/book/163233/1.html")
	if nil != err {
		t.Error(err)
	}
	log.Println(content)
}
