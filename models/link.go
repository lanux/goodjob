package models

type Links struct {
	Model
	Name        string `comment:"名称"`
	Title       string `comment:"标题"`
	Href        string `comment:"链接地址"`
	Icon        string `comment:"图标"`
	BgColor     string `comment:"背景颜色"`
	Color       string `comment:"字体颜色"`
	Description string `comment:"描述"`
	Profile     string `comment:"分组"`
	Status      int8   `comment:"状态(字典)"`
}

type LinkGroup struct {
	Model
	Name   string `comment:"名称"`
	Title  string `comment:"标题"`
	Icon   string `comment:"图标"`
	Status int8   `comment:"状态(字典)"`
}
