module lotus-offline-signature

go 1.16

require (
	github.com/filecoin-project/filecoin-ffi v0.30.4-0.20200910194244-f640612a1a1f
	github.com/filecoin-project/go-address v0.0.5
	github.com/filecoin-project/go-crypto v0.0.0-20191218222705-effae4ea9f03
	github.com/filecoin-project/go-state-types v0.1.0
	github.com/filecoin-project/lotus v1.9.0
	github.com/gin-gonic/gin v1.6.3
	github.com/ipfs/go-block-format v0.0.3
	github.com/ipfs/go-cid v0.0.7
	github.com/ipfs/go-log/v2 v2.1.2-0.20200626104915-0016c0b4b3e4
	github.com/minio/blake2b-simd v0.0.0-20160723061019-3f5f724cb5b1
	github.com/stretchr/testify v1.7.0
	github.com/syndtr/goleveldb v1.0.0
	github.com/whyrusleeping/cbor-gen v0.0.0-20210219115102-f37d292932f2
	golang.org/x/lint v0.0.0-20200302205851-738671d3881b // indirect
	golang.org/x/xerrors v0.0.0-20200804184101-5ec99f83aff1
)

replace github.com/filecoin-project/filecoin-ffi => ./extern/filecoin-ffi
