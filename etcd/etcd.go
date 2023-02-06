package main

import (
	"context"
	"fmt"
	"github.com/coreos/etcd/clientv3"
	"time"
)

func main() {
	var (
		config clientv3.Config
		client *clientv3.Client
		err    error
		kv     clientv3.KV
		putRes *clientv3.PutResponse
		getRes *clientv3.GetResponse
	)

	//配置
	config = clientv3.Config{
		Endpoints:   []string{"127.0.0.1:2379"},
		DialTimeout: 5 * time.Second,
	}

	//建立连接
	if client, err = clientv3.New(config); err != nil {
		fmt.Println("client ok")
		return
	}

	defer client.Close()

	//写etcd中的键值对
	kv = clientv3.NewKV(client)

	if putRes, err = kv.Put(context.TODO(), "/cront/job/1", "ok", clientv3.WithPrevKV()); err != nil {
		fmt.Println(err)
	} else {
		//putRes.PrevKv输出如下：
		/*
			key:"/cron/jobs/job1" create_revision:43 mod_revision:46 version:4 value:"{\"name\":\"job1\",\"command\":\"echo hello\",\"crontab\":\"\"}"
		*/
		fmt.Println(putRes.Header.Revision) //输出Revision
		fmt.Println(putRes.PrevKv.Value)    //输出修改前的值
	}

	//读取etcd的键值
	if getRes, err = kv.Get(context.TODO(), "/cront/job/1"); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(getRes.Kvs)
	}
}
