# 设置要遍历的目录路径
dic=`pwd`
for file in "$dic"/*; do
    if [ -f "$file" ]; then
        edge-tts --voice zh-CN-YunxiNeural  --rate +40% -f $file --write-media $file.mp3 --write-subtitles $file.vtt;
        # 在这里可以对文件执行你想要的操作
    fi
done
