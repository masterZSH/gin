
用于绑定各种数据类型


```go
// Binding 接口
type Binding interface {
	Name() string
	Bind(*http.Request, interface{}) error
}

// 比上面多了一个BindBody方法
type BindingBody interface {
	Binding
	BindBody([]byte, interface{}) error
}

// 绑定Uri接口
type BindingUri interface {
	Name() string
	BindUri(map[string][]string, interface{}) error
}
```


```go
// 结构体验证接口
type StructValidator interface {
	// ValidateStruct can receive any kind of type and it should never panic, even if the configuration is not right.
	// If the received type is not a struct, any validation should be skipped and nil must be returned.
	// If the received type is a struct or pointer to a struct, the validation should be performed.
	// If the struct is not valid or the validation itself fails, a descriptive error should be returned.
	// Otherwise nil must be returned.
	ValidateStruct(interface{}) error

	// Engine returns the underlying validator engine which powers the
	// StructValidator implementation.
	Engine() interface{}
}
```