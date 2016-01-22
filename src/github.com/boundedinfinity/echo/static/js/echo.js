var socket;
var items = [];

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
        $("#count").text(0);
    });

})(document);
