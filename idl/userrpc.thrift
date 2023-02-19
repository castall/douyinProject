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
    1: i64 id;
    2: string name;
    3: i64 follow_count;
    4: i64 follower_count;
    5: bool is_follow;
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
