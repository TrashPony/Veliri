function CreateTerrain(coordinate, x, y, q, r) {

    let floorSprite = game.floorLayer.create(x, y, coordinate.texture_flore);
    floorSprite.anchor.setTo(0.5);
    floorSprite.scale.set(0.5);

    floorSprite.inputEnabled = true; // включаем ивенты на спрайт
    floorSprite.input.pixelPerfectOver = true;   // уберает ивенты наведения на пустую зону спрайта
    floorSprite.input.pixelPerfectClick = true;  // уберает ивенты кликов на пустую зону спрайта

    floorSprite.events.onInputOut.add(TipOff, floorSprite);
    floorSprite.z = 0;
    coordinate.sprite = floorSprite;

    if (game && game.typeService === "battle") {
        let fogSprite = game.fogOfWar.create(x, y, 'FogOfWar');
        fogSprite.anchor.setTo(0.5);
        fogSprite.scale.set(0.5);
        floorSprite.events.onInputDown.add(RemoveSelect);
        coordinate.fogSprite = fogSprite;
        coordinate.fogSprite = fogSprite;
    }

    coordinate.coordinateText = {};

    if (game && game.typeService !== "battle") {
        let style = {font: "12px Arial", fill: "#606060", align: "center"};
        coordinate.coordinateText.qr = game.add.text(x - 10, y - 40, q + "," + r, style, game.redactorMetaText);

        if (metaAlpha === 0){
            coordinate.coordinateText.qr.alpha = metaAlpha;
        }

        let allow = {font: "12px Arial", fill: "#150bff", align: "center"};
        let noAllow = {font: "12px Arial", fill: "#ff2821", align: "center"};

        if (!(coordinate.move && coordinate.view && coordinate.attack)) {
            if (coordinate.move) {
                coordinate.coordinateText.move = game.add.text(floorSprite.x - 40, floorSprite.y - 25, 'm', allow, game.redactorMetaText);
            } else {
                coordinate.coordinateText.move = game.add.text(floorSprite.x - 40, floorSprite.y - 25, 'm', noAllow, game.redactorMetaText);
            }

            if (coordinate.view) {
                coordinate.coordinateText.view = game.add.text(floorSprite.x - 25, floorSprite.y - 25, 'w', allow, game.redactorMetaText);
            } else {
                coordinate.coordinateText.view = game.add.text(floorSprite.x - 25, floorSprite.y - 25, 'w', noAllow, game.redactorMetaText);
            }

            if (coordinate.attack) {
                coordinate.coordinateText.attack = game.add.text(floorSprite.x - 10, floorSprite.y - 25, 'a', allow, game.redactorMetaText);
            } else {
                coordinate.coordinateText.attack = game.add.text(floorSprite.x - 10, floorSprite.y - 25, 'a', noAllow, game.redactorMetaText);
            }
        }

        if (!coordinate.impact) {
            let buttons = createButtons(coordinate);

            coordinate.buttons = buttons;

            floorSprite.events.onInputOver.add(function () {
                buttons.plus.visible = true;
                buttons.minus.visible = true;
                buttons.rotate.visible = true;
            });

            floorSprite.events.onInputOut.add(function () {
                hideButtons();
            });
        }
    }

    if (coordinate.level === 0) {
        let style = {font: "36px Arial", fill: "#bbfff1", align: "center"};
        coordinate.coordinateText.height = game.add.text(floorSprite.x - 50, floorSprite.y - 15, coordinate.level, style, game.redactorMetaText);
    }

    if (coordinate.level === 1) {
        let style = {font: "36px Arial", fill: "#35daff", align: "center"};
        coordinate.coordinateText.height = game.add.text(floorSprite.x - 50, floorSprite.y - 15, coordinate.level, style, game.redactorMetaText);
    }

    if (coordinate.level === 3) {
        let style = {font: "36px Arial", fill: "#68ff59", align: "center"};
        coordinate.coordinateText.height = game.add.text(floorSprite.x - 50, floorSprite.y - 15, coordinate.level, style, game.redactorMetaText);
    }

    if (coordinate.level === 4) {
        let style = {font: "36px Arial", fill: "#fff523", align: "center"};
        coordinate.coordinateText.height = game.add.text(floorSprite.x - 50, floorSprite.y - 15, coordinate.level, style, game.redactorMetaText);
    }

    if (coordinate.level === 5) {
        let style = {font: "36px Arial", fill: "#ff2821", align: "center"};
        coordinate.coordinateText.height = game.add.text(floorSprite.x - 50, floorSprite.y - 15, coordinate.level, style, game.redactorMetaText);
    }
}

function createButtons(coordinate) {
    let buttonPlus = game.redactorButton.create(coordinate.sprite.x - 30, coordinate.sprite.y - 30, 'buttonPlus');
    buttonPlus.scale.set(0.15);
    let buttonMinus = game.redactorButton.create(coordinate.sprite.x + 5, coordinate.sprite.y - 30, 'buttonMinus');
    buttonMinus.scale.set(0.15);
    let buttonRotate = game.redactorButton.create(coordinate.sprite.x - 20, coordinate.sprite.y + 20, 'buttonRotate');
    buttonRotate.scale.set(0.25);

    buttonPlus.inputEnabled = true;
    buttonMinus.inputEnabled = true;
    buttonRotate.inputEnabled = true;

    buttonPlus.events.onInputOver.add(function () {
        buttonPlus.visible = true;
        buttonPlus.events.onInputDown.add(addHeightCoordinate, coordinate);
    });

    buttonMinus.events.onInputOver.add(function () {
        buttonMinus.visible = true;
        buttonMinus.events.onInputDown.add(subtractHeightCoordinate, coordinate);
    });

    buttonRotate.events.onInputOver.add(function () {
        if (coordinate.objectSprite) {
            buttonRotate.visible = true;
            buttonRotate.events.onInputDown.add(ChangeOptionSprite, coordinate);
        }
    });

    hideButtons();
    return {plus: buttonPlus, minus: buttonMinus, rotate: buttonRotate}
}

function hideButtons() {
    if (game.redactorButton) {
        for (let i in game.redactorButton.children) {
            if (game.redactorButton.children.hasOwnProperty(i)) {
                game.redactorButton.children[i].visible = false;
            }
        }
    }
}