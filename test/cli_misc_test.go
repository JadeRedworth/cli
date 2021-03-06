package test

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"strings"
	"testing"

	"github.com/fnproject/cli/testharness"
)

func TestFnVersion(t *testing.T) {
	t.Parallel()

	tctx := testharness.Create(t)
	res := tctx.Fn("version")
	res.AssertSuccess()
}

// this is messy and nasty  as we generate different potential values for FN_API_URL based on its type
func fnApiUrlVariations(t *testing.T) []string {

	srcUrl := os.Getenv("FN_API_URL")

	if srcUrl == "" {
		srcUrl = "http://localhost:8080/"
	}

	if !strings.HasPrefix(srcUrl, "http:") && !strings.HasPrefix(srcUrl, "https:") {
		srcUrl = "http://" + srcUrl
	}
	parsed, err := url.Parse(srcUrl)

	if err != nil {
		t.Fatalf("Invalid/unparsable TEST_API_URL %s: %s", srcUrl, err)
	}

	var cases []string

	if parsed.Scheme == "http" {
		cases = append(cases, "http://"+parsed.Host+parsed.Path)
		cases = append(cases, parsed.Host+parsed.Path)
		cases = append(cases, parsed.Host)
	} else if parsed.Scheme == "https" {
		cases = append(cases, "https://"+parsed.Host+parsed.Path)
		cases = append(cases, "https://"+parsed.Host)
	} else {
		log.Fatalf("Unsupported url scheme for testing %s: %s", srcUrl, parsed.Scheme)
	}

	return cases
}

func TestFnApiUrlSupportsDifferentFormats(t *testing.T) {
	t.Parallel()

	h := testharness.Create(t)
	defer h.Cleanup()

	for _, candidateUrl := range fnApiUrlVariations(t) {
		h.WithEnv("FN_API_URL", candidateUrl)
		h.Fn("list", "apps").AssertSuccess()
	}
}

// Not sure what this test was intending (copied from old test.sh)
func TestSettingTimeoutWorks(t *testing.T) {
	t.Parallel()

	h := testharness.Create(t)
	defer h.Cleanup()
	h.WithEnv("FN_REGISTRY", "some_random_registry")

	appName := h.NewAppName()
	h.Fn("create", "app", appName).AssertSuccess()
	res := h.Fn("list", "apps")

	if !strings.Contains(res.Stdout, fmt.Sprintf("%s", appName)) {
		t.Fatalf("Expecting list apps to contain app name , got %v", res)
	}

	funcName := h.NewFuncName(appName)

	h.MkDir(funcName)
	h.Cd(funcName)
	h.WithMinimalFunctionSource()
	h.FileAppend("func.yaml", "\ntimeout: 50\n\nschema_version: 20180708\n")
	h.Fn("--verbose", "deploy", "--app", appName, "--local").AssertSuccess()
	h.Fn("invoke", appName, funcName).AssertSuccess()

	inspectRes := h.Fn("inspect", "fn", appName, funcName)
	inspectRes.AssertSuccess()
	if !strings.Contains(inspectRes.Stdout, `"timeout": 50`) {
		t.Errorf("Expecting fn inspect to contain CPU %v", inspectRes)
	}

	h.Fn("create", "fn", appName, "another", "some_random_registry/"+funcName+":0.0.2").AssertSuccess()

	h.Fn("invoke", appName, "another").AssertSuccess()
}

//Memory doesn't seem to get persisted/returned
func TestSettingMemoryWorks(t *testing.T) {
	t.Parallel()

	h := testharness.Create(t)
	defer h.Cleanup()
	h.WithEnv("FN_REGISTRY", "some_random_registry")

	appName := h.NewAppName()
	h.Fn("create", "app", appName).AssertSuccess()
	res := h.Fn("list", "apps")


	if !strings.Contains(res.Stdout, fmt.Sprintf("%s", appName)) {
		t.Fatalf("Expecting list apps to contain app name , got %v", res)
	}

	funcName := h.NewFuncName(appName)

	h.MkDir(funcName)
	h.Cd(funcName)
	h.WithMinimalFunctionSource()
	h.FileAppend("func.yaml", "memory: 100\nschema_version: 20180708\n")
	h.Fn("--verbose", "deploy", "--app", appName, "--local").AssertSuccess()
	h.Fn("invoke", appName, funcName).AssertSuccess()

	inspectRes := h.Fn("inspect", "fn", appName, funcName)
	inspectRes.AssertSuccess()
	if !strings.Contains(inspectRes.Stdout, `"memory": 100`) {
		t.Errorf("Expecting fn inspect to contain CPU %v", inspectRes)
	}

	h.Fn("create", "fn", appName, "another", "some_random_registry/"+funcName+":0.0.2").AssertSuccess()

	h.Fn("invoke", appName, "another").AssertSuccess()
}

