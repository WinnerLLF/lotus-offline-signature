## 离线钱包管理服务
    在离线的情况下进行,获取钱包、钱包创建、导出私钥、导入私钥、钱包签名等功能.

## 一、钱包服务
   * 创建钱包
     ````text
     methods:  get
         url:  http://127.0.0.1:8787/wallet/new
       param:  sigType(string) = bls 或者 secp256k1
   * 导出私钥
     ````text
     methods:  get
         url:  http://127.0.0.1:8787/wallet/export
       param:  address(string) = t3qczlflcf45vx3ay67yxvd7gwxm6zfyke7jnskt46phg3bbvmxcmyvl3cd4hac7zjfvcx4wa4xme45sw47lea
   * 导入私钥
     ````text
     methods:  get
         url:  http://127.0.0.1:8787/wallet/import
       param:  privateKey(string) = 7b2254797065223a22626c73222c22507269766174654b6579223a22496f664f724b34545470386a646c716741376d4447655a5a797a39734b62707a486a4159776e53754146593d227d
   * 删除钱包地址
     ````text
     methods:  get
         url:  http://127.0.0.1:8787/wallet/export
       param:  delete(string) = t3qczlflcf45vx3ay67yxvd7gwxm6zfyke7jnskt46phg3bbvmxcmyvl3cd4hac7zjfvcx4wa4xme45sw47lea
   * 获取钱包地址列表
     ````text
     methods:  get 
         url:  http://127.0.0.1:8787/wallet/list
   * 任意签名
        ````text
        methods:  get 
            url:  http://127.0.0.1:8787/wallet/sigAny
          param:  content(string)=签名内容 & address(string) = t3qczlflcf45vx3ay67yxvd7gwxm6zfyke7jnskt46phg3bbvmxcmyvl3cd4hac7zjfvcx4wa4xme45sw47lea
     
## 二、签名服务
   * 离线签名
     ````text
     methods:  get 
         url:  http://127.0.0.1:8787/wallet/sig
       param:  message(string) = 消息体
   * 离线签名消息体[上传需要签名的消息体：message.json]   
     ````text
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
 
## 三、web界面签名
   * 签名界面
     ````text
     methods: get
         url: http://127.0.0.1:8787/index 
    
## 四、服务部署
1. 服务编译
   * CGO_CFLAGS="-D__BLST_PORTABLE__" make
   
2. 服务启动
   * 运行离线签名程序 [nohup ./mwallet-signature &]