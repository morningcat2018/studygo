package main

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"fmt"
	"math/big"
	"strconv"
)

// bytes to hex string
func bytesToHexString(b []byte) string {
	var buf bytes.Buffer
	for _, v := range b {
		t := strconv.FormatInt(int64(v), 16)
		if len(t) > 1 {
			buf.WriteString(t)
		} else {
			buf.WriteString("0" + t)
		}
	}
	return buf.String()
}

// hex string to bytes
func hexStringToBytes(s string) []byte {
	bs := make([]byte, 0)
	for i := 0; i < len(s); i = i + 2 {
		b, _ := strconv.ParseInt(s[i:i+2], 16, 16)
		bs = append(bs, byte(b))
	}
	return bs
}

func fromBase10(base10 string) *big.Int {
	i, ok := new(big.Int).SetString(base10, 10)
	if !ok {
		panic("bad number: " + base10)
	}
	return i
}

func encryptTest(publicKey rsa.PublicKey) {

	text := "你好，GOLANG" //需要加密的字符串
	// encryptedBytes, err := rsa.EncryptOAEP(sha256.New(), rand.Reader, &publicKey, []byte(text), nil)
	encryptedBytes, err := rsa.EncryptPKCS1v15(rand.Reader, &publicKey, []byte(text))
	if err != nil {
		panic(err)
	}
	fmt.Println(bytesToHexString(encryptedBytes))
}

func decryptTest(privateKey *rsa.PrivateKey) {
	//decryptedBytes, err := privateKey.Decrypt(nil, encryptedBytes, &rsa.OAEPOptions{Hash: crypto.SHA256})
	minwen := "930562550d847580790c4b99a1728333762186ae33dcd1341edf4808e9f38e58f70f9974aeeeab4042c6e91a471b7a47286ed561164cddd11eb2001c38d9d845"
	encryptedBytes2 := hexStringToBytes(minwen)
	decryptedBytes, err := rsa.DecryptPKCS1v15(rand.Reader, privateKey, encryptedBytes2)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(decryptedBytes))
}

func getKey() {
	// 生成密钥
	var num int = 1
	privateKey, err := rsa.GenerateKey(rand.Reader, 512+64*num)
	if err != nil {
		panic(err)
	}
	publicKey := privateKey.PublicKey

	// n e d
	n := publicKey.N
	e := publicKey.E
	d := privateKey.D
	fmt.Println(n)
	fmt.Println(e)
	fmt.Println(d)
}

func main() {

	privateKey := &rsa.PrivateKey{
		PublicKey: rsa.PublicKey{
			N: fromBase10("10024046402334537665438324593256381027774491568777420349437153155855551494930992968579095115335998035992292512366201628140916921363490577352034455478837127"),
			E: 65537,
		},
		D:      fromBase10("3547885627180858685742517620049361647927996497083495467072710142411544203677535574073374414767899187836766831505613082822722655998801933606944964650191553"),
		Primes: []*big.Int{
			//fromBase10("16775196964030542637"),
			//fromBase10("17328218193455850539"),
		},
	}

	// 加密
	publicKey := privateKey.PublicKey
	//encryptTest(publicKey)

	// 解密
	decryptTest(privateKey)

	pubKeyBytes := x509.MarshalPKCS1PublicKey(&publicKey)
	fmt.Println(bytesToHexString(pubKeyBytes))

	//pem.EncodeToMemory()

}

func encrypt() {

	publicKeyHex := "305c300d06092a864886f70d0101010500034b00304802410090f2342cac47dfb26c5dd7aeb9ab0967201c57aeaf1943d" +
		"67ceefff57fd30e64929189d16d05aec39852e519145c5c0c2dc24d2105ba89aa22033a95988bf2170203010001"
	publicKeyByte := hexStringToBytes(publicKeyHex)
	fmt.Println(bytesToHexString(publicKeyByte))

	privateKey, _ := rsa.GenerateKey(rand.Reader, 512)
	publicKeyX := privateKey.PublicKey
	publicKeyByte = x509.MarshalPKCS1PublicKey(&publicKeyX)
	fmt.Println(bytesToHexString(publicKeyByte))

	publicKey, err := x509.ParsePKCS1PublicKey(publicKeyByte)
	if err != nil {
		panic(err)
	}

	text := "morningcat" //需要加密的字符串
	// encryptedBytes, err := rsa.EncryptOAEP(sha256.New(), rand.Reader, &publicKey, []byte(text), nil)
	encryptedBytes, err := rsa.EncryptPKCS1v15(rand.Reader, publicKey, []byte(text))
	if err != nil {
		panic(err)
	}
	fmt.Println(bytesToHexString(encryptedBytes))
}
