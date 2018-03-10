## install

```
go get github.com/jjunk1989/go-ffmpeg-web
```

## use

```
go build -i
```

## use ffmpeg cmd line 

```
ffmpeg -t 5 -ss 00:00:00 -i test.gif -i test.mp3 -c:v libx264 -c:a aac -b:a 128k -vf scale=420:-2,format=yuv420p out.mp4
```

## powered 

*[gin](https://github.com/gin-gonic/gin)
*[ffmpeg](http://ffmpeg.org/)
