package main

import "github.com/zeromicro/go-zero/core/logx"

func main() {

	logx.MustSetup(logx.LogConf{
		Mode: "file",
		Path: "./logs",
	})
	logx.Error("测试")

	//for i := 9; i < 100; i++ {
	//	r, _ := http.Post("http://175.178.49.188:30389/base/login", "", bytes.NewBufferString(""))
	//
	//	b, _ := ioutil.ReadAll(r.Body)
	//	fmt.Println(string(b))
	//}
}
