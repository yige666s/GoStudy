package model

import "gorm.io/gen"

// 自定义SQL查询：按对应规则将SQL语句注释到interface的方法上即可。Gen将对其进行解析，并为应用的结构生成查询API。
// 通常建议将自定义查询方法添加到model模块下。

type Querier interface {
	// SELECT * FROM @@table WHERE id=@id
	GetById(id int) (ret gen.T, err error)            // 返回结构体T
	GetByIdReturnMap(id int) (ret gen.M, err error)   // 返回Map
	GetBooksByAuthor(author string) ([]*gen.T, error) //返回切片
}
