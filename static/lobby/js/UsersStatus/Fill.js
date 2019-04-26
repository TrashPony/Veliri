function FillUserStatus(player) {
    userStat = player;

    if (document.getElementById("userAvatar"))
        document.getElementById('userAvatar').style.backgroundImage = "url(" + player.avatar_icon + ")";

    if (document.getElementById("userBiography"))
        document.getElementById('userBiography').innerHTML = player.biography;

    if (document.getElementById("userName"))
        document.getElementById('userName').innerHTML = player.login;

    if (document.getElementById('userTitle'))
        document.getElementById('userTitle').innerHTML = player.title;

    if (document.getElementById('scientific_points_points'))
        document.getElementById('scientific_points_points').innerHTML = player.scientific_points;

    if (document.getElementById('attack_points_points'))
        document.getElementById('attack_points_points').innerHTML = player.attack_points;

    if (document.getElementById('production_points_points'))
        document.getElementById('production_points_points').innerHTML = player.production_points;

}