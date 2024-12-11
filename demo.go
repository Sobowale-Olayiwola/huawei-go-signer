/*
 *  Copyright (c) Huawei Technologies Co., Ltd. 2017-2023. All rights reserved.
 */

package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	demoAppApigw()
}

// https://console-intl.huaweicloud.com/apiexplorer/#/endpoint/IAM --> a toggle section to see endpoints for different Cloud Services Huawei provide

func demoAppApigw() {
	// 认证用的ak和sk硬编码到代码中或者明文存储都有很大的安全风险，建议在配置文件或者环境变量中密文存放，使用时解密，确保安全；
	// 本示例以ak和sk保存在环境变量中为例，运行本示例前请先在本地环境中设置环境变量HUAWEICLOUD_SDK_AK和HUAWEICLOUD_SDK_SK。
	// s := Signer{
	// 	Key:    os.Getenv("HUAWEICLOUD_SDK_AK"),
	// 	Secret: os.Getenv("HUAWEICLOUD_SDK_SK"),
	// }
	s := Signer{
		Key:    "KEY",
		Secret: "SECRET",
	}
	// r, err := http.NewRequest("GET", "https://30030113-3657-4fb6-a7ef-90764239b038.apigw.cn-north-1.huaweicloud.com/app1?a=1&b=2",
	// ioutil.NopCloser(bytes.NewBuffer([]byte("foo=bar"))))
	r, err := http.NewRequest("GET", "https://iam.af-south-1.myhuaweicloud.com/v3.0/OS-USER/users/6f00f987c12a4e2885c42b872320453f", nil) // ap-southeast-1
	//https://iam.ae-ad-1.myhuaweicloud.com/v3.0/OS-MFA/virtual-mfa-devices
	//https://iam.af-south-1.myhuaweicloud.com/v3.0/OS-MFA/virtual-mfa-devices
	//https://iam.af-south-1.myhuaweicloud.com/v3/endpoints
	//https://iam.af-south-1.myhuaweicloud.com/v3.0/OS-USER/users/{user_id}
	//https://ecs.af-north-1.myhuaweicloud.com/v1/{project_id}/cloudservers/os-server-groups
	//https://iam.af-south-1.myhuaweicloud.com/v3.0/OS-USER/users/6f00f987c12a4e2885c42b872320453f
	//https://ecs.ap-southeast-1.myhuaweicloud.com/v1/fecac407622c4d20a4f0a57603ee003a/cloudservers/os-server-groups

	if err != nil {
		fmt.Println(err)
		return
	}

	r.Header.Add("content-type", "application/json; charset=utf-8")
	r.Header.Add("x-stage", "RELEASE")
	s.Sign(r) // signer signing the request
	fmt.Println(r.Header)
	client := http.DefaultClient
	resp, err := client.Do(r)
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(body))
}
