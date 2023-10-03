// Create and maintain by Chaiyapong Lapliengtrakul (chaiyapong@3dsinteractive.com), All right reserved (2021 - Present)
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"gopkg.in/yaml.v2"
)

func createNamespaceIfNotExists(
	ms IMicroservice,
	ctx IContext,
	kubeCfg string,
	ns string) error {

	// Create namespace
	cmd := fmt.Sprintf("create ns %s", ns)
	cmds := createKubectlCommand(cmd, kubeCfg, "", "")
	_, _, err := execCommands(ms, ctx, cmds)
	if err != nil {
		ms.Log("CMD", err.Error())
		return nil
	}
	return nil
}

func execCommands(
	ms IMicroservice,
	ctx IContext,
	commands []string) (string /*std out*/, string /* std err */, error) {

	if len(commands) == 0 {
		return "", "", nil
	}

	ms.Log("CMD", "commands="+strings.Join(commands, " "))
	res, err := execShell(commands)
	ms.Log("CMD", "stdOut="+res.StdOut)
	ms.Log("CMD", "stdErr="+res.StdErr)

	if err != nil {
		ms.Log("CMD", err.Error())
		return res.StdOut, res.StdErr, err
	}
	return res.StdOut, res.StdErr, nil
}

// ExecShell exec shell and print to standard output
func execShell(cmd []string) (*ShellResult, error) {
	if len(cmd) <= 0 {
		return nil, nil
	}

	sh := exec.Command(cmd[0], cmd[1:]...)
	stdOut := &bytes.Buffer{}
	stdErr := &bytes.Buffer{}
	sh.Stdout = stdOut
	sh.Stderr = stdErr

	err := sh.Run()
	if err != nil {
		return &ShellResult{
			StdOut: stdOut.String(),
			StdErr: stdErr.String(),
		}, err
	}
	return &ShellResult{
		StdOut: stdOut.String(),
		StdErr: stdErr.String(),
	}, nil
}

func createKubectlCommand(
	kubeCmd string,
	kubeConfig string,
	file string,
	namespace string,
	moreArgs ...string) []string {

	kubeCmds := strings.Split(kubeCmd, " ")
	cmds := []string{"microk8s", "kubectl"}
	for _, cmd := range kubeCmds {
		if len(cmd) == 0 {
			continue
		}
		cmds = append(cmds, cmd)
	}
	if len(file) > 0 {
		cmds = append(cmds, "-f", file)
	}
	if len(kubeConfig) > 0 {
		cmds = append(cmds, "--kubeconfig", kubeConfig)
	}
	if len(namespace) > 0 {
		cmds = append(cmds, "-n", namespace)
	}
	cmds = append(cmds, moreArgs...)
	return cmds
}

func generateDeploymentFile(
	ms IMicroservice,
	ctx IContext,
	buildsDir string,
	svcFilePath string,
	envFilePath string,
	deployTypes string, // deployTypes is comma separated value of service (prefix or full name) that want to create
	ns string,
	outFile string,
	tmplURLPrefix string,
	externalEnvs map[string]string) error {

	file := NewFile()
	// Append backslash "/" after template directory if not exists
	if !strings.HasSuffix(buildsDir, "/") {
		buildsDir = buildsDir + "/"
	}

	// Read all services
	svcFile, err := readServiceFile(ms, ctx, svcFilePath)
	if err != nil {
		ms.Log("CMD", err.Error())
		return err
	}

	// Read all envs
	envFile, err := readEnvFile(ms, ctx, envFilePath)
	if err != nil {
		ms.Log("CMD", err.Error())
		return err
	}

	// Build each service deployment if match with deploy
	sb := &strings.Builder{}

	// deployTypes is comma separated value of service (prefix or full name) that want to create
	// eg, connect,mkt-mq,mkt-ptask
	deploys := strings.Split(deployTypes, ",")

	for _, service := range svcFile.Services {
		name := service["name"]
		for _, deploy := range deploys {
			if strings.HasPrefix(name, deploy) {

				// Create deployment file by downloading from deployment template on github
				// then replace variable from envFile
				dploy, err := readDeployment(ms, ctx, service, buildsDir, envFile.Envs, ns, tmplURLPrefix, externalEnvs)
				if err != nil {
					ms.Log("CMD", err.Error())
					return err
				}
				sb.WriteString(dploy)
				sb.WriteString("\n---\n")
			}
		}
	}

	output := sb.String()
	if len(output) == 0 {
		err = fmt.Errorf("no output deployment")
		ms.Log("CMD", err.Error())
		return err
	}
	err = file.Write(outFile, output)
	if err != nil {
		ms.Log("CMD", err.Error())
		return err
	}
	return nil
}

