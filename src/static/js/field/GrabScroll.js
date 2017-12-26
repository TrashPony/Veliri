var dragHandler = {
    lastClientX: 0,
    start: function (e) {
        if (e.button === 2) {
            window.addEventListener('mousemove', dragHandler.drag);
            dragHandler.lastClientX = e.clientX;
            dragHandler.lastClientY = e.clientY;
            e.preventDefault();
        }
    },
    end: function (e) {
        if (e.button === 2) {
            window.removeEventListener('mousemove', dragHandler.drag);
        }
    },
    drag: function (e) {
        var deltaX = e.clientX - dragHandler.lastClientX;
        var deltaY = e.clientY - dragHandler.lastClientY;
        window.scrollTo(window.scrollX - deltaX, window.scrollY - deltaY);
        dragHandler.lastClientX = e.clientX;
        dragHandler.lastClientY = e.clientY;
        e.preventDefault();
    }
};
