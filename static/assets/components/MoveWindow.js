function moveWindow(event, id) {
    let window = document.getElementById(id);

    $(window).draggable({
        disabled: false,
        stop: function (event, ui) {
            setState(window.id, $(window).position().left, $(window).position().top, $(window).height(), $(window).width(), true);
        }
    });
    this.onmouseup = function () {
        $(window).draggable({
            disabled: true,
        });
    }
}