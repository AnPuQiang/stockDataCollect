package handler

import (
	"crypto/sha1"
	"fmt"
	"github.com/go-zoo/bone"
	"io"
	"net/http"
	"sort"
	"strings"
)

var Token = "sofkgodfkgijinsn"

func (p *ServiceHandler) GetWXConnect() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//微信接入验证 这是首次对接微信 填写url后 微信服务器会发一个请求过来
		//c.Ctx.Request.URL-------------wx_connect?signature=038d75ed5485b9881a01b3b93e85f9fff28ea739&echostr=5756456183388806654&timestamp=1476173150&nonce=1093541731

		//开发者提交信息(包括URL、Token)后，微信服务器将发送Http Get请求到填写的URL上，
		//GET请求携带四个参数：signature、timestamp、nonce和echostr。公众号服务程序应该按如下要求进行接入验证
		timestamp := bone.GetValue(r, "timestamp")
		nonce := bone.GetValue(r, "nonce")
		signatureIn := bone.GetValue(r, "signature")

		fmt.Printf("timestamp=%s, nonce=%s, signatureIn=%s \r\n", timestamp, nonce, signatureIn)

		signatureGen := makeSignature(timestamp, nonce)

		//将加密后获得的字符串与signature对比，如果一致，说明该请求来源于微信
		if signatureGen != signatureIn {
			fmt.Printf("signatureGen != signatureIn signatureGen=%s,signatureIn=%s\n", signatureGen, signatureIn)
			w.Write([]byte(""))

		} else {
			//如果请求来自于微信，则原样返回echostr参数内容 以上完成后，接入验证就会生效，开发者配置提交就会成功。
			echostr := bone.GetValue(r, "echostr")
			w.Write([]byte(echostr))
		}

	})
}

func makeSignature(timestamp, nonce string) string {

	//1. 将 plat_token、timestamp、nonce三个参数进行字典序排序
	sl := []string{Token, timestamp, nonce}
	sort.Strings(sl)
	//2. 将三个参数字符串拼接成一个字符串进行sha1加密
	s := sha1.New()
	io.WriteString(s, strings.Join(sl, ""))

	return fmt.Sprintf("%x", s.Sum(nil))
}
