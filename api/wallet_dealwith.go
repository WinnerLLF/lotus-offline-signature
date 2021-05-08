package api

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"lotus-offline-signature/types"
	"lotus-offline-signature/utils"
	"lotus-offline-signature/wallet"
	"github.com/filecoin-project/go-address"
	"github.com/gin-gonic/gin"
	"strings"
)

const (
	success_code = 200
	error_code   = 404
)

type walletDealWith struct {
}

var wt *wallet.Wallet

func init() {
	w, err := wallet.NewWallet(wallet.NewMemKeyStore())
	if err != nil {
		fmt.Printf("NewWallet err:%v\n", err)
	}
	wt = w
}

func InitCache() {
	keyList := utils.ReadListDB()

	for _, value := range keyList {
		var ki types.KeyInfo
		data, err := hex.DecodeString(strings.TrimSpace(value))
		if err != nil {
			continue
		}
		if err := json.Unmarshal(data, &ki); err != nil {
			continue
		}
		addr, err := wt.Import(&ki)
		if err != nil {
			continue
		}
		fmt.Printf("import walletAddress info:%v\n", addr)
	}
}

/**
 * @Description: wallet new [sigType = bls、secp256k1]
 * @receiver wd
 * @param c
 */
func (wd *walletDealWith) newWallet(c *gin.Context) {
	typeSig := c.Query("sigType")
	if len(wd.replace(typeSig)) == 0 {
		c.JSON(error_code, typeSig)
		return
	}

	wg, err := wt.GenerateKey(wallet.ActSigType(typeSig))
	if err != nil {
		c.JSON(error_code, err.Error())
		return
	}

	// get privateKey
	wi, err := wt.Export(wg)
	if err != nil {
		c.JSON(error_code, err.Error())
		return
	}
	b, err := json.Marshal(wi)
	if err != nil {
		c.JSON(error_code, err.Error())
		return
	}
	keyInfo := hex.EncodeToString(b)
	if err := utils.WriteDB(wg.String(), keyInfo); err != nil {
		c.JSON(error_code, err.Error())
		return
	}

	c.JSON(success_code, wg.String())
}

/**
 * @Description: wallet list
 * @receiver wd
 * @param c
 */
func (wd *walletDealWith) walletList(c *gin.Context) {
	addrList, err := wt.ListAddrs()
	if err != nil {
		c.JSON(error_code, err.Error())
		return
	}

	var walletList []string
	for _, value := range addrList {
		walletList = append(walletList, value.String())
	}

	c.JSON(success_code, walletList)
}

/**
 * @Description: wallet import [privateKey = xxx]
 * @receiver wd
 * @param c
 */
func (wd *walletDealWith) walletImport(c *gin.Context) {
	privateKey := c.Query("privateKey")
	if len(wd.replace(privateKey)) == 0 {
		c.JSON(error_code, privateKey)
		return
	}

	var ki types.KeyInfo
	data, err := hex.DecodeString(strings.TrimSpace(privateKey))
	if err != nil {
		c.JSON(error_code, err)
		return
	}
	if err := json.Unmarshal(data, &ki); err != nil {
		c.JSON(error_code, err)
		return
	}
	addr, err := wt.Import(&ki)
	if err != nil {
		c.JSON(error_code, err)
		return
	}

	c.JSON(success_code, addr.String())
}

/**
 * @Description: wallet export [address = xxx]
 * @receiver wd
 * @param c
 */
func (wd *walletDealWith) walletExport(c *gin.Context) {
	addWallet := c.Query("address")
	if len(wd.replace(addWallet)) == 0 {
		c.JSON(error_code, addWallet)
		return
	}

	addr, err := address.NewFromString(addWallet)
	if err != nil {
		c.JSON(error_code, err.Error())
		return
	}

	wi, err := wt.Export(addr)
	if err != nil {
		c.JSON(error_code, err.Error())
		return
	}

	b, err := json.Marshal(wi)
	if err != nil {
		c.JSON(error_code, err.Error())
		return
	}
	keyInfo := hex.EncodeToString(b)

	c.JSON(success_code, keyInfo)
}

/**
 * @Description: wallet delete [address = xxx]
 * @receiver wd
 * @param c
 */
func (wd *walletDealWith) walletDelete(c *gin.Context) {
	addWallet := c.Query("address")
	if len(wd.replace(addWallet)) == 0 {
		c.JSON(error_code, addWallet)
		return
	}

	addr, err := address.NewFromString(addWallet)
	if err != nil {
		c.JSON(error_code, err.Error())
		return
	}

	// delete key
	if err := wt.DeleteKey(addr); err != nil {
		c.JSON(error_code, err.Error())
		return
	}

	// delete cache db
	if err := utils.DeleteByKey(addr.String()); err != nil {
		c.JSON(error_code, err.Error())
		return
	}

	c.JSON(success_code, "delete walletAddress success!")
}

/**
 * @Description: replace string
 * @receiver wd
 * @param str
 * @return string
 */
func (wd *walletDealWith) replace(str string) string {
	return strings.Replace(str, " ", "", -1)
}
