package das

/*		---->DasSource	//是所有取数源的父类
*     Das
*		---->DasDest	//是所有写数的父类
 */

type DasBase struct {
	AppStatus int //运行状态
	Parent    interface{}
}

type DasSource struct {
	DasBase
	ConfFile string //配置文件
}

type DasDest struct {
	DasBase
	ConfFile string //配置文件
}
