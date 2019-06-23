function GetUserAvatar(userID) {
    return axios.get('/avatar?user_id=' + userID);
}