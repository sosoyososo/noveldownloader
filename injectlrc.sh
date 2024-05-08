# 遍历当前所有的mp3文件，注入字幕
dic=`pwd`
for file in "$dic"/*.txt; do
    if [ -f "$file" ]; then
        ffmpeg -i $file.mp3 -metadata LYRICS="$(<$file.vvt.lrc)" $file.mp3.mp3
    fi
done