package main

import (
	"flag"
	"fmt"
	"log"
	bi24cc "noveldownload/bi24.cc"
	mbiqukevip "noveldownload/m.biquke.vip"
	xianqihaotianmiorg "noveldownload/xianqihaotianmi.org"
	"os"
)

type Downloader interface {
	ChapterURLList(entryURL string) ([]string, error)
	ChapterDetail(url string) (string, error)
}

var (
	downloaderList = map[string]Downloader{
		"m.biquke.vip":        mbiqukevip.Downloader{},
		"bi24.cc":             bi24cc.Downloader{},
		"xianqihaotianmi.org": xianqihaotianmiorg.Downloader{},
	}
)

func main() {
	var entryUrl string
	var downloaderType string
	var skipCount int
	flag.StringVar(&downloaderType, "downloader", "", "下载器的类型")
	flag.StringVar(&entryUrl, "entry", "", "获取目录的入口地址")
	flag.IntVar(&skipCount, "skip", 0, "跳过的章节数")
	flag.Parse()

	if len(entryUrl) <= 0 {
		log.Panicf("请输入入口地址")
	}
	if len(downloaderType) <= 0 {
		log.Panicf("下载器类型")
	}

	downloader := downloaderList[downloaderType]
	if nil == downloader {
		log.Panicf("no such downloader")
	}
	chapters, err := downloader.ChapterURLList(entryUrl)
	if nil != err {
		log.Panicf("get chapter list failed %v", err)
	}

	for index, chapterURL := range chapters {
		if index <= skipCount {
			continue
		}
		content, err := downloader.ChapterDetail(chapterURL)
		if nil != err {
			log.Panicf("download chapter content fail %v", err)
		}
		if err := saveFile(content, fmt.Sprintf("./%v.txt", index+1)); nil != err {
			log.Panicf("save chapter content fail %v", err)
		}
	}
}

func saveFile(content, path string) error {
	f, err := os.Create(path)
	if nil != err {
		return err
	}
	defer f.Close()
	l, err := f.WriteString(content)
	if nil != err {
		return err
	}
	log.Printf("save file %v [%v byte]", path, l)
	return nil
}
