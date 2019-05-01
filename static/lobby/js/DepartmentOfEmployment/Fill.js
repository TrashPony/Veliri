function FillDepartment(dialogPage, action, mission) {
    console.log(dialogPage)
    DialogAction(action);

    if ((action === "close" || action === "start_training" || action === "miss_training") && document.getElementById('DepartmentOfEmployment')) {
        document.getElementById('DepartmentOfEmployment').remove();
        return;
    }

    if (!document.getElementById('DepartmentOfEmployment')) {
        InitDepartmentOfEmployment(dialogPage, action, mission);
        return
    }

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