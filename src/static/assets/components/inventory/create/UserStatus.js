function CreateUserStatus() {
    let userStat = document.getElementById("userStatus");

    let userIcon = document.createElement("div");
    userIcon.id = "userIcon";
    userStat.appendChild(userIcon);

    let userStatBlock = document.createElement("div");
    userStatBlock.id = "userStatBlock";

    let name = document.createElement("div");
    name.id = "inventoryUserName";
    userStatBlock.appendChild(name);

    let credits = document.createElement("div");
    credits.id = "credits";
    userStatBlock.appendChild(credits);

    let expPoints = document.createElement("div");
    expPoints.id = "expPoints";
    userStatBlock.appendChild(expPoints);

    userStat.appendChild(userStatBlock);

    let skills = document.createElement("div");
    skills.id = "skills";
    userStat.appendChild(skills);

}