package cas

import (
	"GoodJob/config"
	"crypto/tls"
	"encoding/xml"
	"errors"
	"fmt"
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
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
	// 认证前执行的方法，返回false结束认证操作
	preAuthentication func(ctx iris.Context) bool
	// 认证后处理方法
	postAuthentication func(ctx iris.Context, u interface{})
}

func New(preAuthentication func(ctx iris.Context) bool, postAuthentication func(ctx iris.Context, u interface{})) context.Handler {
	c := &Client{
		preAuthentication:  preAuthentication,
		postAuthentication: postAuthentication,
	}
	return c.Authentication
}

func (c *Client) Authentication(ctx iris.Context) {
	tk := ctx.URLParam("ticket")
	if c.preAuthentication != nil && !c.preAuthentication(ctx) {
		ctx.Next()
		return
	}
	if len(tk) <= 0 {
		c.RedirectToLogin(ctx)
	} else {
		c.validateTicket(tk, ctx)
	}
	ctx.Next()
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

func (c *Client) RedirectToLogout(ctx iris.Context) {
	u, err := url.Parse(config.Global.Cas.CasServerUrlPrefix + LogoutPath)
	if err != nil {
		panic(err)
	}
	q := u.Query()
	q.Add("service", config.Global.Cas.ServerName)
	u.RawQuery = q.Encode()
	ctx.Redirect(u.String(), http.StatusFound)
}

// RedirectToLogout replies to the request with a redirect URL to authenticate with CAS.
func (c *Client) RedirectToLogin(ctx iris.Context) {
	u, err := url.Parse(config.Global.Cas.CasServerLoginUrl)
	if err != nil {
		panic(err)
	}
	q := u.Query()
	q.Add("service", config.Global.Cas.ServerName)
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
	u, err := url.Parse(config.Global.Cas.CasServerUrlPrefix + ValidatePath)
	if err != nil {
		panic(err)
	}
	q := u.Query()
	q.Add("ticket", ticket)
	q.Add("service", config.Global.Cas.ServerName)
	u.RawQuery = q.Encode()
	user, err := GetResponseBody(u.String())
	r := &Response{}
	xml.Unmarshal([]byte(user), &r)
	if c.postAuthentication != nil {
		c.postAuthentication(ctx, r)
	}
	return nil
}
