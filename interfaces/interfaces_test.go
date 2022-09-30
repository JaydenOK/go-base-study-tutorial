package interfaces

import (
	"fmt"
	"testing"
)

type EventHandler interface {
	Handle(params interface{})
}

type OrderEventHandler struct {
}

func (orderEventHandle *OrderEventHandler) Handle(params interface{}) {

}

// 事件管理器
type EventManager struct {
	events map[string]EventHandler
}

// 绑定事件
func (eventManager *EventManager) Bind(eventName string, eventHandle EventHandler) (manager *EventManager) {
	isBind := false
	for name, _ := range eventManager.events {
		if eventName == name {
			isBind = true
		}
	}
	if !isBind {
		eventManager.events[eventName] = eventHandle
	}
	manager = eventManager
	return
}

// 触发事件
func (eventManager *EventManager) Trigger(eventName string, params interface{}) {
	for name, event := range eventManager.events {
		if eventName == name {
			//event.handle(params)
			fmt.Println(event)
		}
	}
}

func Test_interfaces(t *testing.T) {
	//interface 是一个接口语法，而interface{}是一个类型
	eventManager := EventManager{}
	orderEvent := OrderEventHandler{}
	eventManager.Bind("EventOrderAdd", &orderEvent).Trigger("EventOrderAdd", "123456")
}
