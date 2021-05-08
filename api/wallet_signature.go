package api

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"lotus-offline-signature/types"
	"github.com/filecoin-project/go-address"
	"github.com/gin-gonic/gin"
	"os"
)

/**
 * @Description:
 * @return types.Message
 */
func readMessage() types.Message {
	type message struct {
		Version    int64
		To         string
		From       string
		Nonce      uint64
		Value      string
		GasLimit   uint64
		GasFeeCap  string
		GasPremium string
		Method     uint64
		Params     []byte
	}
	filePtr, err := os.Open(fmt.Sprintf("%v", "./json/message.json"))
	if err != nil {
		fmt.Printf("Open err:%v\n", err)
		return types.Message{}
	}
	defer filePtr.Close()

	decoder := json.NewDecoder(filePtr)
	var msg message
	err = decoder.Decode(&msg)
	if err != nil {
		return types.Message{}
	}

	to, err := address.NewFromString(msg.To)
	if err != nil {
		return types.Message{}
	}
	from, err := address.NewFromString(msg.From)
	if err != nil {
		return types.Message{}
	}

	msge := types.Message{
		Version:    msg.Version,
		To:         to,
		From:       from,
		Nonce:      msg.Nonce,
		Value:      Bigint_Transform(msg.Value),
		GasLimit:   int64(msg.GasLimit),
		GasFeeCap:  Bigint_Transform(msg.GasFeeCap),
		GasPremium: Bigint_Transform(msg.GasPremium),
		Method:     0,
		Params:     nil,
	}

	return msge
}

/**
 * @Description:
 * @param msg
 */
func writeMessage(msg string) {
	f, err := os.Create(fmt.Sprintf("%v", "./json/message.json"))
	if err != nil {
		fmt.Println("create err：%v\n", err)
		return
	}

	_, err = f.WriteString(msg)
	if err != nil {
		fmt.Println("write err：%v\n", err)
		f.Close()
		return
	}
	err = f.Close()
	if err != nil {
		fmt.Println("clone err：%v\n", err)
		return
	}
}

/**
 * @Description: msg sig
 * @receiver wd
 * @param c
 */
func (wd *walletDealWith) sig(c *gin.Context) {
	// write message
	addWallet := c.Query("message")
	writeMessage(addWallet)
	// read message
	msg := readMessage()

	// wallet sig
	signature, err := wt.Sign(context.TODO(), msg.From, msg.Cid().Bytes())
	if err != nil {
		fmt.Printf("Sign err：%v\n", err)
		c.JSON(error_code, err)
		return
	}

	signedMessage := &types.SignedMessage{
		Message:   msg,
		Signature: *signature,
	}

	c.JSON(success_code, signedMessage)
}

/**
 * @Description:
 * @receiver wd
 * @param c
 */
func (wd *walletDealWith) sigAny(c *gin.Context) {
	content := c.Query("content")
	walletAdd := c.Query("address")

	addr, err := address.NewFromString(walletAdd)
	if err != nil {
		fmt.Printf("NewFromString err：%v\n", err)
		c.JSON(error_code, err)
		return
	}

	msg, err := hex.DecodeString(content)
	if err != nil {
		c.JSON(error_code, err)
		return
	}

	signature, err := wt.Sign(context.TODO(), addr, msg)
	if err != nil {
		fmt.Printf("Sign err：%v\n", err)
		c.JSON(error_code, err)
		return
	}

	sigBytes := append([]byte{byte(signature.Type)}, signature.Data...)

	c.JSON(success_code, hex.EncodeToString(sigBytes))
}

/**
 * @Description:
 * @param value
 * @return types.BigInt
 */
func Bigint_Transform(value string) types.BigInt {
	bigInt, err := types.BigFromString(value)
	if err != nil {
		return bigInt
	}
	return bigInt
}
