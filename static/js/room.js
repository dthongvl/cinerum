function chat() {
    var messageInput = document.getElementById("message-input");
    ws.send(JSON.stringify({
        data: messageInput.value,
        username: "dthongvl",
        type: "chat",
        roomId: roomId
    }));
    messageInput.value = "";
}