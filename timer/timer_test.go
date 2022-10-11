package timer

import (
	"testing"
)

//channel 在 Golang 中是一等公民，它是线程安全的，面对并发问题，应首先想到 channel
//关闭一个未初始化的 channel 会产生 panic
//重复关闭同一个 channel 会产生 panic
//向一个已关闭的 channel 发送消息会产生 panic
//从已关闭的 channel 读取消息不会产生 panic，且能读出 channel 中还未被读取的消息，若消息均已被读取，则会读取到该类型的零值。
//从已关闭的 channel 读取消息永远不会阻塞，并且会返回一个为 false 的值，用以判断该 channel 是否已关闭（x,ok := <- ch）
//关闭 channel 会产生一个广播机制，所有向 channel 读取消息的 goroutine 都会收到消息

func Test_channel(t *testing.T) {

}
