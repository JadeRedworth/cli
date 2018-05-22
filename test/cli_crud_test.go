package test

import (
	"encoding/json"
	"fmt"
	"strings"
	"testing"

	"github.com/fnproject/cli/testharness"
	"github.com/jmoiron/jsonq"
)

// TODO: These are both  Super minimal
func TestFnAppUpdateCycle(t *testing.T) {
	t.Parallel()

	h := testharness.Create(t)
	defer h.Cleanup()

	appName := h.NewAppName()

	// can't create an app twice
	h.Fn("create", "apps", appName).AssertSuccess()
	h.Fn("create", "apps", appName).AssertFailed()
	h.Fn("list", "apps", appName).AssertSuccess().AssertStdoutContains(appName)
	h.Fn("inspect", "apps", appName).AssertSuccess().AssertStdoutContains(fmt.Sprintf(`"name": "%s"`, appName))
	h.Fn("set", "config", "apps", appName, "fooConfig", "barval").AssertSuccess()
	h.Fn("get", "config", "apps", appName, "fooConfig").AssertSuccess().AssertStdoutContains("barval")
	h.Fn("list", "config", "apps", appName).AssertSuccess().AssertStdoutContains("fooConfig=barval")
	h.Fn("unset", "config", "apps", appName, "fooConfig").AssertSuccess()
	h.Fn("get", "config", "apps", appName, "fooConfig").AssertFailed()
	h.Fn("list", "config", "apps", appName).AssertSuccess().AssertStdoutEmpty()
}

func TestSimpleFnRouteUpdateCycle(t *testing.T) {
	t.Parallel()

	h := testharness.Create(t)
	defer h.Cleanup()
	appName1 := h.NewAppName()
	h.Fn("create", "routes", appName1, "myroute", "--image", "foo/duffimage:0.0.1").AssertSuccess()
	h.Fn("create", "routes", appName1, "myroute", "--image", "foo/duffimage:0.0.1").AssertFailed()
	h.Fn("inspect", "routes", appName1, "myroute").AssertSuccess().AssertStdoutContains(`"path": "/myroute"`)
	h.Fn("update", "routes", "-i", "bar/duffbeer:0.1.2", appName1, "myroute").AssertSuccess()
	h.Fn("set", "config", "routes", appName1, "myroute", "confA", "valB").AssertSuccess()
	h.Fn("get", "config", "routes", appName1, "myroute", "confA").AssertSuccess().AssertStdoutContains("valB")
	h.Fn("list", "config", "routes", appName1, "myroute").AssertSuccess().AssertStdoutContains("confA=valB")
	h.Fn("unset", "config", "routes", appName1, "myroute", "confA").AssertSuccess()
	h.Fn("get", "config", "routes", appName1, "myroute", "confA").AssertFailed()
}

func TestRouteUpdateValues(t *testing.T) {
	t.Parallel()

	validCases := []struct {
		args   []string
		query  []string
		result interface{}
	}{
		{[]string{"--image", "baz/fooimage:1.0.0"}, []string{"image"}, "baz/fooimage:1.0.0"},
		{[]string{"--memory", "129"}, []string{"memory"}, 129.0},
		{[]string{"--type", "async"}, []string{"type"}, "async"},
		{[]string{"--headers", "foo=bar"}, []string{"headers", "foo", "0"}, "bar"},
		{[]string{"--format", "default"}, []string{"format"}, "default"},
		{[]string{"--timeout", "111"}, []string{"timeout"}, 111.0},
		{[]string{"--idle-timeout", "128"}, []string{"idle_timeout"}, 128.0},
	}

	for i, tcI := range validCases {
		tc := tcI
		t.Run(fmt.Sprintf("Valid Case %d", i), func(t *testing.T) {
			t.Parallel()
			h := testharness.Create(t)
			defer h.Cleanup()
			appName1 := h.NewAppName()
			h.Fn("create", "routes", appName1, "myroute", "--image", "foo/someimage:0.0.1").AssertSuccess()

			h.Fn(append([]string{"update", "routes", appName1, "myroute"}, tc.args...)...).AssertSuccess()
			resJson := h.Fn("insepct", "routes", appName1, "myroute").AssertSuccess()

			routeObj := map[string]interface{}{}
			err := json.Unmarshal([]byte(resJson.Stdout), &routeObj)
			if err != nil {
				t.Fatalf("Failed to parse routes inspect as JSON %v, %v", err, resJson)
			}

			q := jsonq.NewQuery(routeObj)
			val, err := q.Interface(tc.query...)
			if err != nil {
				t.Fatalf("Failed to find path %v in json body %v", tc.query, resJson.Stdout)
			}

			if val != tc.result {
				t.Fatalf("Expected %s to be %s  after running %s but was %s, %v", strings.Join(tc.query, "."), tc.result, strings.Join(tc.args, " "), val, resJson)
			}

		})
	}

	invalidCases := [][]string{
		{"--image", "fooimage:1.0.0"}, // image with no registry
		//	{"--memory", "0"},  bug?
		{"--memory", "wibble"},
		{"--type", "blancmange"},
		{"--headers", "potatocakes"},
		{"--format", "myharddisk"},
		{"--timeout", "86400"},
		{"--timeout", "sit_in_the_corner"},
		{"--idle-timeout", "86000"},
		{"--idle-timeout", "yawn"},
	}

	for i, tcI := range invalidCases {
		tc := tcI
		t.Run(fmt.Sprintf("Invalid Case %d", i), func(t *testing.T) {
			t.Parallel()
			h := testharness.Create(t)
			defer h.Cleanup()
			appName1 := h.NewAppName()
			h.Fn("create", "routes", appName1, "myroute", "--image", "foo/someimage:0.0.1").AssertSuccess()

			h.Fn(append([]string{"update", "routes", appName1, "myroute"}, tc...)...).AssertFailed()
		})
	}

}
