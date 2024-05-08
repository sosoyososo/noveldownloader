# 把当前目录下所有的vtt文件转成lrc
dic=`pwd`
for file in "$dic"/*.vtt; do
    if [ -f "$file" ]; then
        vtttolrc -vtt $file -lrc $file.lrc
    fi
done