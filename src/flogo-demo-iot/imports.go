package main

import (
	_ "github.com/TIBCOSoftware/flogo-contrib/action/flow"
	_ "github.com/TIBCOSoftware/flogo-contrib/activity/actreply"
	_ "github.com/TIBCOSoftware/flogo-contrib/activity/log"
	_ "github.com/TIBCOSoftware/flogo-contrib/trigger/rest"
	_ "github.com/debovema/flogo-opentracing-listener"
	_ "github.com/square-it/flogo-contrib-activities/command"
	_ "github.com/square-it/flogo-contrib-activities/copyfile"
)
