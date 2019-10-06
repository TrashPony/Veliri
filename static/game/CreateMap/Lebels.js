function CreateLabels(coordinate, x, y) {
    if (coordinate.coordinateText) {
        for (let text in coordinate.coordinateText) {
            coordinate.coordinateText[text].destroy();
        }
    } else {
        coordinate.coordinateText = {};
    }

    if (game && game.typeService === "mapEditor") {

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

function CreateLabelEntry(entryPoints) {
    for (let i of entryPoints) {
        for (let position of i.positions) {
            let baseResp = game.icon.create(position.x, position.y, 'baseResp');
            baseResp.angle = position.resp_rotate;
            baseResp.anchor.setTo(0.5);
            baseResp.scale.setTo(0.05);
        }
    }
}

function CreateLabelBase(bases) {
    for (let i in bases) {
        if (bases.hasOwnProperty(i)) {

            let base = game.icon.create(bases[i].x, bases[i].y, 'baseIcon');
            base.anchor.setTo(0.5);
            base.scale.setTo(0.1);

            for (let j in bases[i].respawns) {
                let respPount = bases[i].respawns[j];
                let baseResp = game.icon.create(respPount.x, respPount.y, 'baseResp');

                baseResp.angle = respPount.resp_rotate;
                baseResp.anchor.setTo(0.5);
                baseResp.scale.setTo(0.05);
            }
        }
    }
}