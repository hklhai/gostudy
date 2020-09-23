package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"fmt"
)

//输入明文，输出密文
func aesCtrEncrypt(plainText, key []byte) ([]byte, error) {
	// 第一步：创建aes密码接口
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	//第二步：创建分组模式ctr
	// iv 要与算法长度一致，16字节
	// 使用bytes.Repeat创建一个切片，长度为blockSize()，16个字符"1"
	iv := bytes.Repeat([]byte("1"), block.BlockSize())
	stream := cipher.NewCTR(block, iv)

	//第三步：加密
	dst := make([]byte, len(plainText))
	stream.XORKeyStream(dst, plainText)

	return dst, nil
}

//输入密文，得到明文
func aesCtrDecrypt(encryptData, key []byte) ([]byte, error) {
	return aesCtrEncrypt(encryptData, key)
}

// 堆成加密
// https://blog.csdn.net/hbshhb/article/details/93404539
func main() {
	//明文，需要加密的数据
	src := "teamA-01"

	//对称秘钥，aes，16字节
	key := "12345678876543AD" //16
	encryptData, err := aesCtrEncrypt([]byte(src), []byte(key))
	if err != nil {
		fmt.Println("加密出错err:", err)
		return
	}

	fmt.Printf("encryptData： %x\n", encryptData)

	//调用解密函数
	plainText, err := aesCtrDecrypt(encryptData, []byte(key))
	if err != nil {
		fmt.Println("err:", err)
		return
	}

	fmt.Printf("解密后的数据: %s\n", plainText)
}
