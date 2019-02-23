function Anomaly(squad) {
    let display = document.getElementById("anomalyDisplay");

    let displayBody = document.createElement("div");
    let weapon = "";

    for (let j in  squad.mather_ship.body.weapons) {
        if ( squad.mather_ship.body.weapons.hasOwnProperty(j) &&  squad.mather_ship.body.weapons[j].weapon) {
            weapon = "url(/assets/units/weapon/" +  squad.mather_ship.body.weapons[j].weapon.name
                + ".png) center center / 25px no-repeat";
        }
    }

    if (weapon !== "") {
        displayBody.id = "anomalyBody";
        displayBody.style.background = weapon + ", url(/assets/units/body/" + squad.mather_ship.body.name + ".png)" +
            " center center / 50px no-repeat";
    }

    display.appendChild(displayBody);
}