package auth

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"section84/clock"
	"section84/entity"
	"section84/store"
	"section84/testutil/fixture"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/lestrrat-go/jwx/v2/jwa"
	"github.com/lestrrat-go/jwx/v2/jwk"
	"github.com/lestrrat-go/jwx/v2/jwt"
)

func TestEmbed(t *testing.T) {
	//public.pem 파일이 잘 embed 되었는지 확인
	want := []byte("-----BEGIN PUBLIC KEY-----")
	if !bytes.Contains(rawPubKey, want) {
		t.Errorf("want %s, but got %s", want, rawPubKey)
	}

	//secret.pem 파일이 잘 embed 되었는지 확인
	want = []byte("-----BEGIN PRIVATE KEY-----")
	if !bytes.Contains(rawPrivKey, want) {
		t.Errorf("want %s, but got %s", want, rawPrivKey)
	}
}

func TestJWTer_GenJWT(t *testing.T) {
	ctx := context.Background()
	wantID := entity.UserID(20)
	u := fixture.User(&entity.User{ID: wantID})

	moq := &StoreMock{}
	moq.SaveFunc = func(ctx context.Context, key string, userID entity.UserID) error {
		if userID != wantID {
			t.Errorf("want %d, but got %d", wantID, userID)
		}
		return nil
	}
	sut, err := NewJWTer(moq, clock.RealClocker{})
	if err != nil {
		t.Fatal(err)
	}
	got, err := sut.GenerateToken(ctx, *u)
	if err != nil {
		t.Fatalf("not want err: %v", err)
	}
	if len(got) == 0 {
		t.Errorf("token is empty")
	}
}

func TestJWTer_GetJWT(t *testing.T) {
	t.Parallel()

	// JWT 생성
	c := clock.FixedClocker{}
	want, err := jwt.NewBuilder().
		JwtID(uuid.New().String()).
		Issuer(`https://github.com/seoppak/GoLang`).
		Subject("access_token").
		IssuedAt(c.Now()).
		Expiration(c.Now().Add(30*time.Minute)).
		Claim(RoleKey, "user").
		Claim(UserNameKey, "test").
		Build()
	if err != nil {
		t.Fatal(err)
	}

	pkey, err := jwk.ParseKey(rawPrivKey, jwk.WithPEM(true))
	if err != nil {
		t.Fatal(err)
	}

	// JWT 서명
	signed, err := jwt.Sign(want, jwt.WithKey(jwa.RS256, pkey))
	if err != nil {
		t.Fatal(err)
	}

	userID := entity.UserID(20)
	ctx := context.Background()
	moq := &StoreMock{}
	moq.LoadFunc = func(ctx context.Context, key string) (entity.UserID, error) {
		return userID, nil
	}

	//sut 는 Mock JWTer
	sut, err := NewJWTer(moq, c)
	if err != nil {
		t.Fatal(err)
	}

	req := httptest.NewRequest(
		http.MethodGet,
		`https://github.com/seoppak`,
		nil,
	)

	// 토큰을 헤더에 추가
	req.Header.Set(`Authorization`, fmt.Sprintf(`Bearer %s`, signed))

	// 토큰을 추출
	got, err := sut.GetToken(ctx, req)
	if err != nil {
		t.Fatalf("want no error, but got %v", err)
	}

	// want는 원본 토큰
	// got은 전자서명 후 저장 -> 가져와서 복호화 한 토큰
	if !reflect.DeepEqual(got, want) {
		t.Errorf("GetToken() got = %v, want %v", got, want)
	}

}

type FixedTomorrowClocker struct{}

func (c FixedTomorrowClocker) Now() time.Time {
	return clock.FixedClocker{}.Now().Add(24 * time.Hour)
}

func TestJWTer_GetJWT_NG(t *testing.T) {
	t.Parallel()

	// JWT 생성
	c := clock.FixedClocker{}
	tok, err := jwt.NewBuilder().
		JwtID(uuid.New().String()).
		Issuer(`github.com/seoppak/GoLang`).
		Subject("access_token").
		IssuedAt(c.Now()).
		Expiration(c.Now().Add(30*time.Minute)).
		Claim(RoleKey, "test").
		Claim(UserNameKey, "test_user").
		Build()
	if err != nil {
		t.Fatal(err)
	}

	pkey, err := jwk.ParseKey(rawPrivKey, jwk.WithPEM(true))
	if err != nil {
		t.Fatal(err)
	}
	// JWT 서명
	signed, err := jwt.Sign(tok, jwt.WithKey(jwa.RS256, pkey))
	if err != nil {
		t.Fatal(err)
	}

	type moq struct {
		userID entity.UserID
		err    error
	}
	tests := map[string]struct {
		c   clock.Clocker
		moq moq
	}{
		"expire": {
			c: FixedTomorrowClocker{},
		},
		"notFoundInStore": {
			c: clock.FixedClocker{},
			moq: moq{
				err: store.ErrNotFound,
			},
		},
	}

	for n, tt := range tests {
		tt := tt
		t.Run(n, func(t *testing.T) {
			t.Parallel()
			ctx := context.Background()
			moq := &StoreMock{}
			moq.LoadFunc = func(ctx context.Context, key string) (entity.UserID, error) {
				return tt.moq.userID, tt.moq.err
			}

			//sut 는 Mock JWTer
			sut, err := NewJWTer(moq, tt.c)
			if err != nil {
				t.Fatal(err)
			}

			req := httptest.NewRequest(
				http.MethodGet,
				`https://github.com/seoppak`,
				nil,
			)
			req.Header.Set(`Authorization`, fmt.Sprintf(`Bearer %s`, signed))

			// 토큰을 추출
			got, err := sut.GetToken(ctx, req)
			if err == nil {
				t.Errorf("want error, but got nil")
			}
			if got != nil {
				t.Errorf("want nil, but got %v", got)
			}
		})
	}
}
