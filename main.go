/*!
 * Copyright 2019 - Author gauravm.git@gmail.com
 */

package main

import (
	"github.com/gauravmahto/go-start/utils"
	"github.com/gauravmahto/go-webserver/framework/web"
)

func main() {

	utils.PrintLine("App started")

	web.Server()

}
