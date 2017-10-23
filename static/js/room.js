$(document).ready(function () {
    var options = {
        hls: {
            withCredentials: true
        }
    };

    videojs("room-video", {flash: options, html5: options});

    var conn = new WebSocket("ws://" + document.location.host + "/" + roomID + "/ws");
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
        console.log(event.data);
        var message = JSON.parse(event.data);
        onChatCommand(message);
    };

    function chat() {
        var messageInput = document.getElementById("message-input");
        conn.send(JSON.stringify({
            data: messageInput.value,
            username: username,
            roomID: roomID
        }));
        messageInput.value = "";
    }

    function onChatCommand(message) {
        var li = document.createElement("li");
        li.innerHTML = '<b style="color:' + getRandomColor() + '">' + message.username + ': </b><small>' + message.data + '</small>';
        document.getElementById("chat-box").appendChild(li);
    }

    function getRandomColor() {
        var letters = '0123456789ABCDEF';
        var color = '#';
        for (var i = 0; i < 6; i++) {
            color += letters[Math.floor(Math.random() * 16)];
        }
        return color;
    }
});