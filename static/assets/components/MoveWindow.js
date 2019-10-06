function moveWindow(event, id) {
    let mWindow = document.getElementById(id);

    $(mWindow).draggable({
        disabled: false,
        stop: function (event, ui) {
            if (window.location.pathname !== "/editors/map/") {
                setState(mWindow.id, $(mWindow).position().left, $(mWindow).position().top, $(mWindow).height(), $(mWindow).width(), true);
            }
        }
    });
    this.onmouseup = function () {
        $(mWindow).draggable({
            disabled: true,
        });
    }
}