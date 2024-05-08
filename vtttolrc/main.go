package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/abhishekkr/gowebvtt"
)

func main() {
	var fileName string
	var lrcFile string
	flag.StringVar(&fileName, "vtt", "", "vtt file to convert")
	flag.StringVar(&lrcFile, "lrc", "", "lrc file")
	flag.Parse()
	if len(fileName) <= 0 {
		panic("no vtt file")
	}

	f, err := os.Create(lrcFile)
	if nil != err {
		log.Panicf("create lrc file fail %v", err)
	}
	defer f.Close()

	vttOpts := gowebvtt.VttOptions{Enabled: true, MaxLinesPerScene: 2}
	vtt, err := gowebvtt.ParseFileWithOptions(fileName, vttOpts)
	if err != nil {
		panic(err)
	}

	lines := []string{}
	for _, item := range vtt.Scenes {
		millSec := item.StartMilliSec % 1000
		sec := item.StartMilliSec / 1000 % 60
		min := item.StartMilliSec / 1000 / 60
		line := fmt.Sprintf("[%v:%v.%v]%v", min, sec, millSec, strings.Join(item.Transcript, " "))
		lines = append(lines, line)
	}
	lrcContent := strings.Join(lines, "\n")

	if _, err := f.WriteString(lrcContent); nil != err {
		log.Panicf("write lrc file fail %v", err)
	}
}
