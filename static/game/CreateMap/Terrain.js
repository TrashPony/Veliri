function CreateTerrain(coordinate, x, y, q, r) {
    if (coordinate.coordinateText) {
        for (let text in coordinate.coordinateText) {
            coordinate.coordinateText[text].destroy();
        }
    } else {
        coordinate.coordinateText = {};
    }

    if (game && game.typeService === "mapEditor") {
        let style = {font: "12px Arial", fill: "#ffed00", align: "center"};

        if (metaAlpha && metaAlpha === 0) {
            coordinate.coordinateText.qr.alpha = metaAlpha;
            coordinate.coordinateText.qr = game.add.text(x - 10, y - 15, q + "," + r, style, game.redactorMetaText);
        }

        let allow = {font: "12px Arial", fill: "#150bff", align: "center"};
        let noAllow = {font: "12px Arial", fill: "#ff2821", align: "center"};

        if (!(coordinate.move && coordinate.view && coordinate.attack)) {
            if (coordinate.move) {
                coordinate.coordinateText.move = game.add.text(x - 20, y - 15, 'm', allow, game.redactorMetaText);
            } else {
                coordinate.coordinateText.move = game.add.text(x - 20, y - 15, 'm', noAllow, game.redactorMetaText);
            }

            if (coordinate.view) {
                coordinate.coordinateText.view = game.add.text(x - 12, y - 15, 'w', allow, game.redactorMetaText);
            } else {
                coordinate.coordinateText.view = game.add.text(x - 12, y - 15, 'w', noAllow, game.redactorMetaText);
            }

            if (coordinate.attack) {
                coordinate.coordinateText.attack = game.add.text(x - 5, y - 15, 'a', allow, game.redactorMetaText);
            } else {
                coordinate.coordinateText.attack = game.add.text(x - 5, y - 15, 'a', noAllow, game.redactorMetaText);
            }
        }


        if (coordinate.transport) {
            let transportIcon = game.redactorMetaText.create(x + 10, y - 10, 'transportIcon');
            transportIcon.anchor.setTo(0.5);
            transportIcon.scale.set(0.5);
        }

        if (coordinate.handler === 'sector') {
            let transportIcon = game.redactorMetaText.create(x + 10, y - 10, 'sectorOutIcon');
            transportIcon.anchor.setTo(0.5);
            transportIcon.scale.set(0.5);
        }

        if (coordinate.handler === 'base') {
            let transportIcon = game.redactorMetaText.create(x + 10, y - 10, 'baseInIcon');
            transportIcon.anchor.setTo(0.5);
            transportIcon.scale.set(0.3);
        }

        if (coordinate.level === 0) {
            let style = {font: "24px Arial", fill: "#bbfff1", align: "center"};
            coordinate.coordinateText.height = game.add.text(x - 5, y - 5, coordinate.level, style, game.redactorMetaText);
        }

        if (coordinate.level === 1) {
            let style = {font: "24px Arial", fill: "#35daff", align: "center"};
            coordinate.coordinateText.height = game.add.text(x - 5, y - 5, coordinate.level, style, game.redactorMetaText);
        }

        if (coordinate.level === 2) {
            let style = {font: "24px Arial", fill: "#68ff59", align: "center"};
            coordinate.coordinateText.height = game.add.text(x - 5, y - 5, coordinate.level, style, game.redactorMetaText);
        }

        if (coordinate.level === 4) {
            let style = {font: "24px Arial", fill: "#fff523", align: "center"};
            coordinate.coordinateText.height = game.add.text(x - 5, y - 5, coordinate.level, style, game.redactorMetaText);
        }

        if (coordinate.level === 5) {
            let style = {font: "24px Arial", fill: "#ff2821", align: "center"};
            coordinate.coordinateText.height = game.add.text(x - 5, y - 5, coordinate.level, style, game.redactorMetaText);
        }
    }
}