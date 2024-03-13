package main

import (
	"bytes"
	"fmt"
	"net/http"
	"time"

	"github.com/aws/aws-sdk-go/aws/credentials"
	v4 "github.com/aws/aws-sdk-go/aws/signer/v4"
	"github.com/zoncoen/scenarigo/plugin"
)

func init() {
	plugin.RegisterSetupEachScenario(setRunID)
}

func AWSSigV4() (X_Amz_Date, Authorization string) {
	profileName := "default"
	// AWS認証情報
	creds := credentials.NewSharedCredentials("/Users/taichi/.aws/credentials", profileName)

	// サービス名とリージョン
	service := "execute-api"
	region := "ap-northeast-1"

	// HTTPリクエスト情報
	method := "GET"
	url := ""
	body := []byte(``)
	headers := map[string][]string{}

	// HTTPリクエストを作成
	req, err := http.NewRequest(method, url, bytes.NewReader(body))
	if err != nil {
		fmt.Println("Error creating HTTP request:", err)
		return
	}

	// HTTPヘッダーを設定
	for key, values := range headers {
		for _, value := range values {
			req.Header.Add(key, value)
		}
	}

	// AWS Signature Version 4の署名を行う
	signer := v4.NewSigner(creds)
	_, err = signer.Sign(req, bytes.NewReader(body), service, region, time.Now())
	if err != nil {
		fmt.Println("Error signing HTTP request:", err)
		return
	}

	return req.Header.Get("X-Amz-Date"), req.Header.Get("Authorization")
}

func setRunID(ctx *plugin.Context) (*plugin.Context, func(*plugin.Context)) {
	// シナリオファイルのパスを取得する
	// filePath := ctx.ScenarioFilepath()

	// シナリオファイルの内容を読み取る
	// content, err := ioutil.ReadFile(filePath)
	// if err != nil {
	// 	fmt.Println("Error reading scenario file:", err)
	// }

	// YAMLデコードして構造体にマッピングする
	// var scenarioData map[string]interface{}
	// if err := yaml.Unmarshal(content, &scenarioData); err != nil {
	// 	fmt.Println("Error decoding scenario file:", err)
	// }
	// steps := scenarioData["steps"].([]interface{})
	// for _, step := range steps {
	// 	stepData := step.(map[interface{}]interface{})
	// }

	x_aws_date, authorization := AWSSigV4()

	return ctx.WithVars(map[string]string{
		"X-Amz-Date":    x_aws_date,
		"Authorization": authorization,
	}), nil
}
