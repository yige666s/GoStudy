package main

// 我们的Go程序编译之后会得到一个可执行的二进制文件，
// 其实在最终的镜像中是不需要go编译器的，也就是说我们只需要一个运行最终二进制文件的容器即可。

// Docker的最佳实践之一是通过仅保留二进制文件来减小镜像大小，为此，我们将使用一种称为多阶段构建的技术，这意味着我们将通过多个步骤构建镜像
// docker build . -t bubble_app
// 我们这里运行bubble_app容器的时候需要使用--link的方式与上面的mysql8019容器关联起来，具体命令如下：
// docker run bubble_app -p 8888:8888 --link=mysql8019:mysql8019

// 使用Compose基本上是一个三步过程：
// 使用Dockerfile定义你的应用环境以便可以在任何地方复制。
// 定义组成应用程序的服务，docker-compose.yml 以便它们可以在隔离的环境中一起运行。
// 执行 docker-compose up命令来启动并运行整个应用程序。