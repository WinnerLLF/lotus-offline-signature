## Offline wallet management service
    It can obtain the wallet, create the wallet, export the private key, import the private key, sign the wallet and so on.

## 一、Wallet service
   * Create a wallet
     ````text
     methods:  get
         url:  http://127.0.0.1:8787/wallet/new
       param:  sigType(string) = bls 或者 secp256k1
   * Export the private key
     ````text
     methods:  get
         url:  http://127.0.0.1:8787/wallet/export
       param:  address(string) = t3qczlflcf45vx3ay67yxvd7gwxm6zfyke7jnskt46phg3bbvmxcmyvl3cd4hac7zjfvcx4wa4xme45sw47lea
   * Import the private key
     ````text
     methods:  get
         url:  http://127.0.0.1:8787/wallet/import
       param:  privateKey(string) = 7b2254797065223a22626c73222c2250726545470386a646c716741376d4447655a5a797a39734b62707a486a4159776e53754146593d227d
   * Delete wallet address
     ````text
     methods:  get
         url:  http://127.0.0.1:8787/wallet/export
       param:  delete(string) = t3qczlflcf45vx3ay67yxvd7gwxm6zfyke7jnskt46phg3bbvmxcmyvl3cd4hac7zjfvcx4wa4xme45sw47lea
   * Get a list of wallet addresses
     ````text
     methods:  get 
         url:  http://127.0.0.1:8787/wallet/list
   * Any signature
        ````text
        methods:  get 
            url:  http://127.0.0.1:8787/wallet/sigAny
          param:  content(string)=签名内容 & address(string) = t3qczlflcf45vx3ay67yxvd7gwxm6zfyke7jnskt46phg3bbvmxcmyvl3cd4hac7zjfvcx4wa4xme45sw47lea
     
## 二、Signature services
   * Off-line signature
     ````text
     methods:  get 
         url:  http://127.0.0.1:8787/wallet/sig
       param:  message(string) = 消息体
   * Sign the body of the message offline [upload the body of the message that needs to be signed: message.json]  
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
 
## 三、The Web world signs in person
   * The signature interface
     ````text
     methods: get
         url: http://127.0.0.1:8787/index 
