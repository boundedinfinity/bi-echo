(function(document) {
    'use strict';
    console.log("app.js loaded")

    var app = document.querySelector('#app');

    app.addEventListener('dom-change', function() {
        console.log('dom-change fired');
    });

    window.addEventListener('WebComponentsReady', function() {
        console.log('WebComponentsReady fired');
    });

})(document);

function submitForm() {
    document.getElementById('form').submit();
}
