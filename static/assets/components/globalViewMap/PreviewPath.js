function PreviewMapPath(path) {

    initCanvasMap('GlobalMapPathCanvas');

    for (let i in path) {
        let xCell = 10 + (path[i].Map.x_global * gridSize);
        let yCell = 10 + (path[i].Map.y_global * gridSize);

        if (path.hasOwnProperty(Number(i) + 1)) {
            let toX = 10 + (path[Number(i) + 1].Map.x_global * gridSize);
            let toY = 10 + (path[Number(i) + 1].Map.y_global * gridSize);
            CanvasGlobalPathXY_To_XY(xCell + 20, yCell + 20, toX + 20, toY + 20)
        }
    }
}

function CanvasGlobalPathXY_To_XY(startX, startY, endX, endY) {
    const canvas = document.getElementById('GlobalMapPathCanvas');
    const ctx = canvas.getContext('2d');

    ctx.strokeStyle = "rgba(0, 255, 240, 0.9)";
    ctx.lineWidth = 3;
    ctx.shadowBlur = 1;
    ctx.shadowColor = 'black';

    ctx.beginPath();
    ctx.moveTo(startX, startY);
    ctx.lineTo(endX, endY);
    ctx.stroke();
}