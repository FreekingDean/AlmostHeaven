package auth

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

type JWTClaims struct {
	jwt.StandardClaims
	//Standard JWT Claim Info
	JTI        string `json:"jti"` //JWT ID
	IAT        int64  `json:"iat"` //Issued at
	Expiration int64  `json:"exp"` //Expiration

	// JWT Extended Info
	ApplicationID string   `json:"application_id"` //App ID this is used for
	Scope         string   `json:"scope"`          //"basic"
	RefreshToken  string   `json:"refresh_token"`  //uuid
	Roles         []string `json:"roles"`          // ["user"]
	Admin         bool     `json:"admin"`

	//Account Basic Info
	ID              string `json:"id"`                //Account ID
	MasterAccountID string `json:"master_account_id"` //Account ID
	FirstName       string `json:"first_name"`
	LastName        string `json:"last_name"`
	Username        string `json:"username"`
	EmailAddress    string `json:"email_address"`

	//Account Extended Info
	GeneratedUsername        bool        `json:"generated_username"` //false?
	Limited                  bool        `json:"limited"`            //Limited access account
	ThirdPartyMarketingOptIn bool        `json:"third_party_marketing_opt_in"`
	CoppaCompliant           bool        `json:"coppa_compliant"`
	Verified                 bool        `json:"verified"`
	AccountState             string      `json:"account_state"`    // "IDENTIFIED"
	ExternalAccount          *struct{}   `json:"external_account"` //Seems to always be empty object
	NewsOffersOptIn          bool        `json:"news_offers_opt_in"`
	OTPDisabled              bool        `json:"otp_disabled"` //MFA
	Language                 string      `json:"language"`     //"en" others??
	Country                  string      `json:"country"`
	Secret1Question          *int        `json:"secret_1_question"` //Needs to be null-able
	Secret2Question          *int        `json:"secret_2_question"`
	Status                   *UserStatus `json:"status"`
}
type UserStatus struct {
	OutsideGracePeriod      bool   `json:"outside_grace_period"` //false
	Verified                bool   `json:"verified"`
	NeedsDeviceVerification bool   `json:"needs_device_verification"` //false?
	AccountType             string `json:"account_type"`              //"full"
	MissingLegalDocuments   bool   `json:"missing_legal_documents"`
}

func (u *User) JWT() *JWTClaims {
	//Build JWT using library & return as string
	secret := 1
	return &JWTClaims{
		JTI:        "2edd3c0e-db9d-4256-83bf-7fd064122948",
		IAT:        time.Now().Unix(),
		Expiration: time.Now().Add(time.Second * 7200).Unix(),

		ApplicationID: "2edd3c0e-db9d-4256-83bf-7fd064122948",
		Scope:         "basic",
		RefreshToken:  "2edd3c0e-db9d-4256-83bf-7fd064122948",
		Roles:         []string{"user"},
		Admin:         false,

		ID:              "2edd3c0e-db9d-4256-83bf-7fd064122948",
		MasterAccountID: "2edd3c0e-db9d-4256-83bf-7fd064122948",
		FirstName:       u.FirstName,
		LastName:        u.LastName,
		Username:        u.Username,
		EmailAddress:    u.EmailAddress,

		GeneratedUsername:        false,
		Limited:                  false,
		ThirdPartyMarketingOptIn: false,
		CoppaCompliant:           true,
		Verified:                 true,
		AccountState:             "IDENTIFIED",
		ExternalAccount:          nil,
		NewsOffersOptIn:          false,
		OTPDisabled:              false,
		Language:                 "en",
		Country:                  "US",
		Secret1Question:          &secret,
		Secret2Question:          nil,
		Status: &UserStatus{
			OutsideGracePeriod:      false,
			Verified:                true,
			NeedsDeviceVerification: false,
			AccountType:             "full",
			MissingLegalDocuments:   false,
		},
	}
}

func (c *JWTClaims) String() (string, error) {
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		return "", err
	}
	token := jwt.NewWithClaims(jwt.SigningMethodES256, c)
	return token.SignedString(privateKey)
}
