package handler

import (
	"context"
	"encoding/json"
	oteldemo "github.com/phbpx/otel-demo"
	"github.com/uptrace/opentelemetry-go-extra/otelutil"
	"io"
	"net/http"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
)

func decode(r *http.Request, into interface{}) error {
	rawJson, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}
	return json.Unmarshal(rawJson, into)
}

func respond(ctx context.Context, status int, keyValue attribute.KeyValue) {
	ctx, span := otel.GetTracerProvider().Tracer("").Start(ctx, "handler.respond")
	span.SetAttributes(attribute.Int("http.status", status))
	span.SetAttributes(keyValue)
	span.End()
}

func respondWriter(rw http.ResponseWriter, status int, data *oteldemo.Lead) {

	if status == http.StatusNoContent || data == nil {
		rw.WriteHeader(status)
		return
	}

	rawJson, err := json.Marshal(data)
	if err != nil {
		panic("respond-json-marshal:" + err.Error())
	}

	rw.Header().Add("Content-Type", "application-json")
	rw.WriteHeader(status)
	rw.Write(rawJson)

}
func respondErr(ctx context.Context, rw http.ResponseWriter, status int, err error) {
	keyValue := otelutil.Attribute("err", err)
	respond(ctx, status, keyValue)
	respondWriter(rw, status, nil)
}
