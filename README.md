# 参考链接

* [ffmpeg](https://www.ffmpeg.org/documentation.html)

* [使用 ffmpeg 实现 MP4 与 GIF 的互转](http://note.rpsh.net/posts/2015/04/21/mac-osx-ffmpeg-mp4-gif-convert/)

ffmpeg -i input.gif -vf scale=420:-2,format=yuv420p out.mp4

* [FFmpeg](http://blog.csdn.net/john_crash/article/details/48676229)

# 命令行

```
ffmpeg -t 5 -ss 00:00:00 -i test1.gif -i test1.mp3 -c:v libx264 -c:a aac -b:a 128k -vf scale=420:-2,format=yuv420p out.mp4
```# go-ffmpeg-web
