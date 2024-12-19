package utils

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
)

const privateKey = `-----BEGIN RSA PRIVATE KEY-----
MIIEogIBAAKCAQEA2QmHhpye89O66rJDEh7K8izGEPukMGi1qhinKlFEbMR7Z5Hi
rrqRPmFxqDSlzhZnR7c+U1ljMz2PD1OZlXWg+HlRlI9kt8Cu+3L1oyJFLvVio/+X
8eQ9jWiFMlkhG1P0RIUhPRDLegk9TGqBo/8eVB8XcjuFN9w0DQkPkfjw+0L4AbtZ
L9wrH4K8WoZt2/v0WpiUc5ozDklRz9Xj5leDavyPQhI+SyST3NyXeO/rdlCHgUi3
4QF+tWYPyIMNVTu+psirKs8RtnXlMNJlYpuhg5yVH4ZidWrLwPtR0wfd+yK5WTR6
uN7+T5c1Zl9Xs1JefMEA0Z99p9+/0/0H3v3c3wIDAQABAoIBAQCFgpovEYuqXFYA
aBZgQDcB0M8qxVHUQaV29GFo48Me0aNEK/bxG22bnR/opDJW22s/L5cUeMlQTQ8D
E59H1KtoDmFD7Q29B7fckHRQnRVH9MVwMxvKPmgq4+AD/9Yg1H1P5UF0Ki+xa5uc
tWv4cUi74MBwkfK6UShn2GoOlr/PH9mT0TFMMlcf3SF3VVYwqnZH0+Uqe0XWKnAi
BFtsddYy+KLViwcB9EZBNu9nLS0K52IGxPPYrRfMjYnf9HsCFyLtfpuYiMLbNEDh
B7/EFrdr8b17EbBi+1f0P6zJ+YmGdsgJKBgSbktpROwk7+Vjt0Yp6bLzKGMX/c5Q
4c6dMYAxAoGBAN1oz+Zn1SpAkFtX3W0VmRX5FSd6ZDX+gM9XDdwNaoEpinarG9cc
eTpvTMY/+o0kZSeOgHSZOYISKMIwU85rnI+Wg3ecvhAEufSkYzSqPGhFvsdC1PC4
3uK2gELj3w1s/XhezP6spu+EZnI7Kur9rh+obD1A+jV1iFAeOYo1NPl3AoGBAPrx
2lP7PgsJMrIpDbEJvigSHKhMkIFzK+37DGKyqLq/5Oe5sQaf2Y6RZmLQf1VXMY4s
5CXoXdf83HNFruTAcpGUZ2iBs2Lcso9nkZIKjo8/hdKBUM6jlh8bfQOOqBC8ERVD
4rLJplBQ0vBHryr+9FNyWGDA8QktSZ+IpwotBpHZAoGBAN0mN6J3y98sJgRWuwCe
Ng5QgOSxxy53SsnOtjU54Uup5nepaWFvO1c2nITYmrYnyDG+kGfmNB7LbJKqGstv
iS2StTXgBaeSUPsex0hvUW8FWw1En9Thkx74Exy85qpOVs7IQhGO+h77LCilGLUn
NajBXXvBTynef40m42o8wvD5AoGAJdqh0pVNsKB8kz96H0CW/LhutjyGd+CFuepq
3eCRb0pPH111hMhMKyNnzHQjIR58DTRcXhOFRHWSU5kpXrxhC/DVAIDALKOaE+PN
o7gP9S/h4fU72U/3701YOFcmpw+XjBuncMzWV7s5dqF5nQTEmmnfRwkT0wyP1I7q
k7BjE1ECfx59vf7vd43eRIU8iTfnWu0w5F45bA8HkDvczG8cWTrbV4EFtQQodbUp
4jJxOGaTMnThw3qMvIuHl7AwxLQqNUyszjC3Y2CTqH97A8MzQ9hMFPeA8jvM6dqN
Gxta9OhlzN0x5yKAXHV1QDB0y0kfL8JgAhu/qIhQbDf90ufkhQ4=
-----END RSA PRIVATE KEY-----`

const publicKey = `-----BEGIN RSA PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA2QmHhpye89O66rJDEh7K
8izGEPukMGi1qhinKlFEbMR7Z5HirrqRPmFxqDSlzhZnR7c+U1ljMz2PD1OZlXWg
+HlRlI9kt8Cu+3L1oyJFLvVio/+X8eQ9jWiFMlkhG1P0RIUhPRDLegk9TGqBo/8e
VB8XcjuFN9w0DQkPkfjw+0L4AbtZL9wrH4K8WoZt2/v0WpiUc5ozDklRz9Xj5leD
avyPQhI+SyST3NyXeO/rdlCHgUi34QF+tWYPyIMNVTu+psirKs8RtnXlMNJlYpuh
g5yVH4ZidWrLwPtR0wfd+yK5WTR6uN7+T5c1Zl9Xs1JefMEA0Z99p9+/0/0H3v3c
3wIDAQAB
-----END RSA PUBLIC KEY-----`

func RSAEncrypt(text string) string {
	pubBlock, _ := pem.Decode([]byte(publicKey))
	pubKeyValue, _ := x509.ParsePKIXPublicKey(pubBlock.Bytes)
	pub := pubKeyValue.(*rsa.PublicKey)
	ciphertext, _ := rsa.EncryptPKCS1v15(rand.Reader, pub, []byte(text))
	return base64.StdEncoding.EncodeToString(ciphertext)
}

func RSADecrypt(base64Text string) string {
	ciphertext, _ := base64.StdEncoding.DecodeString(base64Text)
	priBlock, _ := pem.Decode([]byte(privateKey))
	priKey, _ := x509.ParsePKCS1PrivateKey(priBlock.Bytes)
	text, _ := rsa.DecryptPKCS1v15(rand.Reader, priKey, ciphertext)
	return string(text)
}
