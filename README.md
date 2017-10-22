# cinerum

### Dependencies

- Golang with [Glide](https://github.com/Masterminds/glide)
- [Nginx](https://nginx.org/) with [nginx-ts-module](https://github.com/arut/nginx-ts-module)

### First time run:
- Run
> glide install
- Replace **nginx.conf**
- Run nginx (only when reset computer)

### To test
- Run command:
```
ffmpeg -re -i movie.ext -c:v libx264 -preset veryfast -maxrate 3000k -bufsize 6000k -pix_fmt yuv420p -g 50 -b:a 160k -ac 2 -ar 44100 -f mpegts http://127.0.0.1:8000/publish/ffmw
```
- Run main.go
- [Enjoy](http://localhost:3000)