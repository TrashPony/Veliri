function GameView(jsonMessage) {
    var func = function () {
        sendJoinToLobbyGame(this.id);
    };

    var onmouse = function () {
        MouseOverMap(this.map);
    };

    var offmouse= function () {
        MouseOutMap()
    };

    var newGame = Object();

    newGame.Name = JSON.parse(jsonMessage).name_game;
    newGame.Map  = JSON.parse(jsonMessage).map;
    newGame.Creator = JSON.parse(jsonMessage).creator;
    newGame.Copasity = JSON.parse(jsonMessage).players;
    newGame.Players = JSON.parse(jsonMessage).num_of_players;

    var gameLine = document.getElementById(JSON.parse(jsonMessage).name_game);
    if (gameLine) {
        gameLine.remove();
    }

    CreateGame('Game', 'Menu', 'Select Menu', JSON.parse(jsonMessage).name_game, func, onmouse, offmouse, newGame);
}

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

    if (list && gameContent === "Game") {
        tdName.appendChild(document.createTextNode(game.Name));
        tdID.appendChild(document.createTextNode(game.Map.Name));
        tdStep.appendChild(document.createTextNode(game.Creator));
        tdPhase.appendChild(document.createTextNode(game.Players));
        tdMyStep.appendChild(document.createTextNode(game.Copasity));

        tdName.className = "Value";
    }

    tr.appendChild(tdName);
    tr.appendChild(tdID);
    tr.appendChild(tdStep);
    tr.appendChild(tdPhase);
    tr.appendChild(tdMyStep);

    list.appendChild(tr);
}