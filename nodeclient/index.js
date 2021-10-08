const WebSocket = require('ws');

const ws = new WebSocket('ws://127.0.0.1:8080/ws/getCustomers');

ws.on('open', function open() {
    ws.send('something');
});

ws.on('message', function incoming(message) {
    console.log('received: %s', message);
});

ws.on('error', function incoming(err) {
    console.log('Error: %s', err);
});