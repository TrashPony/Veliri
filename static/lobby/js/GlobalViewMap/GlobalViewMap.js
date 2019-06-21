function GlobalViewMap() {
    if (document.getElementById('GlobalViewMap')) {
        let jBox = $('#GlobalViewMap');
        setState('GlobalViewMap', jBox.position().left, jBox.position().top, jBox.height(), jBox.width(), false);
        return;
    }

    let GlobalViewMap = document.createElement('div');
    GlobalViewMap.id = 'GlobalViewMap';
    document.body.appendChild(GlobalViewMap);

    GlobalViewMap.innerHTML = `
        <div id="GlobalMapWrapper"></div>
    `;

    if (!allMaps) {
        chat.send(JSON.stringify({
            event: "openMapMenu",
        }));
    } else {
        FillGlobalMap(allMaps, SectorID)
    }

    let buttons = CreateControlButtons("2px", "31px", "-3px", "29px", "Карта мира", "105px");
    $(buttons.move).mousedown(function (event) {
        moveWindow(event, 'GlobalViewMap')
    });
    $(buttons.close).mousedown(function () {
        setState(GlobalViewMap.id, $(GlobalViewMap).position().left, $(GlobalViewMap).position().top, $(GlobalViewMap).height(), $(GlobalViewMap).width(), false);
    });
    GlobalViewMap.appendChild(buttons.move);
    GlobalViewMap.appendChild(buttons.close);
    GlobalViewMap.appendChild(buttons.head);

    openWindow(GlobalViewMap.id, GlobalViewMap)
}

function initCanvasMap(id) {
    const canvas = document.getElementById(id);
    const mapWrapper = document.getElementById('GlobalMapWrapper');
    canvas.width = mapWrapper.offsetWidth;
    canvas.height = mapWrapper.offsetHeight;
}