func TestAllMainCommandsExist(t *testing.T) {
	t.Parallel()

	h := testharness.Create(t)
	defer h.Cleanup()

	testCommands := []string{
		"build",
		"bump",
		"call",
		"create",
		"delete",
		"deploy",
		"get",
		"init",
		"inspect",
		"list",
		"push",
		"run",
		"set",
		"start",
		"test",
		"unset",
		"update",
		"use",
		"version",
	}

	for _, cmd := range testCommands {
		res := h.Fn(cmd)
		if strings.Contains(res.Stderr, "command not found") {
			t.Errorf("Expected command %s to exist", cmd)
		}
	}
}

func TestAppYamlDeploy(t *testing.T) {
	t.Parallel()

	h := testharness.Create(t)
	defer h.Cleanup()

	appName := h.NewAppName()
	fnName := h.NewFuncName(appName)
	h.WithFile("app.yaml", fmt.Sprintf(`name: %s`, appName), 0644)
	h.MkDir(fnName)
	h.Cd(fnName)
	h.WithMinimalFunctionSource()
	h.Cd("")
	h.Fn("deploy", "--all", "--local").AssertSuccess()
	h.Fn("invoke", appName, fnName).AssertSuccess()
	h.Fn("deploy", "--all", "--local").AssertSuccess()
	h.Fn("invoke", appName, fnName).AssertSuccess()
}

func TestBump(t *testing.T) {
	t.Parallel()

	h := testharness.Create(t)
	defer h.Cleanup()

	expectFuncYamlVersion := func(v string) {
		funcYaml := h.GetFile("func.yaml")
		if !strings.Contains(funcYaml, fmt.Sprintf("version: %s", v)) {
			t.Fatalf("Exepected version to be %s but got %s", v, funcYaml)
		}

	}

	appName := h.NewAppName()
	fnName := h.NewFuncName(appName)
	h.MkDir(fnName)
	h.Cd(fnName)
	h.WithMinimalFunctionSource()

	expectFuncYamlVersion("0.0.1")

	h.Fn("bump").AssertSuccess()
	expectFuncYamlVersion("0.0.2")

	h.Fn("bump", "--major").AssertSuccess()
	expectFuncYamlVersion("1.0.0")

	h.Fn("bump").AssertSuccess()
	expectFuncYamlVersion("1.0.1")

	h.Fn("bump", "--minor").AssertSuccess()
	expectFuncYamlVersion("1.1.0")

	h.Fn("deploy", "--local", "--app", appName).AssertSuccess()
	expectFuncYamlVersion("1.1.1")

	h.Fn("i", "function", appName, fnName).AssertSuccess().AssertStdoutContains(fmt.Sprintf(`%s:1.1.1`, fnName))

	h.Fn("deploy", "--local", "--no-bump", "--app", appName).AssertSuccess()
	expectFuncYamlVersion("1.1.1")

	h.Fn("i", "function", appName, fnName).AssertSuccess().AssertStdoutContains(fmt.Sprintf(`%s:1.1.1`, fnName))

}

// test fn list id value matches fn inspect id value for app, function & trigger
func TestListID(t *testing.T) {
	t.Parallel()

	h := testharness.Create(t)
	defer h.Cleanup()

	// execute create app, func & trigger commands
	appName := h.NewAppName()
	funcName := h.NewFuncName(appName)
	triggerName := h.NewTriggerName(appName, funcName)
	h.Fn("create", "app", appName).AssertSuccess()
	h.Fn("create", "function", appName, funcName, "foo/duffimage:0.0.1").AssertSuccess()
	h.Fn("create", "trigger", appName, funcName, triggerName, "--type", "http", "--source", "/mytrigger").AssertSuccess()


	h.Fn("list", "apps", appName).AssertSuccess().AssertStdoutContains("ID")
	h.Fn("list", "fn", appName).AssertSuccess().AssertStdoutContains("ID")
	h.Fn("list", "triggers", appName).AssertSuccess().AssertStdoutContains("ID")

}