func readDeployment(
	ms IMicroservice,
	ctx IContext,
	service map[string]string,
	buildsDir string,
	allEnvs []*Env,
	namespace string,
	tmplURLPrefix string,
	externalEnvs map[string]string) (string, error) {

	// If TmplFile is URL we use requester, otherwise we use file
	imageName := service["image"]
	fileName := imageName + ".yml"

	body, err := readTemplateBody(ms, tmplURLPrefix, fileName)
	if err != nil {
		ms.Log("CMD", err.Error())
		return "", err
	}

	image, err := readImage(ms, ctx, imageName, buildsDir)
	if err != nil {
		ms.Log("CMD", err.Error())
		return "", err
	}

	// If alias is set, use alias instead of name
	name := service["name"]
	alias := service["alias"]
	if len(alias) > 0 {
		name = alias
	}

	envs, err := readEnvs(ms, imageName, allEnvs, tmplURLPrefix, externalEnvs)
	if err != nil {
		ms.Log("CMD", err.Error())
		return "", err
	}

	svcAttrs, err := readSvcAttrs(ms, tmplURLPrefix)
	if err != nil {
		ms.Log("CMD", err.Error())
		return "", err
	}

	for _, attr := range svcAttrs {
		varName := fmt.Sprintf("<%s>", attr.Alias)
		if attr.Replace {
			val := service[attr.Name]
			body = setVar(body, varName, val, attr.DoubleQuote)
			continue
		} else {
			// Custom attribute
			switch attr.Name {
			case "name":
				body = setVar(body, varName, name, attr.DoubleQuote)
			case "envs":
				body = setVar(body, varName, envs, attr.DoubleQuote)
			case "image":
				body = setVar(body, varName, image, attr.DoubleQuote)
			case "namespace":
				body = setVar(body, varName, namespace, attr.DoubleQuote)
			}
		}
	}

	return body, nil
}

func setVar(body string, key string, value string, addQuoted bool) string {
	if addQuoted {
		return strings.ReplaceAll(body, key, fmt.Sprintf(`"%s"`, value))
	}
	return strings.ReplaceAll(body, key, value)
}

func readSvcAttrs(
	ms IMicroservice,
	tmplURLPrefix string) ([]*ServiceAttr, error) {

	body, err := readTemplateBody(ms, tmplURLPrefix, "svc-attrs.yml")
	if err != nil {
		ms.Log("CMD", err.Error())
		return nil, err
	}

	attrFile := &ServiceAttrFile{}
	err = yaml.Unmarshal([]byte(body), &attrFile)
	if err != nil {
		ms.Log("CMD", err.Error())
		return nil, err
	}

	return attrFile.Attrs, nil
}

