function CreateUserLine(user) {
    var oldUser = document.getElementById(user.Name);

    if (oldUser) {
        return;
    }

    var list = document.getElementById('gameInfo');
    var tr = document.createElement('tr');
    var tdName = document.createElement('td');
    var tdRespawn = document.createElement('td');
    var tdReady = document.createElement('td');

    tr.style.wordWrap = 'break-word';
    tr.className = "User";
    tr.align = "center";
    tr.id = user.Name;

    tdName.appendChild(document.createTextNode(user.Name));
    tdName.className = "Value";

    if (user.Ready) {
        tdReady.innerHTML = "Готов.";
        tdReady.className = "Success";
        tdRespawn.innerHTML = user.Respawn.Name;
    } else {
        tdReady.innerHTML = "Не готов.";
        tdReady.className = "Failed";
        tdRespawn.innerHTML = "";
    }

    tr.appendChild(tdName);
    tr.appendChild(tdRespawn);
    tr.appendChild(tdReady);

    list.appendChild(tr);
}