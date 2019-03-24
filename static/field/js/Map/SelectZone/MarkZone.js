function MarkZone(xy, placeCoordinates, q, r, selectClass, addEmpty, typeLine, selector, typeSelect, drawLine) {
    let topLeft = false;
    let topRight = false;
    let left = false;
    let right = false;
    let botLeft = false;
    let botRight = false;

    let sprite;

    if (!xy) {
        return
    }

    /*
        соседи гексов беруться по разному в зависимости от четности строки
        // even {Q,R}

           {0,-1}  {+1,-1}
        {-1,0} {0,0} {+1,0}
           {0,+1}  {+1,+1}

        // odd
          {-1,-1}  {0,-1}
        {-1,0} {0,0} {+1,0}
          {-1,+1}  {0,+1}
    */

    if (placeCoordinates.hasOwnProperty(Number(q) - 1)) {
        if (placeCoordinates[Number(q) - 1].hasOwnProperty(Number(r))) {
            left = true;
        }
    }

    if (placeCoordinates.hasOwnProperty(Number(q) + 1)) {
        if (placeCoordinates[Number(q) + 1].hasOwnProperty(Number(r))) {
            right = true;
        }
    }

    if (r % 2 !== 0) {
        if (placeCoordinates.hasOwnProperty(Number(q))) {
            if (placeCoordinates[Number(q)].hasOwnProperty(Number(r) - 1)) {
                topLeft = true;
            }
        }

        if (placeCoordinates.hasOwnProperty(Number(q) + 1)) {
            if (placeCoordinates[Number(q) + 1].hasOwnProperty(Number(r) - 1)) {
                topRight = true;
            }
        }

        if (placeCoordinates.hasOwnProperty(Number(q))) {
            if (placeCoordinates[Number(q)].hasOwnProperty(Number(r) + 1)) {
                botLeft = true;
            }
        }

        if (placeCoordinates.hasOwnProperty(Number(q) + 1)) {
            if (placeCoordinates[Number(q) + 1].hasOwnProperty(Number(r) + 1)) {
                botRight = true;
            }
        }
    } else {
        if (placeCoordinates.hasOwnProperty(Number(q) - 1)) {
            if (placeCoordinates[Number(q) - 1].hasOwnProperty(Number(r) - 1)) {
                topLeft = true;
            }
        }

        if (placeCoordinates.hasOwnProperty(Number(q))) {
            if (placeCoordinates[Number(q)].hasOwnProperty(Number(r) - 1)) {
                topRight = true;
            }
        }

        if (placeCoordinates.hasOwnProperty(Number(q) - 1)) {
            if (placeCoordinates[Number(q) - 1].hasOwnProperty(Number(r) + 1)) {
                botLeft = true;
            }
        }

        if (placeCoordinates.hasOwnProperty(Number(q))) {
            if (placeCoordinates[Number(q)].hasOwnProperty(Number(r) + 1)) {
                botRight = true;
            }
        }
    }

    if (addEmpty) {
        if (selector === "move" || selector === "place") {
            sprite = typeSelect.create(xy.x, xy.y, 'selectEmpty');
            sprite.anchor.setTo(0.5);
        }
        if (selector === "target") {
            sprite = typeSelect.create(xy.x, xy.y, 'selectTarget');
            sprite.anchor.setTo(0.5);
        }
        sprite.scale.set(0.5);
    }

    if (drawLine) {
        if (!left) {
            let line = typeLine.create(xy.x, xy.y, 'line' + selectClass, 4);
            line.anchor.setTo(0.5);
            line.scale.set(0.5)
        }

        if (!right) {
            let line = typeLine.create(xy.x, xy.y, 'line' + selectClass, 1);
            line.anchor.setTo(0.5);
            line.scale.set(0.5)
        }

        if (!topLeft) {
            let line = typeLine.create(xy.x, xy.y, 'line' + selectClass, 5);
            line.anchor.setTo(0.5);
            line.scale.set(0.5)
        }

        if (!topRight) {
            let line = typeLine.create(xy.x, xy.y, 'line' + selectClass, 0);
            line.anchor.setTo(0.5);
            line.scale.set(0.5)
        }

        if (!botLeft) {
            let line = typeLine.create(xy.x, xy.y, 'line' + selectClass, 3);
            line.anchor.setTo(0.5);
            line.scale.set(0.5)
        }

        if (!botRight) {
            let line = typeLine.create(xy.x, xy.y, 'line' + selectClass, 2);
            line.anchor.setTo(0.5);
            line.scale.set(0.5)
        }
    }

    return sprite
}