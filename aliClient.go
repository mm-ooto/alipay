package alipay

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/mm-ooto/alipay/consts"
	"github.com/mm-ooto/alipay/utils"
	"io/ioutil"
	"net/http"
	"net/url"
	"sort"
	"strings"
	"sync"
	"time"
)

var (
	signDataIsEmptyErr         = errors.New("check sign Fail! The reason : signData is Empty")
	encryptContentIsEmptyErr   = errors.New("the content to be encrypted is empty")
	encryptKeyOrTypeIsEmptyErr = errors.New("the encryption type and key cannot be empty")
	encryptTypeErr             = errors.New("the encryption type can only be AES")
)

type AliClient struct {
	appId         string          // 支付宝分配给开发者的应用ID
	format        string          // (可不设置) 仅支持JSON
	charset       string          // 请求使用的编码格式，如utf-8,gbk,gb2312等
	signType      string          // 商户生成签名字符串所使用的签名算法类型，目前支持RSA2和RSA，推荐使用RSA2
	version       string          // (可不设置) 调用的接口版本，固定为：1.0
	appPrivateKey *rsa.PrivateKey // 应用私钥，开发者自己生成
	aliPublicKey  *rsa.PublicKey  // 支付宝公钥（公钥模式下设置，证书模式下无需设置），创建支付宝应用之后，从支付宝后台获取
	Client        *http.Client    // http client
	gatewayUrl    string          // 支付宝网关地址
	encryptKey    string          // 加密密钥
	encryptType   string          // 加密类型，默认AES

	mutex     sync.Mutex // 互斥锁
	appCertSN string     // 应用公钥证书序列号SN（证书模式下设置，公钥模式下无需设置）
	// 注意：如果使用公钥证书签名则需要在请求参数中将"app_cert_sn"和"alipay_root_cert_sn"传入，
	// 序列号SN 值是通过解析 X.509 证书文件中签发机构名称（name）以及内置序列号（serialNumber），将二者拼接后的字符串计算 MD5 值获取
	alipayRootCertSn        string                    // 支付宝根证书序列号SN（证书模式下设置，公钥模式下无需设置）
	aliCertSN               string                    // 支付宝公钥证书序列号SN（证书模式下设置，公钥模式下无需设置），主要用于验签，参考：https://opendocs.alipay.com/common/02mse7
	certSnRelationPublicKey map[string]*rsa.PublicKey // 证书序列号对应的公钥

	location     *time.Location
	isProduction bool // 是否是生产环境
}

type OptionFunc func(c *AliClient)

// AddClient 添加指定的Client
func AddClient(client *http.Client) OptionFunc {
	return func(c *AliClient) {
		c.Client = client
	}
}

// AddEncryptKey 添加加密密钥
func (a *AliClient) AddEncryptKey(encryptKey string) {
	a.encryptType = consts.EncryptTypeAes
	a.encryptKey = encryptKey
}

// NewAliClient 初始化支付宝客户端
func NewAliClient(appId, aliPublicKey, appPrivateKey, signType string, isProduction bool, opts ...OptionFunc) (aliClient *AliClient, err error) {
	aliClient = &AliClient{
		appId:        appId,
		format:       consts.FormatJson,
		charset:      consts.CharSetUTF8,
		signType:     signType,
		version:      consts.ApiVersion,
		Client:       http.DefaultClient,
		isProduction: isProduction,
	}
	if len(aliPublicKey) > 0 {
		aliClient.aliPublicKey, err = utils.ParsePKIXPublicKey(utils.GetPemPublic(aliPublicKey))
		if err != nil {
			return
		}
	}
	if len(appPrivateKey) > 0 {
		aliClient.appPrivateKey, err = utils.ParsePKCS1PrivateKey(utils.GetPemPrivate(appPrivateKey))
		if err != nil {
			return
		}
	}
	if isProduction {
		aliClient.gatewayUrl = consts.GateWalProdUrl
	} else {
		aliClient.gatewayUrl = consts.GateWalSandboxUrl
	}
	aliClient.location, _ = time.LoadLocation("Local")

	for _, opt := range opts {
		opt(aliClient)
	}
	return
}

