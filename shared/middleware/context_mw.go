package middleware

import (
	"context"
	"encoding/json"
	"mc-stove/shared/constant"
	"mc-stove/shared/helper"
	"mc-stove/shared/port"
	"mc-stove/shared/types"
	"net/http"

	"github.com/gorilla/mux"
)

func ManagerContext(secretKey string, logger port.ILogger) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			managerContext := types.NewManagerContext()
			ctxt := context.WithValue(r.Context(), constant.CONTEXT_ID, managerContext)
			nReq := r.WithContext(ctxt)

			if accessToken := r.Header.Get("Access-Token"); accessToken != "" {

				jwtUserData := &types.JwtUserData{}
				jwt := helper.NewJwtHelper(secretKey)

				if err := jwt.Load(accessToken); err != nil {
					logger.Error(err.Error())
				} else if jwt.TokenIsValid() {
					if err = json.Unmarshal([]byte(jwt.GetPayload("user").(string)), &jwtUserData); err != nil {
						logger.Error(err.Error())
					} else {
						managerContext.User = types.NewUserContext().LoadFromJwt(jwtUserData)
					}
				}
			}

			next.ServeHTTP(w, nReq)
		})
	}
}
