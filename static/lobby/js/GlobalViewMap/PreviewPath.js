function PreviewPath(path) {
    console.log(path)
    initCanvasMap('GlobalMapPathCanvas');

    for (let i in path) {
        let xCell = 10 + (path[i].Map.x_global * 60);
        let yCell = 10 + (path[i].Map.y_global * 60);

        if (path.hasOwnProperty(Number(i) + 1)) {
            let toX = 10 + (path[Number(i) + 1].Map.x_global * 60);
            let toY = 10 + (path[Number(i) + 1].Map.y_global * 60);
            CanvasGlobalPathXY_To_XY(xCell + 20, yCell + 20, toX + 20, toY + 20)
        }
    }
}

function CanvasGlobalPathXY_To_XY(startX, startY, endX, endY) {
    const canvas = document.getElementById('GlobalMapPathCanvas');
    const ctx = canvas.getContext('2d');

    ctx.strokeStyle = "#00f1f9";
    ctx.lineWidth = 3;
    ctx.shadowBlur = 1;
    ctx.shadowColor = 'black';

    ctx.beginPath();
    ctx.moveTo(startX, startY);
    ctx.lineTo(endX, endY);
    ctx.stroke();
}