// HandlerRequest 处理请求
// httpMethod 请求方法 GET,POST,PUT...
// apiName 接口名
// requestParams 请求的参数struct
func (a *AliClient) HandlerRequest(httpMethod string, requestParams RequestParams, result interface{}) (err error) {
	var urlValues url.Values
	urlValues, err = a.handlerParams(requestParams)
	if err != nil {
		return
	}

	//fmt.Println("HandlerRequest 请求参数：", urlValues.Encode())

	var req *http.Request
	req, err = http.NewRequest(httpMethod, a.gatewayUrl, strings.NewReader(urlValues.Encode()))
	if err != nil {
		return
	}
	req.Header.Set("Content-Type", consts.ContentType)
	var resp *http.Response
	resp, err = a.Client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	var data []byte
	data, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	resContent := string(data)
	//fmt.Println("响应数据:", resContent)
	// 对返回结果验签
	apiMethodName := requestParams.GetOtherParams().Get(consts.ApiMethodNameFiled)
	_, err = a.SyncVerifySign(resContent, apiMethodName)
	if err != nil {
		return
	}

	// todo 对内容解密，这一块有问题
	if requestParams.GetNeedEncrypt() {
		resContent, err = a.decryptJSONSignSource(apiMethodName, resContent)
		if err != nil {
			return
		}
	}
	if err = json.Unmarshal([]byte(resContent), &result); err != nil {
		return
	}

	return
}

// HandlerSDKRequest 生成用于调用收银台SDK的字符串
func (a *AliClient) HandlerSDKRequest(requestParams RequestParams) (result string, err error) {
	var urlValues url.Values
	urlValues, err = a.handlerParams(requestParams)
	if err != nil {
		return
	}
	// url encode
	result = urlValues.Encode()
	return
}

// HandlerPageRequest 页面提交执行方法
// result：构建好的、签名后的最终跳转URL（GET）或String形式的form（POST）
func (a *AliClient) HandlerPageRequest(httpMethod string, requestParams RequestParams) (result string, urlResult *url.URL, err error) {
	var urlValues url.Values
	urlValues, err = a.handlerParams(requestParams)
	if err != nil {
		return
	}
	if strings.ToUpper(httpMethod) == "GET" {
		// 拼接GET请求串字符 & 将字符解析为URL对象
		rawUrl := fmt.Sprintf("%s?%s", a.gatewayUrl, urlValues.Encode())
		urlResult, err = url.Parse(rawUrl)
	} else {
		// 拼接表单字符串
		result = a.buildRequestForm(urlValues)
	}
	return
}

// 建立请求，以表单HTML形式构造（默认）
// urlValues: 请求参数
func (a *AliClient) buildRequestForm(urlValues url.Values) (fromHtml string) {
	// 响应为表单格式，可嵌入页面，具体以返回的结果为准
	fromHtml = "<form id='alipaySubmit' name='alipaySubmit' action='" + a.gatewayUrl + "?charset=" + a.charset + "' method='POST'>"
	for field, values := range urlValues {
		if len(values) == 0 {
			continue
		}
		fromHtml += "<input type='hidden' name='" + field + "' value='" + values[0] + "'/>"
	}
	fromHtml += "<input type='submit' value='ok' style='display:none;''></form>"
	fromHtml += "<script>document.forms['alipaySubmit'].submit();</script>"
	return
}

