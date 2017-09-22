window.onload = function () {
    var conn;
    var roomId = document.URL.substring(document.URL.lastIndexOf("/") + 1, document.URL.length)
    $("#modal").modal("show");
    var username = "dthongvl";

    conn = new WebSocket("ws://" + document.location.host + "/ws?roomId=" + roomId);
    conn.onopen = function (event) {
        document.getElementById("message-input").addEventListener("keyup", function (event) {
            event.preventDefault();
            if (event.keyCode === 13 && document.getElementById("message-input").value !== "") {
                chat();
            }
        });
    };

    conn.onclose = function (event) {
    };

    conn.onmessage = function (event) {
        var message = JSON.parse(event.data);
        if (message['type'] === "chat") {
            onChat(message);
        }
    };

    function chat() {
        var messageInput = document.getElementById("message-input");
        conn.send(JSON.stringify({
            data: messageInput.value,
            username: username,
            type: "chat",
            roomId: roomId
        }));
        messageInput.value = "";
    }
};