# Logs with file rotation

Example:
```go
l, err := logs.New(&logs.Config{
    App:      "test",
    FilePath: "main.log",
})
if err != nil {
    t.Fatal(err)
}
l.Info("123")
```
Result
```json
{"level":"info","app":"test","time":1597143777,"datetime":"11.08.2020 16:02:57.9183141","message":"123"}
```








---
Wrapper for https://github.com/rs/zerolog  
Added file rotation