// handlerParams 处理请求参数
// requestParams 请求的参数struct
func (a *AliClient) handlerParams(requestParams RequestParams) (url.Values, error) {
	// biz_content,notify_url,return_url,app_auth_token,method 这几个参数需要在调用该方法的时候就传入requestParams中
	urlValues := requestParams.GetOtherParams()
	// 系统参数数据组装
	urlValues.Add("app_id", a.appId)
	urlValues.Add("format", a.format)
	urlValues.Add("charset", a.charset)
	urlValues.Add("sign_type", a.signType)
	urlValues.Add("timestamp", time.Unix(time.Now().In(a.location).Unix(), 0).Format(consts.RequestTimestampFormat))
	urlValues.Add("version", a.version)
	if a.appCertSN != "" {
		urlValues.Add("app_cert_sn", a.appCertSN)
	}
	if a.alipayRootCertSn != "" {
		urlValues.Add("alipay_root_cert_sn", a.alipayRootCertSn)
	}

	rawBizContent := urlValues.Get(consts.BizContentFiled)
	if requestParams.GetNeedEncrypt() && rawBizContent != "" {
		// AES加密
		encryptBizContent, _ := a.encryptContent(rawBizContent)
		urlValues.Set(consts.BizContentFiled, encryptBizContent)
		urlValues.Add(consts.EncryptTypeField, a.encryptType)
	}

	// 获取签名
	sign, err := a.getSign(urlValues)
	if err != nil {
		return urlValues, err
	}
	urlValues.Add(consts.SignFiled, sign)
	return urlValues, nil
}

// AsyncNotifyVerifySign 异步通知验签，公钥、公钥证书两种模式下，异步通知验签方式相同。
// isLifeIsNo 该异步通知是否是生活号，生活号异步通知组成的待验签串里需要保留sign_type参数
// 其验签步骤为：
// 第一步：在通知返回参数列表中，除去sign、sign_type两个参数外，凡是通知返回回来的参数皆是待验签的参数。
// 第二步：将剩下参数进行url_decode, 然后进行字典排序，组成字符串，得到待签名字符串：
// 第三步：将签名参数（sign）使用base64解码为字节码串。
// 第四步：使用RSA的验签方法，通过签名字符串、签名参数（经过base64解码）及支付宝公钥验证签名。
// 第五步：在步骤四验证签名正确后，必须再严格按照如下描述校验通知数据的正确性
func (a *AliClient) AsyncNotifyVerifySign(urlValues url.Values, isLifeIsNo bool) (result bool, err error) {
	// 除去sign、sign_type两个参数&进行字典排序
	keys := make([]string, 0)
	for k, _ := range urlValues {
		if k == consts.SignFiled || !isLifeIsNo && k == consts.SignTypeFiled {
			continue
		}
		keys = append(keys, k)
	}
	sort.Strings(keys)
	// 组成字符串
	var valueList = make([]string, 0, 0)
	for _, key := range keys {
		valueList = append(valueList, key+"="+urlValues.Get(key))
	}
	// 待签名字符串
	var strParams = strings.Join(valueList, "&")
	// 获取异步通知返回的签名和签名算法类型
	sign, signType := urlValues.Get(consts.SignFiled), urlValues.Get(consts.SignTypeFiled)
	// base64解码
	signStr := base64.StdEncoding.EncodeToString([]byte(sign))
	// 签名验证
	if err = utils.RSAVerify(strParams, a.aliPublicKey, signStr, signType); err != nil {
		return
	}
	result = true
	return
}

