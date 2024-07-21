package middlewares

import (
	"net/http"
	"time"

	"github.com/alpardfm/e-commerce/src/utils/config"
	"github.com/alpardfm/go-toolkit/appcontext"
	"github.com/alpardfm/go-toolkit/codes"
	"github.com/alpardfm/go-toolkit/errors"
	"github.com/alpardfm/go-toolkit/header"
	"github.com/alpardfm/go-toolkit/log"
)

func ECMiddlware(ecCfg config.Application, log log.Interface) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		handler := func(w http.ResponseWriter, r *http.Request) {

			getTime := func() time.Time {
				location, err := time.LoadLocation("Asia/Jakarta")
				if err != nil {
					log.Warn(r.Context(), errors.NewWithCode(codes.CodeInvalidValue, "Failed Load Location Time, %v", err))
				}

				return time.Now().UTC().In(location)
			}

			r = r.WithContext(appcontext.SetRequestStartTime(r.Context(), getTime()))
			r = r.WithContext(appcontext.SetServiceVersion(r.Context(), ecCfg.Meta.Version))
			r = r.WithContext(appcontext.SetUserAgent(r.Context(), header.KeyUserAgent))
			r = r.WithContext(appcontext.SetRequestId(r.Context(), header.KeyRequestID))

			next.ServeHTTP(w, r)
		}

		return http.HandlerFunc(handler)
	}
}
