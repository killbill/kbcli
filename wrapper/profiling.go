package killbill

import (
	"context"
	"github.com/go-openapi/runtime"
)

const KB_PROFILING_RESP = "X-Killbill-Profiling-Resp"

type profCtxKey struct {}
type ProfCtxValue struct {
	profInfo *string
	profRes string
}
// Caller specifies profiling info (request) through a new context
func WithProfilingInfo(parent context.Context, profInfo string) context.Context {
	return context.WithValue(parent, profCtxKey{}, &ProfCtxValue{profInfo: &profInfo})
}
// Caller retrieves the profiling result through this method
func GetProfilingRes(pctx context.Context) string {
	return getProfValue(pctx).profRes
}

// Internal for api calls
func getProfValue(ctx context.Context) *ProfCtxValue {
	prof, ok := ctx.Value(profCtxKey{}).(*ProfCtxValue)
	if !ok {
		return &ProfCtxValue{}
	}
	return prof
}

func setProfRes(val *ProfCtxValue, resp runtime.ClientResponse) {
	val.profRes = resp.GetHeader(KB_PROFILING_RESP)
}
