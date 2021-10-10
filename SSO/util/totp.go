package util

import (
	"time"

	"github.com/pquerna/otp/totp"
)

func GenerateTOTPSharedKey() (string, error) {
	key, err := totp.Generate(totp.GenerateOpts{
		Issuer:      "hustunique.com",
		AccountName: "root@hustunique.com",
	})
	if err != nil {
		return "", err
	}
	return key.Secret(), nil
}

func GenerateTOTPPasscode(shardKey string) (string, error) {
	return totp.GenerateCode(shardKey, time.Now())
}

// TODO
func ValidateTOTPPasscode(passcode, shardKey string) bool {
	return totp.Validate(passcode, shardKey)
}
