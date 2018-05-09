package client

import (
	"crypto/rsa"
	"fmt"
	"io/ioutil"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"syscall"

	"github.com/fnproject/cli/config"
	"github.com/fnproject/cli/utils"
	fnclient "github.com/fnproject/fn_go/client"
	openapi "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/oracle/oci-go-sdk/common"
	"github.com/spf13/viper"
	"golang.org/x/crypto/ssh/terminal"
)

func Host() string {
	hostURL := HostURL()
	return hostURL.Host
}

func HostURL() *url.URL {
	return hostURL(viper.GetString(config.EnvFnAPIURL))
}

func hostURL(urlStr string) *url.URL {
	if !strings.Contains(urlStr, "://") {
		urlStr = fmt.Sprint("http://", urlStr)
	}

	url, err := url.Parse(urlStr)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unparsable FN API Url: %s. Error: %s \n", urlStr, err)
		os.Exit(1)
	}

	if url.Port() == "" {
		if url.Scheme == "http" {
			url.Host = fmt.Sprint(url.Host, ":80")
		}
		if url.Scheme == "https" {
			url.Host = fmt.Sprint(url.Host, ":443")
		}
	}

	//maintain backwards compatibility with first version FN_API_URL env vars
	if url.Path == "" || url.Path == "/" {
		url.Path = "/v1"
	}

	return url
}

func defaultProvider(transport *openapi.Runtime) {
	if token := viper.GetString(config.EnvFnToken); token != "" {
		transport.DefaultAuthentication = openapi.BearerToken(token)
	}
}

func challengeForPKeyPassword() string {
	fmt.Print("Private Key Phrase: ")
	bytePassword, err := terminal.ReadPassword(int(syscall.Stdin))
	if err != nil {
		panic(fmt.Sprintf("%s", err))
	}
	password := string(bytePassword)
	fmt.Println()

	return password
}

func privateKey(pkeyFilePath string) *rsa.PrivateKey {
	keyBytes, err := ioutil.ReadFile(pkeyFilePath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to load private key from file: %s. Error: %s \n", pkeyFilePath, err)
		os.Exit(1)
	}

	pKeyPword := viper.GetString(config.OracleKeyPassword)
	if pKeyPword == "" {
		pKeyPword = challengeForPKeyPassword()
	}

	key, err := common.PrivateKeyFromBytes(keyBytes, common.String(pKeyPword))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to load private key from file bytes: %s. Error: %s \n", pkeyFilePath, err)
		os.Exit(1)
	}
	return key
}

func oracleProvider(transport *openapi.Runtime) {

	// Load configuration from .oci directory which has secrets (key-file, key-password and fingerprint)
	viper.AddConfigPath(filepath.Join(utils.GetHomeDir(), ".oci"))
	viper.SetConfigName("oci-config")
	viper.MergeInConfig()

	compartmentID := viper.GetString(config.OracleCompartmentID)
	tenancyID := viper.GetString(config.OracleTenancyID)
	userID := viper.GetString(config.OracleUserID)
	fingerprint := viper.GetString(config.OracleFingerprint)

	keyID := tenancyID + "/" + userID + "/" + fingerprint

	pKey := privateKey(viper.GetString(config.OracleKeyFile))

	if viper.GetBool(config.OracleDisableCerts) {
		transport.Transport = InsecureRoundTripper(transport.Transport)
	}

	transport.Transport =
		NewCompartmentIDRoundTripper(
			compartmentID,
			NewOCISigningRoundTripper(
				keyID,
				pKey,
				transport.Transport))
}

func GetTransportAndRegistry() (*openapi.Runtime, strfmt.Registry) {
	hostURL := HostURL()
	transport := openapi.New(hostURL.Host, hostURL.Path, []string{hostURL.Scheme})

	switch viper.GetString(config.ContextProvider) {
	case "default":
		defaultProvider(transport)
	case "oracle":
		oracleProvider(transport)
	default:
		defaultProvider(transport)
	}

	return transport, strfmt.Default
}

func APIClient() *fnclient.Fn {
	return fnclient.New(GetTransportAndRegistry())
}
