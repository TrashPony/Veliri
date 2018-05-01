function NotEndGame(jsonMessage) {
    var func = function () {
        JoinToGame(this.id);
    };

    var game = Object();

    game.Name = JSON.parse(jsonMessage).name_game;
    game.Id = JSON.parse(jsonMessage).id_game;
    game.Step = JSON.parse(jsonMessage).step_game;
    game.Phase = JSON.parse(jsonMessage).phase_game;
    game.Ready = !JSON.parse(jsonMessage).ready;

    CreateGame('NotEndGame', 'SubMenu', 'Select SubMenu', JSON.parse(jsonMessage).id_game, func, null, null, game);
}

function CreateGame(gameContent, menu, className, id, func, funcMouse, funcOutMouse, game) {
    var list = document.getElementById(menu);
    var tr = document.createElement('tr');
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
    tr.map = game.Map;

    if (list && gameContent === "NotEndGame") {
        tdName.appendChild(document.createTextNode(game.Name));
        tdID.appendChild(document.createTextNode(game.Map.Name));
        tdStep.appendChild(document.createTextNode(game.Step));
        tdPhase.appendChild(document.createTextNode(game.Phase));
        tdMyStep.appendChild(document.createTextNode(game.Ready));

        tdName.className = "Value";
    }

    tr.appendChild(tdName);
    tr.appendChild(tdID);
    tr.appendChild(tdStep);
    tr.appendChild(tdPhase);
    tr.appendChild(tdMyStep);

    list.appendChild(tr);
}