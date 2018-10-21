function MarkZone(cellSprite, placeCoordinates, q, r, selectClass, addEmpty, typeLine, selector, typeSelect, drawLine) {
    let topLeft = false;
    let topRight = false;
    let left = false;
    let right = false;
    let botLeft = false;
    let botRight = false;

    let sprite;

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
        if (selector === "move" || selector === "place") sprite = typeSelect.create(cellSprite.x, cellSprite.y, 'selectEmpty');
        if (selector === "target") sprite = typeSelect.create(cellSprite.x, cellSprite.y, 'selectTarget');
    }

    if (drawLine) {
        if (!left) {
            typeLine.create(cellSprite.x, cellSprite.y, 'line' + selectClass, 4);
        }

        if (!right) {
            typeLine.create(cellSprite.x, cellSprite.y, 'line' + selectClass, 1);
        }

        if (!topLeft) {
            typeLine.create(cellSprite.x, cellSprite.y, 'line' + selectClass, 5);
        }

        if (!topRight) {
            typeLine.create(cellSprite.x, cellSprite.y, 'line' + selectClass, 0);
        }

        if (!botLeft) {
            typeLine.create(cellSprite.x, cellSprite.y, 'line' + selectClass, 3);
        }

        if (!botRight) {
            typeLine.create(cellSprite.x, cellSprite.y, 'line' + selectClass, 2);
        }
    }

    return sprite
}