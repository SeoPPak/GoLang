package auth

import (
	"context"
	_ "embed"
	"fmt"
	"log"
	"net/http"
	"section85/clock"
	"section85/entity"
	"time"

	"github.com/google/uuid"
	"github.com/lestrrat-go/jwx/v2/jwa"
	"github.com/lestrrat-go/jwx/v2/jwk"
	"github.com/lestrrat-go/jwx/v2/jwt"
)

// secret.pem과 public.pem 파일을 담을 변수
//
//go:embed cert/secret.pem
var rawPrivKey []byte

//go:embed cert/public.pem
var rawPubKey []byte

// JWT를 다루는 구조체
type JWTer struct {
	// 라이브러리 사용할 수 있도록 jwk.Key 타입으로 파싱
	PrivateKey, PublicKey jwk.Key

	// 초기화 함수 외부에서 파라미터로 받아온다.
	Store Store

	// JWT 발행 시간 설정을 위한 Clocker 인터페이스
	Clocker clock.Clocker
}

//go:generate go run github.com/matryer/moq -out moq_test.go . Store
type Store interface {
	Save(ctx context.Context, key string, userID entity.UserID) error
	Load(ctx context.Context, key string) (entity.UserID, error)
}

// JWTer 생성자
func NewJWTer(s Store, c clock.Clocker) (*JWTer, error) {
	j := &JWTer{Store: s}

	// secret.pem 파일을 파싱
	privKey, err := parse(rawPrivKey)
	if err != nil {
		return nil, fmt.Errorf("failed in NewJWTer: private key: %w", err)
	}

	// public.pem 파일을 파싱
	pubKey, err := parse(rawPubKey)
	if err != nil {
		return nil, fmt.Errorf("failed in NewJWTer: public key: %w", err)
	}

	// JWTer에 키를 저장
	j.PrivateKey = privKey
	j.PublicKey = pubKey
	j.Clocker = c

	return j, nil
}

// rawKey를 파싱하는 함수
func parse(rawKey []byte) (jwk.Key, error) {
	key, err := jwk.ParseKey(rawKey, jwk.WithPEM(true))
	if err != nil {
		return nil, err
	}
	log.Printf("key: %v", key)
	return key, nil
}

const (
	RoleKey     = "role"
	UserNameKey = "user_name"
)

// JWT토큰 발행
func (j *JWTer) GenerateToken(ctx context.Context, u entity.User) ([]byte, error) {
	//Issuer 필드에 발급자 정보를 추가 -> 발급자를 식별/ 토큰 출처 구분/ 토큰 신뢰성 검증
	//Subject 필드에 무엇에 대한 토큰인지 저장(토큰의 용도)
	//Claim 필드에 추가적인 정보 저장
	tok, err := jwt.NewBuilder().
		JwtID(uuid.New().String()).
		Issuer(`"go_todo_app"`).
		Subject("access_token").
		IssuedAt(j.Clocker.Now()).
		Expiration(j.Clocker.Now().Add(30*time.Minute)).
		Claim(RoleKey, u.Role).
		Claim(UserNameKey, u.Name).
		Build()
	if err != nil {
		return nil, fmt.Errorf("GenerateToken: failed to build token: %w", err)
	}
	if err := j.Store.Save(ctx, tok.JwtID(), u.ID); err != nil {
		return nil, err
	}

	signed, err := jwt.Sign(tok, jwt.WithKey(jwa.RS256, j.PrivateKey))
	if err != nil {
		return nil, err
	}
	return signed, nil
}

// 토큰을 파싱하고 유효성 검사
func (j *JWTer) GetToken(ctx context.Context, r *http.Request) (jwt.Token, error) {
	token, err := jwt.ParseRequest(
		r,
		jwt.WithKey(jwa.RS256, j.PublicKey),
		jwt.WithValidate(false),
	)

	if err != nil {
		return nil, err
	}

	if err := jwt.Validate(token, jwt.WithClock(j.Clocker)); err != nil {
		return nil, fmt.Errorf("GetToken: failed to validate token: %w", err)
	}

	if _, err := j.Store.Load(ctx, token.JwtID()); err != nil {
		return nil, fmt.Errorf("GetToken: %q expired: %w", token.JwtID(), err)
	}
	return token, nil
}

// 컨텍스트에 userID 저장
// 저장된 userID 꺼내오기
// 토큰에 저장된 추가정보 꺼내 context에 저장
type userIDKey struct{}
type roleKey struct{}

func SetUserID(ctx context.Context, uid entity.UserID) context.Context {
	return context.WithValue(ctx, userIDKey{}, uid)
}

func GetUserID(ctx context.Context) (entity.UserID, bool) {
	id, ok := ctx.Value(userIDKey{}).(entity.UserID)
	return id, ok
}

func SetRole(ctx context.Context, tok jwt.Token) context.Context {
	get, ok := tok.Get(RoleKey)
	if !ok {
		return context.WithValue(ctx, roleKey{}, "")
	}
	return context.WithValue(ctx, roleKey{}, get)
}

func GetRole(ctx context.Context) (string, bool) {
	role, ok := ctx.Value(roleKey{}).(string)
	return role, ok
}

// Request에 저장된 Token에서 uid, role 추출하고 context에 저장
func (j *JWTer) FillContext(r *http.Request) (*http.Request, error) {
	token, err := j.GetToken(r.Context(), r)
	if err != nil {
		return nil, err
	}
	uid, err := j.Store.Load(r.Context(), token.JwtID())
	if err != nil {
		return nil, err
	}

	ctx := SetUserID(r.Context(), uid)
	ctx = SetRole(ctx, token)
	clone := r.Clone(ctx)
	return clone, nil
}

func IsAdmin(ctx context.Context) bool {
	role, ok := GetRole(ctx)
	if !ok {
		return false
	}
	return role == "admin"
}
