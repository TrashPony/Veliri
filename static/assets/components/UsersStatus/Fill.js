function FillUserStatus(player, skill, userID) {

    userStat = player;

    if (document.getElementById("userAvatar")) {
        GetUserAvatar(userID).then(function (response) {
            document.getElementById('userAvatar').style.backgroundImage = "url('" + response.data.avatar + "')";
        });
    }

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

    if (document.getElementById('ScySkills')) {
        let ScySkills = document.getElementById('ScySkills');
        fillSkills(ScySkills, "scientific", player)
    }

    if (document.getElementById('AttackSkills')) {
        let AttackSkills = document.getElementById('AttackSkills');
        fillSkills(AttackSkills, "attack", player)
    }

    if (document.getElementById('IndustrySkills')) {
        let IndustrySkills = document.getElementById('IndustrySkills');
        fillSkills(IndustrySkills, "production", player)
    }

    if (skill) {
        let experiencePoint;
        if (skill.level > 0) experiencePoint = 200;
        if (skill.level > 1) experiencePoint = 400;
        if (skill.level > 2) experiencePoint = 800;
        if (skill.level > 3) experiencePoint = 1600;
        if (skill.level > 4) experiencePoint = '-';

        SelectSkill(skill.id, skill.level, skill.name, skill.specification, skill.icon, experiencePoint, skill.type)
    }
}

function fillSkills(skillBlock, typeSkill, player) {
    skillBlock.innerHTML = ``;

    for (let i in player.current_skills) {
        let skill = player.current_skills[i];

        if (skill.type === typeSkill) {

            let experiencePoint = 100;

            let back1, back2, back3, back4, back5;

            if (skill.level > 0) {
                back1 = "#00f1ff;";
                experiencePoint = 200;
            }
            if (skill.level > 1) {
                back2 = "#00f1ff;";
                experiencePoint = 400;
            }

            if (skill.level > 2) {
                back3 = "#00f1ff;";
                experiencePoint = 800;
            }

            if (skill.level > 3) {
                back4 = "#00f1ff;";
                experiencePoint = 1600;
            }

            if (skill.level > 4) {
                back5 = "#00f1ff;";
                experiencePoint = '-'
            }

            skillBlock.innerHTML += (
                `<div class="skill" style="background-image: url('${skill.icon}')"
                        onclick="SelectSkill('${skill.id}', '${skill.level}', '${skill.name}', '${skill.specification}', '${skill.icon}', '${experiencePoint}', '${skill.type}')">
                     <div class="skillPrice">${experiencePoint}</div>
                       <div class="skillLvl">
                           <div style="background: ` + back1 + `"></div>
                           <div style="background: ` + back2 + `"></div>
                           <div style="background: ` + back3 + `"></div>
                           <div style="background: ` + back4 + `"></div>
                           <div style="background: ` + back5 + `"></div>
                       </div>
                     </div>`)
        }
    }
}

function SelectSkill(id, level, name, specification, icon, needPrice, type) {

    document.getElementById('skillName').innerHTML = name;
    document.getElementById('skillDescription').innerHTML = specification;
    document.getElementById('needPrice').innerHTML = needPrice;

    if (type === 'scientific') {
        document.getElementById('needPrice').style.color = 'cornflowerblue';
    } else if (type === 'attack') {
        document.getElementById('needPrice').style.color = 'crimson';
    } else {
        document.getElementById('needPrice').style.color = 'chartreuse';
    }

    document.getElementById('skillIcon').style.backgroundImage = 'url(' + icon + ')';

    let back1, back2, back3, back4, back5;
    if (level > 0) back1 = "#00f1ff;";
    if (level > 1) back2 = "#00f1ff;";
    if (level > 2) back3 = "#00f1ff;";
    if (level > 3) back4 = "#00f1ff;";
    if (level > 4) back5 = "#00f1ff;";

    document.getElementById('skillLvl').innerHTML = (
        `<div style="background: ` + back1 + `"></div>
         <div style="background: ` + back2 + `"></div>
         <div style="background: ` + back3 + `"></div>
         <div style="background: ` + back4 + `"></div>
         <div style="background: ` + back5 + `"></div>`);

    document.getElementById('upperSkill').onclick = function () {
        chat.send(JSON.stringify({
            event: "upSkill",
            id: Number(id),
        }));
    }
}

function FillOtherUserStat(stat) {
    if (document.getElementById("userName"))
        document.getElementById('userName').innerHTML = stat.user_name;

    if (document.getElementById("userAvatar")) {
        GetUserAvatar(stat.user_id).then(function (response) {
            document.getElementById('userAvatar').style.backgroundImage = "url('" + response.data.avatar + "')";
        });
    }

    if (document.getElementById("userBiography"))
        document.getElementById('userBiography').innerHTML = stat.biography;

    if (document.getElementById('userTitle'))
        document.getElementById('userTitle').innerHTML = stat.title;
}