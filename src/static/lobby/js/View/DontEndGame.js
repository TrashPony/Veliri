function NotEndGame(jsonMessage) {
    var games = JSON.parse(jsonMessage).dont_end_games;

    for(var i = 0; i < games.length; i ++) {
        CreateGame(games[i]);
    }
}

function CreateGame(game) {
    var list = document.getElementById('SubMenu');
    var tr = document.createElement('tr');
    var tdName = document.createElement('td');
    var tdId = document.createElement('td');
    var tdStep = document.createElement('td');
    var tdPhase = document.createElement('td');
    var tdMyStep = document.createElement('td');

    tr.className = 'Select SubMenu';
    tr.id = game.Id;
    tr.onclick = function () {
        JoinToGame(this.id);
    };

    tdName.appendChild(document.createTextNode(game.Name));
    tdId.appendChild(document.createTextNode(game.Id));
    tdStep.appendChild(document.createTextNode(game.Step));
    tdPhase.appendChild(document.createTextNode(game.Phase));
    tdMyStep.appendChild(document.createTextNode(game.Ready));

    tdName.className = "Value";

    tr.appendChild(tdName);
    tr.appendChild(tdId);
    tr.appendChild(tdStep);
    tr.appendChild(tdPhase);
    tr.appendChild(tdMyStep);

    list.appendChild(tr);
}