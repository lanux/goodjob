module github.com/lanux/goodjob

go 1.12

require (
	github.com/BurntSushi/toml v0.3.1
	github.com/casbin/casbin v0.0.0-20190327153408-c87db9db0776
	github.com/gorilla/securecookie v1.1.1
	github.com/iris-contrib/middleware v0.0.0-20181021162728-8bd5d51b3c2e
	github.com/jinzhu/gorm v0.0.0-20190310121721-8b07437717e7
	github.com/juju/loggo v0.0.0-20190526231331-6e530bcce5d8 // indirect
	github.com/juju/testing v0.0.0-20190429233213-dfc56b8c09fc // indirect
	github.com/k0kubun/colorstring v0.0.0-20150214042306-9440f1994b88 // indirect
	github.com/kataras/iris v0.0.0-20190410115758-5c6b23c07981
	github.com/kr/pretty v0.1.0 // indirect
	github.com/mattn/go-colorable v0.1.2 // indirect
	github.com/onsi/ginkgo v1.8.0 // indirect
	github.com/onsi/gomega v1.5.0 // indirect
	github.com/smartystreets/goconvey v0.0.0-20190330032615-68dc04aab96a // indirect
	github.com/yudai/pp v2.0.1+incompatible // indirect
	gopkg.in/mgo.v2 v2.0.0-20180705113604-9856a29383ce // indirect
)

replace (
	cloud.google.com/go => github.com/googleapis/google-cloud-go v0.39.0
	golang.org/x/crypto => github.com/golang/crypto v0.0.0-20190513172903-22d7a77e9e5f
	golang.org/x/lint => github.com/golang/lint v0.0.0-20190409202823-959b441ac422
	golang.org/x/oauth2 => github.com/golang/oauth2 v0.0.0-20190523182746-aaccbc9213b0
	golang.org/x/sync => github.com/golang/sync v0.0.0-20190423024810-112230192c58
	golang.org/x/sys => github.com/golang/sys v0.0.0-20190526052359-791d8a0f4d09
	golang.org/x/text => github.com/golang/text v0.3.2
	golang.org/x/time => github.com/golang/time v0.0.0-20190308202827-9d24e82272b4
	golang.org/x/tools => github.com/golang/tools v0.0.0-20190525145741-7be61e1b0e51
)
