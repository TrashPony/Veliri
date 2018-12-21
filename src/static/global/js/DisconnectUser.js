function DisconnectUser(jsonData) {
    for (let i = 0; i < game.otherUsers.length; i++) {
        if (game.otherUsers[i].user_name === jsonData.other_user.user_name) {
            game.otherUsers[i].sprite.destroy();
        }
    }
}