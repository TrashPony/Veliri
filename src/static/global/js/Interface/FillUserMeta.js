function FillUserMeta(credits, experience) {
    let creditBlock = document.getElementById("Credits");
    creditBlock.innerHTML = "" +
        "<div style='position: relative'>" +
            "<span style='position: absolute;'>Кредиты: </span>" +
            "<span style='position: absolute; right: 10px; color: #00fdff'>" + credits + "</span>" +
        "</div>";
    let experienceBlock = document.getElementById("Experience");
    experienceBlock.innerHTML = "" +
        "<div style='position: relative'>" +
            "<span style='position: absolute;'>Опыт: </span>" +
            "<span style='position: absolute; right: 10px; color: #00fdff'>" + experience + "</span>" +
        "</div>";
}