// SyncVerifySign 同步返回验签，参考：https://opendocs.alipay.com/common/02mse7
// 开发者只对支付宝返回的 JSON 中 xxx_response 的值做验签（xxx 代表接口名），公钥、公钥证书两种模式下，异步通知验签方式不相同。
// 公钥模式说明：
// 1.xxx_response 的 JSON 值内容需要包含首尾的 { 和 } 两个尖括号，双引号也需要参与验签。
// 2.如果字符串中包含 http:// 的正斜杠，需要先将正斜杠做转义，默认打印出来的字符串是已经做过转义的。建议验签不通过时将正斜杠转义一次后再做一次验签。
// 针对公钥、公钥证书两种不同的签名模式，开放平台网关的响应报文有所区别，下面将向您分别介绍两种模式下，如何验签。
// 公钥证书模式说明：
// 1.公钥证书模式下，开放平台网关的同步响应报文中，会多一个响应参数 alipay_cert_sn（支付宝公钥证书序列号），与 xxx_repsose、sign 平级，该参数表示开发者需要使用该 SN 对应的支付宝公钥证书验签。详情请参考 常见问题。
// 2.支付宝公钥证书由于证书到期等原因，会重新签发新的证书（证书中密钥内容不变），开发者在自行实现的验签逻辑中需要判断当前使用的支付宝公钥证书 SN 与网关响应报文中的 SN 是否一致。若不一致，开发者需先调用 支付宝公钥证书下载接口 下载对应的支付宝公钥证书，再做验签。
func (a *AliClient) SyncVerifySign(rawData, apiMethodName string) (result bool, err error) {
	var resContent, signStr, alipayCertSn string
	resContent, signStr, alipayCertSn = a.parseJSONSource(apiMethodName, rawData)

	if signStr == "" {
		err = signDataIsEmptyErr
		return
	}
	//fmt.Println("返回的待签名数据为：", resContent)
	//fmt.Println("返回的签名为：", signStr)
	// 目前只考虑使用公钥模式签名

	var aliPublicKey *rsa.PublicKey // 支付宝公钥

	// 如果使用了公钥证书模式签名则就从支付宝证书中提取公钥
	if len(alipayCertSn) != 0 {
		// 当前使用的支付宝公钥证书 SN 与网关响应报文中的 SN 是否一致。若不一致，开发者需先调用 支付宝公钥证书下载接口 下载对应的支付宝公钥证书，再做验签
		if a.aliCertSN != alipayCertSn || a.certSnRelationPublicKey[alipayCertSn] == nil {
			var responseParam OpenAppAlipaycertDownloadResponseParams
			responseParam, err = a.OpenAppAlipaycertDownloadRequest(OpenAppAlipaycertDownloadRequestParams{})
			if err != nil {
				return
			}
			// 对公钥证书进行base64解码
			var alipayCertContent []byte
			alipayCertContent, err = base64.StdEncoding.DecodeString(responseParam.Data.AlipayCertContent)
			if err != nil {
				return
			}
			// 提取公钥证书中的公钥
			var x509Cert *x509.Certificate
			aliPublicKey, x509Cert, err = utils.GetPublicKeyFromCertContent(string(alipayCertContent))
			if err != nil {
				return
			}
			// 计算新的证书序列号&保存
			certSN := utils.Md5(x509Cert.Issuer.String() + x509Cert.SerialNumber.String())
			a.mutex.Lock()
			a.certSnRelationPublicKey[certSN] = aliPublicKey
			a.mutex.Unlock()
		} else {
			aliPublicKey = a.certSnRelationPublicKey[alipayCertSn]
		}
	} else {
		// 说明签名方式是公钥模式则直接取支付宝公钥即可
		aliPublicKey = a.aliPublicKey
	}
	// 签名验证
	if len(signStr) != 0 {
		if err = utils.RSAVerify(resContent, aliPublicKey, signStr, a.signType); err != nil {
			return
		}
	}
	result = true
	return
}

// 获取未加密内容
func (a *AliClient) parseJSONSource(apiMethodName, responseRawData string) (resContent, sign, alipayCertSn string) {
	rootNodeName := strings.Replace(apiMethodName, ".", "_", -1) + consts.ResponseSuffix
	rootIndex := strings.LastIndex(responseRawData, rootNodeName)
	errorIndex := strings.LastIndex(responseRawData, consts.ErrorResponse)
	if rootIndex > 0 {
		return a.parseJSONSourceItem(responseRawData, rootNodeName, rootIndex)
	} else if errorIndex > 0 {
		return a.parseJSONSourceItem(responseRawData, consts.ErrorResponse, errorIndex)
	}
	return
}

