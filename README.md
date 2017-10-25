# cinerum

### Dependencies

- [Golang](http://golang.org/) with [Glide](https://github.com/Masterminds/glide)
- [Nginx](https://nginx.org/) with [nginx-rtmp-module](https://github.com/arut/nginx-rtmp-module)
- FFmpeg

### Preconfigure:
- Run
> glide install
- Replace **nginx.conf**

### To test
- Run nginx (only when reset computer)
- Run command to push stream:
```
ffmpeg -re -i [fileName] -strict experimental -c:v libx264 -preset veryfast -maxrate 3000k -bufsize 6000k -pix_fmt yuv420p -g 50 -c:a aac -b:a 160k -ac 2 -ar 44100 -f flv rtmp://localhost:1935/app/key[roomName]
```
- Run main.go
- [Enjoy](http://localhost:3000)

### TODO
- [ ] Adjust chat box and player size
- [ ] Show auth modal when unauthenticated user chat
- [ ] Register feature