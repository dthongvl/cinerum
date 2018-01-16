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
            var messageInput = document.getElementById("message-input");
            if (event.keyCode === 13) {
                if (messageInput.value !== "\n") {
                    chat(messageInput.value);
                }
                messageInput.value = "";
            }
        });
    };

    conn.onclose = function (event) {
    };

    conn.onmessage = function (event) {
        console.log(event.data);
        var message = JSON.parse(event.data);
        if (message.type === "message") {
            onChatCommand(message);
        } else if(message.type === "online") {
            onUpdateTotalOnline(message)
        }
    };

    function chat(data) {
        conn.send(JSON.stringify({
            data: data,
            username: username,
            roomID: roomID,
            type: "message"
        }));
    }

    function onChatCommand(message) {
        var li = document.createElement("li");
        li.className += 'list-group-item';
        li.innerHTML = '<b style="color:' + getRandomColor() + '">' + message.username + ': </b><small>' + message.data + '</small>';
        var chatBox = document.getElementById("chat-box");
        chatBox.appendChild(li);
        chatBox.scrollTop = chatBox.scrollHeight;
    }

    function onUpdateTotalOnline(message) {
        document.getElementById("total-online").innerText = message.data
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