function getTypesCoordinate(menuId) {

    let menu = document.getElementsByClassName('active');
    menu[0].className ="";

    if (menuId === 1) {
        document.getElementById("oneMenu").className = "active";
    }

    if (menuId === 2) {
        document.getElementById("twoMenu").className = "active";
    }

    if (menuId === 3) {
        document.getElementById("threeMenu").className = "active";
    }

    if (menuId === 4) {
        document.getElementById("fourMenu").className = "active";
    }

    mapEditor.send(JSON.stringify({
        event: "getAllTypeCoordinate"
    }));
}

function createCoordinateMenu(typeCoordinates) {
    let menu = document.getElementsByClassName('active');

    let typesCoordinate = document.getElementsByClassName('coordinateType');
    while(typesCoordinate.length > 0){
        typesCoordinate[0].parentNode.removeChild(typesCoordinate[0]);
    }

    if (menu[0]) {
        if (menu[0].dataset.idMenu === "1"){
            ViewPatternCoordinate(typeCoordinates)
        }

        if (menu[0].dataset.idMenu === "2"){
            ViewTerrainCoordinate(typeCoordinates)
        }

        if (menu[0].dataset.idMenu === "3"){
            ViewObjectsCoordinate(typeCoordinates)
        }

        if (menu[0].dataset.idMenu === "4"){
            ViewAnimateObjectsCoordinate(typeCoordinates)
        }
    }
}
