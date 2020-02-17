package main

import (
	"fmt"

	"github.com/golden-corp/openplatform-sdk-go/goland"
)

func main() {
	sdk := goland.NewSdk("fc36541461483b2db498", "d2641bfc30b293505ca2c09560b870aa", "1.0.0", "test")
	items := make([]map[string]interface{}, 1)
	item := make(map[string]interface{})
	item["name"] = "咨询服务"
	item["tax_code"] = "1020202000000000000"
	item["tax_type"] = ""
	item["models"] = "xyz"
	item["unit"] = "个"
	item["total_price"] = 8644
	item["total"] = "5"
	item["price"] = "17.288"
	item["tax_rate"] = 100
	item["tax_amount"] = 864
	item["discount"] = 0
	item["zero_tax_flag"] = ""
	item["preferential_policy_flag"] = ""
	item["vat_special_management"] = ""
	items[0] = item
	var post = map[string]interface{}{
		"seller_name":          "",
		"seller_taxpayer_num":  "111112222233333",
		"seller_address":       "",
		"seller_tel":           "",
		"seller_bank_name":     "",
		"seller_bank_account":  "",
		"title_type":           1,
		"buyer_title":          "我是抬头",
		"buyer_taxpayer_num":   "",
		"buyer_address":        "",
		"buyer_bank_name":      "",
		"buyer_bank_account":   "",
		"buyer_phone":          "",
		"buyer_email":          "",
		"taker_phone":          "",
		"taker_name":           "",
		"order_id":             "test001asdfsadfasdf",
		"invoice_type_code":    "032",
		"callback_url":         "http://www.xx.com",
		"drawer":               "小刘",
		"payee":                "小刘",
		"checker":              "小刘",
		"trade_type":           0,
		"user_openid":          "",
		"special_invoice_kind": "",
		"terminal_code":        "",
		"amount_has_tax":       9508,
		"tax_amount":           864,
		"amount_without_tax":   8644,
		"remark":               "",
		"items":                items,
	}

	r, err := sdk.HttpPost("/invoice/blue", post)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(r))
	}
}
