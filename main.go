package main

import (
	"flag"
	"fmt"
	"os"

	_ "github.com/quasilyte/go-ruleguard/dsl"
	// kit_utils "github.com/sailsforce/gomicro-kit/utils"
	jose "gopkg.in/square/go-jose.v2/jwt"
)

// func init() {
// if err := kit_utils.InitEnv(); err != nil {
// log.Printf("error loading .env file: %v\n", err)
// }
// }

type JwtClaims struct {
	IssuerId          string            `json:"iss"`
	IssuerType        int               `json:"ist"`
	AudienceId        string            `json:"aud"`
	AudienceType      string            `json:"aut"`
	JwtExpiration     *jose.NumericDate `json:"exp"`
	JwtValidNotBefore *jose.NumericDate `json:"nbf"`
	JwtIssuedAt       *jose.NumericDate `json:"iat"`
	JwtId             string            `json:"jti"`
	CallerContext     string            `json:"ctx"`
	SubjectType       string            `json:"sty"`
	Subject           string            `json:"sub"`
	OrchClaims
}

type OrchClaims struct {
	AccountId string `json:"org62.account_id"`
	ContactId string `json:"org62.contact_id"`
}

func main() {
	// pass token using the '-token={jwt token} flag'
	tokenPtr := flag.String("token", "", "jwt token")
	flag.Parse()

	// parse token
	jwtToken, err := jose.ParseSigned(*tokenPtr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}

	// pull out claims
	claims := JwtClaims{}
	if err := jwtToken.UnsafeClaimsWithoutVerification(&claims); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
	fmt.Fprintf(os.Stdout, "claims: %+v\n", claims)

	os.Exit(0)
}