// parseJSONSignSourceItem 解析json
// resContent：待签名数据，sign：签名，alipayCertSn：支付宝证书序列号（加签模式为公钥证书模式时）
func (a *AliClient) parseJSONSourceItem(responseRawData, nodeName string, nodeIndex int) (resContent, sign, alipayCertSn string) {
	var signIndex = strings.LastIndex(responseRawData, "\""+consts.SignFiled+"\"")
	var alipayCertSnIndex = strings.LastIndex(responseRawData, "\""+consts.AlipayCertSnField+"\"") // 如果>0说明签名是公钥证书模式
	var splitStartIndex, splitEndIndex int

	if alipayCertSnIndex > 0 {
		splitEndIndex = alipayCertSnIndex - 1
	} else if signIndex > 0 {
		splitEndIndex = signIndex - 1
	} else {
		splitEndIndex = len(responseRawData) - 1
	}
	splitStartIndex = nodeIndex + len(nodeName) + 2
	if splitEndIndex-splitStartIndex <= 0 {
		return
	}
	// 获取返回值中的 待验签字段数据
	resContent = responseRawData[splitStartIndex:splitEndIndex]
	// 获取返回值中的 签名
	if signIndex > 0 {
		sign = responseRawData[signIndex+len(consts.SignFiled)+4:]
		sign = sign[:strings.LastIndex(sign, "\"")]
	}
	// 获取 返回值中的 alipay_cert_sn
	if alipayCertSnIndex > 0 {
		alipayCertSn = responseRawData[alipayCertSnIndex+len(consts.AlipayCertSnField)+4:]
		alipayCertSn = alipayCertSn[:strings.Index(alipayCertSn, "\"")]
	}
	return
}

// decryptJSONSignSource 解密内容
func (a *AliClient) decryptJSONSignSource(apiMethodName, responseRawData string) (resContent string, err error) {
	rootNodeName := strings.Replace(apiMethodName, ".", "_", -1) + consts.ResponseSuffix
	rootIndex := strings.LastIndex(responseRawData, rootNodeName)
	errorIndex := strings.LastIndex(responseRawData, consts.ErrorResponse)
	var splitStartIndex, splitEndIndex int
	var encryptContent string
	if rootIndex > 0 {
		splitStartIndex, splitEndIndex, encryptContent = a.parseEncryptJSONSourceItem(responseRawData, rootNodeName, rootIndex)
	} else if errorIndex > 0 {
		splitStartIndex, splitEndIndex, encryptContent = a.parseEncryptJSONSourceItem(responseRawData, consts.ErrorResponse, errorIndex)
	}
	bizContent, err := a.decryptContent(encryptContent)
	if err != nil {
		return "", err
	}
	bodyIndexContent := responseRawData[:splitStartIndex-1]
	bodyEndContent := responseRawData[splitEndIndex:]
	resContent = bodyIndexContent + bizContent + bodyEndContent
	return
}

// parseEncryptJSONSourceItem 解析加密的json
// resContent：待签名数据，sign：签名，alipayCertSn：支付宝证书序列号（加签模式为公钥证书模式时）
func (a *AliClient) parseEncryptJSONSourceItem(responseRawData, nodeName string, nodeIndex int) (splitStartIndex, splitEndIndex int, encryptContent string) {
	var signIndex = strings.LastIndex(responseRawData, "\""+consts.SignFiled+"\"")
	var alipayCertSnIndex = strings.LastIndex(responseRawData, "\""+consts.AlipayCertSnField+"\"") // 如果>0说明签名是公钥证书模式

	if alipayCertSnIndex > 0 {
		splitEndIndex = alipayCertSnIndex - 1
	} else if signIndex > 0 {
		splitEndIndex = signIndex - 1
	} else {
		splitEndIndex = len(responseRawData) - 1
	}
	splitStartIndex = nodeIndex + len(nodeName) + 3
	if splitEndIndex-splitStartIndex <= 0 {
		return
	}
	// 获取返回值中的 加密的待验签数据
	encryptContent = responseRawData[splitStartIndex : splitEndIndex-1]
	return
}

