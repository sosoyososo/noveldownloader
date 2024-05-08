#小说下载
单本小说下载，每章单独存储为一个文件。

`text_to_audio.sh` 这个脚本可以在下载好的目录下遍历所有文件，把文件内容转成语音。

## 下载的小说转换成mp3语音文件和附带的字幕
遍历当前目录下所有txt文件，生成对应的mp3和vvt字幕;

1.txt -> 1.txt.vvt & 1.txt.mp3

```
sh ./text_to_audio.sh
```

### `text_to_audio` 的细节

`text_to_audio.sh` 提供的tts服务依赖 [github 开源服务 edge-tts](https://github.com/rany2/edge-tts)，按照文档内容进行安装；

`text_to_audio.sh` 内的参数比较简单，主要是语音和语速，文件里是我自己比较喜欢的节奏，可以按照自己需要进行调节。

语速就简单调整数字，具体的看[edge-tts readme](https://github.com/rany2/edge-tts?tab=readme-ov-file#changing-rate-volume-and-pitch)
下面命令获取中文语音列表
```
edge-tts --list-voices | grep zh-CN 
```

## 字幕格式转换
edge-tts生成的字幕文件是webvtt格式，这种格式不支持嵌入mp3
vtttolrc用来把vtt格式的字幕转成lrc
#### 编译

```
go build -o vtttolrc main.go
```
编译好的文件移动到 **$PATH** 定义的路径中以便可以全局执行

#### 使用
在包含vvt字幕的目录下执行脚本，当前目录下所有vvt会被转换成lrc
```
sh ./vtttolrc.sh
```

## 字幕嵌入mp3
字幕嵌入依赖ffmpeg，mac下可以使用brew进行安装

```
brew install ffmpeg
```

遍历目录下所有的txt文件，把对应的lrc字幕嵌入到mp3中。

比如1.txt对应的mp3是1.txt.mp3，

对应的vvt是1.txt.vvt，

对应的lrc是1.txt.vvt.lrc，

对应处理后的mp3是1.txt.mp3.mp3

```
sh ./injectlrc.sh
```

## 编译
···
go build main.go
···

## m.biquke.vip 小说下载
注意entry需要是章节的**分页列表首页**，
点击章节列表，默认是最新章节列表，需要再次点击查看完整目录，才能跳转到需要的页面，注意辨别格式
···
./main -downloader=m.biquke.vip -entry='https://m.biquke.vip/book/23557/
···
