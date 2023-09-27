# ToDoList
protoc的信息传输流载体，用于连接各个系统

## 生成代码
#### 需要首先编写proto文件、proto文件中需要被调用的方法应与svc中的相对应。
```bash
protoc -I ./proto userService.proto --micro_out ./proto/pb --go_out=./proto/pb
protoc -I ./proto taskService.proto --micro_out ./proto/pb --go_out=./proto/pb
```



