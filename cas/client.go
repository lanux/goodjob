package cas

import (
	"crypto/tls"
	"encoding/xml"
	"errors"
	"fmt"
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
	"github.com/kataras/iris/sessions"
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
	LoginPath    = "/login"
	LogoutPath   = "/logout"
	ValidatePath = "/serviceValidate"
)

var (
	client *Client
)

type Client struct {
	config          Config
	tickets         string ""
	sessionsManager *sessions.Sessions
	Handler         context.Handler
}

type Config struct {
	CasServerLoginUrl  string `desc:"CAS Login url"`
	ServerName         string `desc:"project service url"`
	CasServerUrlPrefix string `desc:"CAS url prefix"`
}

func New(cfg Config, sessionsManager *sessions.Sessions) context.Handler {
	c := &Client{
		config:          cfg.Validate(),
		tickets:         "",
		sessionsManager: sessionsManager,
	}
	return c.Authentication
}

func (c *Client) Authentication(ctx iris.Context) {
	tk := ctx.URLParam("ticket")
	session := c.sessionsManager.Start(ctx)
	user := session.Get("user")
	if user != nil {
		return
	}
	if len(tk) <= 0 {
		c.RedirectToLogin(ctx)
	} else {
		c.validateTicket(tk, ctx)
	}
}

func GetResponseBody(url string) (string, error) {
	client := httpClient()
	response, err := client.Get(url)
	if err != nil {
		return "", err
	}

	if response.StatusCode != 200 {
		errMsg := fmt.Sprintf("response should be 200 but is: %d", response.StatusCode)
		return "", errors.New(errMsg)
	}

	body, err := ioutil.ReadAll(response.Body)
	defer response.Body.Close()

	return string(body), nil
}

func httpClient() *http.Client {
	transport := &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}
	return &http.Client{Transport: transport}
}

func (c Config) Validate() Config {
	if c.CasServerLoginUrl == "" {
		if c.CasServerUrlPrefix == "" {
			panic(c)
		}
		c.CasServerLoginUrl = c.CasServerUrlPrefix + LoginPath
	}
	return c
}

func (c *Client) RedirectToLogout(ctx iris.Context) {
	u, err := url.Parse(c.config.CasServerUrlPrefix + LogoutPath)
	if err != nil {
		panic(err)
	}
	q := u.Query()
	q.Add("service", c.config.ServerName)
	u.RawQuery = q.Encode()
	ctx.Redirect(u.String(), http.StatusFound)
}

// RedirectToLogout replies to the request with a redirect URL to authenticate with CAS.
func (c *Client) RedirectToLogin(ctx iris.Context) {
	u, err := url.Parse(c.config.CasServerLoginUrl)
	if err != nil {
		panic(err)
	}
	q := u.Query()
	q.Add("service", c.config.ServerName)
	u.RawQuery = q.Encode()
	ctx.Redirect(u.String(), http.StatusFound)
}

type Response struct {
	XMLName               xml.Name          `xml:"serviceResponse"`
	AuthenticationFailure string            `xml:"authenticationFailure"`
	AuthSuccess           AuthSuccessStruct `xml:"authenticationSuccess"`
}

type AuthSuccessStruct struct {
	XMLName    xml.Name         `xml:"authenticationSuccess"`
	User       string           `xml:"user"`
	Attributes AttributesStruct `xml:"attributes"`
}

type AttributesStruct struct {
	XMLName       xml.Name `xml:"attributes"`
	UserMobile    string   `xml:"UserMobile"`
	UserName      string   `xml:"UserName"`
	UserTitle     string   `xml:"UserTitle`
	DeptFullName  string   `xml:"DeptFullName`
	UserNum       string   `xml:"UserNum`
	UserNameHex   string   `xml:"UserNameHex`
	OfficeAddress string   `xml:"OfficeAddress`
	UserEmail     string   `xml:"UserEmail`
	ACCOUNT       string   `xml:"ACCOUNT`
	UserBelong    int8     `xml:"UserBelong`
	UserId        int16    `xml:"UserId`
	DeptName      string   `xml:"DeptName`
}

// validateTicket performs CAS ticket validation with the given ticket and service.
//
// If the request returns a 404 then validateTicketCas1 will be returned.
func (c *Client) validateTicket(ticket string, ctx iris.Context) error {
	u, err := url.Parse(c.config.CasServerUrlPrefix + ValidatePath)
	if err != nil {
		panic(err)
	}
	q := u.Query()
	q.Add("ticket", ticket)
	q.Add("service", c.config.ServerName)
	u.RawQuery = q.Encode()
	user, err := GetResponseBody(u.String())
	r := Response{}
	xml.Unmarshal([]byte(user), &r)
	session := c.sessionsManager.Start(ctx)
	session.Set("user", r)
	return nil
}
