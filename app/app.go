// Copyright 2016 by caixw, All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

// app 提供了一些公共的函数、结构体及设置项。
package app

import (
	"time"

	"github.com/issue9/term/colors"
)

// 一些代码级别的配置项。
const (
	// 版本号
	//
	// 版本号按照 http://semver.org/lang/zh-CN/ 中的规则
	Version = "2.11.87+20160616"

	// 程序的正式名称
	Name = "apidoc"

	// 源码仓库地址
	RepoURL = "https://github.com/caixw/apidoc"

	// 官网
	OfficialURL = "https://caixw.github.io/apidoc"

	// 配置文件名称。
	ConfigFilename = ".apidoc.json"

	// 默认的文档标题
	DefaultTitle = "APIDOC"

	// 默认的分组名称，同时也是默认的索引文件名
	DefaultGroupName = "index"

	// 默认的时间格式，仅对 html 中的模板启作用
	TimeFormat = time.RFC3339
)

// Message 向终端输出不同颜色的信息
//
// color 是输出的字体颜色，仅对 prefix
// 参数起作用，其它字符串依然使用系统默认的颜色。
func Message(color colors.Color, prefix string, v ...interface{}) {
	colors.Print(colors.Stdout, color, colors.Default, prefix)
	colors.Println(colors.Stdout, colors.Default, colors.Default, v...)
}

// Warn 输出警告性的信息
func Warn(v ...interface{}) {
	Message(colors.Cyan, "[WARN] ", v...)
}

// Error 输出错误的信息
func Error(v ...interface{}) {
	Message(colors.Red, "[ERROR] ", v...)
}

// Info 输出提示信息
func Info(v ...interface{}) {
	Message(colors.Green, "[INFO] ", v...)
}