package business

import (
	"context"
	"fmt"
	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/apache/rocketmq-client-go/v2/producer"
	"time"
)

type MyTransactionListener struct {
	// 需要实现下面函数
	//ExecuteLocalTransaction(message *primitive.Message) primitive.LocalTransactionState
	//CheckLocalTransaction(*primitive.MessageExt) primitive.LocalTransactionState
}

// ExecuteLocalTransaction 发送成功后调用
func (l *MyTransactionListener) ExecuteLocalTransaction(message *primitive.Message) primitive.LocalTransactionState {
	fmt.Printf("开始执行本地事物: %s\n", message)
	time.Sleep(2 * time.Second)
	fmt.Printf("执行本地事物, 返回状态: %d\n", primitive.UnknowState)
	// 返回提交状态 CommitMessageState: 执行成功提交; RollbackMessageState: 执行失败回滚消息;
	return primitive.UnknowState
}

// CheckLocalTransaction 检查没有响应消息的状态(如果未返回 或者 UnknowState:调用此方法
func (l *MyTransactionListener) CheckLocalTransaction(message *primitive.MessageExt) primitive.LocalTransactionState {
	fmt.Printf("检查本地状态: %s\n", message)
	time.Sleep(5 * time.Second)
	return primitive.UnknowState
}

func init() {

}

// Transaction 发送事务消息
func Transaction() {
	p, err := rocketmq.NewTransactionProducer(
		&MyTransactionListener{},
		producer.WithNameServer([]string{"82.157.232.70:9876"}),
		producer.WithRetry(1),
	)
	if err != nil {
		// 示例化失败
		panic(any(err))
	}
	err = p.Start()
	if err != nil {
		fmt.Printf("start producer error: %s\n", err.Error())
	}

	// 发送事物消息
	message := primitive.NewMessage("hello", []byte("事物消息"))
	res, err := p.SendMessageInTransaction(context.Background(), message)
	if err != nil {
		// 发送失败
		fmt.Println("发送失败")
		panic(any(err))
	}

	fmt.Printf("发送结果: %s\n", res.String())

	time.Sleep(10 * time.Hour)

	if err := p.Shutdown(); err != nil {
		// 结束失败
		panic(any(err))
	}
}
