// @Author: Ciusyan 1/16/24

package reflect_module

import (
	"fmt"
	"log"
	"reflect"
)

type ModuleSample struct {
	ProjectNamespaceId int
	StructName         string
	NickName           string
	Description        string
	SelfParamDesc      map[string]*SelfParam
	DependParamDesc    map[int]*DependParam
	ReturnParamDesc    *DependParam
	Status             int //模版是否可用
}

type SelfParam struct {
	Type        string `json:"type"`         //参数数据类型
	DefultValue string `json:"defult_value"` //参数默认值
}

type DependParam struct {
	Type       string               `json:"type"`        //表示是一个struct，还是一个interface
	StructType string               `json:"struct_type"` //如果是一个struct,那么就是reflect.type值
	Methods    []*DependParamMethod `json:"methods"`     //如果是一个interface，表示包含的方法
}

type DependParamMethod struct {
	Name       string `json:"name"`        //方法的名称
	ReturnType string `json:"return_type"` //方法返回值类型
}

const (
	BASE_TYPE = "base type"
)

func parseInParam(method reflect.Value, moduleSample *ModuleSample) (err error) {
	defer func() {
		if recoverErr := recover(); recoverErr != nil {
			err = fmt.Errorf("%v", recoverErr)
		}
	}()

	// 获取 method 参数的数量
	paramInNumber := method.Type().NumIn()
	// 至少要有两个参数
	if paramInNumber > 2 {
		dependParamDesc := make(map[int]*DependParam)
		// 从第三个参数开始解析，前两个参数固定位 go context 和 strategy context
		for i := 2; i < paramInNumber; i++ {
			dependParam := &DependParam{}
			// 获取当前参数的类型
			kind := method.Type().In(i).Kind()

			// 根据参数类型，进行不同的处理
			if kind == reflect.Ptr {
				// 如果是指针类型，检查指向是否是结构体
				if method.Type().In(i).Elem().Kind() == reflect.Struct {
					dependParam.Type = fmt.Sprintf("%v", reflect.Struct)
				} else {
					// 不是则代表基本类型
					dependParam.Type = BASE_TYPE
				}
			} else if kind == reflect.Interface {
				// 对于接口类型，获取接口的方法数量
				interfaceMethodNumber := method.Type().In(i).NumMethod()
				if interfaceMethodNumber > 0 {
					dependParamMethods := make([]*DependParamMethod, 0)
					// 遍历接口所有的方法
					for j := 0; j < interfaceMethodNumber; j++ {
						dependParamMethod := &DependParamMethod{
							Name:       method.Type().In(i).Method(j).Name,                           // 方法名
							ReturnType: fmt.Sprintf("%v", method.Type().In(i).Method(j).Type.Out(0)), //只会有一个返回值
						}
						dependParamMethods = append(dependParamMethods, dependParamMethod)
					}
					dependParam.Methods = dependParamMethods
				}
				dependParam.Type = fmt.Sprintf("%v", reflect.Interface)
			} else if kind == reflect.Struct {
				// 结构体类型
				dependParam.Type = fmt.Sprintf("%v", reflect.Struct)
			} else {
				// 其余的统统标识为基本类型
				dependParam.Type = BASE_TYPE
			}
			// 设置依赖参数的结构类型
			dependParam.StructType = fmt.Sprintf("%v", method.Type().In(i))
			dependParamDesc[i] = dependParam
		}
		moduleSample.DependParamDesc = dependParamDesc
	}
	return nil
}

