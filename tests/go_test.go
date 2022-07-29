package tests

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestGo(t *testing.T) {
	url := "https://openplatform.shopee.cn/api/v2/order/get_order_list"
	detailUrl := "https: //openplatform.shopee.cn/api/v2/order/get_order_detail"
	s := getOrderList(url)
	s2 := getOrderDetail(detailUrl)
	fmt.Println(s, s2)
}

func getOrderDetail(url string) string {
	resp, _ := http.Get(url)
	defer resp.Body.Close()
	buffer, _ := ioutil.ReadAll(resp.Body)
	return string(buffer)
}

func getOrderList(url string) string {
	resp, _ := http.Get(url)
	defer resp.Body.Close()
	buffer, _ := ioutil.ReadAll(resp.Body)
	return string(buffer)
}
