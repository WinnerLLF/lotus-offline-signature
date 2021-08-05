package api

import (
	"github.com/gin-gonic/gin"
)

func RouterApiServer() *gin.Engine {
	// Create restful handler (using Gin)
	srv := new(walletDealWith)

	router := gin.Default()

	/* Offline Wallet
	1、Generate a wallet
	2、Save wallet address
	3、Save wallet private key
	*/
	router.GET("/wallet/new", srv.newWallet)
	router.GET("/wallet/list", srv.walletList)
	router.GET("/wallet/export", srv.walletExport)
	router.GET("/wallet/import", srv.walletImport)
	router.GET("/wallet/delete", srv.walletDelete)

	/* Off-line signature
	1、Read message body
	2、Message signature
	3、Any signature
	*/
	router.GET("/wallet/sig", srv.sig)
	router.GET("/wallet/sigAny", srv.sigAny)

	return router
}
