Polymer({
    is: 'bi-channel-list',

    data: [],

    ready: function() {
        this.$.grid.items = this.data;
        console.log("bi-channel-list ready");
    },

    addItem: function(item) {
        //console.log("bi-channel-list.addItem" + JSON.stringify(item));
        this.data.push(item);
        this.$.grid.size = this.data.length;
        this.$.grid.refreshItems();
    }
});
