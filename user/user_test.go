package user

import (
	"context"
	"errors"
	"regexp"
	"strings"
	"testing"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type MockUserStore struct {
	CreateUserFunc func(ctx context.Context, user *UserFull) error
	CapturedUser   chan *UserFull
}

func (m *MockUserStore) CreateUser(ctx context.Context, user *UserFull) error {
	select {
	case m.CapturedUser <- user:
	default:
	}

	if m.CreateUserFunc != nil {
		return m.CreateUserFunc(ctx, user)
	}
	return nil
}

var uuidRegex = regexp.MustCompile("^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$")

func TestService_CreateUser_Success(t *testing.T) {
	// Arrange
	const (
		testEmail    = "test.user@example.com"
		testPassword = "SecurePassword12345"
	)

	mockStore := &MockUserStore{
		CreateUserFunc: func(ctx context.Context, user *UserFull) error { return nil },
		CapturedUser:   make(chan *UserFull, 1),
	}
	svc := NewUserService(mockStore)

	// Act
	gotUUID, err := svc.CreateUser(testEmail, testPassword)

	// Assert
	if err != nil {
		t.Fatalf("CreateUser failed unexpectedly: %v", err)
	}

	if !uuidRegex.MatchString(gotUUID) {
		t.Errorf("CreateUser returned invalid UUID format: got %q", gotUUID)
	}

	select {
	case capturedUser := <-mockStore.CapturedUser:
		if capturedUser.Email != testEmail {
			t.Errorf("Captured user email mismatch: got %s, want %s", capturedUser.Email, testEmail)
		}
		if capturedUser.AuthMethod != "PASSWORD" {
			t.Errorf("Captured AuthMethod mismatch: got %s, want PASSWORD", capturedUser.AuthMethod)
		}
		if capturedUser.EmailVerified {
			t.Error("Captured EmailVerified was true, expected false")
		}
		if capturedUser.PasswordHash == nil {
			t.Fatal("Captured PasswordHash was nil")
		}
		if err := bcrypt.CompareHashAndPassword([]byte(*capturedUser.PasswordHash), []byte(testPassword)); err != nil {
			t.Errorf("Hashed password comparison failed. The stored hash is incorrect: %v", err)
		}
		if capturedUser.UUID.String() != gotUUID {
			t.Errorf("Returned UUID (%s) does not match UUID passed to store (%s)", gotUUID, capturedUser.UUID.String())
		}
	case <-time.After(time.Millisecond * 10):
		t.Fatal("MockUserStore CreateUser was not called within timeout")
	}
}

func TestService_CreateUser_StoreFailure(t *testing.T) {
	// Arrange
	testError := errors.New("database connection failure")

	mockStore := &MockUserStore{
		CreateUserFunc: func(ctx context.Context, user *UserFull) error { return testError },
		CapturedUser:   make(chan *UserFull, 1),
	}
	svc := NewUserService(mockStore)

	// Act
	gotUUID, err := svc.CreateUser("error.user@example.com", "Password123")

	// Assert
	if gotUUID != "" {
		t.Errorf("Expected empty UUID string on error, got %s", gotUUID)
	}

	if err == nil {
		t.Fatal("Expected an error from CreateUser, got nil")
	}

	expectedErrMsg := "failed to create user in database"
	if !strings.Contains(err.Error(), expectedErrMsg) {
		t.Errorf("Error message missing expected prefix: got %q, want prefix %q", err.Error(), expectedErrMsg)
	}

	if !errors.Is(err, testError) {
		t.Errorf("Error was not correctly wrapped (errors.Is failed). Got error: %v", err)
	}
	select {
	case <-mockStore.CapturedUser:
	case <-time.After(time.Millisecond * 10):
		t.Fatal("MockUserStore CreateUser was not called within timeout")
	}
}

func Test_hashPassword_Valid(t *testing.T) {
	// Arrange
	plainPassword := "aComplexPassword!@#$"

	// Act
	hashedPassword, err := hashPassword(plainPassword)

	// Assert
	if err != nil {
		t.Fatalf("hashPassword failed unexpectedly: %v", err)
	}

	if len(hashedPassword) == 0 {
		t.Fatal("Returned hash was empty")
	}
	if err := bcrypt.CompareHashAndPassword(hashedPassword, []byte(plainPassword)); err != nil {
		t.Errorf("Hashed password comparison failed. The generated hash is invalid: %v", err)
	}
}

func Test_hashPassword_EmptyString(t *testing.T) {
	// Arrange
	plainPassword := ""

	// Act
	hashedPassword, err := hashPassword(plainPassword)

	// Assert
	if err != nil {
		t.Fatalf("hashPassword failed unexpectedly for empty string: %v", err)
	}

	if len(hashedPassword) == 0 {
		t.Fatal("Returned hash was empty for empty string input")
	}

	if err := bcrypt.CompareHashAndPassword(hashedPassword, []byte(plainPassword)); err != nil {
		t.Errorf("Hashed password comparison failed for empty string. The generated hash is invalid: %v", err)
	}
}
