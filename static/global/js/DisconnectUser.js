function DisconnectUser(jsonData) {
    for (let i = 0; game.otherUsers && i < game.otherUsers.length; i++) {
        if (game.otherUsers[i].user_name === jsonData.other_user.user_name) {
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