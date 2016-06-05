var socket;
var items = [];
var detailsOpenIndex = -1;

(function(document) {
    'use strict';

    window.addEventListener('WebComponentsReady', function() {
        console.log("channel: " + channel);
        socket = new WebSocket('ws://' + window.location.host + '/ws/' + channel);

        socket.onmessage = function (event) {
            var data = JSON.parse(event.data);

            items.push(data);
            grid.size = items.length;
            grid.refreshItems();
            $("#count").text(items.length);
        };

       var jsonRenderer = function(cell) {
             cell.element.innerHTML = '';
             var cellData = JSON.parse(cell.data);
             var cellJson = JSON.stringify(cellData, null, 2);
             cell.element.innerHTML = cellJson;
        };

        var grid = document.getElementById('grid');
        grid.columns[2].renderer = jsonRenderer;
        grid.items = items;

        grid.rowDetailsGenerator = function(rowIndex) {
            var elem = document.createElement('code');

            grid.getItem(rowIndex, function(error, item) {
                if(error) {
                    alert('error: ' + error);
                } else {
                    var json = JSON.parse(item.body);
                    elem.innerHTML = JSON.stringify(json, null, 4);
                }
            });

            return elem;
        };

        grid.addEventListener('selected-items-changed', function() {
            grid.setRowDetailsVisible(detailsOpenIndex, false);
            var selected = grid.selection.selected();

            if (selected.length == 1) {
                grid.setRowDetailsVisible(selected[0], true);
                detailsOpenIndex = selected[0];
            }
        });

        setTimeout(function(){
            $("#count").text(0);
            $("#restUrl").text("http://" + window.location.host + '/rest/' + channel)
        }, 100);

    });

})(document);
