package cmd

import (
	"fmt"
	_ "qcloud/configs"
	"qcloud/internal/DescribeAccountBalance"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var describeaccountbalance = &cobra.Command{
	Use:   "DescribeAccountBalance",
	Short: "查詢帳戶餘額",
	Long:  "用來查詢帳戶餘額使用方法如下 : \nDescribeAccountBalance -a <帳戶名稱> ex: DescribeAccountBalance -a account",
	Run: func(cmd *cobra.Command, args []string) {
		var SecretId = viper.GetString(account + ".SecretId")
		var SecretKey = viper.GetString(account + ".SecretKey")

		request := DescribeAccountBalance.DescribeAccountBalance(SecretId, SecretKey)
		fmt.Printf("%s 帳戶餘額: %f", account, request)
	},
}

func init() {
	describeaccountbalance.Flags().StringVarP(&account, "account", "a", "", "帳號 (require)")
	describeaccountbalance.MarkFlagRequired("account")
}
