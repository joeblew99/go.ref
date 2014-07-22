package blesser

import (
	"fmt"
	"strings"
	"time"

	"veyron/services/identity"
	"veyron/services/identity/googleoauth"
	"veyron2"
	"veyron2/ipc"
	"veyron2/vdl/vdlutil"
)

type googleOAuth struct {
	rt                     veyron2.Runtime
	clientID, clientSecret string
	duration               time.Duration
	domain                 string
}

// NewGoogleOAuthBlesserServer provides an identity.OAuthBlesserService that uses authorization
// codes to obtain the username of a client and provide blessings with that name.
//
// For more details, see documentation on Google OAuth 2.0 flows:
// https://developers.google.com/accounts/docs/OAuth2
//
// Blessings generated by this server expire after duration. If domain is non-empty, then blessings
// are generated only for email addresses from that domain.
func NewGoogleOAuthBlesserServer(rt veyron2.Runtime, clientID, clientSecret string, duration time.Duration, domain string) interface{} {
	return identity.NewServerOAuthBlesser(&googleOAuth{rt, clientID, clientSecret, duration, domain})
}

func (b *googleOAuth) Bless(ctx ipc.ServerContext, authcode, redirectURL string) (vdlutil.Any, error) {
	config := googleoauth.NewOAuthConfig(b.clientID, b.clientSecret, redirectURL)
	name, err := googleoauth.ExchangeAuthCodeForEmail(config, authcode)
	if err != nil {
		return nil, err
	}
	if len(b.domain) > 0 && !strings.HasSuffix(name, "@"+b.domain) {
		return nil, fmt.Errorf("blessings for %q are not allowed", name)
	}
	self := b.rt.Identity()
	// Use the blessing that was used to authenticate with the client to bless it.
	if self, err = self.Derive(ctx.LocalID()); err != nil {
		return nil, err
	}
	// TODO(ashankar,ataly): Use the same set of caveats as is used by the HTTP handler.
	// For example, a third-party revocation caveat?
	return self.Bless(ctx.RemoteID(), name, b.duration, nil)
}
