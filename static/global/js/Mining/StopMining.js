function StopMining(jsonData) {
    if (game.squad && Number(jsonData.other_user.squad_id) === game.squad.id) {
        for (let i in game.squad.miningLaser) {
            if (game.squad.miningLaser[i] && game.squad.miningLaser[i].id === "miningEquip" + jsonData.type_slot + "" + jsonData.slot) {
                game.squad.miningLaser[i].out.destroy();
                game.squad.miningLaser[i].in.destroy();
                game.squad.miningLaser[i] = null;
            }
        }
    } else {
        for (let i = 0; game.otherUsers && i < game.otherUsers.length; i++) {
            if (game.otherUsers[i].user_name === jsonData.other_user.user_name) {

                for (let j in game.otherUsers[i].miningLaser) {
                    if (game.otherUsers[i].miningLaser[j] && game.otherUsers[i].miningLaser[j].id === game.otherUsers[i].user_name + "miningEquip" + jsonData.type_slot + "" + jsonData.slot) {
                        game.otherUsers[i].miningLaser[j].out.destroy();
                        game.otherUsers[i].miningLaser[j].in.destroy();
                        game.otherUsers[i].miningLaser[j] = null;
                    }
                }
            }
        }
    }
}