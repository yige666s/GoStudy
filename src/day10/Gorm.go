package main

import (
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	// 这是老版1.x
	// 这是新版2.x
)

// func GormCase() {
// 	db, err := goa.Open(gormsql.Open("test.db"), &gorm.Config{})
// 	if err != nil {
// 		panic("failed to connect to mysql")
// 	}

// 	db2, err2 := goj.Open("mysql", "user:password@(localhost)/dbname?charset=utf8mb4&parseTime=True&loc=Local")
// 	if err2 != nil {
// 		panic("failed to connect to mysql")
// 	}
// 	defer db2.Close() // goj这个包中有close
// }

// Userinfo 用户信息
// 表名默认就是结构体名称的复数
type Userinfo struct {
	// gorm.Model
	ID     uint
	Name   string `gorm:"default:'jack'"`
	Gender string
}

// 使用指针
type User struct {
	ID   int64
	Name *string `gorm:"default:'小王子'"`
	Age  int64
}

// type Animal struct {
// 	AnimalId int64     `gorm:"column:beast_id`
// 	Birthday time.Time `gorm:"column:day_of_the_beast"`
// 	Age      int64     `gorm:"column:age_of_the_beast`
// }

func WslDBTest() {
	// dsn := "root:root1234@tcp(127.0.0.1:13306)/GormTest?charset=utf8mb4&parseTime=True&loc=Local"
	// db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	// if err != nil {
	// 	panic("open GormTestDB failed")
	// }

	// // 迁移schema,
	// // AutoMigrate: 这是Gorm提供的一个方法，用于自动迁移数据库表结构。它会检查数据库中是否存在与指定结构体对应的表，
	// // 如果不存在则创建表，如果存在则检查表的结构是否与结构体匹配，不匹配则进行修改（例如添加新字段）
	// db.AutoMigrate(&Userinfo{})

	// u1 := &Userinfo{1, "jack", "man", 89}
	// u2 := &Userinfo{2, "lili", "woman", 80}

	// // insert
	// db.Create(&u1)
	// db.Create(&u2)

	// // select
	// var u = new(Userinfo)
	// db.First(u)
	// fmt.Printf("%#v\n", u)

	// var u3 Userinfo
	// db.Find(&u3, "Gender=?", "man")
	// fmt.Printf("%#v\n", u)

	// // update
	// db.Model(&u).Update("score", 140)

	// // delete
	// db.Delete(&u)

	// // 使用UserInfo结构体创建名为`deleted_users`的表
	// db.Table("deleted_users").Migrator().CreateTable(&Userinfo{})
	// var deleted_users []Userinfo
	// // // SELECT * FROM deleted_users;
	// db.Table("deleted_users").Find(&deleted_users)
	// // // Delete From deleted_users where name = "lili"
	// db.Table("deleted_users").Where("name = ?", "lili").Delete(&Userinfo{})
	// // Insert into Delted_users value ("jack","man")
	// db.Create(&Userinfo{Name: "jack", Gender: "man"})

}

