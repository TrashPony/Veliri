let game;

function LoadGame(jsonData) {
    game = CreateGame(jsonData.map);
    game.typeService = "global";

    setTimeout(function () { // todo костыль связаной с прогрузкой карты )
        CreateUser(jsonData.squad);
        game.input.onDown.add(initMove, game);
        CreateBase(jsonData.bases);
        CreateOtherUsers(jsonData.other_users);
        CreateMiniMap(jsonData.map);
    }, 1500);
}

function CreateUser(squad) {
    let x = squad.global_x;
    let y = squad.global_y;
    let weaponName;

    for (let i in squad.mather_ship.body.weapons) {
        if (squad.mather_ship.body.weapons.hasOwnProperty(i) && squad.mather_ship.body.weapons[i].weapon) {
            weaponName = squad.mather_ship.body.weapons[i].weapon.name
        }
    }

    CreateSquad(squad, x, y, squad.mather_ship.body.name, weaponName, squad.mather_ship.rotate + 90, true);
    game.squad = squad;
}

function CreateOtherUsers(otherUsers) {
    game.otherUsers = [];
    for (let i = 0; i < otherUsers.length; i++){
        CreateOtherUser(otherUsers[i])
    }
}

function CreateOtherUser(otherUser) {
    let x = otherUser.x;
    let y = otherUser.y;
    CreateSquad(otherUser, x, y, otherUser.body_name, otherUser.weapon_name,otherUser.rotate + 90);
    game.otherUsers.push(otherUser)
}

function CreateBase(bases) {
    for (let i in bases) {
        if (bases.hasOwnProperty(i)) {
            if (game.map.OneLayerMap.hasOwnProperty(bases[i].q) && game.map.OneLayerMap.hasOwnProperty(bases[i].r)) {
                let coordinate = game.map.OneLayerMap[bases[i].q][bases[i].r];

                coordinate.objectSprite.inputEnabled = true;
                coordinate.objectSprite.input.pixelPerfectOver = true;
                coordinate.objectSprite.input.pixelPerfectClick = true;

                coordinate.objectSprite.events.onInputDown.add(function () {
                    if (game.input.activePointer.leftButton.isDown) {
                        game.squad.toBase = {
                            baseID: bases[i].id,
                            into: true,
                            x: coordinate.sprite.x,
                            y: coordinate.sprite.y
                        }
                    }
                });
            }
        }
    }
}