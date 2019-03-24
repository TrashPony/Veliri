function CreatePreviewPath(jsonMessage) {
    while (game.PreviewPath && game.PreviewPath.children.length > 0) {
        let sprite = game.PreviewPath.children.shift();
        sprite.destroy();
    }

    let path = JSON.parse(jsonMessage).path;

    path.pop(); // удаляем последний элемент что бы не пометился, там своя анимация

    for (let i = 0; i < path.length; i++) {
        let xy = GetXYCenterHex(path[i].q, path[i].r);
        let pathNode = game.PreviewPath.create(xy.x, xy.y, 'pathCell');
        pathNode.anchor.setTo(0.5);
        pathNode.scale.set(0.5);
    }
}