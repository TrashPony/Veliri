function CreateUserLine(userName, ready, respID) {
    let oldUser = document.getElementById(userName);

    if (oldUser) {
        return;
    }

    let list = document.getElementById('gameInfo');
    let tr = document.createElement('tr');
    let tdName = document.createElement('td');
    let tdRespawn = document.createElement('td');
    let tdReady = document.createElement('td');

    tr.style.wordWrap = 'break-word';
    tr.className = "User";
    tr.align = "center";
    tr.id = userName;

    tdName.appendChild(document.createTextNode(userName));
    tdName.className = "Value";

    if (ready) {
        tdReady.innerHTML = "Готов.";
        tdReady.className = "Success";
        tdRespawn.innerHTML = respID;
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