function FillUserMeta(credits, experience, squad) {
    let creditBlock = document.getElementById("Credits");
    if (creditBlock) {
        creditBlock.innerHTML = "" +
            "<div style='position: relative'>" +
            "<span style='position: absolute;'>Кредиты: </span>" +
            "<span style='position: absolute; right: 10px; color: #00fdff'>" + credits + "</span>" +
            "</div>";
    }

    let experienceBlock = document.getElementById("Experience");
    if (experienceBlock) {
        experienceBlock.innerHTML = "" +
            "<div style='position: relative'>" +
            "<span style='position: absolute;'>Опыт: </span>" +
            "<span style='position: absolute; right: 10px; color: #00fdff'>" + experience + "</span>" +
            "</div>";
    }
}

function ChangeGravity(high_gravity) {
    let gravity = document.getElementById("lowGravity");

    if (!high_gravity) {
        gravity.innerHTML = "LOW GRAVITY";
        gravity.style.visibility = "visible";
        gravity.style.color = "#bdbd00";
    } else {
        gravity.innerHTML = "High GRAVITY";
        gravity.style.visibility = "visible";
        gravity.style.color = "#BD2D20";
        setTimeout(function () {
            gravity.style.visibility = "hidden";
        }, 2000)
    }
}