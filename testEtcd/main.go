package main

import (
	"context"
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"log"
	"time"
)

func main() {
	client, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"127.0.0.1:2379"}, // conf.Etcd.Host
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		log.Printf("Connect etcd failed. Error : %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	rsp, _ := client.Put(ctx, "product", "5678")
	fmt.Println(rsp)

	//ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	//defer cancel()
	//rsp, _ := client.Put(ctx, "product","5678")
	//fmt.Println(rsp)

	//ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	//defer cancel()
	//rsp, _ := client.Get(ctx, "product")
	//fmt.Println(rsp)

	//ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	//defer cancel()
	//response, err := client.Delete(ctx, "product")
	//if err != nil {
	//	log.Printf("Delete etcd failed. err: %#v", err)
	//}
	//fmt.Println(response)
}
