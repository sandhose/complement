package oidc

import (
	"encoding/base64"
	"encoding/json"
)


type OIDCTokenIntrospectionResponse struct {
	Active 		bool  `json:"active"`
	Sub    		string `json:"sub"`
	Username  string `json:"username"`
	Scope  		string `json:"scope"`
	Name  		*string `json:"name"`
}

func BuildAccessToken(localpart, deviceID, displayname string, isGuest bool) (string, error) {
	if deviceID == "" {
		deviceID = "ABCDEFGH"
	}

	var apiScope string
	if isGuest {
		apiScope = "urn:matrix:org.matrix.msc2967.client:api:guest"
	} else {
		apiScope = "urn:matrix:org.matrix.msc2967.client:api:*"
	}

	deviceScope := "urn:matrix:org.matrix.msc2967.client:device:" + deviceID

	introspectionResult := OIDCTokenIntrospectionResponse{
		Active: true,
		Username: localpart,
		Sub: 	localpart,
		Scope: 	"openid " + apiScope + " " + deviceScope,
	}

	if displayname != "" {
		introspectionResult.Name = &displayname
	}

	b, err := json.Marshal(introspectionResult)

	if err != nil {
		return "", err
	}

	return base64.RawURLEncoding.EncodeToString(b), nil
}