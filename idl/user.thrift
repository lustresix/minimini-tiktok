namespace go userservice

// 用户结构体
struct User {
    1: i64 id // 用户唯一标识
    2: string name // 用户名
    3: i64 follow_count // 关注数
    4: i64 follower_count // 被关注数
    5: bool is_follow // 是否关注
}

// 用户注册请求
struct UserRegisterReq {
    1: string username
    2: string password
}

// 用户注册响应
struct UserRegisterResp {
    1: i64 status_code
    2: string status_msg
    3: i64 user_id
    4: string token
}


// 用户登录请求
struct UserLoginReq {
    1: string username
    2: string password
}

// 用户登录响应
struct UserLoginResp {
    1: i64 status_code
    2: string status_msg
    3: i64 user_id
    4: string token
}

// 用户粉丝列表请求
struct RelationFollowerListReq {
    1: string user_id
    2: string token
}

// 用户粉丝列表响应
struct RelationFollowerListResp {
    1: i64 status_code
    2: string status_msg
    3: list<User> user_list
}

// 用户好友列表请求
struct RelationFriendListReq {
    1: string user_id
    2: string token
}

// 用户好友列表响应
struct RelationFriendListResp {
    1: i64 status_code
    2: string status_msg
    3: list<User> user_list
}

service Userservice {
    UserRegisterResp Register(1: UserRegisterReq req) // 用户注册
    UserLoginResp Login(1: UserLoginReq req) // 用户登录
    RelationFollowerListResp FollowerList(1: RelationFollowerListReq req) // 用户粉丝列表
    RelationFriendListResp FriendList(1: RelationFriendListReq req) // 用户好友列表
}