// sortParams 参数排序处理最终得到待签名字符串
// 按API要求, 参数名应按照第一个字符的键值 ASCII 码递增排序（字母升序排序）
// 如果遇到相同字符则按照第二个字符的键值 ASCII 码递增排序
func sortParams(urlValues url.Values) (strParams string) {
	// 进行字典排序
	keys := make([]string, 0)
	for k, _ := range urlValues {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	// 组成字符串
	var valueList = make([]string, 0, 0)
	for _, key := range keys {
		valueStr := urlValues.Get(key) // 去除空数据
		if len(valueStr) == 0 {
			urlValues.Del(key)
			continue
		}
		valueList = append(valueList, key+"="+valueStr)
	}
	strParams = strings.Join(valueList, "&")
	return
}

// getSign 获取签名
func (a *AliClient) getSign(urlValues url.Values) (signStr string, err error) {
	var strParams string
	strParams = sortParams(urlValues)
	fmt.Println("strParams:", strParams)
	return utils.RSASign(strParams, a.appPrivateKey, a.signType)
}

// 对bizContent内容进行加密
func (a *AliClient) encryptContent(content string) (encryptContent string, err error) {
	// 检查content是否为空
	if content == "" || strings.Trim(content, " ") == "" {
		err = encryptContentIsEmptyErr
		return
	}
	if a.encryptType == "" || a.encryptKey == "" {
		err = encryptKeyOrTypeIsEmptyErr
		return
	}
	if a.encryptType != consts.EncryptTypeAes {
		err = encryptTypeErr
		return
	}
	encryptContent, err = utils.AesCBCEncrypt(content, []byte(a.encryptKey))
	return
}

// 对bizContent内容进行解密（如果对请求参数做了 AES 加密，针对支付宝回复报文，开发者需要先验签，再解密）
func (a *AliClient) decryptContent(sourceContent string) (decryptContent string, err error) {
	decryptContent, err = utils.AesCBCDecrypt(sourceContent, []byte(a.encryptKey))
	return
}

// SetDataToBizContent 设置业务字段
func (a *AliClient) SetDataToBizContent(structData interface{}, needEncrypt bool) string {
	// 这种是针对公共参数中无biz_content的情况
	if structData == nil {
		return ""
	}
	bodyStr, _ := json.Marshal(structData)
	// todo 内容加密/解密这一块没有调试通================
	// 是否对biz_content内容进行加密，
	if needEncrypt {
		encryptContent, _ := a.encryptContent(string(bodyStr))
		return encryptContent
	}
	return string(bodyStr)
}

// GetCertSNFromPath 从证书中提取序列号
// certPath 证书文件路径
// certSN 返回证书序列号，SN 值是通过解析 X.509 证书文件中签发机构名称（name）以及内置序列号（serialNumber），
// 将二者拼接后的字符串计算 MD5 值获取，可参考开放平台 SDK 源码：
func (a *AliClient) GetCertSNFromPath(certPath string) (certSN string, err error) {
	certPEMBlock, err := ioutil.ReadFile(certPath)
	if err != nil {
		return
	}

	return a.GetCertSNFromContent(string(certPEMBlock))
}

// GetCertSNFromContent 从证书中提取序列号
// certContent 公钥应用证书内容字符串（包含begin，end）
// certSN 返回证书序列号，SN 值是通过解析 X.509 证书文件中签发机构名称（name）以及内置序列号（serialNumber），
// 将二者拼接后的字符串计算 MD5 值获取，可参考开放平台 SDK 源码：
func (a *AliClient) GetCertSNFromContent(certContent string) (certSN string, err error) {

	var x509Cert *x509.Certificate
	var publicKey *rsa.PublicKey

	publicKey, x509Cert, err = utils.GetPublicKeyFromCertContent(certContent)
	if err != nil {
		return
	}

	// 证书序列号的计算
	certSN = utils.Md5(x509Cert.Issuer.String() + x509Cert.SerialNumber.String())
	a.mutex.Lock()
	a.certSnRelationPublicKey[certSN] = publicKey
	a.mutex.Unlock()

	return
}

// GetRootCertSNFromPath 提取根证书序列号
// rootCertPath 根证书文件地址
// certSN 返回证书序列号，SN 值是通过解析 X.509 证书文件中签发机构名称（name）以及内置序列号（serialNumber），
// 将二者拼接后的字符串计算 MD5 值获取，可参考开放平台 SDK 源码：
func (a *AliClient) GetRootCertSNFromPath(rootCertPath string) (rootCertSN string, err error) {
	certPEMBlock, err := ioutil.ReadFile(rootCertPath)
	if err != nil {
		return
	}
	return a.GetRootCertSNFromContent(string(certPEMBlock))
}

// GetRootCertSNFromContent 获取根证书序列号
// rootCertContent 根证书文件内容
// certSN 返回证书序列号，SN 值是通过解析 X.509 证书文件中签发机构名称（name）以及内置序列号（serialNumber），
//// 将二者拼接后的字符串计算 MD5 值获取，可参考开放平台 SDK 源码：
func (a *AliClient) GetRootCertSNFromContent(rootCertContent string) (rootCertSN string, err error) {
	certStrSlice := strings.Split(rootCertContent, consts.CertificateSuffix)
	var rootCertSnSlice []string
	for _, v := range certStrSlice {
		x509Cert, _ := utils.ParseX509Certificate(v + consts.CertificateSuffix)
		if x509Cert == nil || x509Cert.SignatureAlgorithm != x509.SHA1WithRSA && x509Cert.SignatureAlgorithm != x509.SHA256WithRSA {
			continue
		}
		// 证书序列号的计算
		certSN := utils.Md5(x509Cert.Issuer.String() + x509Cert.SerialNumber.String())
		rootCertSnSlice = append(rootCertSnSlice, certSN)
	}
	if len(rootCertSnSlice) > 0 {
		rootCertSN = strings.Join(rootCertSnSlice, "_")
	}
	return
}

// LoadAppCertSN 从应用公钥证书中加载 应用公钥证书序列号SN
// certPath：从证书中提取序列号，certContent：从证书内容中提取序列号
func (a *AliClient) LoadAppCertSN(certPath, certContent string) {
	var certSN string
	if certPath != "" {
		certSN, _ = a.GetCertSNFromPath(certPath)
	} else {
		certSN, _ = a.GetCertSNFromContent(certContent)
	}
	a.appCertSN = certSN
}

// LoadAliCertSN 从支付宝公钥证书中加载 支付宝公钥证书序列号SN
// certPath：从证书中提取序列号，certContent：从证书内容中提取序列号
func (a *AliClient) LoadAliCertSN(certPath, certContent string) {
	var certSN string
	if certPath != "" {
		certSN, _ = a.GetCertSNFromPath(certPath)
	} else {
		certSN, _ = a.GetCertSNFromContent(certContent)
	}
	a.aliCertSN = certSN
}

// LoadAlipayRootCertSN 从支付宝根证书书中加载 支付宝根证书序列号SN
// certPath：从证书中提取序列号，certRootContent：从证书内容中提取序列号
func (a *AliClient) LoadAlipayRootCertSN(certRootPath, certRootContent string) {
	var certRootSN string
	if certRootPath != "" {
		certRootSN, _ = a.GetRootCertSNFromPath(certRootPath)
	} else {
		certRootSN, _ = a.GetRootCertSNFromContent(certRootContent)
	}
	a.alipayRootCertSn = certRootSN
}

// EncodeURLParam 将参数mapParams编码为url编码格式
func EncodeURLParam(mapParams map[string]interface{}) string {
	urlValues := url.Values{}
	for key, value := range mapParams {
		urlValues.Add(key, value.(string))
	}
	return urlValues.Encode()
}
