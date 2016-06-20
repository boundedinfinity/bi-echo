Polymer({
    is: 'bi-channel-list',
    properties: {
        properties: {

        },

        ready: function() {
        },

        onError: function(error) {
            this.fire('onerror', error);
        }
    }
});
