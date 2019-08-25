function FillDepartment(dialogPage, action, mission, userID) {
    DialogAction(action);

    if ((action === "close" || action === "start_training" || action === "miss_training") && document.getElementById('DepartmentOfEmployment')) {
        let doe = document.getElementById('DepartmentOfEmployment');
        setState(doe.id, $(doe).position().left, $(doe).position().top, $(doe).height(), $(doe).width(), false);
        return;
    }

    if (!document.getElementById('DepartmentOfEmployment')) {
        InitDepartmentOfEmployment(dialogPage, action, mission, userID);
        return
    }

    document.getElementById('missionText').innerHTML = dialogPage.text;
    document.getElementById('missionHead').innerHTML = dialogPage.name;
    document.getElementById('missionAsc').innerHTML = '';

    GetDialogPicture(dialogPage.id, userID).then(function (response) {
        if (document.getElementById("missionFace")) document.getElementById("missionFace").style.backgroundImage = "url('" + response.data.picture + "')";
    });

    if (mission) {
        document.getElementById('rewardBlock2').style.visibility = "visible";
        document.getElementById('countRewardCredits').innerHTML = mission.reward_cr;
    } else {
        document.getElementById('rewardBlock2').style.visibility = "hidden";
    }

    CreateAsk(document.getElementById('missionAsc'), dialogPage, false)
}