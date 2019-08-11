function GetUserAvatar(userID) {
    return axios.get('/avatar?user_id=' + userID);
}

function GetGroupAvatar(groupID) {
    return axios.get('/chat_group_avatar?chat_group_id=' + groupID);
}