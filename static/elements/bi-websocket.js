Polymer({
    is: 'bi-websocket',
    properties: {
        socket: null,

        properties: {
            protocol: {
                type: String,
                value: ""
            },
            url: {
                type: String
            }
        },

        ready: function() {
            this.socket = new WebSocket(this.url, this.protocol);
            this.socket.onerror = this.onError.bind(this);
            this.socket.onopen = this.onOpen.bind(this);
            this.socket.onmessage = this.onMessage.bind(this);
        },

        onError: function(error) {
            this.fire('onerror', error);
        },

        onOpen: function(event) {
            this.fire('onopen');
        },

        onMessage: function(event) {
            this.fire('onmessage', event.data);
        },

        send: function(message) {
            this.socket.send(message);
        },

        close: function() {
            this.socket.close();
        }
    }
});
