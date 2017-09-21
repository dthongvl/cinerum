var player = videojs("room-video");

player.ready(function () {
    player.on("play", function () {
        console.log("play");
    });

    player.on("pause", function () {
        console.log("paused");
    });

    player.on("timeupdate", function () {
        console.log("time update" + player.currentTime());
    });

    player.on("seeked", function () {
        console.log("seeked");
    });
});