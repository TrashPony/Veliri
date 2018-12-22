function moveWindow(event, id) {

    let window = document.getElementById(id);

    window.style.marginTop = "0px";

    let coordinates = getCoordinates(window);

    let shiftX = event.pageX - coordinates.left + document.body.scrollLeft;
    let shiftY = event.pageY - coordinates.top + document.body.scrollTop;

    document.body.appendChild(window);
    moveAt(event);

    function moveAt(event) {
        window.style.left = event.clientX - shiftX + 'px';
        window.style.top = event.clientY - shiftY +  'px';
        window.style.position = "fixed";
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
