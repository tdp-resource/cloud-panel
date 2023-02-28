package certmagic

import (
	"context"

	"tdp-cloud/helper/logman"
)

func magicEvent(ctx context.Context, evt string, data map[string]any) error {

	evtlog := logman.Named("cert.event").Sugar()

	evtlog.Infof("Certmagic Event: %s with data: %v\n", evt, data)
	return nil

}
