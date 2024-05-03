package main

import (
	"fmt"
	"os"
	"reflect"
	"strings"
)

// 反射是指在程序运行期间对程序本身进行访问和修改的能力。
// 反射就是在运行时动态的获取一个变量的类型信息和值信息。

func reflectType(x interface{}) {
	t := reflect.TypeOf(x)
	v := reflect.ValueOf(x)
	fmt.Printf("type:%v, value: %v, kind: %v \n", t, v, t.Kind())
}
func reflectTest() {
	var a float32 = 3.14
	reflectType(a) // type:float32
	var b int64 = 100
	reflectType(b) // type:int64
}

type myInt int32

func KindTest() {
	var a *float32 // 指针
	var b myInt    // 自定义类型
	var c rune     // 类型别名
	reflectType(a) // type: kind:ptr
	reflectType(b) // type:myInt kind:int64
	reflectType(c) // type:int32 kind:int32

	type person struct {
		name string
		age  int
	}
	type book struct{ title string }
	var d = person{
		name: "沙河小王子",
		age:  18,
	}
	var e = book{title: "《跟小王子学Go语言》"}
	reflectType(d) // type:person kind:struct
	reflectType(e) // type:book kind:struct
}

func reflectValue1(x interface{}) {
	v := reflect.ValueOf(x)
	k := v.Kind()
	switch k {
	case reflect.Int64:
		// v.Int()从反射中获取整型的原始值，然后通过int64()强制类型转换
		fmt.Printf("type is int64, value is %d\n", int64(v.Int())) // v.Int()从接口中获取原本的int值
	case reflect.Float32:
		// v.Float()从反射中获取浮点型的原始值，然后通过float32()强制类型转换
		fmt.Printf("type is float32, value is %f\n", float32(v.Float()))
	case reflect.Float64:
		// v.Float()从反射中获取浮点型的原始值，然后通过float64()强制类型转换
		fmt.Printf("type is float64, value is %f\n", float64(v.Float()))
	}
}

func reflectTest1() {
	a := 3.134
	var b int64 = 100
	reflectValue1(a)
	reflectValue1(b)
}

func reflectTest2(x interface{}) {
	v := reflect.ValueOf(x)
	if v.Elem().Kind() == reflect.Int64 { // 反射中使用专有的Elem()方法来获取指针对应的值。
		v.Elem().SetInt(200)
	}
}
func reflectTest3() {
	var a int64 = 100
	reflectTest2(&a)
	fmt.Println(a)
}

func IsValueorNil() {
	// *int类型空指针
	var a *int
	fmt.Println("var a *int IsNil:", reflect.ValueOf(a).IsNil())
	// nil值
	fmt.Println("nil IsValid:", reflect.ValueOf(nil).IsValid())
	// 实例化一个匿名结构体
	b := struct{}{}
	// 尝试从结构体中查找"abc"字段
	fmt.Println("不存在的结构体成员:", reflect.ValueOf(b).FieldByName("abc").IsValid())
	// 尝试从结构体中查找"abc"方法
	fmt.Println("不存在的结构体方法:", reflect.ValueOf(b).MethodByName("abc").IsValid())
	// map
	c := map[string]int{}
	// 尝试从map中查找一个不存在的键
	fmt.Println("map中不存在的键:", reflect.ValueOf(c).MapIndex(reflect.ValueOf("娜扎")).IsValid())
}

// 结构体反射
type student struct {
	Name  string `json:"name"`
	Score int    `json:"score"`
}

func StructReflect() {
	stu1 := student{
		Name:  "小王子",
		Score: 90,
	}

	t := reflect.TypeOf(stu1)
	fmt.Println(t.Name(), t.Kind()) // student struct
	// 通过for循环遍历结构体的所有字段信息
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fmt.Printf("name:%s index:%d type:%v json tag:%v\n", field.Name, field.Index, field.Type, field.Tag.Get("json"))
	}

	// 通过字段名获取指定结构体字段信息
	if scoreField, ok := t.FieldByName("Score"); ok {
		fmt.Printf("name:%s index:%d type:%v json tag:%v\n", scoreField.Name, scoreField.Index, scoreField.Type, scoreField.Tag.Get("json"))
	}
}

// 给student添加两个方法 Study和Sleep(注意首字母大写)
func (s student) Study() string {
	msg := "好好学习，天天向上。"
	fmt.Println(msg)
	return msg
}

func (s student) Sleep() string {
	msg := "好好睡觉，快快长大。"
	fmt.Println(msg)
	return msg
}

func printMethod(x interface{}) {
	t, v := reflect.TypeOf(x), reflect.ValueOf(x)
	fmt.Println(t.NumMethod())
	for i := 0; i < v.NumMethod(); i++ {
		fmt.Printf("name : %s\n", t.Method(i).Name)
		fmt.Printf("method: %s\n", v.Method(i).Type())
		// 通过反射调用方法传递的参数必须是 []reflect.Value 类型
		args := []reflect.Value{}
		v.Method(i).Call(args)
	}
}

// 基于反射的代码是极其脆弱的，反射中的类型错误会在真正运行的时候才会引发panic，那很可能是在代码写完的很长时间之后。
// 大量使用反射的代码通常难以理解。
// 反射的性能低下，基于反射实现的代码通常比正常代码运行速度慢一到两个数量级。

//编写代码利用反射实现一个ini文件的解析器程序
// 1. 打开文件；2.读取配置并保存；3.关闭文件

type IniConfig struct {
	Sections map[string]map[string]string `reflect:"sections"`
}

func readIni() {
	data, err := os.ReadFile("config.ini")
	if err != nil {
		fmt.Println(err)
		return
	}

	lines := strings.Split(string(data), "\n")
	config := &IniConfig{Sections: make(map[string]map[string]string)}

	var sectionName string
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "[") && strings.HasSuffix(line, "]") {
			sectionName = strings.TrimPrefix(strings.TrimSuffix(line, "]"), "[")
			config.Sections[sectionName] = make(map[string]string)
			continue
		}
		if strings.Contains(line, "=") {
			parts := strings.Split(line, "=")
			key := strings.TrimSpace(parts[0])
			value := strings.TrimSpace(parts[1])
			config.Sections[sectionName][key] = value
		}
	}

	// Access data using reflection (for learning only)
	v := reflect.ValueOf(config).Elem()
	// sectionType := v.Type().Field(0).Type
	for _, sectionName := range v.MapKeys() {
		sectionValue := v.MapIndex(sectionName)
		sectionMap := sectionValue.Interface().(map[string]string)
		fmt.Printf("[Section]: %s\n", sectionName.String())
		for key, value := range sectionMap {
			fmt.Printf("  %s = %s\n", key, value)
		}
	}
}
