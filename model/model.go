package model

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`

}


type LoginResponse struct {
	Code int `json:"code"`
	Data LoginDataResponse `json:"data"`
}

type LoginDataResponse struct {
	Token    string `json:"token"`
	OrgExist bool  `json:"org_exist"`
}

type InvokeChaincodeCfg struct {
	ChannelId string   `json:"channel_id"`
	CcName    string   `json:"cc_name"`
	CcFcn     string   `json:"cc_fcn"`
	CcArg     string   `json:"cc_arg"`
	CcArgs    []string `json:"cc_args"`
}
