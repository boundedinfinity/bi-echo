Polymer({
    is: 'bi-channel-input',

    properties: {
        channelName: {
            type: String
        }
    },

    ready: function() {
        console.log("bi-channel-input ready");
    },

    handleTap: function() {
        if(this.channelName) {
            this.fire('item', { name: this.channelName} );
        } else {
            this.displayMessage("channel cannot be empty");
        }
    },

    displayMessage: function(message) {
        this.$.toast.show({
            text: message,
            duration: 3000
        });
    }
});
