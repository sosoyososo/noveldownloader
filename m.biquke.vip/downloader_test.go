package mbiqukevip

import (
	"log"
	"testing"
)

func TestDownloadChapters(t *testing.T) {
	urls, err := Downloader{}.ChapterURLList("https://m.biquke.vip/chapters/23557?sort=1&page=1")
	if nil != err {
		t.Error(err)
	}
	t.Log(urls)
}

func TestDownloadChapterContent(t *testing.T) {
	content, err := Downloader{}.ChapterDetail("https://m.biquke.vip/book/23557/19491694.html")
	if nil != err {
		t.Error(err)
	}
	log.Println(content)
}
