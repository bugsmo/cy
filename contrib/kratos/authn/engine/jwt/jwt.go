package jwt

import (
	"context"
	"errors"

	jwtSdk "github.com/golang-jwt/jwt/v5"

	"github.com/bugsmo/cy/contrib/kratos/authn/engine"
	"github.com/bugsmo/cy/contrib/kratos/authn/engine/utils"
)

type Authenticator struct {
	options *Options
}

var _ engine.Authenticator = (*Authenticator)(nil)

func NewAuthenticator(opts ...Option) (engine.Authenticator, error) {
	auth := &Authenticator{
		options: &Options{},
	}

	for _, o := range opts {
		o(auth.options)
	}

	if auth.options.signingMethod == nil {
		auth.options.signingMethod = jwtSdk.SigningMethodHS256
	}

	return auth, nil
}

func (jwt *Authenticator) Authenticate(ctx context.Context, contextType engine.ContextType) (*engine.AuthClaims, error) {
	tokenString, err := utils.AuthFromMD(ctx, utils.BearerWord, contextType)
	if err != nil {
		return nil, engine.ErrMissingBearerToken
	}

	return jwt.AuthenticateToken(tokenString)
}

func (jwt *Authenticator) AuthenticateToken(token string) (*engine.AuthClaims, error) {
	jwtToken, err := jwt.parseToken(token)

	switch {
	case errors.Is(err, jwtSdk.ErrTokenMalformed):
		return nil, engine.ErrInvalidToken
	case errors.Is(err, jwtSdk.ErrTokenSignatureInvalid):
		return nil, engine.ErrSignTokenFailed
	case errors.Is(err, jwtSdk.ErrTokenExpired) || errors.Is(err, jwtSdk.ErrTokenNotValidYet):
		return nil, engine.ErrTokenExpired
	case !jwtToken.Valid:
		return nil, engine.ErrInvalidToken
	case jwtToken.Method != jwt.options.signingMethod:
		return nil, engine.ErrUnsupportedSigningMethod
	case jwtToken.Claims == nil:
		return nil, engine.ErrInvalidClaims
	}

	claims, ok := jwtToken.Claims.(jwtSdk.MapClaims)
	if !ok {
		return nil, engine.ErrInvalidClaims
	}

	authClaims, err := utils.MapClaimsToAuthClaims(claims)
	if err != nil {
		return nil, err
	}

	return authClaims, nil
}

func (jwt *Authenticator) CreateIdentityWithContext(ctx context.Context, contextType engine.ContextType, claims engine.AuthClaims) (context.Context, error) {
	strToken, err := jwt.CreateIdentity(claims)
	if err != nil {
		return ctx, err
	}

	ctx = utils.MDWithAuth(ctx, utils.BearerWord, strToken, contextType)

	return ctx, nil
}

func (jwt *Authenticator) CreateIdentity(claims engine.AuthClaims) (string, error) {
	jwtToken := jwtSdk.NewWithClaims(jwt.options.signingMethod, utils.AuthClaimsToJwtClaims(claims))

	strToken, err := jwt.generateToken(jwtToken)
	if err != nil {
		return "", err
	}

	return strToken, nil
}

func (jwt *Authenticator) Close() {}

func (jwt *Authenticator) parseToken(token string) (*jwtSdk.Token, error) {
	if jwt.options.keyFunc == nil {
		return nil, engine.ErrMissingKeyFunc
	}

	return jwtSdk.Parse(token, jwt.options.keyFunc)
}

func (jwt *Authenticator) generateToken(jwtToken *jwtSdk.Token) (string, error) {
	if jwt.options.keyFunc == nil {
		return "", engine.ErrMissingKeyFunc
	}

	key, err := jwt.options.keyFunc(jwtToken)
	if err != nil {
		return "", engine.ErrGetKeyFailed
	}

	strToken, err := jwtToken.SignedString(key)
	if err != nil {
		return "", engine.ErrSignTokenFailed
	}

	return strToken, nil
}
