function CreatePreviewPath(jsonMessage) {
    let path = JSON.parse(jsonMessage).path;

    path.pop(); // удаляем последний элемент что бы не пометился, там своя анимация

    for (let i = 0; i < path.length; i++) {
        let cellSprite = game.map.OneLayerMap[path[i].q][path[i].r].sprite;

        let pathSprite = game.SelectRangeLayer.create(cellSprite.x, cellSprite.y, 'pathCell');
        if (cellSprite) {
            //pathSprite.animations.add('previewPath'); todo
            //pathSprite.animations.play('previewPath', 2, true);
        }
    }
}