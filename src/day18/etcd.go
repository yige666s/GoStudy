package main

import "go.etcd.io/etcd/clientv3"

// etcd是近几年比较火热的一个开源的、分布式的键值对数据存储系统，
// 提供共享配置、服务的注册和发现，本文主要介绍etcd的安装和使用。
// etcd是近几年比较火热的一个开源的、分布式的键值对数据存储系统，提供共享配置、服务的注册和发现，本文主要介绍etcd的安装和使用。

func Etcd() {
	cli, err := clientv3.New()
}
