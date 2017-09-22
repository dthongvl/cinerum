window.onload = function () {
    var conn;
    var roomId = document.URL.substring(document.URL.lastIndexOf("/") + 1, document.URL.length)

    conn = new WebSocket("ws://" + document.location.host + "/ws?roomId=" + roomId);
    conn.onopen = function (evt) {

    };

    conn.onclose = function (evt) {
    };

    conn.onmessage = function (evt) {
        console.log(evt);
    };
};