package auth

import (
	jwt "github.com/dgrijalva/jwt-go"
)

const (
	jwtSecret = "no secret"
)

type JWTClaims struct {
	jwt.StandardClaims
	//Standard JWT Claim Info
	JTI        string `json:"jti"` //JWT ID
	IAT        int    `json:"iat"` //Issued at
	Expiration int    `json:"exp"` //Expiration

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
	GeneratedUsername        bool     `json:"generated_username"` //false?
	Limited                  bool     `json:"limited"`            //Limited access account
	ThirdPartyMarketingOptIn bool     `json:"third_party_marketing_opt_in"`
	CoppaCompliant           bool     `json:"coppa_compliant"`
	Verified                 bool     `json:"verified"`
	AccountState             string   `json:"account_state"`    // "IDENTIFIED"
	ExternalAccount          struct{} `json:"external_account"` //Seems to always be empty object
	NewsOffersOptIn          bool     `json:"news_offers_opt_in"`
	OTPDisabled              bool     `json:"otp_disabled"` //MFA
	Language                 string   `json:"language"`     //"en" others??
	Country                  string   `json:"country"`
	Secret1Question          *int     `json:"secret_1_question"` //Needs to be null-able
	Secret2Question          *int     `json:"secret_2_question"`
	Status                   struct {
		OutsideGracePeriod      bool   `json:"outside_grace_period"` //false
		Verified                bool   `json:"verified"`
		NeedsDeviceVerification bool   `json:"needs_device_verification"` //false?
		AccountType             string `json:"account_type"`              //"full"
		MissingLegalDocuments   bool   `json:"missing_legal_documents"`
	} `json:"status"`
}

func (u *User) JWT() *JWTClaims {
	//Build JWT using library & return as string
	return &JWTClaims{}
}

func (c *JWTClaims) String() (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodES256, c)
	return token.SignedString(jwtSecret)
}
