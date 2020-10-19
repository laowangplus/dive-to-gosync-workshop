package main

import (
	"context"
	"fmt"
)

//源代码来看，context.Background 和 context.TODO 函数其实也只是互为别名，没有太大的差别。
// 它们只是在使用和语义上稍有不同：
//
//context.Background 是上下文的默认值，所有其他的上下文都应该从它衍生（Derived）出来；
//context.TODO 应该只在不确定应该使用哪种上下文时使用；

func main() {
	//ctx := context.Background()
	ctx := context.TODO()
	ctx = context.WithValue(ctx, "key1", "0001")
	ctx = context.WithValue(ctx, "key2", "0001")
	ctx = context.WithValue(ctx, "key3", "0001")
	ctx = context.WithValue(ctx, "key4", "0004")

	fmt.Println(ctx.Value("key4"))
}
