function GlobalViewMap() {
    if (document.getElementById('GlobalViewMap')) {
        document.getElementById('GlobalViewMap').remove();
        return;
    }

    let GlobalViewMap = document.createElement('div');
    GlobalViewMap.id = 'GlobalViewMap';
    document.body.appendChild(GlobalViewMap);

    GlobalViewMap.innerHTML = `
        <div id="GlobalMapWrapper"></div>
    `;

    if (!allMaps) {
        lobby.send(JSON.stringify({
            event: "openMapMenu",
        }));
    } else {
        FillGlobalMap(allMaps, SectorID)
    }

    let buttons = CreateControlButtons("2px", "31px", "-3px", "29px", "", "145px");
    $(buttons.move).mousedown(function (event) {
        moveWindow(event, 'GlobalViewMap')
    });
    $(buttons.close).mousedown(function () {
        GlobalViewMap.remove();
    });
    GlobalViewMap.appendChild(buttons.move);
    GlobalViewMap.appendChild(buttons.close);
}

function initCanvasMap(id) {
    const canvas = document.getElementById(id);
    const mapWrapper = document.getElementById('GlobalMapWrapper');
    canvas.width = mapWrapper.offsetWidth;
    canvas.height = mapWrapper.offsetHeight;
}