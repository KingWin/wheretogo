namespace go sys.account

struct LoginReq {
    1: string Username (api.body="username");
    2: string Password (api.body="password");
}

struct LoginResp {
    1: string Resp;
}

struct LogoutReq {
    1: string Logout (api.body="username");
}

struct LogoutResp {
    1: string Resp;
}

struct AuthReq {
    1: string Auth (api.body="username");
}

struct AuthResp {
    1: string Resp;
}

service LoginService {
    LoginResp LoginMethod(1: LoginReq request) (api.post="/login");
    LogoutResp LogoutMethod(1: LogoutReq request) (api.post="/logout");
}

service AuthService {
    AuthResp AuthMethod(1: AuthReq request) (api.post="/auth");
}
