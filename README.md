# go-api-module
存放go中的模型，让这些模型可以被多个项目直接使用

## 模型结构
+ 1层：对应项目
+ + 2层：对应server层
+ + + 3层：对应模型版本

例如：
`iam\apiserver\v1`中，`iam`是项目，`apiserver`是对应的服务，`v1`是第一版本