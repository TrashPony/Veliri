function AttackUnit(jsonMessage) {
    var attackX = JSON.parse(jsonMessage).x;
    var attackY = JSON.parse(jsonMessage).y;
    var toX = JSON.parse(jsonMessage).to_x;
    var toY = JSON.parse(jsonMessage).to_y;

    var attackID = attackX + ":" + attackY;
    var cell = document.getElementById(attackID);
    cell.innerHTML = "ПЫЩЬ1";

    var targetID = toX + ":" + toY;
    var targeCell = document.getElementById(targetID);
    targeCell.innerHTML = "БдфЩь!";
}