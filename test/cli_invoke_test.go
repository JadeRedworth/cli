package test

import (
	"github.com/fnproject/cli/testharness"
	"testing"
	"strings"
)

func TestFnInvokeInvalidImage(t *testing.T) {
	t.Parallel()

	h := testharness.Create(t)
	defer h.Cleanup()
	appName1 := h.NewAppName()
	funcName1 := h.NewFuncName(appName1)
	h.Fn("create", "app", appName1).AssertSuccess()
	h.Fn("create", "function", appName1, funcName1, "foo/duffimage:0.0.1").AssertSuccess()
	h.Fn("invoke", appName1, funcName1).AssertFailed()
}

func TestFnInvokeValidImage(t *testing.T) {
	t.Parallel()

	h := testharness.Create(t)
	defer h.Cleanup()
	appName1 := h.NewAppName()
	funcName1 := h.NewFuncName(appName1)
	h.Fn("create", "app", appName1).AssertSuccess()
	h.Fn("create", "function", appName1, funcName1, "fnproject/hello:latest").AssertSuccess()
	h.Fn("invoke", appName1, funcName1).AssertSuccess()
}

func TestFnInvokeViaDirectUrl(t *testing.T) {
	t.Parallel()

	h := testharness.Create(t)
	defer h.Cleanup()
	appName1 := h.NewAppName()
	funcName1 := h.NewFuncName(appName1)
	h.Fn("create", "app", appName1).AssertSuccess()
	h.Fn("create", "function", appName1, funcName1, "fnproject/hello:latest").AssertSuccess()

	res := h.Fn("inspect", "function", appName1, funcName1, "--endpoint").AssertSuccess()

	url := strings.TrimSpace(res.Stdout)

	h.Fn("invoke", "--endpoint", url).AssertSuccess()
}
