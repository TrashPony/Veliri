function DisconnectUser(jsonData) {
    console.log(jsonData)
    if (!game.otherUsers) return;

    for (let i = 0; game.otherUsers && i < game.otherUsers.length; i++) {
        if (game.otherUsers[i].squad_id === jsonData.other_user.squad_id) {
            if (game.otherUsers[i].sprite) {
                game.otherUsers[i].sprite.destroy();
            }
            if (game.otherUsers[i].colision) {
                game.otherUsers[i].colision.destroy();
            }
            game.otherUsers.splice(i, 1);
        }
    }
}