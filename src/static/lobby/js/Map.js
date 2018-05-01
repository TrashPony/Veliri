function MapView(jsonMessage) {
    var map = JSON.parse(jsonMessage).map;

    var mapLine = document.getElementById(map.Name);

    if (mapLine) {
        mapLine.remove();
    }

    CreateMapLine(map);
}

function CreateMapLine(map) {
    var list = document.getElementById('SubMenu');
    var tr = document.createElement('tr');
    var tdName = document.createElement('td');
    var tdPhase = document.createElement('td');

    tr.style.wordWrap = 'break-word';
    tr.className = 'Select SubMenu';
    tr.align = "center";
    tr.id = map.Id;

    tr.onclick = function () {
        CreateLobbyGame(this.id);
    };

    tr.onmouseover = function () {
        MouseOverMap(this.id);
    };

    tr.onmouseout = function () {
        MouseOutMap()
    };

    tdName.appendChild(document.createTextNode(map.Name));
    tdPhase.appendChild(document.createTextNode(map.Respawns));

    tdName.className = "Value";
    tr.map = map;

    tr.appendChild(tdName);
    tr.appendChild(tdPhase);

    list.appendChild(tr);
}

function MouseOverMap(id) {
    var map = document.getElementById(id).map;

    var info = document.getElementById('SelectInfo');
    info.innerHTML = "<span class='Value'>" + map.Name + "</span>";

    var div = document.createElement('div');
    div.style.wordWrap = 'break-word';
    div.className = "infoMap";
    div.style.backgroundImage = "url(/assets/" + map.Name + ".png)";
    div.id = "infoImage";
    info.appendChild(div);

    var div2 = document.createElement('div');
    div2.style.wordWrap = 'break-word';
    div2.className = "infoMap";
    div2.innerHTML = "<span class='Value'>Описание: </span> <br>" + map.Specification;

    info.appendChild(div2);
}

function MouseOutMap() {
    var info = document.getElementById('SelectInfo');
    info.innerHTML = "";
    DelElements("infoMap");
}