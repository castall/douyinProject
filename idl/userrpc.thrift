namespace go user

struct RegistReq {
    1: string username ;
    2: string password;
}

struct RegistResp {
    1: i32 status_code;
    2: string status_msg;
    3: i64 user_id;
    4: string token;
}


struct LoginReq {
    1: string username;
    2: string password;
}

struct LoginResp {
    1: i32 status_code;
    2: string status_msg;
    3: i64 user_id;
    4: string token;
}

struct InfoReq {
    1: i64 user_id;
    2: string token;
}

struct InfoResp {
    1: i32 status_code;
    2: string status_msg;
    3: User user;
}


struct User {
        1: required i64 id; // 用户id
        2: required string name; // 用户名称
        3: optional i64 follow_count; // 关注总数
        4: optional i64 follower_count; // 粉丝总数
        5: required bool is_follow; // true-已关注，false-未关注
        6: optional string avatar; //用户头像
        7: optional string background_image; //用户个人页顶部大图
        8: optional string signature; //个人简介
        9: optional i64 total_favorited; //获赞数量
        10: optional i64 work_count; //作品数量
        11: optional i64 favorite_count; //点赞数量
}


struct AuthReq {
    1: i64 user_id;
    2: string token;
}

struct AuthResp {
        1: i32 status_code;
        2: string status_msg;
        3: i32 status;
}


service UserService {
    RegistResp RegistMethod(1: RegistReq request);
    LoginResp LoginMethod(1: LoginReq request) ;
    InfoResp InfoMethod(1: InfoReq request) ;
    AuthResp AuthMethod(1: AuthReq request);
}
