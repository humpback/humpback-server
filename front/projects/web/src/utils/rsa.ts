import JSEncrypt from "jsencrypt"

const publicKey = `-----BEGIN RSA PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA2QmHhpye89O66rJDEh7K
8izGEPukMGi1qhinKlFEbMR7Z5HirrqRPmFxqDSlzhZnR7c+U1ljMz2PD1OZlXWg
+HlRlI9kt8Cu+3L1oyJFLvVio/+X8eQ9jWiFMlkhG1P0RIUhPRDLegk9TGqBo/8e
VB8XcjuFN9w0DQkPkfjw+0L4AbtZL9wrH4K8WoZt2/v0WpiUc5ozDklRz9Xj5leD
avyPQhI+SyST3NyXeO/rdlCHgUi34QF+tWYPyIMNVTu+psirKs8RtnXlMNJlYpuh
g5yVH4ZidWrLwPtR0wfd+yK5WTR6uN7+T5c1Zl9Xs1JefMEA0Z99p9+/0/0H3v3c
3wIDAQAB
-----END RSA PUBLIC KEY-----`

export function RSAEncrypt(text: string): string {
  const encryptor = new JSEncrypt()
  encryptor.setPublicKey(publicKey)
  return encryptor.encrypt(text) as string
}
