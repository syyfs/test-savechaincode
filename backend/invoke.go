package backend

import (
	"net/http"
	"fmt"
	"io/ioutil"
	"encoding/json"
	"bytes"
	"github.com/syyfs/test-savechaincode/common"
	"strconv"
	"github.com/syyfs/test-savechaincode/model"
)

func InvokeSaveValue(token string , i int)  {
	client := &http.Client{}
	invokeChaincodeCfg := &model.InvokeChaincodeCfg{
		ChannelId: common.GetChannelName(),
		CcName: common.GetChaincodeName(),
		CcFcn: common.GetChaincodeFunc(),
		CcArgs: []string{strconv.Itoa(i)},
	}
	byteInvoke , err := json.Marshal(invokeChaincodeCfg)
	if err != nil {
		fmt.Errorf("** InvokeChaincodeCfg Marshal Faild !!!**\n")
	}

	url := fmt.Sprintf("http://%s:%s/chaincode/execute", common.Getclientcommont() , common.GetClientPort())

	result, err := clientDo(token, client , "POST", url, byteInvoke)
	if err != nil {
		panic(err)
	}
	fmt.Printf("----- 执行结果：%s ------\n", result)
}

func clientDo(token string,client *http.Client ,method string, url string, cfgBytes []byte ) (string , error) {

	cibody := bytes.NewBuffer(cfgBytes)
	req, err := http.NewRequest(method, url , cibody)
	if err != nil {
		return	"", fmt.Errorf(" ==== [clientDo] NewRequest faild ! err is %s \n ", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("accept", "application/json")
	req.Header.Set("X-Auth-Token", token)

	resp, err := client.Do(req)
	if err != nil {
		return "",fmt.Errorf(" ==== [clientDo] client.Do(req) faild ! err is %s \n ", err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		return "",fmt.Errorf(" ==== [clientDo] ReadAll faild ! err is %s \n ", err)
	}
	bodymap := make(map[string]interface{})
	err = json.Unmarshal(body,&bodymap)
	if err != nil {
		fmt.Printf("\n %c[1;40;32m  err:%s\n %c[0m\n\n", 0x1B, err,0x1B)
		return "",fmt.Errorf(" ==== [clientDo] ReadAll faild ! err is %s \n ", err)
	}
	val , ok := bodymap["code"]
	if !ok {
		panic(fmt.Errorf(" ==== [clientDo] response code faild !\n "))
	}
	if val.(float64) != 200 {
		panic(fmt.Errorf(" ==== [clientDo] response code faild !\n "))
	}

	return bodymap["data"].(string) , err
}

