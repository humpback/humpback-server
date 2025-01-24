import JSEncrypt from "jsencrypt"
import { RSAPublicKey } from "@/models"

export function RSAEncrypt(text: string): string {
  const encryptor = new JSEncrypt()
  encryptor.setPublicKey(RSAPublicKey)
  return encryptor.encrypt(text) as string
}
