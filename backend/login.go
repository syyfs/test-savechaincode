package backend

import (
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/syyfs/test-savechaincode/common"
	"bytes"
	"io/ioutil"
	"github.com/syyfs/test-savechaincode/model"
)

func Login() (string , error) {
	loginReq := model.LoginRequest{
		Username:"admin",
		Password:"brilliance",
	}

	loginByte, err := json.Marshal(loginReq)
	if err != nil {
		fmt.Errorf("** Login Model Marshal faild !!! **\n")
	}

	client := http.Client{}
	url := fmt.Sprintf("http://%s:%s/admin/login",common.Getclientcommont() , common.GetClientPort())
	resp , err := client.Post(url, "application/json" , bytes.NewBuffer(loginByte))
	defer resp.Body.Close()

	respByte , err := ioutil.ReadAll(resp.Body)
	loginResp := model.LoginResponse{}
	err = json.Unmarshal(respByte,&loginResp)
	if err != nil {
		fmt.Errorf("** Login Model UnMarshal faild !!! **\n")
	}

	return loginResp.Data.Token , err
}
