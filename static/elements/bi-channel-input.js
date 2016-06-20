Polymer({
    is: 'bi-channel-input',

    properties: {
        channelName: {
            type: String
        }
    },

    ready: function() { },

    handleTap: function() {
        if(this.channelName) {
            this.displayMessage("channel name is " + this.channelName);
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