func readEnvs(
	ms IMicroservice,
	image string,
	allEnvs []*Env,
	tmplURLPrefix string,
	externalEnvs map[string]string) (string, error) {

	// If TmplFile is URL we use requester, otherwise we use file
	fileNames, err := getPlatformConfigFileNames(ms, image, tmplURLPrefix)
	if err != nil {
		ms.Log("CMD", err.Error())
		return "", err
	}

	envs := []*Env{}
	for _, fileName := range fileNames {
		body, err := readTemplateBody(ms, tmplURLPrefix, fileName)
		if err != nil {
			ms.Log("CMD", err.Error())
			return "", err
		}

		cfgTmpl := &CfgTmpl{}
		err = yaml.Unmarshal([]byte(body), &cfgTmpl)
		if err != nil {
			ms.Log("CMD", err.Error())
			return "", err
		}

		for _, cfg := range cfgTmpl.Configs {
			// Replace name with alias if exists
			cfgName := cfg.Name
			if len(cfg.Alias) > 0 {
				cfgName = cfg.Alias
			}

			value, found := getEnv(cfg.Name, cfg.Alias, allEnvs)
			if !found && cfg.Required {
				return "", fmt.Errorf("configuration %s is required", cfgName)
			}
			if found {
				envs = append(envs, &Env{Name: cfgName, Value: value})
			}
		}
	}

	for extEnvKey, extEnvVal := range externalEnvs {
		envs = append(envs, &Env{Name: extEnvKey, Value: extEnvVal})
	}

	if len(envs) == 0 {
		return "", nil
	}

	res, err := json.Marshal(envs)
	if err != nil {
		ms.Log("CMD", err.Error())
		return "", err
	}

	// Remove "[" and "]" before return
	return "," + string(res[1:len(res)-1]), nil
}

func getEnv(name string,
	alias string,
	allEnvs []*Env) (string, bool) {

	for _, env := range allEnvs {
		// Find with name first then alias
		if env.Name == name {
			return env.Value, true
		}
		if env.Name == alias {
			return env.Value, true
		}
	}
	return "", false
}

func getPlatformConfigFileNames(
	ms IMicroservice,
	image string,
	tmplURLPrefix string) ([]string, error) {

	body, err := readTemplateBody(ms, tmplURLPrefix, "cfgmaps.yml")
	if err != nil {
		ms.Log("CMD", err.Error())
		return nil, err
	}

	mappings := map[string]interface{}{}
	err = yaml.Unmarshal([]byte(body), &mappings)
	if err != nil {
		ms.Log("CMD", err.Error())
		return nil, err
	}

	files := []string{}
	for file, imgs := range mappings {
		images, ok := imgs.([]interface{})
		if !ok {
			err := fmt.Errorf("invalid cfgmaps.yml format")
			ms.Log("CMD", err.Error())
			return nil, err
		}
		for _, img := range images {
			imgStr := fmt.Sprintf("%v", img)
			if imgStr == image {
				files = append(files, file+".yml")
			}
		}
	}

	return files, nil
}

func readImage(
	ms IMicroservice,
	ctx IContext,
	image string,
	tmplDir string) (string, error) {

	file := NewFile()

	buildFileName := filepath.Join(tmplDir, image+"-build.yml")

	exists, err := file.Exists(buildFileName)
	if err != nil {
		ms.Log("CMD", err.Error())
		return "", err
	}
	if !exists {
		return "", nil
	}

	body, err := file.Read(buildFileName)
	if err != nil {
		ms.Log("CMD", err.Error())
		return "", err
	}
	buildFile := &BuildFile{}
	err = yaml.Unmarshal([]byte(body), &buildFile)
	if err != nil {
		ms.Log("CMD", err.Error())
		return "", err
	}
	if len(buildFile.Images) == 0 {
		err := fmt.Errorf("no image in build file: %s", buildFileName)
		ms.Log("CMD", err.Error())
		return "", err
	}
	return buildFile.Images[0].Image, nil
}

func readTemplateBody(
	ms IMicroservice,
	tmplURLPrefix string,
	fileName string) (string, error) {

	body := ""
	var err error

	if strings.HasPrefix(tmplURLPrefix, "http://") || strings.HasPrefix(tmplURLPrefix, "https://") {
		rq := NewRequester(tmplURLPrefix, 5*time.Second, ms)
		body, err = rq.Get(fileName, nil)
		if err != nil {
			ms.Log("CMD", err.Error())
			return "", err
		}
	} else {
		file := NewFile()
		body, err = file.Read(tmplURLPrefix + fileName)
		if err != nil {
			ms.Log("CMD", err.Error())
			return "", err
		}
	}

	return body, nil
}

