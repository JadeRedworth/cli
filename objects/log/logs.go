package log

import (
	"context"
	"fmt"

	fnclient "github.com/fnproject/fn_go/clientv2"
	apicall "github.com/fnproject/fn_go/clientv2/operations"
	"github.com/urfave/cli"
)

type logsCmd struct {
	client *fnclient.Fn
}

func (l *logsCmd) get(ctx *cli.Context) error {
	callID := ctx.Args().Get(0)

	// if callID == "last" || callID == "l" {
	// 	params := ccall.GetCallsCallIDParams{
	// 		AppID:   &a.ID,
	// 		CallID:  callID,
	// 		Context: context.Background(),
	// 	}
	// 	resp, err := l.client.Call.GetCallsCallID(&params)
	// 	if err != nil {
	// 		switch e := err.(type) {
	// 		case *ccall.GetCallsCallIDNotFound:
	// 			return errors.New(e.Payload.Message)
	// 		default:
	// 			return err
	// 		}
	// 	}
	// 	calls := resp.Payload
	// 	// if len(calls) > 0 {
	// 	// 	callID = calls[0].ID
	// 	// } else {
	// 	// 	return errors.New("No previous calls found.")
	// 	// }
	// }

	params := apicall.GetCallLogsParams{
		CallID:  &callID,
		Context: context.Background(),
	}

	resp, err := l.client.Operations.GetCallLogs(&params)
	fmt.Println("Resp: ", resp)
	if err != nil {
		switch e := err.(type) {
		case *apicall.GetCallLogsNotFound:
			return fmt.Errorf("%v", e.Payload.Message)
		default:
			return err
		}
	}
	fmt.Print(resp.Payload.Log)
	return nil
}
