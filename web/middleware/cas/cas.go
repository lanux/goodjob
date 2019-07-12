package cas

import (
	"encoding/xml"
	"errors"
	"fmt"
	"github.com/kataras/iris"
	"github.com/lanux/goodjob/v1/common/consts"
	"github.com/lanux/goodjob/v1/common/logger"
	"github.com/lanux/goodjob/v1/config"
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
	LoginPath    = "/login"
	LogoutPath   = "/logout"
	ValidatePath = "/serviceValidate"
)

var C *Client

func InitCas(app *iris.Application, i Interceptor) {
	C = &Client{i}
	app.Use(C.Authentication)
}

type Client struct {
	i Interceptor
}

func (c *Client) Authentication(ctx iris.Context) {
	if c.i != nil && c.i.PreAuthentication(ctx) {
		ctx.Next()
		return
	}
	tk := ctx.URLParam(consts.CAS_TICKET)
	if len(tk) <= 0 {
		RedirectToLogin(ctx)
		ctx.StatusCode(http.StatusFound)
		ctx.StopExecution()
		return
	} else {
		u, err := validateTicket(tk)
		if err != nil {
			RedirectToLogin(ctx)
			ctx.StopExecution()
			return
		}
		if err == nil && c.i != nil {
			c.i.PostAuthentication(ctx, u.AuthSuccess)
		}
	}
	ctx.Next()
}

func GetResponseBody(url string) (string, error) {
	response, err := http.Get(url)
	if err != nil {
		return "", err
	}
	if response.StatusCode != http.StatusOK {
		errMsg := fmt.Sprintf("response should be 200 but is: %d", response.StatusCode)
		return "", errors.New(errMsg)
	}

	body, err := ioutil.ReadAll(response.Body)
	defer response.Body.Close()

	return string(body), nil
}

func (c *Client) RedirectToLogout(ctx iris.Context) {
	if c.i != nil {
		c.i.BeforeLogout(ctx)
	}
	u, err := url.Parse(config.Global.Cas.CasServerUrlPrefix + LogoutPath)
	if err != nil {
		logger.Panic()
	}
	q := u.Query()
	q.Add(consts.CAS_SERVICE, config.Global.Cas.ServerName)
	u.RawQuery = q.Encode()
	ctx.StatusCode(http.StatusFound)
	ctx.Redirect(u.String(), http.StatusFound)
	ctx.StopExecution()
}

// RedirectToLogout replies to the request with a redirect URL to authenticate with CAS.
func RedirectToLogin(ctx iris.Context) {
	u, err := url.Parse(config.Global.Cas.CasServerLoginUrl)
	if err != nil {
		panic(err)
	}
	q := u.Query()
	q.Add(consts.CAS_SERVICE, config.Global.Cas.ServerName)
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
func validateTicket(ticket string) (*Response, error) {
	validReq, err := url.Parse(config.Global.Cas.CasServerUrlPrefix + ValidatePath)
	if err != nil {
		return nil, err
	}
	q := validReq.Query()
	q.Add(consts.CAS_TICKET, ticket)
	q.Add(consts.CAS_SERVICE, config.Global.Cas.ServerName)
	validReq.RawQuery = q.Encode()
	user, err := GetResponseBody(validReq.String())
	if err != nil {
		return nil, err
	}
	r := &Response{}
	xml.Unmarshal([]byte(user), &r)
	if r.AuthenticationFailure != "" {
		return nil, errors.New(r.AuthenticationFailure)
	}
	return r, nil
}
