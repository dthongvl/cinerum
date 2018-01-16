$(document).ready(function () {
    $('#message-input').click(function() {
        if (username === "") {
            $('#authModal').modal('show');
            $('#message-input').blur();
        }
    });

    var options = {
        hls: {
            withCredentials: true
        }
    };
    videojs("room-video", {flash: options, html5: options});

    var conn = new WebSocket("ws://" + document.location.host + "/" + roomID + "/ws");
    conn.onopen = function (event) {
        $('#message-input').keyup(function (keyup) {
            keyup.preventDefault();
            var messageInput = $('#message-input');
            if (keyup.keyCode === 13) {
                if (messageInput.val() !== "\n") {
                    chat(messageInput.val());
                }
                messageInput.val("");
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
        li.innerHTML = '<b style="color:' + getRandomColor() + '">' + message.username + ': </b><small style="word-wrap: break-word;">' + message.data + '</small>';
        var chatBox = $('#chat-box');
        chatBox.append(li);
        chatBox.scrollTop(chatBox.prop('scrollHeight'));
    }

    function onUpdateTotalOnline(message) {
        $('#total-online').val(message.data)
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