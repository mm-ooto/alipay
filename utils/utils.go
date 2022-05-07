package utils

import (
	"crypto"
	"crypto/md5"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/hex"
	"encoding/pem"
	"io/ioutil"
)

// ParsePKCS1PrivateKey 解析私钥
func ParsePKCS1PrivateKey(privateKeyPemStr string) (privateKey *rsa.PrivateKey, err error) {
	// pem解码
	block, _ := pem.Decode([]byte(privateKeyPemStr))
	// x509解码
	// ParsePKCS1PrivateKey解析ASN.1 PKCS#1 DER编码的rsa私钥。
	privateKey, err = x509.ParsePKCS1PrivateKey(block.Bytes)
	return
}

// ParsePKIXPublicKey 解析公钥
func ParsePKIXPublicKey(publicKeyPemStr string) (publicKey *rsa.PublicKey, err error) {
	// pem解码
	block, _ := pem.Decode([]byte(publicKeyPemStr))
	// x509解码
	// ParsePKIXPublicKey解析一个DER编码的公钥。这些公钥一般在以"BEGIN PUBLIC KEY"出现的PEM块中
	publicKeyInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	// 类型断言
	publicKey = publicKeyInterface.(*rsa.PublicKey)
	return
}

// ParseX509Certificate 解析X.509编码的证书
func ParseX509Certificate(certPemStr string) (x509Cert *x509.Certificate, err error) {
	// pem解码
	block, _ := pem.Decode([]byte(certPemStr))
	x509Cert, err = x509.ParseCertificate(block.Bytes)
	if err != nil {
		return
	}
	return
}

// RSASign 签名
// data 排序后的待签名字符串
// rsaType签名算法类型：可选值（RSA,RSA2）
func RSASign(data string, rsaPrivateKey *rsa.PrivateKey, rsaType string) (string, error) {
	hashP := crypto.SHA256
	if rsaType == "RSA" {
		hashP = crypto.SHA1
	}
	hash := hashP.New()
	hash.Write([]byte(data))
	sign, err := rsa.SignPKCS1v15(rand.Reader, rsaPrivateKey, hashP, hash.Sum(nil))
	if err != nil {
		return "", err
	}
	//  对签名进行base64编码
	signByte := base64.StdEncoding.EncodeToString(sign)
	return signByte, nil
}

// RSAVerify 验签
// data 排序后的待签名字符串
func RSAVerify(data string, rsaPublicKey *rsa.PublicKey, signData string, rsaType string) (err error) {
	hashP := crypto.SHA256
	if rsaType == "RSA" {
		hashP = crypto.SHA1
	}
	// 对签名进行base64解码
	sign, err := base64.StdEncoding.DecodeString(signData)
	if err != nil {
		return err
	}
	hash := hashP.New()
	hash.Write([]byte(data))
	return rsa.VerifyPKCS1v15(rsaPublicKey, hashP, hash.Sum(nil), sign)
}

// GetPemPublic 将公钥字符串转换为RSA公钥格式
func GetPemPublic(rawPublicKey string) string {
	publicPemStr := "-----BEGIN PUBLIC KEY-----\n"
	strlen := len(rawPublicKey)
	for i := 0; i < strlen; i += 64 {
		if i+64 >= strlen {
			publicPemStr += rawPublicKey[i:] + "\n"
		} else {
			publicPemStr += rawPublicKey[i:i+64] + "\n"
		}
	}
	publicPemStr += "-----END PUBLIC KEY-----"
	return publicPemStr
}

// GetPemPrivate 将私钥字符串转换为RSA私钥格式
func GetPemPrivate(rawPrivateKey string) string {
	privatePemStr := "-----BEGIN RSA PRIVATE KEY-----\n"
	strlen := len(rawPrivateKey)
	for i := 0; i < strlen; i += 64 {
		if i+64 >= strlen {
			privatePemStr += rawPrivateKey[i:] + "\n"
		} else {
			privatePemStr += rawPrivateKey[i:i+64] + "\n"
		}
	}
	privatePemStr += "-----END RSA PRIVATE KEY-----"
	return privatePemStr
}

// GetPemCert 将证书字符串转换为Cert证书格式
func GetPemCert(rawCert string) string {
	certPemStr := "-----BEGIN CERTIFICATE-----\n"
	strlen := len(rawCert)
	for i := 0; i < strlen; i += 76 {
		if i+76 >= strlen {
			certPemStr += rawCert[i:] + "\n"
		} else {
			certPemStr += rawCert[i:i+76] + "\n"
		}
	}
	certPemStr += "-----END CERTIFICATE-----"
	return certPemStr
}

// GetPublicKeyFromCertPath 从证书中提取公钥
// certPath 证书文件路径
func GetPublicKeyFromCertPath(certPath string) (publicKey *rsa.PublicKey, x509Cert *x509.Certificate, err error) {
	certPEMBlock, err := ioutil.ReadFile(certPath)
	if err != nil {
		return
	}
	x509Cert, err = ParseX509Certificate(string(certPEMBlock))
	if err != nil {
		return
	}
	// 获取该证书里面的公钥
	var ok bool
	publicKey, ok = x509Cert.PublicKey.(*rsa.PublicKey)
	if !ok {
		return
	}
	return
}

// GetPublicKeyFromCertContent 从证书content中提取公钥
// certContent 公钥应用证书内容字符串（包含begin，end）
func GetPublicKeyFromCertContent(certContent string) (publicKey *rsa.PublicKey, x509Cert *x509.Certificate, err error) {

	x509Cert, err = ParseX509Certificate(certContent)
	if err != nil {
		return
	}
	// 获取该证书里面的公钥
	var ok bool
	publicKey, ok = x509Cert.PublicKey.(*rsa.PublicKey)
	if !ok {
		return
	}
	return
}

func Md5(plainText string) string {
	h := md5.New()
	_, err := h.Write([]byte(plainText))
	if err != nil {
		panic(err)
	}
	md5String := hex.EncodeToString(h.Sum(nil))
	return md5String
}
