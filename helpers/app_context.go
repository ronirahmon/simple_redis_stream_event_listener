package helpers

import (
	"context"
	"time"
)

func AppContextWithTimeOut() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), 20*time.Second)
}
