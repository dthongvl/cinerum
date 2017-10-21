# cinerum

### Dependencies

- Golang with [Glide](https://github.com/Masterminds/glide)
- [Nginx](https://nginx.org/) with [nginx-ts-module](https://github.com/arut/nginx-ts-module)

### How to run:
- Replace nginx.conf (only first time)
- Run nginx (only when reset computer)
- cd template and run **hero** (anytime you edit html)
- Run command:
```
ffmpeg -re -i movie.ext -c:v libx264 -preset veryfast -maxrate 3000k -bufsize 6000k -pix_fmt yuv420p -g 50 -b:a 160k -ac 2 -ar 44100 -f mpegts http://127.0.0.1:8000/publish/ffmw
```
- Run main.go
- [Enjoy](http://localhost:3000)