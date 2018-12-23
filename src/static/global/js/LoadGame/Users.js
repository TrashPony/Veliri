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