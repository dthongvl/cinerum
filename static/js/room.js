$(document).ready(function () {
    var conn = new WebSocket("ws://" + document.location.host + "/ws?roomID=" + roomID);
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
        console.log("ahihi");
        console.log(event.data);
        var message = JSON.parse(event.data);
        onChatCommand(message);
    };

    var messageFormat = '<div class="row">\n' +
        '                        <div class="col-1">\n' +
        '                            <img class="user-avatar rounded" width="36px" height="36px"\n' +
        '                                 src="/static/img/default-avatar.jpeg"\n' +
        '                                 alt="...">\n' +
        '                        </div>\n' +
        '                        <div class="col-10 message-area">\n' +
        '                            <div class="message-header">\n' +
        '                                <span class="user-name">{{username}}</span><span class="message-time">{{time}}</span>\n' +
        '                            </div>\n' +
        '                            <div class="message">\n' +
        '                                {{message}}\n' +
        '                            </div>\n' +
        '                        </div>\n' +
        '                    </div>';

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
        var newMessage = messageFormat.replace("{{message}}", message.data);
        newMessage = newMessage.replace("{{username}}", message.username);
        var date = new Date();
        newMessage = newMessage.replace("{{time}}", date.toLocaleTimeString(navigator.language, {
            hour: '2-digit',
            minute: '2-digit'
        }));
        var li = document.createElement("li");
        li.innerHTML = newMessage;
        document.getElementById("chat-box").appendChild(li);
    }

    var options = {
        hls: {
            withCredentials: true
        }
    };

    videojs("room-video", {flash: options, html5: options});
});