// Create and maintain by Chaiyapong Lapliengtrakul (chaiyapong@3dsinteractive.com), All right reserved (2021 - Present)
package main

import (
	"fmt"
)

func main() {

	ms := NewMicroservice()

	ms.CommandLine(
		"setup",
		"Setup Application",
		"Setup Application",
		[]*CommandArg{},
		[]*CommandArg{
			ArgVaultToken, // vault token
			ArgDeploy,     // prefix of service name to deploy
			ArgBaseDir,    // base directory
		},
		func(ctx IContext) error {
			// Deploy is prefix of service name to deploy
			deploy := ctx.Param(ArgDeploy.Name)
			if len(deploy) == 0 {
				deploy = "tcir"
			}
			deploy = escapeName(deploy)

			baseDir := ctx.Param(ArgBaseDir.Name)
			if len(baseDir) == 0 {
				baseDir = "/root/automation-tech-k8s/automation-technology/app/"
			}
			// Setup path to read file (this can be read from env)
			buildsDir := baseDir + "builds"
			svcFilePath := baseDir + "automation-technology-svc.yml"
			envFilePath := baseDir + "out-env.yml"
			ns := "tcir-app"
			kubeCfg := "" // microk8s no config
			// kubeCfg := "/root/.kube/config"

			// Create out-xxx deployment files
			outFile := fmt.Sprintf("out-%s.yml", deploy)
			tmplURLPrefix := "https://raw.githubusercontent.com/3dsinteractive/automation-tech-deployment-tmpl/master/"
			err := generateDeploymentFile(ms, ctx, buildsDir, svcFilePath, envFilePath, deploy, ns, outFile, tmplURLPrefix)
			if err != nil {
				ms.Log("CMD", err.Error())
				return err
			}

			// Create namespace if not exists
			createNamespaceIfNotExists(ms, ctx, kubeCfg, ns)
			if err != nil {
				ms.Log("CMD", err.Error())
				return err
			}

			// Apply out file
			cmds := createKubectlCommand("apply", kubeCfg, outFile, ns)
			_, _, err = execCommands(ms, ctx, cmds)
			if err != nil {
				ms.Log("CMD", err.Error())
				return err
			}

			response := map[string]interface{}{
				"success": true,
			}
			ctx.Response(200, response)
			ms.Log("CMD", "exit")
			return nil
		})
}
