package DescribeAccountBalance

import (
	billing "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/billing/v20180709"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

func DescribeAccountBalance(SecretId string, SecretKey string) float64 {

	credential := common.NewCredential(
		SecretId,
		SecretKey,
	)
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "billing.tencentcloudapi.com"
	client, _ := billing.NewClient(credential, "", cpf)

	request := billing.NewDescribeAccountBalanceRequest()

	response, err := client.DescribeAccountBalance(request)
	/*if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return
	}*/
	if err != nil {
		panic(err)
	}

	s := float64(*response.Response.Balance) / 100
	//fmt.Printf("%s", response.ToJsonString())

	return s
}
