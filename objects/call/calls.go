package call

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"

	apps "github.com/fnproject/cli/objects/app"
	fns "github.com/fnproject/cli/objects/fn"
	fnclient "github.com/fnproject/fn_go/clientv2"
	apicall "github.com/fnproject/fn_go/clientv2/call"
	models "github.com/fnproject/fn_go/modelsv2"
	"github.com/go-openapi/strfmt"
	"github.com/urfave/cli"
)

type callsCmd struct {
	client *fnclient.Fn
}

// getMarshalableCall returns a call struct that we can marshal to JSON and output
func getMarshalableCall(call *models.Call) interface{} {
	if call.Error != "" {
		return struct {
			ID          string          `json:"id"`
			AppID       string          `json:"appId"`
			Path        string          `json:"path"`
			CreatedAt   strfmt.DateTime `json:"createdAt"`
			StartedAt   strfmt.DateTime `json:"startedAt"`
			CompletedAt strfmt.DateTime `json:"completedAt"`
			Status      string          `json:"status"`
			ErrorReason string          `json:"errorReason"`
		}{
			call.ID,
			call.AppID,
			call.Path,
			call.CreatedAt,
			call.StartedAt,
			call.CompletedAt,
			call.Status,
			call.Error,
		}
	}

	return struct {
		ID          string          `json:"id"`
		AppID       string          `json:"appId"`
		Path        string          `json:"path"`
		CreatedAt   strfmt.DateTime `json:"createdAt"`
		StartedAt   strfmt.DateTime `json:"startedAt"`
		CompletedAt strfmt.DateTime `json:"completedAt"`
		Status      string          `json:"status"`
	}{
		call.ID,
		call.AppID,
		call.Path,
		call.CreatedAt,
		call.StartedAt,
		call.CompletedAt,
		call.Status,
	}
}

func printCalls(c *cli.Context, calls []*models.Call) error {
	outputFormat := strings.ToLower(c.String("output"))
	if outputFormat == "json" {
		var allCalls []interface{}
		for _, call := range calls {
			c := getMarshalableCall(call)
			allCalls = append(allCalls, c)
		}
		b, err := json.MarshalIndent(allCalls, "", "    ")
		if err != nil {
			return err
		}
		fmt.Fprint(os.Stdout, string(b))
	} else {
		for _, call := range calls {
			fmt.Println(fmt.Sprintf(
				"ID: %v\n"+
					"App Id: %v\n"+
					"Route: %v\n"+
					"Created At: %v\n"+
					"Started At: %v\n"+
					"Completed At: %v\n"+
					"Status: %v\n",
				call.ID, call.AppID, call.Path, call.CreatedAt,
				call.StartedAt, call.CompletedAt, call.Status))
			if call.Error != "" {
				fmt.Println(fmt.Sprintf("Error reason: %v\n", call.Error))
			}
		}
	}
	return nil
}

func (c *callsCmd) get(ctx *cli.Context) error {
	callID := ctx.Args().Get(0)

	params := apicall.GetCallsCallIDParams{

		CallID:  callID,
		Context: context.Background(),
	}
	resp, err := c.client.Call.GetCallsCallID(&params)
	if err != nil {
		switch e := err.(type) {
		case *apicall.GetCallsCallIDNotFound:
			return errors.New(e.Payload.Message)
		default:
			return err
		}
	}
	printCalls(ctx, []*models.Call{resp.Payload})
	return nil
}

func (c *callsCmd) list(ctx *cli.Context) error {
	app := ctx.Args().Get(0)
	fn := ctx.Args().Get(1)
	a, err := apps.GetAppByName(app)
	if err != nil {
		return err
	}

	f, err := fns.GetFnByName(c.client, a.ID, fn)
	if err != nil {
		return err
	}
	params := apicall.GetCallsParams{
		AppID:   &a.ID,
		FnID:    &f.ID,
		Context: context.Background(),
	}

	if ctx.String("cursor") != "" {
		cursor := ctx.String("cursor")
		params.Cursor = &cursor
	}
	if ctx.String("from-time") != "" {
		fromTime := ctx.String("from-time")
		fromTime_int64, err := time.Parse(time.RFC3339, fromTime)
		if err != nil {
			return err
		}
		res := fromTime_int64.Unix()
		params.FromTime = &res

	}

	if ctx.String("to-time") != "" {
		toTime := ctx.String("to-time")
		toTime_int64, err := time.Parse(time.RFC3339, toTime)
		if err != nil {
			return err
		}
		res := toTime_int64.Unix()
		params.ToTime = &res
	}

	n := ctx.Int64("n")
	if n < 0 {
		return errors.New("Number of calls: negative value not allowed")
	}

	var resCalls []*models.Call
	for {
		resp, err := c.client.Call.GetCalls(&params)
		fmt.Println("Resp: ", resp)
		if err != nil {
			switch e := err.(type) {
			case *apicall.GetCallsNotFound:
				return errors.New(e.Payload.Message)
			default:
				return err
			}
		}

		resCalls = append(resCalls, resp.Payload.Items...)
		howManyMore := n - int64(len(resCalls)+len(resp.Payload.Items))
		if howManyMore <= 0 || resp.Payload.NextCursor == "" {
			break
		}

		params.Cursor = &resp.Payload.NextCursor
	}

	printCalls(ctx, resCalls)
	return nil
}
