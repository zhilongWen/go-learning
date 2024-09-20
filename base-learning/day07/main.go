package main

import (
	"fmt"
	"reflect"
)

func main() {

	// https://segmentfault.com/a/1190000044313900

	println("hello world...")

	var x float64 = 3.14159
	fmt.Println("type : ", reflect.TypeOf(x)) // type :  float64

	p := &Person{Name: "tom"}
	value := reflect.ValueOf(p)
	method := value.MethodByName("Talk")

	method.Call(nil)

	t := reflect.TypeOf(x)
	v := reflect.ValueOf(x)
	fmt.Println("Type:", t)  // Type: float64
	fmt.Println("Value:", v) // Value: 3.14159

	LoadConfig("{\"field1\":\"1\",\"field2\":\"2\"}", &Config{})
}

type Person struct {
	Name string
}

func (p *Person) Talk() {
	fmt.Println("hi , my name is ", p.Name)
}

// Config 配置化
type Config struct {
	Field1 string `json:"field1"`
	Field2 string `json:"field2"`
}

func LoadConfig(str string, config *Config) {
	v := reflect.ValueOf(config).Elem()
	t := v.Type()
	// 省略JSON解析步骤
	// 动态设置字段值
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		// 使用jsonTag从JSON数据中获取相应的值，并设置到结构体字段
		jsonTag := field.Tag.Get("json")
		fmt.Println(jsonTag)
	}
}

// Plugin 插件化 ：动态加载插件
type Plugin interface {
	PerformAction()
}

func LoadPlugin(pluginName string) Plugin {
	// 使用反射来动态创建插件实例

	return nil
}
