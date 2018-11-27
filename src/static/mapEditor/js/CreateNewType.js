function CreateNewTerrain() {

    let formData = new FormData(document.forms.uploadNewTerrain);

    let terrainName;

    if (formData.get("terrainTexture").name !== "") {
        terrainName = formData.get("terrainTexture").name;
    }

    mapEditor.send(JSON.stringify({
        event: "loadNewTypeTerrain",
        terrain_name: terrainName.substr(0, terrainName.lastIndexOf('.')) || terrainName
    }));
}

function CreateNewObject() {
    let formData = new FormData(document.forms.uploadNewObject);

    let objectName = "";
    let animateName = "";

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
        event: "loadNewTypeObject",
        object_name: objectName.substr(0, objectName.lastIndexOf('.')) || objectName,
        animate_name: animateName.substr(0, animateName.lastIndexOf('.')) || animateName,
        move: Move,
        watch: Watch,
        attack: Attack,
        radius: Number(Radius)
    }));
}