function CreateUserLine(user) {
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
    tdReady.appendChild(document.createTextNode(user.Ready));

    tdName.className = "Value";

    if (user.Ready !== " Готов") {
        tdReady.className = "Failed";
    } else {
        tdReady.className = "Success";
    }

    tr.appendChild(tdName);
    tr.appendChild(tdRespawn);
    tr.appendChild(tdReady);

    list.appendChild(tr);
}