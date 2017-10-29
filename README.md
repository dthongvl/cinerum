# cinerum

### Dependencies

- [Golang](http://golang.org/) with [Dep](https://github.com/golang/dep)
- *Linux only*: [Nginx](https://nginx.org/) with [nginx-rtmp-module](https://github.com/arut/nginx-rtmp-module)
- *Windows only*: [Mingw-w64](https://sourceforge.net/projects/mingw-w64/)
- FFmpeg

### Preconfigure:
- Run
```
dep ensure
```
- *Linux*: Replace **nginx.conf**

### To test
- Run nginx (everytime reset computer)
    - *Window*: run **nginx.exe** in nginx-window folder
- Run ffmpeg to push stream:
```
ffmpeg -re -i [fileName] -strict experimental -c:v libx264 -preset veryfast -maxrate 3000k -bufsize 6000k -pix_fmt yuv420p -g 50 -c:a aac -b:a 160k -ac 2 -ar 44100 -f flv rtmp://localhost:1935/app/[streamKey]
```
- Run main.go
- [Enjoy](http://localhost:3000)

### TODO
- [ ] Adjust chat box and player size
- [ ] Show auth modal when unauthenticated user chat
- [ ] Register feature
- [ ] Video player

### Credits and References
- https://benwilber.github.io/streamboat.tv/nginx/rtmp/streaming/2016/10/22/implementing-stream-keys-with-nginx-rtmp-and-django.html
- https://github.com/illuspas/nginx-rtmp-win32