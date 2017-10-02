var SizeUnit = 100;
function SizeMap(params) {
    var div = document.getElementsByClassName("fieldUnit");
    if (params === 1) SizeUnit = SizeUnit + 30;
    if (SizeUnit > 45) {
        if (params === 2) SizeUnit = SizeUnit - 30;
    }

    for (var i = 0; 0 < div.length; i++) {
        if (params === 1) {
            div[i].style.height = SizeUnit + "px";
            div[i].style.width = SizeUnit + "px";
        }

        if (params === 2) {
            div[i].style.height = SizeUnit + "px";
            div[i].style.width = SizeUnit + "px";
        }

    }
}
function Rotate(params) {
    var div = document.getElementById('main');
    if(params === 0) {
        div.style.transition = "5s all";
        div.style.boxShadow = "25px 25px 20px  rgba(0,0,0,0.5)";
        div.style.transform = "rotateX(13deg) translate(0px, -250px) rotate(0deg)";
    }
    if(params === 90) {
        div.style.transition = "5s all";
        div.style.boxShadow = "25px -25px 20px  rgba(0,0,0,0.5)";
        div.style.transform = "rotateX(13deg) translate(0px, -250px) rotate(90deg)";
    }
    if(params === 180) {
        div.style.transition = "5s all";
        div.style.boxShadow = "-25px -25px 20px  rgba(0,0,0,0.5)";
        div.style.transform = "rotateX(13deg) translate(0px, -250px) rotate(180deg)";
    }
    if(params === 270) {
        div.style.transition = "5s all";
        div.style.boxShadow = "-25px 25px 20px  rgba(0,0,0,0.5)";
        div.style.transform = "rotateX(13deg) translate(0px, -250px) rotate(270deg)";
    }
}
