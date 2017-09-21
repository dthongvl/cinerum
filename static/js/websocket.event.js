var loc = window.location;
var uri = 'ws:';

if (loc.protocol === 'https:') {
    uri = 'wss:';
}
uri += '//' + loc.host + "/ws";

ws = new WebSocket(uri);

ws.onopen = function() {
    console.log('Connected')
};

ws.onmessage = function(evt) {
    console.log("message")
};
