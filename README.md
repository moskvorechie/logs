# Logs with file rotation

Example:
```go
l, err := logs.New(&logs.Config{
    App:      "test",
    FilePath: "main.log",
    Clear:    true,
    ToFileOnly: false,
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

You can add some fields
```
l.Logger().Info().Str("doc_type", "realization").Msg("123")
```
Result
```json
{"level":"info","app":"test","doc_type":"realization","time":1597143777,"datetime":"11.08.2020 16:02:57.9183141","message":"123"}
```







---
Wrapper for https://github.com/rs/zerolog  
Added file rotation