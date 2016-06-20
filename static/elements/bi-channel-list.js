Polymer({
    is: 'bi-channel-list',

    data: [],

    ready: function() {
        this.$.grid.items = this.data;
        console.log("bi-channel-list ready");
    },

    addItem: function(item) {
        console.log("bi-channel-list.addItem");
        this.data.push(item);
        this.$.grid.size = data.length;
        this.refreshItems();
    }
});
