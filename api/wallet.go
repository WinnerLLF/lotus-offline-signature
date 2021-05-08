package api

import (
	"github.com/gin-gonic/gin"
)

/**
 * @Description:
 * @return *gin.Engine
 */
func RouterApiServer() *gin.Engine {
	// Create restful handler (using Gin)
	srv := new(walletDealWith)

	router := gin.Default()

	/* 离线钱包
	1、生成钱包
	2、保存钱包地址
	3、保存钱包私钥
	*/
	router.GET("/wallet/new", srv.newWallet)
	router.GET("/wallet/list", srv.walletList)
	router.GET("/wallet/export", srv.walletExport)
	router.GET("/wallet/import", srv.walletImport)
	router.GET("/wallet/delete", srv.walletDelete)

	/* 离线签名
	1、读取消息体
	2、消息签名
	3、任意签名
	*/
	router.GET("/wallet/sig", srv.sig)
	router.GET("/wallet/sigAny", srv.sigAny)

	return router
}
