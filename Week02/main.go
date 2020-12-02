package main

import (
	"database/sql"
	"fmt"

	"github.com/pkg/errors"
)

//## Week02 作业题目：
//1. 我们在数据库操作的时候，比如 dao 层中当遇到一个 sql.ErrNoRows 的时候，是否应该 Wrap 这个 error，抛给上层。为什么，应该怎么做请写出代码？

func main() {
	err := CallXXXService()
	// Cause 方法用来判断底层错误
	if errors.Cause(err) == sql.ErrNoRows {
		//fmt.Printf("%v\n", err) // 使用 %v 作为格式化参数，那么错误信息会保持一行
		fmt.Printf("%+v\n", err) // 使用 %+v 会输出完整的调用栈信息
		return
	}

	if err != nil {
		//fmt.Println(err)
	}

	fmt.Println("ok")
}

func GetXXXDao() error {
	//return nil
	//return errors.WithStack(sql.ErrNoRows)
	//return errors.WithMessage(sql.ErrNoRows, "GetDao failed")

	// Wrap: 方法用来包装底层错误，增加上下文文本信息并附加调用栈
	return errors.Wrap(sql.ErrNoRows, "GetXXXDao failed")
}

func CallXXXService() error {
	// WithMessage: 方法仅增加上下文文本信息，不附加调用栈
	return errors.WithMessage(GetXXXDao(), "CallXXXService failed")
}
