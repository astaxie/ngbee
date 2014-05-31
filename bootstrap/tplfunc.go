package bootstrap

import (
	"github.com/astaxie/beego"
	"github.com/beego/compress"
)

func i18n(lang, key string) string {
	if kv, ok := I18N[key]; ok {
		if v, ok := kv[lang]; ok {
			return v
		}
	}
	return key
}

func config(key string) string {
	return beego.AppConfig.String("key")
}

func docs(lang, path string) string {

}

func init() {
	setting, err := compress.LoadJsonConf("conf/compress.json", true, "/")
	if err != nil {
		panic("LoadJsonConf compress.json err:" + err.Error())
	}

	setting.RunCommand()

	setting.RunCompress(true, false, true)

	beego.AddFuncMap("compress_js", setting.Js.CompressJs)
	beego.AddFuncMap("compress_css", setting.Css.CompressCss)
	beego.AddFuncMap("i18n", i18n)
	beego.AddFuncMap("config", config)
	beego.AddFuncMap("docs", docs)
}
