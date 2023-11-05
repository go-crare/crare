package updates

import (
	"sync"

	"gopkg.in/crare.v1/pkg/params"
)

var paramsPool = sync.Pool{
	New: func() any {
		return &params.Updates{}
	},
}

func AcquireParams() *params.Updates {
	return paramsPool.Get().(*params.Updates)
}

func ReleaseParams(v *params.Updates) {
	v.Limit = 0
	v.Offset = 0
	clear(v.AllowedUpdates)
	v.Timeout = 0
	paramsPool.Put(v)
}
