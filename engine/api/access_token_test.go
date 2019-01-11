package api

import (
	"encoding/json"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ovh/cds/engine/api/group"
	"github.com/ovh/cds/engine/api/test"
	"github.com/ovh/cds/engine/api/test/assets"
	"github.com/ovh/cds/sdk"
)

func TestAPI_TokenHandlers(t *testing.T) {
	api, db, router, end := newTestAPI(t)
	defer end()

	grp := sdk.Group{Name: sdk.RandomString(10)}
	user, password := assets.InsertLambdaUser(db, &grp)
	test.NoError(t, group.SetUserGroupAdmin(db, grp.ID, user.ID))

	// Test a call with a JWT Token
	jwt, err := assets.NewJWTToken(t, db, *user, grp)
	test.NoError(t, err)
	uri := router.GetRoute("POST", api.postNewAccessTokenHandler, nil)
	req := assets.NewJWTAuthentifiedRequest(t, jwt, "POST", uri,
		sdk.AccessTokenRequest{
			Origin:                "test",
			Description:           "test",
			ExpirationDelaySecond: 3600,
			GroupsIDs:             []int64{grp.ID},
		},
	)
	w := httptest.NewRecorder()
	router.Mux.ServeHTTP(w, req)
	assert.Equal(t, 201, w.Code)

	// Test a call with a JWT Token and an XSFR Token
	jwtxsrf, xsrf, err := assets.NewJWTTokenWithXSRF(t, db, api.Cache, *user, grp)
	test.NoError(t, err)
	uri = router.GetRoute("POST", api.postNewAccessTokenHandler, nil)
	req = assets.NewXSRFJWTAuthentifiedRequest(t, jwtxsrf, xsrf, "POST", uri,
		sdk.AccessTokenRequest{
			Origin:                "test",
			Description:           "testxsrf",
			ExpirationDelaySecond: 3600,
			GroupsIDs:             []int64{grp.ID},
		},
	)
	w = httptest.NewRecorder()
	router.Mux.ServeHTTP(w, req)
	assert.Equal(t, 201, w.Code)

	jwtToken := w.Header().Get("X-CDS-JWT")
	t.Logf("jwt token is %v...", jwtToken[:12])

	var accessToken sdk.AccessToken
	test.NoError(t, json.Unmarshal(w.Body.Bytes(), &accessToken))

	vars := map[string]string{
		"id": accessToken.ID,
	}
	uri = router.GetRoute("PUT", api.putRegenAccessTokenHandler, vars)
	req = assets.NewAuthentifiedRequest(t, user, password, "PUT", uri,
		sdk.AccessTokenRequest{
			Origin:                "test",
			Description:           "test",
			ExpirationDelaySecond: 3600,
			GroupsIDs:             []int64{grp.ID},
		},
	)
	w = httptest.NewRecorder()
	router.Mux.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)

	jwtToken = w.Header().Get("X-CDS-JWT")
	t.Logf("jwt token is %v...", jwtToken[:12])
	t.Logf("access token is : %s", w.Body.String())

	// Test_getAccessTokenByGroupHandler
	vars = map[string]string{
		"id": strconv.FormatInt(grp.ID, 10),
	}
	uri = router.GetRoute("GET", api.getAccessTokenByGroupHandler, vars)
	req = assets.NewAuthentifiedRequest(t, user, password, "GET", uri, nil)
	w = httptest.NewRecorder()
	router.Mux.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
	t.Logf("getAccessTokenByGroupHandler result is : %s", w.Body.String())

	// Test_getAccessTokenByUserHandler
	vars = map[string]string{
		"id": strconv.FormatInt(user.ID, 10),
	}
	uri = router.GetRoute("GET", api.getAccessTokenByUserHandler, vars)
	req = assets.NewAuthentifiedRequest(t, user, password, "GET", uri, nil)
	w = httptest.NewRecorder()
	router.Mux.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
	t.Logf("getAccessTokenByUserHandler result is : %s", w.Body.String())

}
