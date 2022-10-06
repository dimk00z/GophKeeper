package utils_test

import (
	"testing"
	"time"

	"github.com/dimk00z/GophKeeper/internal/utils"
	"github.com/google/uuid"
)

func TestCrypto(t *testing.T) {
	phrase := "This is top secret"
	secretKey := "secretKey"
	encryptedString := utils.Encrypt(secretKey, phrase)
	decryptedString := utils.Decrypt(secretKey, encryptedString)

	if phrase != decryptedString {
		t.Errorf("got %q, wanted %q", decryptedString, phrase)
	}
}

func TestHash(t *testing.T) {
	password := "TestPassword"
	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		t.Errorf("got %v error", err)
	}
	if utils.VerifyPassword(hashedPassword, password) != nil {
		t.Errorf("got %v error", err)
	}
}

func TestToken(t *testing.T) {
	testPrivateKey := "LS0tLS1CRUdJTiBSU0EgUFJJVkFURSBLRVktLS0tLQpNSUlCUEFJQkFBSkJBTzVIKytVM0xrWC91SlRvRHhWN01CUURXSTdGU0l0VXNjbGFFKzlaUUg5Q2VpOGIxcUVmCnJxR0hSVDVWUis4c3UxVWtCUVpZTER3MnN3RTVWbjg5c0ZVQ0F3RUFBUUpCQUw4ZjRBMUlDSWEvQ2ZmdWR3TGMKNzRCdCtwOXg0TEZaZXMwdHdtV3Vha3hub3NaV0w4eVpSTUJpRmI4a25VL0hwb3piTnNxMmN1ZU9wKzVWdGRXNApiTlVDSVFENm9JdWxqcHdrZTFGY1VPaldnaXRQSjNnbFBma3NHVFBhdFYwYnJJVVI5d0loQVBOanJ1enB4ckhsCkUxRmJxeGtUNFZ5bWhCOU1HazU0Wk1jWnVjSmZOcjBUQWlFQWhML3UxOVZPdlVBWVd6Wjc3Y3JxMTdWSFBTcXoKUlhsZjd2TnJpdEg1ZGdjQ0lRRHR5QmFPdUxuNDlIOFIvZ2ZEZ1V1cjg3YWl5UHZ1YStxeEpXMzQrb0tFNXdJZwpQbG1KYXZsbW9jUG4rTkVRdGhLcTZuZFVYRGpXTTlTbktQQTVlUDZSUEs0PQotLS0tLUVORCBSU0EgUFJJVkFURSBLRVktLS0tLQ=="
	testPublicKey := "LS0tLS1CRUdJTiBQVUJMSUMgS0VZLS0tLS0KTUZ3d0RRWUpLb1pJaHZjTkFRRUJCUUFEU3dBd1NBSkJBTzVIKytVM0xrWC91SlRvRHhWN01CUURXSTdGU0l0VQpzY2xhRSs5WlFIOUNlaThiMXFFZnJxR0hSVDVWUis4c3UxVWtCUVpZTER3MnN3RTVWbjg5c0ZVQ0F3RUFBUT09Ci0tLS0tRU5EIFBVQkxJQyBLRVktLS0tLQ=="
	testToken, err := utils.CreateToken(
		time.Hour,
		uuid.New(),
		testPrivateKey)
	if err != nil {
		t.Errorf("got %v error", err)
	}
	if _, err := utils.ValidateToken(testToken, testPublicKey); err != nil {
		t.Errorf("got %v error", err)
	}
}