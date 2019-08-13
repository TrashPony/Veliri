function GetUserAvatar(userID) {
    return axios.get('/avatar?user_id=' + userID);
}

function GetGroupAvatar(groupID) {
    return axios.get('/chat_group_avatar?chat_group_id=' + groupID);
}

function GetDialogPicture(dialog_page_id, userID) {
    return axios.get('/get_picture_dialog?dialog_page_id=' + dialog_page_id + "&user_id=" + userID);
}