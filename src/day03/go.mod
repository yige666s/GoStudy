module day03

go 1.22.2

require github.com/go-sql-driver/mysql v1.8.1

// 行尾的indirect表示该依赖包为间接依赖，说明在当前程序中的所有 import 语句中没有发现引入这个包
require filippo.io/edwards25519 v1.1.0 // indirect	
