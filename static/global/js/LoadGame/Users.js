function CreateUser(squad) {
    let x = squad.global_x;
    let y = squad.global_y;
    let weapon;

    for (let i in squad.mather_ship.body.weapons) {
        if (squad.mather_ship.body.weapons.hasOwnProperty(i) && squad.mather_ship.body.weapons[i].weapon) {
            weapon = squad.mather_ship.body.weapons[i]
        }
    }

    CreateSquad(squad, x, y, squad.mather_ship.body, weapon, squad.mather_ship.rotate,
        squad.mather_ship.body_color_1, squad.mather_ship.body_color_2,
        squad.mather_ship.weapon_color_1, squad.mather_ship.weapon_color_2);

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

    let weapon;
    for (let i in otherUser.body.weapons) {
        if (otherUser.body.weapons.hasOwnProperty(i) && otherUser.body.weapons[i].weapon) {
            weapon = otherUser.body.weapons[i]
        }
    }

    if (!find) game.otherUsers.push(otherUser);
    if (!sprite) {
        CreateSquad(otherUser, x, y, otherUser.body, weapon, otherUser.rotate,
            otherUser.body_color_1, otherUser.body_color_2,
            otherUser.weapon_color_1, otherUser.weapon_color_2);
        if (game && debug) {
            CreateOtherMSGeo(otherUser);
        }
    }
}