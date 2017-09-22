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

function onChat(message) {
    var li = document.createElement("li");
    var newMessage = messageFormat.replace("{{message}}", message.data);
    newMessage = newMessage.replace("{{username}}", message.username);
    var date = new Date();
    newMessage = newMessage.replace("{{time}}", date.toLocaleTimeString(navigator.language, {hour: '2-digit', minute:'2-digit'}));
    li.innerHTML = newMessage;
    var ul = document.getElementById("chat-box");
    ul.appendChild(li);
}