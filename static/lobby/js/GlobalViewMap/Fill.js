let allMaps;
let SectorID;
let gridSize = 100;

function FillGlobalMap(maps, userSectorID) {
    SectorID = userSectorID;
    allMaps = maps;

    const mapWrapper = document.getElementById('GlobalMapWrapper');
    mapWrapper.innerHTML = '<canvas id="GlobalMapCanvas"></canvas><canvas id="GlobalMapPathCanvas"></canvas>';

    initCanvasMap('GlobalMapCanvas');

    let xGridMax = 0, yGridMax = 0;

    for (let i in maps) {

        let fractionIcon = '../assets/logo/' + maps[i].fraction + '.png';
        if (maps[i].fraction === '')
            fractionIcon = 'https://img.icons8.com/color/48/000000/storage.png';

        if (xGridMax < maps[i].x_global) xGridMax = maps[i].x_global;
        if (yGridMax < maps[i].y_global) yGridMax = maps[i].y_global;

        let xCell = 10 + (maps[i].x_global * gridSize);
        let yCell = 10 + (maps[i].y_global * gridSize);

        let cell = document.createElement('div');
        cell.className = 'MapPoint';
        cell.innerHTML = `
                <div class="animateAura" onmouseover="previewPath('${maps[i].id}')"></div>
                <div class="fractionIcon" style="background-image: url(' ${fractionIcon} ')"></div>
                <div class="endPoint"></div>
                <div class="sectorName">${maps[i].Name}</div>`;
        cell.style.left = xCell + 'px';
        cell.style.top = yCell + 'px';

        cell.onmouseout = function () {
            initCanvasMap('GlobalMapPathCanvas');
        };

        if (maps[i].id === userSectorID) cell.className += ' User';

        mapWrapper.appendChild(cell);

        for (let j in maps[i].handlers_coordinates) {
            if (maps[i].handlers_coordinates[j].handler === 'sector') {

                let toX = 10 + getMapByID(maps, maps[i].handlers_coordinates[j].to_map_id).x_global * gridSize;
                let toY = 10 + getMapByID(maps, maps[i].handlers_coordinates[j].to_map_id).y_global * gridSize;

                CanvasGlobalLineXY_To_XY(xCell + 20, yCell + 20, toX + 20, toY + 20);
            }
        }
    }

    // нельзя делать в цикле который сверху иначе пути будут перекривать стрелки
    for (let i in maps) {
        let xCell = 10 + (maps[i].x_global * gridSize);
        let yCell = 10 + (maps[i].y_global * gridSize);

        for (let j in maps[i].handlers_coordinates) {
            if (maps[i].handlers_coordinates[j].handler === 'sector') {
                let toX = 10 + getMapByID(maps, maps[i].handlers_coordinates[j].to_map_id).x_global * gridSize;
                let toY = 10 + getMapByID(maps, maps[i].handlers_coordinates[j].to_map_id).y_global * gridSize;
                CanvasArrowPath(xCell + 20, yCell + 20, toX + 20, toY + 20);
            }
        }
    }

    for (let i = 0; i <= xGridMax; i++) {
        for (let j = 0; j <= yGridMax; j++) {

            let xCell = (i * gridSize) - 20;
            let yCell = (j * gridSize) - 20;

            let cell = document.createElement('div');
            cell.className = 'GridBox';
            cell.style.left = xCell + 'px';
            cell.style.top = yCell + 'px';
            cell.style.height = gridSize + 'px';
            cell.style.width = gridSize + 'px';
            mapWrapper.appendChild(cell);
        }
    }
}

function previewPath(id) {
    lobby.send(JSON.stringify({
        event: "previewPath",
        id: Number(id)
    }));
}

function getMapByID(maps, id) {
    for (let i in maps) {
        if (maps[i].id === id) {
            return maps[i]
        }
    }
}

function CanvasArrowPath(startX, startY, endX, endY, headLength = 7) {

    const canvas = document.getElementById('GlobalMapCanvas');
    const ctx = canvas.getContext('2d');

    let points = [];

    let time = Math.max(Math.abs(startX - endX), Math.abs(startY - endY));//для повышения точности выбираем что сильнее изменилось, для ускорения выбор наоборот min
    for (let i = 0; i <= time; i++) {
        let delta = i / time;
        let a = delta * (endX - startX) + startX;
        let b = delta * (endY - startY) + startY;
        points.push({x: a, y: b});
    }

    let pointNumber = Math.round((points.length / 100) * 40);
    let centerPathX = points[pointNumber].x;
    let centerPathY = points[pointNumber].y;

    ctx.strokeStyle = "#02ff0e";
    ctx.lineWidth = 3;
    ctx.shadowColor = 'black';
    ctx.shadowBlur = 2;

    // constants (could be declared as globals outside this function)
    let PI = Math.PI;
    let degreesInRadians225 = 225 * PI / 180;
    let degreesInRadians135 = 135 * PI / 180;

    // calc the angle of the line
    let dx = endX - startX;
    let dy = endY - startY;
    let angle = Math.atan2(dy, dx);

    // calc arrowhead points
    let x225 = centerPathX + headLength * Math.cos(angle + degreesInRadians225);
    let y225 = centerPathY + headLength * Math.sin(angle + degreesInRadians225);
    let x135 = centerPathX + headLength * Math.cos(angle + degreesInRadians135);
    let y135 = centerPathY + headLength * Math.sin(angle + degreesInRadians135);

    ctx.beginPath();

    ctx.moveTo(points[pointNumber + 1].x, points[pointNumber + 1].y);
    ctx.lineTo(centerPathX, centerPathY);

    ctx.moveTo(centerPathX, centerPathY);
    ctx.lineTo(x225, y225);
    // draw partial arrowhead at 135 degrees
    ctx.moveTo(centerPathX, centerPathY);
    ctx.lineTo(x135, y135);
    // stroke the line and arrowhead
    ctx.stroke();
}

function CanvasGlobalLineXY_To_XY(startX, startY, endX, endY) {
    const canvas = document.getElementById('GlobalMapCanvas');
    const ctx = canvas.getContext('2d');

    ctx.strokeStyle = "rgba(127, 127, 127, 0.8)";
    ctx.lineWidth = 3;
    ctx.shadowBlur = 1;
    ctx.shadowColor = 'white';

    ctx.beginPath();
    ctx.moveTo(startX, startY);
    ctx.lineTo(endX, endY);
    ctx.stroke();
}