package control_structure

/**
if-else 结构
switch 结构
select 结构
for (range) 结构

break continue return goto

Go 完全省略了结构中条件语句两侧的括号 使代码更加简洁
*/

/*
if condition1 {
    // do something
} else if condition2 {
    // do something else
} else {
    // catch-all or default
}
*/

func IfDemo() string {
	b1 := true
	if b1 {
		return "true e"
	} else {
		return "false e"
	}
}
