package main

import (
	"day10/dal/model"

	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
)

// gorm gen configure

const MYSQLDNS = "root:root1234@tcp(127.0.0.1:13306)/GormTest?charset=utf8mb4&parseTime=True"

func ConnetDB(dsn string) *gorm.DB {
	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		panic("DB connet failed!")
	}
	return db
}

func Generate() {
	// 指定生成代码的具体相对目录(相对当前文件)，默认为：./query
	// 默认生成需要使用WithContext之后才可以查询的代码，但可以通过设置gen.WithoutContext禁用该模式
	g := gen.NewGenerator(gen.Config{
		// 默认会在 OutPath 目录生成CRUD代码，并且同目录下生成 model 包
		// 所以OutPath最终package不能设置为model，在有数据库表同步的情况下会产生冲突
		// 若一定要使用可以通过ModelPkgPath单独指定model package的名称
		OutPath: "./dal/query",
		// ModelPkgPath: "./model",
		// gen.WithoutContext：禁用WithContext模式
		// gen.WithDefaultQuery：生成一个全局Query对象Q
		// gen.WithQueryInterface：生成Query接口
		Mode: gen.WithDefaultQuery | gen.WithQueryInterface,
	})

	// 通常复用项目中已有的SQL连接配置db(*gorm.DB)非必需，但如果需要复用连接时的gorm.Config或需要连接数据库同步表信息则必须设置
	g.UseDB(ConnetDB(MYSQLDNS))

	// 从连接的数据库为所有表生成Model结构体和CRUD代码,可以手动指定需要生成代码的数据表
	g.ApplyBasic(g.GenerateAllTable()...)

	g.GenerateModel("users")               // users表生成UseraModel
	g.GenerateModelAs("users", "Employee") // users表生成EmployeeModel
	g.GenerateModel("users", gen.FieldIgnore("address"), gen.FieldType("id", "int64"))

	// 通过ApplyInterface为book表添加自定义方法
	g.ApplyInterface(func(model.Querier) {}, g.GenerateModel("book"))

	// 执行并生成代码
	g.Execute()

}

type CommonMethod struct {
	ID   int32
	Name *string
}

func (m *CommonMethod) IsEmpty() bool {
	if m == nil {
		return true
	}
	return m.ID == 0
}
func (m *CommonMethod) GetName() string {
	if m == nil || m.Name == nil {
		return ""
	}
	return *m.Name
}

// 当生成 `People` 结构体时添加 IsEmpty 方法
g.GenerateModel("people", gen.WithMethod(CommonMethod{}.IsEmpty))

// 生成`User`结构体时添加 `CommonMethod` 的所有方法
g.GenerateModel("user", gen.WithMethod(CommonMethod{}))

// 可以自行指定字段类型和数据库列类型之间的数据类型映射。
// 在某些业务场景下，这个功能非常有用，例如，我们希望将数据库中数字列在生成结构体时都定义为int64类型
var dataMap = map[string]func(gorm.ColumnType) (dataType string){
  // int mapping
  "int": func(columnType gorm.ColumnType) (dataType string) {
    if n, ok := columnType.Nullable(); ok && n {
      return "*int32"
    }
    return "int32"
  },

  // bool mapping
  "tinyint": func(columnType gorm.ColumnType) (dataType string) {
    ct, _ := columnType.ColumnType()
    if strings.HasPrefix(ct, "tinyint(1)") {
      return "bool"
    }
    return "byte"
  },
}

g.WithDataTypeMap(dataMap)
