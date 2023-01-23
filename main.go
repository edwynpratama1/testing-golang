package main

import "serverLocalGo/Routes"

func mainan() {
	r := Routes.SetupRouter()
	r.Run()
}
