package token

import (
	"time"

	"errors"

	"git.containerum.net/ch/auth/pkg/utils"
)

type mockTokenRecord struct {
	IssuedAt time.Time
}

type mockIssuerValidator struct {
	returnedLifeTime time.Duration
	issuedTokens     map[string]mockTokenRecord
}

// NewMockIssuerValidator sets up a mock object used for testing purposes
func NewMockIssuerValidator(returnedLifeTime time.Duration) IssuerValidator {
	return &mockIssuerValidator{
		returnedLifeTime: returnedLifeTime,
		issuedTokens:     make(map[string]mockTokenRecord),
	}
}

func (m *mockIssuerValidator) IssueTokens(extensionFields ExtensionFields) (accessToken, refreshToken *IssuedToken, err error) {
	tokenID := utils.NewUUID()
	accessToken = &IssuedToken{
		Value:    "a" + tokenID,
		LifeTime: m.returnedLifeTime,
		ID:       tokenID,
	}
	now := time.Now().UTC()
	m.issuedTokens[tokenID] = mockTokenRecord{
		IssuedAt: now,
	}
	refreshToken = &IssuedToken{
		Value:    "r" + tokenID,
		LifeTime: m.returnedLifeTime,
		ID:       tokenID,
		IssuedAt: now,
	}
	return
}

func (m *mockIssuerValidator) ValidateToken(token string) (result *ValidationResult, err error) {
	rec, present := m.issuedTokens[token[1:]]
	var kind Kind
	switch token[0] {
	case 'a':
		kind = KindAccess
	case 'r':
		kind = KindRefresh
	default:
		return nil, errors.New("invalid token received")
	}
	return &ValidationResult{
		Valid: present && time.Now().Before(rec.IssuedAt.Add(m.returnedLifeTime)),
		Kind:  kind,
		ID:    token[1:],
	}, nil
}

func (m *mockIssuerValidator) Now() time.Time {
	return time.Now().UTC()
}
