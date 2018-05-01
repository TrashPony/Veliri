function GameView(jsonMessage) {

    var game = JSON.parse(jsonMessage).game;

    var gameLine = document.getElementById(game.Name);
    if (gameLine) {
        gameLine.remove();
    }

    CreateLobyGame(game);
}

function CreateLobyGame(game) {
    var list = document.getElementById('Menu');
    if (list) {
        var tr = document.createElement('tr');
        var tdName = document.createElement('td');
        var tdID = document.createElement('td');
        var tdStep = document.createElement('td');
        var tdPhase = document.createElement('td');
        var tdMyStep = document.createElement('td');

        tr.style.wordWrap = 'break-word';
        tr.className = 'Select Menu';
        tr.align = "center";
        tr.id = game.Name;

        tr.onclick = function () {
            sendJoinToLobbyGame(this.id);
        };
        tr.onmouseover = function () {
            MouseOverMap(this.map);
        };
        tr.onmouseout = function () {
            MouseOutMap()
        };

        tr.map = game.Map;

        tdName.appendChild(document.createTextNode(game.Name));
        tdID.appendChild(document.createTextNode(game.Map.Name));
        tdStep.appendChild(document.createTextNode(game.Creator));
        tdPhase.appendChild(document.createTextNode(game.Map.Respawns));
        tdMyStep.appendChild(document.createTextNode(game.Users));
        console.log(game.Users);// todo

        tdName.className = "Value";

        tr.appendChild(tdName);
        tr.appendChild(tdID);
        tr.appendChild(tdStep);
        tr.appendChild(tdPhase);
        tr.appendChild(tdMyStep);

        list.appendChild(tr);
    }
}