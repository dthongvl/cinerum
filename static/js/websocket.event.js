$(document).ready(function () {
        $(".modal").modal("show");

        var conn;
        var roomId = document.URL.substring(document.URL.lastIndexOf("/") + 1, document.URL.length);
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
                onChatCommand(message);
            }
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

        function setUsername() {
            $(".modal").modal("hide");
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

        function onPlayerCommand(message) {

        }

        function onRTCCommand(message) {

        }

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
    }
);
