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
        option.innerHTML = maps[i].Name;
        mapSelect.appendChild(option);
    }
}

function selectMap() {
    let mapSelector = document.getElementById("mapSelector");
    let selectedValue = mapSelector.options[mapSelector.selectedIndex].value;

    mapEditor.send(JSON.stringify({
        event: "SelectMap",
        id: Number(selectedValue)
    }));
}

function createGame(jsonMessage) {
    game = CreateGame(JSON.parse(jsonMessage).map);
    appendRedactorEventsToFloor(game)
}

function appendRedactorEventsToFloor(game) {
    let map = game.map.OneLayerMap;

    setTimeout(function () {
        //костыль, если без таймаута то карта не успевает заполниться спрайтами
        for (let q in map) {
            if (map.hasOwnProperty(q)) {
                for (let r in map[q]) {
                    if (map[q].hasOwnProperty(r)) {

                        let buttonPlus = game.redactorButton.create(map[q][r].sprite.x + 33, map[q][r].sprite.y + 30, 'buttonPlus');
                        buttonPlus.scale.set(0.15);
                        let buttonMinus = game.redactorButton.create(map[q][r].sprite.x + 68, map[q][r].sprite.y + 30, 'buttonMinus');
                        buttonMinus.scale.set(0.15);

                        buttonPlus.alpha = 0;
                        buttonMinus.alpha = 0;


                        buttonPlus.inputEnabled = true;
                        buttonMinus.inputEnabled = true;

                        buttonPlus.events.onInputOver.add(function () {
                            buttonPlus.alpha = 1;
                            buttonMinus.alpha = 1;
                            buttonPlus.events.onInputDown.add(addHeightCoordinate, map[q][r]);
                        });

                        buttonMinus.events.onInputOver.add(function () {
                            buttonMinus.alpha = 1;
                            buttonPlus.alpha = 1;
                            buttonMinus.events.onInputDown.add(subtractHeightCoordinate, map[q][r]);
                        });

                        map[q][r].sprite.events.onInputOver.add(function () {
                            buttonPlus.alpha = 1;
                            buttonMinus.alpha = 1;
                        });

                        map[q][r].sprite.events.onInputOut.add(function () {
                            buttonPlus.alpha = 0;
                            buttonMinus.alpha = 0;
                            buttonPlus.events.onInputDown.removeAll();
                            buttonMinus.events.onInputDown.removeAll();
                        });
                    }
                }
            }
        }
    }, 1000)
}

function addHeightCoordinate() {
    mapEditor.send(JSON.stringify({
        event: "addHeightCoordinate",
        q: Number(this.q),
        r: Number(this.r)
    }));
}

function subtractHeightCoordinate() {
    mapEditor.send(JSON.stringify({
        event: "subtractHeightCoordinate",
        q: Number(this.q),
        r: Number(this.r)
    }));
}