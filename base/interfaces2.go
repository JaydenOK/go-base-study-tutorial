package main

import (
	"fmt"
)

type SalaryCalculator interface {
	CalculateSalary() int
}

type Permanent struct {
	empId    int
	basicpay int
	pf       int
}

type Contract struct {
	empId    int
	basicpay int
}

//给接收器Permanent 加上SalaryCalculator接口所有的方法, 即为实现了接口
//salary of permanent employee is sum of basic pay and pf，
func (p Permanent) CalculateSalary() int {
	return p.basicpay + p.pf
}

//salary of contract employee is the basic pay alone
func (c Contract) CalculateSalary() int {
	return c.basicpay
}

/**
传入接口SalaryCalculator
*/
func totalExpense(s []SalaryCalculator) {
	expense := 0
	for _, v := range s {
		expense = expense + v.CalculateSalary()
	}
	fmt.Printf("Total Expense Per Month $%d", expense)
}

func main() {
	pemp1 := Permanent{1, 5000, 20}
	pemp2 := Permanent{2, 6000, 30}
	cemp1 := Contract{3, 3000}
	//定义接口切片 (切片为不定长度的数组)
	employees := []SalaryCalculator{pemp1, pemp2, cemp1}
	totalExpense(employees)

}

/////////////////////////////////////////////////////////////
//空接口
//没有包含方法的接口称为空接口。空接口表示为 interface{}。由于空接口没有方法，因此所有类型都实现了空接口
//describe(i interface{}) 函数接收空接口作为参数，因此，可以给这个函数传递任何类型。
func describe(i interface{}) {
	fmt.Printf("Type = %T, value = %v\n", i, i)
}
