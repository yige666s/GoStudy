package main

import (
	"day10/dal"
	"day10/dal/query"
)

const MYSQLDSN = "root:root1234@tcp(127.0.0.1:13306)/GormTest?charset=utf8mb4&parseTime=True"

func init() {
	dal.DB = dal.ConnetDB(MYSQLDSN).Debug()
}

func GenTest() {
	// 设置默认DB对象
	query.SetDefault(dal.DB)

	// 创建
	// b1 := model.Book{
	// 	Title:       "<C++ Primer>",
	// 	Author:      "Mr.C",
	// 	PublishDate: time.Date(2024, 6, 17, 0, 0, 0, 0, time.UTC),
	// 	Price:       98,
	// }
	// err := query.Book.WithContext(context.Background()).Create(&b1)
	// if err != nil {
	// 	fmt.Printf("create book fail, err :%v\n", err)
	// 	return
	// }

	// // 更新
	// ret, err := query.Book.WithContext(context.Background()).
	// 	Where(query.Book.ID.Eq(1)).
	// 	Update(query.Book.Price, 200)
	// if err != nil {
	// 	fmt.Printf("update book failed,err:%v\n", err)
	// 	return
	// }
	// fmt.Printf("RowAffected:%v\n", ret.RowsAffected)

	// 查询
	// book, err := query.Book.WithContext(context.Background()).First()
	// if err != nil {
	// 	fmt.Printf("Select book failed,err:%v", err)
	// 	return
	// }
	// fmt.Printf("Book : %v\n", book)

	// 删除
	// ret, err := query.Book.WithContext(context.Background()).
	// 	Where(query.Book.ID.Eq(2)).Delete()
	// if err != nil {
	// 	fmt.Printf("delete book failed,err :%v\n", err)
	// 	return
	// }
	// fmt.Printf("RowAffected:%v\n", ret.RowsAffected)

	// // 使用自定义接口
	// rets, err := query.Book.WithContext(context.Background()).GetBooksByAuthor("jack")
	// if err != nil {
	// 	fmt.Printf("GetBooksByAuther fail,err: %v\n", err)
	// 	return
	// }
	// for i, b := range rets {
	// 	fmt.Printf("%d:%v\n", i, b)
	// }

}

// 自定义SQL查询