func readEnvFile(
	ms IMicroservice,
	ctx IContext,
	envFilePath string) (*EnvFile, error) {

	file := NewFile()
	envBody, err := file.Read(envFilePath)
	if err != nil {
		ms.Log("CMD", err.Error())
		return nil, err
	}

	kv := map[string]string{}
	err = yaml.Unmarshal([]byte(envBody), &kv)
	if err != nil {
		ms.Log("CMD", err.Error())
		return nil, err
	}

	envFile := &EnvFile{}
	envFile.Version = "v1"
	for k, v := range kv {
		envFile.Envs = append(envFile.Envs, &Env{
			Name:  k,
			Value: v,
		})
	}

	return envFile, nil
}

func readServiceFile(
	ms IMicroservice,
	ctx IContext,
	svcFilePath string) (*ServiceFile, error) {

	file := NewFile()
	svcBody, err := file.Read(svcFilePath)
	if err != nil {
		ms.Log("CMD", err.Error())
		return nil, err
	}
	svcFile := &ServiceFile{}
	err = yaml.Unmarshal([]byte(svcBody), &svcFile)
	if err != nil {
		ms.Log("CMD", err.Error())
		return nil, err
	}

	return svcFile, nil
}

type BuildFile struct {
	Version string   `json:"version" yaml:"version"`
	Images  []*Build `json:"images" yaml:"images"`
}

type Build struct {
	Name  string `json:"name" yaml:"name"`
	Image string `json:"image" yaml:"image"`
}

type ServiceFile struct {
	Version  string              `json:"version" yaml:"version"`
	Services []map[string]string `json:"services" yaml:"services"`
}

type EnvFile struct {
	Version string `json:"version" yaml:"version"`
	Envs    []*Env `json:"envs" yaml:"envs"`
}

type Env struct {
	Name  string `json:"name" yaml:"name"`
	Value string `json:"value" yaml:"value"`
}

type ShellResult struct {
	StdOut string `json:"std_out"`
	StdErr string `json:"std_err"`
}

type CfgTmpl struct {
	Version string `json:"version" yaml:"version"`
	Configs []*Cfg `json:"configs" yaml:"configs"`
}

// Cfg is configuration
type Cfg struct {
	Name     string  `json:"name" yaml:"name"`         // We will use Name if match, otherwise will use Alias if alias match
	Alias    string  `json:"alias" yaml:"alias"`       // Alias will be a string that write to configmaps file
	Required bool    `json:"required" yaml:"required"` // If required = true, we will not use Default
	Default  string  `json:"default" yaml:"default"`   // Default only work if Required = false
	Type     CfgType `json:"type" yaml:"type"`         // Use Type for validation
	Sample   string  `json:"sample" yaml:"sample"`     // Sample config
}

type CfgType string

const (
	CfgTypeString      CfgType = "s"
	CfgTypeBool        CfgType = "b"
	CfgTypeNumber      CfgType = "n"
	CfgTypeURL         CfgType = "u"
	CfgTypeDockerImage CfgType = "d"
	CfgTypeTimestamp   CfgType = "t"
)

type ServiceAttr struct {
	Name        string `json:"name" yaml:"name"`
	Alias       string `json:"alias" yaml:"alias"`
	DoubleQuote bool   `json:"doublequote" yaml:"doublequote"`
	Replace     bool   `json:"replace" yaml:"replace"`
}

type ServiceAttrFile struct {
	Version string         `json:"version" yaml:"version"`
	Attrs   []*ServiceAttr `json:"attrs" yaml:"attrs"`
}

var (
	ArgVaultEndpoint *CommandArg = &CommandArg{
		Name:        "vault",
		ShortName:   "v",
		Description: "Vault endpoint",
		IsRequired:  false,
	}
	ArgVaultToken *CommandArg = &CommandArg{
		Name:        "token",
		ShortName:   "t",
		Description: "Vault token",
		IsRequired:  false,
	}
	ArgDeploy *CommandArg = &CommandArg{
		Name:        "deploy",
		ShortName:   "d",
		Description: "Deploy name",
		IsRequired:  false,
	}
	ArgBaseDir *CommandArg = &CommandArg{
		Name:        "base_dir",
		ShortName:   "b",
		Description: "Base Directory",
		IsRequired:  false,
	}
)
