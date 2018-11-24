function CreateNewCoordinate() {

    let formData = new FormData(document.forms.uploadNewCoordinate);

    let terrainName;
    let objectName;
    let animateName;

    if (formData.get("terrainTexture").name !== "") {
        terrainName = formData.get("terrainTexture").name;
    }

    if (formData.get("objectTexture").name !== "") {
        objectName = formData.get("objectTexture").name;
    }

    if (formData.get("animateSprite").name !== "") {
        animateName = formData.get("animateSprite").name;
    }

    if (formData.get("objectTexture").name !== "" && formData.get("animateSprite").name !== "") {
        alert("Нельзя одновременно выбрать Текустуру обьекта и Анимацию спрайта");
        return
    }

    let Move = false;
    let Watch = false;
    let Attack = false;

    if (formData.get("move")) {
        Move = true;
    }

    if (formData.get("watch")) {
        Watch = true;
    }

    if (formData.get("attack")) {
        Attack = true;
    }

    let Radius = formData.get("radius");

    mapEditor.send(JSON.stringify({
        event: "loadNewTypeCoordinate",
        terrain_name: terrainName,
        object_name: objectName,
        animate_name: animateName,
        move: Move,
        watch: Watch,
        attack: Attack,
        radius: Number(Radius)
    }));
}

function sendFiles() {
    // todo вызвать этот метод только когда по сокетам вернулось что все окей
    // загрузка файлов новой координаты
    let formData = new FormData(document.forms.uploadNewCoordinate);
    let xhr = new XMLHttpRequest();
    xhr.open("POST", "http://642e0559eb9c.sn.mynetname.net:8080/upload");
    xhr.send(formData);
}