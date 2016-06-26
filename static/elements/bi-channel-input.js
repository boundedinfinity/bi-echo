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
            this.fire('item', this.createItem());
        } else {
            this.displayMessage("channel cannot be empty");
        }
    },

    createItem: function() {
        return {
            name: this.channelName,
            url: 'http://' + window.location.host + '/ws/' + this.channelName
        }
    },

    displayMessage: function(message) {
        this.$.toast.show({
            text: message,
            duration: 3000
        });
    }
});
