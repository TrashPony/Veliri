function CreateUser(squad) {
    let x = squad.global_x;
    let y = squad.global_y;
    let weaponName;

    for (let i in squad.mather_ship.body.weapons) {
        if (squad.mather_ship.body.weapons.hasOwnProperty(i) && squad.mather_ship.body.weapons[i].weapon) {
            weaponName = squad.mather_ship.body.weapons[i].weapon.name
        }
    }

    CreateSquad(squad, x, y, squad.mather_ship.body.name, weaponName, squad.mather_ship.rotate + 90);
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
    if (!sprite) CreateSquad(otherUser, x, y, otherUser.body_name, otherUser.weapon_name, otherUser.rotate + 90);
}