#小说下载
单本小说下载，每章单独存储为一个文件。

`text_to_audio.sh` 这个脚本可以在下载好的目录下遍历所有文件，把文件内容转成语音。

## `text_to_audio.sh` 使用
`text_to_audio.sh` 提供的tts服务依赖 [github 开源服务 edge-tts](https://github.com/rany2/edge-tts)，按照文档内容进行安装；

`text_to_audio.sh` 内的参数比较简单，主要是语音和语速，文件里是我自己比较喜欢的节奏，可以按照自己需要进行调节。

语速就简单调整数字，具体的看[edge-tts readme](https://github.com/rany2/edge-tts?tab=readme-ov-file#changing-rate-volume-and-pitch)
下面命令获取中文语音列表
```
edge-tts --list-voices | grep zh-CN 
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
