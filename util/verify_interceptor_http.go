package util

import (
	"net/http"
	"strconv"
)

// VerifyInterceptorHTTP ...
func VerifyInterceptorHTTP(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		accessDetails, err := ExtractFromRedis(r)
		if err != nil {
			ResponseWithError(w, http.StatusUnauthorized, "Verify Token failure. Reason: "+err.Error())
			return
		}
		// fmt.Printf("VerifyInterceptorHTTP prinn: %v", accessDetails)
		s := strconv.FormatInt(accessDetails.UserID, 10)

		r.Header.Set("userId", s)
		next.ServeHTTP(w, r)

	})
}
