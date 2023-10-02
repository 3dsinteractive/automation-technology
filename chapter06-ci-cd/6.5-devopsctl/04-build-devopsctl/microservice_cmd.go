// Create and maintain by Chaiyapong Lapliengtrakul (chaiyapong@3dsinteractive.com), All right reserved (2021 - Present)
package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/thatisuday/commando"
)

const (
	// When define non required flag with empty default value, the library need one value instead of accept empty default value
	// So we replace empty value with this UUID, and will replace it with empty value when get with ctx.Param("name")
	// We expect you don't use this value as the value of flag, If you use this value, please change!
	EMPTY_DEFAULT_VALUE = "22c211c6-8774-11eb-8dcd-0242ac130003"
)

// CommandArg is arg or flag to command
type CommandArg struct {
	Name         string
	ShortName    string
	Description  string
	DefaultValue string
	IsRequired   bool
}

// CommandLine is one time runner task
type CommandLine struct {
	Handler          ServiceHandleFunc
	Command          string
	Description      string
	ShortDescription string
	Args             []*CommandArg
	Flags            []*CommandArg
}

// StartCommandLines will start all schedulers
func (ms *Microservice) CommandLine(
	command string,
	description string,
	shortDesc string,
	args []*CommandArg,
	flags []*CommandArg,
	handler ServiceHandleFunc) error {

	registry := commando.NewCommandRegistry()
	cmd := &CommandLine{
		Handler:          handler,
		Command:          command,
		Description:      description,
		ShortDescription: shortDesc,
		Args:             args,
		Flags:            flags,
	}

	var cli *commando.Command
	if len(cmd.Command) > 0 {
		// Sub command eg. (program command -flag1 -flag2)
		cli = registry.Register(cmd.Command)
	} else {
		// Root command (program -flag1 -flag2)
		cli = registry.Register(nil)
	}
	if len(cmd.Description) > 0 {
		cli.SetDescription(cmd.Description)
	}
	if len(cmd.ShortDescription) > 0 {
		cli.SetShortDescription(cmd.ShortDescription)
	}

	for _, arg := range cmd.Args {
		if len(arg.Name) == 0 {
			ms.Log("CMD", "Command line argument name must not empty.")
			continue
		}
		if len(arg.ShortName) > 0 {
			if arg.IsRequired {
				cli.AddArgument(fmt.Sprintf("%s,%s", arg.Name, arg.ShortName), arg.Description, "")
			} else {
				cli.AddArgument(fmt.Sprintf("%s,%s", arg.Name, arg.ShortName), arg.Description, arg.DefaultValue)
			}
		} else {
			if arg.IsRequired {
				cli.AddArgument(arg.Name, arg.Description, "")
			} else {
				cli.AddArgument(arg.Name, arg.Description, arg.DefaultValue)
			}
		}
	}

	for _, flag := range cmd.Flags {
		if len(flag.Name) == 0 {
			ms.Log("CMD", "Command line flag must not empty.")
			continue
		}
		if len(flag.ShortName) > 0 {
			if flag.IsRequired {
				cli.AddFlag(fmt.Sprintf("%s,%s", flag.Name, flag.ShortName), flag.Description, commando.String, nil)
			} else {
				if len(flag.DefaultValue) == 0 {
					// When define non required flag with empty default value, the library need one value instead of accept empty default value
					// So we replace empty value with this UUID, and will replace it with empty value when get with ctx.Param("name")
					// We expect you don't use this value as the value of flag, If you use this value, please change!
					// EMPTY_DEFAULT_VALUE = "22c211c6-8774-11eb-8dcd-0242ac130003"
					cli.AddFlag(fmt.Sprintf("%s,%s", flag.Name, flag.ShortName), flag.Description, commando.String, EMPTY_DEFAULT_VALUE)
				} else {
					cli.AddFlag(fmt.Sprintf("%s,%s", flag.Name, flag.ShortName), flag.Description, commando.String, flag.DefaultValue)
				}
			}
		} else {
			if flag.IsRequired {
				cli.AddFlag(flag.Name, flag.Description, commando.String, nil)
			} else {
				if len(flag.DefaultValue) == 0 {
					// When define non required flag with empty default value, the library need one value instead of accept empty default value
					// So we replace empty value with this UUID, and will replace it with empty value when get with ctx.Param("name")
					// We expect you don't use this value as the value of flag, If you use this value, please change!
					// EMPTY_DEFAULT_VALUE = "22c211c6-8774-11eb-8dcd-0242ac130003"
					cli.AddFlag(flag.Name, flag.Description, commando.String, EMPTY_DEFAULT_VALUE)
				} else {
					cli.AddFlag(flag.Name, flag.Description, commando.String, flag.DefaultValue)
				}
			}
		}
	}

	// Capture variable to use in closure
	cmd = cmd
	action := func(args map[string]commando.ArgValue, flags map[string]commando.FlagValue) {
		if cmd.Handler == nil {
			ms.Log("CMD", "Command line handler must not empty.")
			return
		}

		inputs := make(map[string]string)
		for k, v := range args {
			inputs[k] = v.Value
		}
		for k, v := range flags {
			inputs[k] = fmt.Sprintf("%v", v.Value)
		}

		ctx := NewCmdContext(ms, inputs)
		err := cmd.Handler(ctx)
		if err != nil {
			ms.Log("CMD", err.Error())
			ms.Stop()
		}
	}
	cli.SetAction(action)

	// Capture Ctrl+C to terminate program
	go func() {
		osQuit := make(chan os.Signal, 1)
		signal.Notify(osQuit, syscall.SIGTERM, syscall.SIGINT)

		<-osQuit

		ms.Log("CMD", "Receive SigTerm (source=Engine.StartCommandLines)")

		ms.Cleanup()
		// When exit from command line we use os.Exit
		os.Exit(1)
	}()

	// execute command
	registry.Parse(nil)

	return nil
}