func CRUD() {
	dsn := "root:root1234@tcp(127.0.0.1:13306)/GormTest?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("open GormTestDB failed")
	}

	// 迁移schema,
	// AutoMigrate: 这是Gorm提供的一个方法，用于自动迁移数据库表结构。它会检查数据库中是否存在与指定结构体对应的表，
	// 如果不存在则创建表，如果存在则检查表的结构是否与结构体匹配，不匹配则进行修改（例如添加新字段）
	db.AutoMigrate(&Userinfo{})

	// Userinfo1 := &Userinfo{Name: "m1", Gender: "man"}
	// if Userinfo1.ID == 0 {
	// 	log.Println("the primary key is null")
	// }
	// db.Create(&Userinfo1)
	// if Userinfo1.ID != 0 {
	// 	log.Println("the primary key has been created")
	// }

	// user := &User{Name: new(string), Age: 18))}
	// db.Create(&user)  // 此时数据库中该条记录name字段的值就是''

	//// SELECT * FROM users ORDER BY id LIMIT 1;
	// db.First(&Userinfo{})

	// //// SELECT * FROM users LIMIT 1;
	// db.Take(&Userinfo{})

	// //// SELECT * FROM users ORDER BY ID DESC LIMIT 1;
	// db.Find(&Userinfo{})

	// //// SELECT * FROM users;
	// db.Find(&Userinfo{})

	// //// SELECT * FROM users WHERE id = 10;
	// db.First(&Userinfo{}, 10)

	// //// SELECT * FROM users WHERE name = 'jinzhu' limit 1;
	// db.Where("name = ?", "Megan Jackson").First(&Userinfo{})
	// db.Where("name = ?", "Megan Jackson").Find(&Userinfo{})
	// db.Where("name <> ?", "Megan Jackson").Find(&Userinfo{})
	// db.Where("name IN (?)", []string{"jack", "lili"}).Find(&Userinfo{})
	// db.Where("name LIKE ?", "Jac%").Find(&Userinfo{})
	// db.Where("name = ? AND Gender = ?", "lili", "women").Find(&Userinfo{})
	// // db.Where("updated_at > ?", lastWeek).Find(&Userinfo{})
	// // // SELECT * FROM users WHERE updated_at > '2000-01-01 00:00:00';
	// // db.Where("created_at BETWEEN ? AND ?", lastWeek, today).Find(&Userinfo{})
	// // // SELECT * FROM users WHERE created_at BETWEEN '2000-01-01 00:00:00' AND '2000-01-08 00:00:00';

	// // Struct
	// // SELECT * FROM users WHERE name = "jinzhu" AND age = 20 LIMIT 1;
	// db.Where(&Userinfo{Name: "jack", Gender: "man"}).First(&Userinfo{})

	// // Map
	// // SELECT * FROM users WHERE name = "jinzhu" AND age = 20;
	// db.Where(map[string]interface{}{"name": "jack", "Gender": "man"}).Find(&Userinfo{})

	// // 主键的切片
	// // SELECT * FROM users WHERE id IN (20, 21, 22);
	// db.Where([]int32{1, 2, 3, 4, 5}).Find(&Userinfo{})

	// // 如果字段值为0，''，false或者其他零值时，将不会被用于构建查询条件
	// // SELECT * FROM userinfos WHERE Name = 'jack';
	// db.Where(&Userinfo{Name: "jack", Gender: ""}).Find(&Userinfo{})
	// db.Not("name", "jack").First(&Userinfo{})
	// db.Not("name", []string{"jack", "lili"}).Find(&Userinfo{})
	// db.Not([]int{1, 2, 3, 4, 5}).Find(&Userinfo{})

	// // Plain SQL,这里的 NOT 等同于 '!='
	// db.Not("name = ?", "jinzhu").First(&user)
	// //// SELECT * FROM users WHERE NOT(name = "jinzhu");
	// db.Not(&Userinfo{Name: "jack"}).First(&Userinfo{})

	// // SELECT * FROM userinfos WHERE name = 'jack' or Gender = 'male';
	// db.Where("name = ?", "jack").Or("Gender = ?", "male").Find(&Userinfo{})
	// db.Where("name = ? ", "jack").Or(&Userinfo{Name: "jack"}).Find(&Userinfo{})
	// db.Where("name = ?", "jack").Or(map[string]interface{}{"name": "lili"}).Find(&Userinfo{})

	// // 根据主键获取记录 (只适用于整形主键)
	// db.First(&user, 23)
	// // // SELECT * FROM users WHERE id = 23 LIMIT 1;
	// // 根据主键获取记录, 如果它是一个非整形主键
	// db.First(&user, "id = ?", "string_primary_key")
	// //// SELECT * FROM users WHERE id = 'string_primary_key' LIMIT 1;

	// // Plain SQL
	// db.Find(&user, "name = ?", "jinzhu")
	// //// SELECT * FROM users WHERE name = "jinzhu";

	// db.Find(&users, "name <> ? AND age > ?", "jinzhu", 20)
	// //// SELECT * FROM users WHERE name <> "jinzhu" AND age > 20;

	// // Struct
	// db.Find(&users, User{Age: 20})
	// //// SELECT * FROM users WHERE age = 20;

	// // Map
	// db.Find(&users, map[string]interface{}{"age": 20})
	// // // SELECT * FROM users WHERE age = 20;

	// SELECT * FROM userinfos WHERE id = 10 FOR UPDATE;
	db.Set("gorm:query_option", "FOR UPDATE").First(&Userinfo{}, 10)

	// 查询符合条件的记录，找到则返回否则创建
	db.Where(&Userinfo{Name: "jack"}).FirstOrInit(&Userinfo{})

	// 如果记录未找到，将使用参数初始化 struct.
	db.Where(&Userinfo{Name: "lucy"}).Attrs(&Userinfo{Gender: "fmale"}).FirstOrInit(&Userinfo{})

	// 不管记录是否找到，都将参数赋值给 struct.
	db.Where(&Userinfo{Name: "jack"}).Assign(&Userinfo{Gender: "fmale"}).FirstOrInit(&Userinfo{})

	// 	FirstOrInit：如果记录存在，则返回该记录；如果不存在，则初始化一个新实例，但不保存到数据库。
	// FirstOrCreate：如果记录存在，则返回该记录；如果不存在，则创建一个新记录并保存到数据库。

	// 基于 *gorm.expr 的子查询
	// 	subQuery := db.Table("orders").Select("AVG(amount)").Where("state = >", "paid").SubQuery()
	// 	db.Where("amount > ?", subQuery).Find(&orders{})

	// Select，指定你想从数据库中检索出的字段，默认会选择全部字段。
	db.Select("name").Find(&Userinfo{})
	db.Select([]string{"name", "Gender"}).Find(&Userinfo{})
	// COALESCE函数会返回其第一个非NULL的参数值。因此，如果age列的值为NULL，则返回42；否则，返回age列的实际值。
	db.Table("userinfos").Select("COALESCE(age,?)", 42).Rows()

	//Order，指定从数据库中检索出记录的顺序。设置第二个参数 reorder 为 true ，可以覆盖前面定义的排序条件。
	// SELECT * FROM users ORDER BY age desc, name;
	db.Order("age desc, name").Find(&Userinfo{})
	db.Order("age desc").Order("name").Find(&Userinfo{})

	// Limit，指定从数据库检索出的最大记录数。
	db.Limit(5).Find(&Userinfo{})
	// limit(-1)取消返回数量限制
	db.Limit(10).Find(&Userinfo{}).Limit(-1).Find(&Userinfo{})

	// Offset，指定开始返回记录前要跳过的记录数
	db.Offset(5).Find(&Userinfo{})
	// -1 取消 Offset 条件
	db.Offset(10).Find(&Userinfo{}).Offset(-1).Find(&Userinfo{})
	//// SELECT * FROM users OFFSET 10; (users1)
	//// SELECT * FROM users; (users2)

	var count int64
	//// SELECT * from USERS WHERE name = 'jinzhu' OR name = 'jinzhu 2'; (users)
	//// SELECT count(*) FROM users WHERE name = 'jinzhu' OR name = 'jinzhu 2'; (count)
	db.Where("name = ?", "jack").Or("Gender = ?", "fmale").Find(&Userinfo{}).Count(&count)
	//// SELECT count(*) FROM users WHERE name = 'jinzhu'; (count)
	db.Model(&Userinfo{}).Where("Gender = ?", "male").Count(&count)
	//// SELECT count(*) FROM deleted_users;
	db.Table("deleted_users").Count(&count)
	//// SELECT count( distinct(name) ) FROM deleted_users; (count)
	db.Table("deleted_users").Select("count (distinct(name))").Count(&count)
	//Count 必须是链式查询的最后一个操作 ，因为它会覆盖前面的 SELECT，但如果里面使用了 count 时不会覆盖

	type Result struct {
		Date   time.Time
		Toatal int
	}
	var rets []Result
	rows, _ := db.Table("orders").Select("date(created_at) as date, sum(amount) as total").Group("date(created_at)").Having("sum(amount) > ?", 100).Scan(&rets).Rows()
	for rows.Next() {
		// do somthing
	}

	// Joins，指定连接条件
	rows2, _ := db.Table("users").Select("users.name, emails.email").Joins("left join emails on emails.user_id = users.id").Rows()
	for rows2.Next() {
	}
	// 多连接及参数
	db.Joins("JOIN emails ON emailsuser_id = user.id AND emails.email = ?", "mytest@gmail.com").Joins("JOIN credit_card ON Credit_card.user_id = user.id AND credit_card.number = ?", "123456789").Find(&User{})

	// Pluck，查询 model 中的一个列作为切片，如果您想要查询多个列，您应该使用 Scan
	var Genders []int64
	var userinfos []Userinfo
	// 首先加载所有Userinfo记录到内存中，然后再从这些记录中提取Gender字段的值,适用于需要获取所有记录并进行进一步操作的场景。
	db.Find(&userinfos).Pluck("Gender", &Genders)
	// 适用于仅需要特定字段值的场景，且不需要加载所有记录,直接从数据库中提取Gender字段的值，而不需要先加载所有Userinfo记录到内存中，因此更高效
	db.Model(&Userinfo{}).Pluck("Gender", &Genders)
	// 适用于操作没有模型定义的表或特殊表的场景,直接操作表名而不是模型的情况，通常用于查询非标准表或没有模型定义的表。
	db.Table("deleted_users").Pluck("Gender", &Genders)

	// Scan，扫描结果保存至一个 struct.
	db.Raw("SELECT name,Gender FROM users WHERE Gender = ?", "male").Scan(&[]Userinfo{})

	// Method Chaining，Gorm 实现了链式操作接口，所以你可以把代码写成这样：

	// 创建一个查询
	// tx := db.Where("name = ?", "jinzhu")

	// // 添加更多条件
	// if someCondition {
	//   tx = tx.Where("age = ?", 20)
	// } else {
	//   tx = tx.Where("age = ?", 30)
	// }

	// if yetAnotherCondition {
	//   tx = tx.Where("active = ?", 1)
	// }

	// 在调用立即执行方法前不会生成Query语句，借助这个特性你可以创建一个函数来处理一些通用逻辑。
	// 立即执行方法是指那些会立即生成SQL语句并发送到数据库的方法, 他们一般是CRUD方法，比如：
	// Create, First, Find, Take, Last, Save, FirstOrInit, FirstOrCreate, UpdateXXX, Delete, Scan, Row, Rows, Count
	// 在Gorm中，钩子方法（Hook Methods）是指在模型生命周期的特定时刻自动触发的函数。
	// 这些钩子方法允许开发者在数据库操作（如创建、更新、删除等）前后执行特定的逻辑。
	// 通过使用钩子方法，可以在执行数据库操作时自动进行数据验证、日志记录、缓存清理等操作。

	// Scope是建立在链式操作的基础之上的。基于它，你可以抽取一些通用逻辑，写出更多可重用的函数库

	// 在 GORM 中使用多个立即执行方法时，后一个立即执行方法会复用前一个立即执行方法的条件 (不包括内联条件) 。
	db.Where("name LIKE ?", "jinzhu%").Find(&Userinfo{}, "id IN (?)", []int{1, 2, 3}).Count(&count)

	// Save()默认会更新该对象的所有字段，即使你没有赋值。
	// var user Userinfo
	// db.First(&user)
	// user.Name = "lucy"
	// user.Gender = "male"
	// db.Save(&user)

	// 更新单个属性，如果它有变化
	// db.Model(&user).Update("name", "hello")
	// //// UPDATE users SET name='hello', updated_at='2013-11-17 21:34:10' WHERE id=111;

	// // 根据给定的条件更新单个属性
	// db.Model(&user).Where("active = ?", true).Update("name", "hello")
	// //// UPDATE users SET name='hello', updated_at='2013-11-17 21:34:10' WHERE id=111 AND active=true;

	// // 使用 map 更新多个属性，只会更新其中有变化的属性
	// db.Model(&user).Updates(map[string]interface{}{"name": "hello", "age": 18, "active": false})
	// //// UPDATE users SET name='hello', age=18, active=false, updated_at='2013-11-17 21:34:10' WHERE id=111;

	// // 使用 struct 更新多个属性，只会更新其中有变化且为非零值的字段
	// db.Model(&user).Updates(User{Name: "hello", Age: 18})
	// //// UPDATE users SET name='hello', age=18, updated_at = '2013-11-17 21:34:10' WHERE id = 111;

	// // 警告：当使用 struct 更新时，GORM只会更新那些非零值的字段
	// // 对于下面的操作，不会发生任何更新，"", 0, false 都是其类型的零值
	// db.Model(&user).Updates(User{Name: "", Age: 0, Active: false})

	// db.Model(&user).Select("name").Updates(map[string]interface{}{"name": "hello", "age": 18, "active": false})
	// UPDATE users SET name='hello', updated_at='2013-11-17 21:34:10' WHERE id=111;

	// 上面的更新操作会自动运行 model 的 BeforeUpdate, AfterUpdate 方法，更新 UpdatedAt 时间戳,
	// 在更新时保存其 Associations,如果你不想调用这些方法，你可以使用 UpdateColumn， UpdateColumns
	// db.Model(&User{}).UpdateColumn("name","hello")
	// db.Model(&Userinfo{}).UpdateColumns(Userinfo{Name: "jack", Gender: "male"})

	// 批量更新时Hooks（钩子函数）不会运行。
	//// UPDATE users SET name='hello', age=18 WHERE id IN (10, 11);
	db.Table("users").Where("in in (?)", []int64{1, 2, 3, 4, 5}).Updates(map[string]interface{}{"name": "test", "Gender": "Unknown"})
	rownum := db.Model(Userinfo{}).Updates(&Userinfo{Name: "hello", Gender: "world"}).RowsAffected
	fmt.Println(rownum)

	var user User
	db.First(&user)
	//// UPDATE "users" SET "age" = age - 1 WHERE "id" = '1' AND quantity > 10;
	db.Model(&user).Where("age > 10").UpdateColumn("age", gorm.Expr("age - ?", 1))

	// 删除记录时，请确保主键字段有值，GORM 会通过主键去删除记录，如果主键为空，GORM 会删除该 model 的所有记录。
	// 如果一个 model 有 DeletedAt 字段，他将自动获得软删除的功能！
	db.Delete(&user)                                       // 通过实例删除
	db.Delete(&User{}, user.ID)                            // 通过主键删除
	db.Where("Gender = ?", "fmale").Delete(&User{})        // 条件批量删除
	db.Unscoped().Where("Gender = male").Find(&Userinfo{}) // Unscoped 方法可以查询被软删除的记录
	db.Unscoped().Delete(&user)                            // Unscoped 方法可以物理删除记录

}
