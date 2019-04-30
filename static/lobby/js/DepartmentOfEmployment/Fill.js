function FillDepartment(dialogPage, action, mission) {
    console.log(dialogPage, action);

    if (action === "close" && document.getElementById('DepartmentOfEmployment')) {
        document.getElementById('DepartmentOfEmployment').remove();
        return;
    }

    if (!document.getElementById('DepartmentOfEmployment')) {
        InitDepartmentOfEmployment(dialogPage, action, mission);
        return
    }

    DialogAction(action);

    document.getElementById('missionText').innerHTML = dialogPage.text;
    document.getElementById('missionHead').innerHTML = dialogPage.name;
    document.getElementById('missionAsc').innerHTML = '';

    if (mission) {
        document.getElementById('rewardBlock2').style.visibility = "visible";
        document.getElementById('countRewardCredits').innerHTML = mission.reward_cr;

    } else {
        document.getElementById('rewardBlock2').style.visibility = "hidden";
    }

    CreateAsk(document.getElementById('missionAsc'), dialogPage, false)
}