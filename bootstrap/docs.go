package bootstrap

import (
	"strconv"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/toolbox"
)

func init() {
	updateTask := toolbox.NewTask("check file update", "0 */5 * * * *", checkFileUpdates)

	if needCheckUpdate() {
		if err := updateTask.Run(); err != nil {
			beego.Error(err)
		}

		beego.AppConfig.Set("app::update_check_time", strconv.Itoa(int(time.Now().Unix())))
	}

	// ATTENTION: you'd better comment following code when developing.
	toolbox.AddTask("check file update", updateTask)
	toolbox.StartTask()
}
