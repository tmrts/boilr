package main

import (
	"fmt"

	"github.com/Wattpad/boilr/pkg/boilr"
	"github.com/Wattpad/boilr/pkg/cmd"
	"github.com/Wattpad/boilr/pkg/util/exit"
	"github.com/Wattpad/boilr/pkg/util/osutil"
)

func main() {
	if exists, err := osutil.DirExists(boilr.Configuration.TemplateDirPath); !exists {
		if err := osutil.CreateDirs(boilr.Configuration.TemplateDirPath); err != nil {
			exit.Error(fmt.Errorf("tried to initialise your template directory, but it has failed: %s", err))
		}
	} else if err != nil {
		exit.Error(fmt.Errorf("failed to init: %s", err))
	}

	cmd.Run()
}
