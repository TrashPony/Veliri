function EquippingParse(jsonMessage) {

    var equipping = JSON.parse(jsonMessage).equipping;
    var tableEquip = document.getElementById("tableEquip");

    console.log(equipping);

    for (var i = 0; i < equipping.length; i++) {
        var tr = CreateEquipRow(equipping[i].type, equipping[i].id, equipping[i].specification);
        tableEquip.appendChild(tr);
    }
}
// TODO передовать в параметры выбранный слот в шипе что бы знать куда положить/заменить модуль
function CreateEquipRow(type, id, value) {
    var tr = document.createElement("tr");
    tr.className = "EquipRow";
    var tdIcon = document.createElement("td");
    tdIcon.className = "tdIconEquip";
    tdIcon.style.backgroundImage = "url(/lobby/img/" + type + ".png)";

    var tdSpecification = document.createElement("td");
    tdSpecification.innerHTML = "<spen class='Value'>" + type + "</spen><br>" + value + "<br>";

    var acceptButton = document.createElement("input");
    acceptButton.type = "button";
    acceptButton.value = "Выбрать";
    acceptButton.className = "EquipButton";
    acceptButton.id = id+":equip";
    //acceptButton.onclick = EquipBackToLobby;

    tdSpecification.appendChild(acceptButton);
    tr.appendChild(tdIcon);
    tr.appendChild(tdSpecification);

    return tr;
}