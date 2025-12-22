package event

import (
	"context"
	"encoding/json"
)

type KeyContextValue string

func BindContexData(ctx context.Context, bind any) error {
	ctxData := ctx.Value(KeyContextValue("data"))

	jsonString, err := json.Marshal(ctxData)
	if err != nil {
		return err
	}

	json.Unmarshal(jsonString, bind)
	return nil
}
