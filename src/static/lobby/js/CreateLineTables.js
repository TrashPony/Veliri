function CreateLobbyLine(gameContent, menu, className, id, func, funcMouse, funcOutMouse, text, owned) {
    var list = document.getElementById(menu);
    var tr = document.createElement('tr');
        // todo это немного пиздец надо переделать)
    var tdName = document.createElement('td');
    var tdID = document.createElement('td');
    var tdStep = document.createElement('td');
    var tdPhase = document.createElement('td');
    var tdMyStep = document.createElement('td');

    tr.style.wordWrap = 'break-word';
    tr.className = className;
    tr.align = "center";
    tr.id = id;
    tr.onclick = func;
    tr.onmouseover = funcMouse;
    tr.onmouseout = funcOutMouse;

    if (list && gameContent === "User") {

        tdName.appendChild(document.createTextNode(text.Name));
        tdStep.appendChild(document.createTextNode(text.Respawn));
        tdPhase.appendChild(document.createTextNode(text.Ready));

        tdName.className = "Value";

        if (text.Ready !== " Готов") {
            tdPhase.className = "Failed";
        } else {
            tdPhase.className = "Success";
        }

        tr.appendChild(tdName);
        tr.appendChild(tdStep);
        tr.appendChild(tdPhase);

        list.appendChild(tr);

        if (id === owned && text.Respawn === "") {
            CreateSelectRespawn(id);
            Respawn();
        }
    }
}

function CreateSelectRespawn(id) {
    var user = document.getElementById(id).cells[1];
    var selectList = document.createElement("select");
    selectList.id = "RespawnSelect";
    selectList.className = "RespawnSelect";
    user.appendChild(selectList);
}