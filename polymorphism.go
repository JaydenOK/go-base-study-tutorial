package main

import "fmt"

//继承，继承面向接口编程

//IOrder接口 ，函数名、返回值需一致
type IOrder interface {
	GetName() string
	GetOrderNo() string
	GetTotalPrice() float32
}

//Order父类
type Order struct {
	name       string
	orderNo    string
	totalPrice float32
}

//
func (o Order) getOrderInfo(order IOrder) {
	fmt.Printf("订单类型：%s；订单号：%s； 金额：%f \n", order.GetName(), order.GetOrderNo(), order.GetTotalPrice())
}

//AmazonOrder结构体 实现Order接口
type AmazonOrder struct {
	name       string
	orderNo    string
	totalPrice float32
	Order
}

func (o AmazonOrder) GetName() string {
	return o.name
}
func (o AmazonOrder) GetOrderNo() string {
	return o.orderNo
}
func (o AmazonOrder) GetTotalPrice() float32 {
	return o.totalPrice
}

//EbayOrder结构体 实现Order接口
type EbayOrder struct {
	name       string
	orderNo    string
	totalPrice float32
	Order
}

func (o EbayOrder) GetName() string {
	return o.name
}
func (o EbayOrder) GetOrderNo() string {
	return o.orderNo
}
func (o EbayOrder) GetTotalPrice() float32 {
	return o.totalPrice
}

func main() {
	amazonOrder := AmazonOrder{
		name:       "亚马逊订单",
		orderNo:    "A124354354654",
		totalPrice: 100.3,
	}
	ebayOrder := EbayOrder{
		name:       "Ebay订单",
		orderNo:    "E2134325345436546",
		totalPrice: 200.23,
	}

	amazonOrder.getOrderInfo(amazonOrder)
	ebayOrder.getOrderInfo(ebayOrder)

}
