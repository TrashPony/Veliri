function GameView(jsonMessage) {

    let game = JSON.parse(jsonMessage).game;

    let gameLine = document.getElementById(game.Name);
    if (gameLine) {
        gameLine.remove();
    }

    CreateLobyGame(game);
}

function CreateLobyGame(game) {
    let list = document.getElementById('Menu');
    if (list) {
        let tr = document.createElement('tr');
        let tdName = document.createElement('td');
        let tdID = document.createElement('td');
        let tdStep = document.createElement('td');
        let tdPhase = document.createElement('td');
        let tdMyStep = document.createElement('td');

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

        let countUsers = 0;

        for (let key in game.Users) {
            countUsers++;
        }

        tdName.appendChild(document.createTextNode(game.Name));
        tdID.appendChild(document.createTextNode(game.Map.Name));
        tdStep.appendChild(document.createTextNode(game.Creator));
        tdPhase.appendChild(document.createTextNode(game.Map.Respawns));
        tdMyStep.appendChild(document.createTextNode(countUsers));

        tdName.className = "Value";

        tr.appendChild(tdName);
        tr.appendChild(tdID);
        tr.appendChild(tdStep);
        tr.appendChild(tdPhase);
        tr.appendChild(tdMyStep);

        list.appendChild(tr);
    }
}