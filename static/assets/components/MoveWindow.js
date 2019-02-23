function moveWindow(event, id) {
    let window = document.getElementById(id);

    $(window).draggable({
        disabled: false
    });
    this.onmouseup = function () {
        $(window).draggable({
            disabled: true
        });
    }
}