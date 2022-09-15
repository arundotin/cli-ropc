package utils

import (
	"encoding/base64"
	"encoding/json"
	"strings"
)

type AccessTokenClaims struct {
	CommonName string
	Mail string `json:"mail"`
	DisplayName string `json:"displayName"`
}

func ExtractFullNameFromUserToken(accessToken string) string  {
	accessTokenPayload := GetPayloadFromAccessToken(accessToken)
	var accessTokenClaims AccessTokenClaims

	payloadJson,_ := base64.RawStdEncoding.DecodeString(accessTokenPayload)

	json.Unmarshal(payloadJson, &accessTokenClaims)

	return accessTokenClaims.DisplayName

}

func GetPayloadFromAccessToken(accessToken string) string {
	// JWT Access-Token Format is  HEADER.PAYLOAD.SIGNATURE. I am gonna just grab the PAYLOAD section that has the claims

	accessTokenPieces:= strings.Split(accessToken,".")

	return accessTokenPieces[1]
}
