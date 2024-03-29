package ibmppcsession

import (
	"fmt"
	"strings"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/IBM-Cloud/ppc-aas-go-client/helpers"
	"github.com/IBM-Cloud/ppc-aas-go-client/ppc-aas/client"
	"github.com/IBM/go-sdk-core/v5/core"
)

// IBMPPCSession ...
type IBMPPCSession struct {
	CRNFormat string
	Power     *client.PowerPrivateCloudAsaServiceAPI
	Options   *IBMPPCOptions
}

// SPOptions
type IBMPPCOptions struct {
	// The authenticator implementation to be used by the
	// service instance to authenticate outbound requests
	// Required
	Authenticator core.Authenticator

	// Enable/Disable http transport debugging log
	Debug bool

	// Region of the Power Cloud Service Instance
	// For generating the default endpoint
	// Deprecated: Region is deprecated, the URL is auto generated based on Zone when not provided.
	Region string

	// Power Virtual Server host or URL endpoint
	// This will be used instead of generating the default host
	// eg: dal.ppc-aas.cloud.ibm.com
	URL string

	// Account id of the Power Cloud Service Instance
	// It will be part of the CRN string
	// Required
	UserAccount string

	// Zone of the Power Cloud Service Instance
	// It will be part of the CRN string
	// Required
	Zone string
}

// Create a IBMPPCSession
func NewIBMPPCSession(o *IBMPPCOptions) (*IBMPPCSession, error) {
	if core.IsNil(o) {
		return nil, fmt.Errorf("options is required")
	}

	if core.IsNil(o.Authenticator) {
		return nil, fmt.Errorf("option Authenticator is required")
	}

	if o.UserAccount == "" {
		return nil, fmt.Errorf("option UserAccount is required")
	}

	if o.Zone == "" {
		return nil, fmt.Errorf("option Zone is required")
	}

	region := o.Region
	if region == "" {
		region = costructRegionFromZone(o.Zone)
	}

	var serviceURL string
	if o.URL != "" {
		serviceURL = o.URL
	} else {
		// Check in env
		serviceURL = helpers.GetPowerEndPoint()
		// If not set in env use prod endpoint
		if serviceURL == "" {
			serviceURL = "ppc-aas.cloud.ibm.com"
		}
	}

	// Prepend region to endpoint if not present
	if strings.HasPrefix(serviceURL, "ppc-aas.") {
		serviceURL = region + "." + serviceURL
	}

	// We need just the server host from the URL
	var host, scheme string
	if strings.HasPrefix(serviceURL, "https://") {
		scheme = SCHEME_HTTPS
		host = strings.TrimPrefix(serviceURL, "https://")
	} else if strings.HasPrefix(serviceURL, "http://") {
		scheme = SCHEME_HTTP
		host = strings.TrimPrefix(serviceURL, "http://")
	} else {
		// by default we use "https"
		scheme = SCHEME_HTTPS
		host = serviceURL
	}

	return &IBMPPCSession{
		CRNFormat: crnBuilder(o.UserAccount, o.Zone, host),
		Options:   o,
		Power:     getSPClient(o.Debug, host, scheme),
	}, nil
}

// authInfo ...
func (s *IBMPPCSession) AuthInfo(cloudInstanceID string) runtime.ClientAuthInfoWriter {
	return runtime.ClientAuthInfoWriterFunc(func(r runtime.ClientRequest, _ strfmt.Registry) error {
		auth, err := fetchAuthorizationData(s.Options.Authenticator)
		if err != nil {
			return err
		}
		if err := r.SetHeaderParam("Authorization", auth); err != nil {
			return err
		}
		return r.SetHeaderParam("CRN", fmt.Sprintf(s.CRNFormat, cloudInstanceID))
	})
}
