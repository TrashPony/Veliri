function CreateUnits(units) {
    for (let i in units) {
        if (units.hasOwnProperty(i)) {
            console.log(units[i]);
            game.units[i] = CreateUnit(
                units[i],
                units[i].x,
                units[i].y,
                units[i].rotate,
                units[i].body_color_1,
                units[i].body_color_2,
                units[i].weapon_color_1,
                units[i].weapon_color_2,
                units[i].owner_id,
                'MySelectUnit',
                false,
            );
        }
    }
}

function CreateUser(squad) {

    let x = squad.mather_ship.global_x;
    let y = squad.mather_ship.global_y;

    squad.body = squad.mather_ship.body;
    CreateUnit(squad, x, y, squad.mather_ship.rotate,
        squad.mather_ship.body_color_1, squad.mather_ship.body_color_2,
        squad.mather_ship.weapon_color_1, squad.mather_ship.weapon_color_2, squad.user_id, 'MySelectUnit', false);

    if (game && debug) {
        CreateMSGeo(squad);
    }
    game.squad = squad;
}

function CreateOtherUsers(otherUsers) {
    if (!game.otherUsers) game.otherUsers = [];
    for (let i = 0; i < otherUsers.length; i++) { // создаем новых
        CreateOtherUser(otherUsers[i])
    }

    for (let i = 0; i < game.otherUsers.length; i++) { // докидываем тех кто долетел до загрузки и не смог создатся т.к. небыло группы
        CreateOtherUser(game.otherUsers[i])
    }
}

function CreateOtherUser(otherUser) {
    let x = otherUser.x;
    let y = otherUser.y;

    // куда ж без пары костылей
    if (!game) return;
    if (game.squad && Number(otherUser.squad_id) === game.squad.id) return;

    if (!game.otherUsers) game.otherUsers = [];
    let find = false;
    let sprite = false;

    for (let i = 0; i < game.otherUsers.length; i++) {
        if (game.otherUsers[i].squad_id === otherUser.squad_id) {
            find = true;
            if (game.otherUsers[i].sprite !== undefined) {
                sprite = true
            }
        }
    }

    if (!find) game.otherUsers.push(otherUser);
    if (!sprite) {
        CreateUnit(otherUser, x, y, otherUser.rotate,
            otherUser.body_color_1, otherUser.body_color_2,
            otherUser.weapon_color_1, otherUser.weapon_color_2, otherUser.user_id, 'MySelectUnit', false);
        if (game && debug) {
            CreateOtherMSGeo(otherUser);
        }
    }
}