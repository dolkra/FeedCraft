package craft

// craft template
// craft模版, 如limit, translate等, 定义不同的参数可以派生出多种craft atom, 如limit5, translate-cn等

type CraftTemplate struct {
	Name                string
	Description         string
	ParamTemplateDefine []ParamTemplate // param 格式, 主要给用户填写时提供参考
	OptionFunc          func(map[string]string) []CraftOption
}

func (tmpl CraftTemplate) GetOptions(params map[string]string) []CraftOption {
	return tmpl.OptionFunc(params)
}

// ParamTemplate 后续传递给craft template 参数全部使用 map[string]string, 这里的option template 是记录每个字段的名字和含义
type ParamTemplate struct {
	Key         string
	Description string
	Default     string
}

// CraftAtom craft atom 由 craft template 派生出来
type CraftAtom struct {
	Name         string
	Description  string
	TemplateName string
	Params       map[string]string
}