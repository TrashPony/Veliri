function GetMapList() {
    mapEditor.send(JSON.stringify({
        event: "getMapList"
    }));

    mapEditor.send(JSON.stringify({
        event: "getAllTypeCoordinate"
    }));
}

let game;

function CreateMapList(jsonMessage) {

    let maps = JSON.parse(jsonMessage).maps;

    let mapSelect = document.getElementById("mapSelector");

    for (let i = 0; i < maps.length; i++) {
        if (document.getElementById(maps[i].Id)) {
            continue;
        }

        let option = document.createElement("option");
        option.id = maps[i].Id;
        option.value = maps[i].Id;
        if (maps[i].global) {
            option.innerHTML = maps[i].Name + "<span style='color: red'> Глоб.</span>";
        } else {
            option.innerHTML = maps[i].Name;
        }
        mapSelect.appendChild(option);
    }
}

function selectMap() {
    let selectedValue = document.getElementById("mapSelector").options[document.getElementById("mapSelector").selectedIndex].value;

    mapEditor.send(JSON.stringify({
        event: "SelectMap",
        id: Number(selectedValue)
    }));
}

function createGame(jsonMessage) {
    if (game) {
        UpdateMap(JSON.parse(jsonMessage).map, game);
    } else {
        game = CreateGame(JSON.parse(jsonMessage).map);
        appendRedactorEventsToFloor(game)
    }
}

function appendRedactorEventsToFloor(game) {
    let map = game.map.OneLayerMap;

    setTimeout(function () {
        //костыль, если без таймаута то карта не успевает заполниться спрайтами
        addButtons(map)
    }, 2500)
}

function addButtons(map) {
    for (let q in map) {
        if (map.hasOwnProperty(q)) {
            for (let r in map[q]) {
                if (map[q].hasOwnProperty(r)) {

                    if (map[q][r].impact) {
                        continue
                    }

                    map[q][r].sprite.events.onInputOver.add(function () {
                        createButtons(map[q][r].sprite)
                    });

                    map[q][r].sprite.events.onInputOut.add(function () {
                        // TODO срабатывает когда наводишь на кнопку т.к. мыш уходит с спрайта земли
                        removeButtons();
                    });
                }
            }
        }
    }
}

function createButtons(coordinate) {
    let buttonPlus = game.redactorButton.create(coordinate.x - 30, coordinate.y - 30, 'buttonPlus');
    buttonPlus.scale.set(0.15);
    let buttonMinus = game.redactorButton.create(coordinate.x + 5, coordinate.y - 30, 'buttonMinus');
    buttonMinus.scale.set(0.15);
    let buttonRotate = game.redactorButton.create(coordinate.x - 20, coordinate.y + 20, 'buttonRotate');
    buttonRotate.scale.set(0.25);

    buttonPlus.inputEnabled = true;
    buttonMinus.inputEnabled = true;
    buttonRotate.inputEnabled = true;

    buttonPlus.events.onInputOver.add(function () {
        buttonPlus.events.onInputDown.add(addHeightCoordinate, coordinate);
    });

    buttonMinus.events.onInputOver.add(function () {
        buttonMinus.events.onInputDown.add(subtractHeightCoordinate, coordinate);
    });

    buttonRotate.events.onInputOver.add(function () {
        if (coordinate.objectSprite) {
            buttonRotate.events.onInputDown.add(ChangeOptionSprite, coordinate);
        }
    });
}

function removeButtons() {
    if (game.redactorButton) {
        for (let i in game.redactorButton.children) {
            if (game.redactorButton.children.hasOwnProperty(i)) {
                game.redactorButton.children[i].kill()
            }
        }
    }
}