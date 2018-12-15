function CreatePreviewPath(jsonMessage) {
    let path = JSON.parse(jsonMessage).path;

    path.pop(); // удаляем последний элемент что бы не пометился, там своя анимация

    for (let i = 0; i < path.length; i++) {
        let cellSprite = game.map.OneLayerMap[path[i].q][path[i].r].sprite;
        game.SelectRangeLayer.create(cellSprite.x, cellSprite.y, 'pathCell').anchor.setTo(0.5);
    }
}