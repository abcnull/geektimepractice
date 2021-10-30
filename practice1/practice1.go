package main

import (
	"database/sql"
	"fmt"
	"github.com/pkg/errors"
)

/*
日期：
2021/10/31

问题：
1.我们在数据库操作的时候，比如 dao 层中当遇到一个 sql.ErrNoRows 的时候，是否应该 Wrap 这个 error，抛给上层。为什么，应该怎么做请写出代码？

回答：
需要 Wrap 住 error 保留堆栈信息往上抛
*/

// Foo 使用 Wrap 包装 sql.ErrNoRows
func Foo() error {
	// Wrap 包装根错误，保留堆栈信息
	return errors.Wrap(sql.ErrNoRows, "Foo() failed")
}

// Bar 上层调用使用 WithMessage 进一步封装
func Bar() error {
	return errors.WithMessage(Foo(), "Bar() failed")
}

// main 主函数
func main() {
	// 上层来调用 Bar()
	err := Bar()
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			// 如果匹配到根错误，打印完整堆栈信息
			fmt.Printf("打印错误信息: %+v\n", err)
			return
		} else {
			// todo: sth else
			return
		}
	}

	// todo: sth
}
