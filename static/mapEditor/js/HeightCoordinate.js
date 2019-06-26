function HeightCoordinate(add) {
    if (add) {
        let callBack = function (q, r) {
            addHeightCoordinate(q, r)
        };
        SelectedSprite(event, 0, callBack, null, null, null, true)
    } else {
        let callBack = function (q, r) {
            subtractHeightCoordinate(q, r)
        };
        SelectedSprite(event, 0, callBack, null, null, null, true)
    }
}