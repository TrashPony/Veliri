
function moveWindow(event, id) {

    var window = document.getElementById(id);

    window.style.marginTop = "0px";

    var coordinates = getCoordinates(window);

    var shiftX = event.pageX - coordinates.left;
    var shiftY = event.pageY - coordinates.top;

    document.body.appendChild(window);
    moveAt(event);

    function moveAt(event) {
        window.style.left = event.pageX - shiftX + 'px';
        window.style.top = event.pageY - shiftY + 'px';
    }

    document.onmousemove = function(event) {
        moveAt(event);
    };

    document.onmouseup = function() {
        document.onmousemove = null;
        window.onmouseup = null;
    };

    window.ondragstart = function() {
        return false;
    };

    function getCoordinates(window) {
        var box = window.getBoundingClientRect();
        return {
            top: box.top + pageYOffset,
            left: box.left + pageXOffset
        };
    }
}
