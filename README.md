# protoc-gen-go-errors
用于生成在grpc调用时，需要被处理的错误的创建和判断方法；为每一个错误携带 `http status code` 以及 `biz code(业务错误码)`.

## 介绍
通过如下定义
```proto
enum TestErrorReason {
    option (errors.settings) = {
        default_http_code: 500
        start_biz_code: 100001
    };

    TestNotFound = 0 [(errors.code) = {http_code:404}];
    TestBusy = 1;
    TestIncrease = 2 [(errors.code) = {http_code:502 biz_code:100010}];
    TestRedirect = 3 [(errors.code) = {http_code:302}];
}
```

可以生成如下go文件

```go

var bizErrorCodeMap map[string]int = map[string]int{

	"errors.test_TestErrorReason_TestNotFound": 100001,
	"errors.test_TestErrorReason_TestBusy":     100002,
	"errors.test_TestErrorReason_TestIncrease": 100010,
	"errors.test_TestErrorReason_TestRedirect": 100011,
}

func IsTestnotfound(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == "errors.test_TestErrorReason_TestNotFound" && e.Code == 404
}

func ErrorTestnotfound(format string, args ...interface{}) *errors.Error {
	return errors.New(404, "errors.test_TestErrorReason_TestNotFound", fmt.Sprintf(format, args...))
}

// ...

func BizErrorCode(err error) int {
	if err == nil {
		return 0
	}
	e := errors.FromError(err)
	return bizErrorCodeMap[e.Reason]
}

```

## 使用案例

1. build the binary util
```shell
# 下载依赖
make init

# 安装可执行文件
make build
```
2. use it

```shell
make test
```