// 将 method 中的返回值，放置在 moduleSample 中
func parseOutParam(method reflect.Value, moduleSample *ModuleSample) (err error) {
	defer func() {
		if recoverErr := recover(); recoverErr != nil {
			err = fmt.Errorf("%v", recoverErr)
		}
	}()

	returnParam := &DependParam{}

	// 只需要解析第一个返回值即可，第二个参数互撕固定不变的
	if method.Type().Out(0).Kind() == reflect.Ptr {
		// 说明是指针类型，需要看看是否是结构体指针
		if method.Type().Out(0).Elem().Kind() == reflect.Struct {
			// 说明是结构体指针，将指针对应结构体的类型传入进行解析
			parseOutParamMethod(method.Type().Out(0).Elem(), returnParam)
		} else {
			returnParam.Type = BASE_TYPE
		}
	} else if method.Type().Out(0).Kind() == reflect.Struct {
		// 说明返回值直接是一个结构体类型，直接将其类型传入下面的方法去解析
		parseOutParamMethod(method.Type().Out(0), returnParam)
	} else {
		returnParam.Type = BASE_TYPE
	}

	returnParam.StructType = fmt.Sprintf("%v", method.Type().Out(0))

	moduleSample.ReturnParamDesc = returnParam

	return nil
}

// 将 methodType 结构体的参数，解析后放置在 returnParam 中
func parseOutParamMethod(methodType reflect.Type, returnParam *DependParam) {
	returnParam.Type = fmt.Sprintf("%v", reflect.Struct)
	// 创建一个新的类型来模拟结构体的类型
	clone := reflect.New(methodType)
	// 看看这个结构体都实现了什么方法
	if clone.NumMethod() > 0 { //如果返回的是一个struct，且实现了一些方法，保存这些方法的信息
		dependParamMethods := make([]*DependParamMethod, 0)
		for i := 0; i < clone.NumMethod(); i++ {
			dependParamMethod := &DependParamMethod{
				Name:       clone.Type().Method(i).Name,                      // 方法名
				ReturnType: fmt.Sprintf("%v", clone.Method(i).Type().Out(0)), //只会有一个返回值
			}
			dependParamMethods = append(dependParamMethods, dependParamMethod)
		}
		returnParam.Methods = dependParamMethods
	}
}

func parseSelfParam(reflectValue reflect.Value, moduleSample *ModuleSample) (err error) {
	defer func() {
		if recoverErr := recover(); recoverErr != nil {
			err = fmt.Errorf("%v", recoverErr)
		}
	}()

	// 存储参数描述的 Map
	selfParamDesc := make(map[string]*SelfParam)

	// 获取 module 的类型
	reflectType := reflect.TypeOf(reflectValue.Elem().Interface())
	// 获取实际的值对象
	v := reflectValue.Elem()
	// 遍历所有的字段
	for i := 0; i < reflectType.NumField(); i++ {
		// 获取字段信息
		field := reflectType.Field(i)
		// 获取字段的值
		fieldValue := v.FieldByName(field.Name)
		// 验证值是否有效
		if !fieldValue.IsValid() {
			return fmt.Errorf("RegisterModule.fieldValueIsValid")
		}

		// 特殊处理 "NickName" 和 "Description"，获取对应的 default_value tag 的值
		if field.Name == "NickName" {
			moduleSample.NickName = field.Tag.Get("default_value")
			continue
		}

		if field.Name == "Description" {
			moduleSample.Description = field.Tag.Get("default_value")
			continue
		}

		// 其他字段创建 SelfParam 添加到映射中
		selfParam := &SelfParam{
			Type:        fieldValue.Type().String(),
			DefultValue: field.Tag.Get("default_value"),
		}
		selfParamDesc[field.Name] = selfParam
	}

	if len(selfParamDesc) > 0 {
		moduleSample.SelfParamDesc = selfParamDesc
	}
	return nil
}

func Do(moduleSampleObj interface{}) {
	//解析内存中的module信息
	var err error

	reflectValue := reflect.ValueOf(moduleSampleObj)
	moduleSample := &ModuleSample{
		ProjectNamespaceId: 1,
		StructName:         "test",
	}
	//解析module本身参数信息
	err = parseSelfParam(reflectValue, moduleSample)
	if err != nil {
		log.Fatal(11)
	}

	// 获取 module 的 DoAction 方法
	method := reflectValue.MethodByName("DoAction")
	//解析输入参数信息
	err = parseInParam(method, moduleSample)
	if err != nil {
		log.Fatal(11)
	}

	//解析输出参数信息
	err = parseOutParam(method, moduleSample)
	if err != nil {
		log.Fatal(11)
	}

	//将解析后的module信息,上报给管理平台
	fmt.Println(moduleSample)
}
