package zkp

import (
	"math/big"

	"github.com/ninjadotorg/constant/privacy-protocol"

	blake2b "github.com/minio/blake2b-simd"
)

// GenerateChallenge get hash of n points in G append with input values
// return blake_2b(G[0]||G[1]||...||G[CM_CAPACITY-1]||<values>)
// G[i] is list of all generator point of Curve
func GenerateChallenge(values [][]byte) *big.Int {
	appendStr := privacy.PedCom.G[0].Compress()
	for i := 1; i < CM_CAPACITY; i++ {
		appendStr = append(appendStr, privacy.PedCom.G[i].CompressPoint()...)
	}
	for i := 0; i < len(values); i++ {
		appendStr = append(appendStr, values[i]...)
	}
	hashFunc := blake2b.New256()
	hashFunc.Write(appendStr)
	hashValue := hashFunc.Sum(nil)
	result := big.NewInt(0)
	result.SetBytes(hashValue)
	result.Mod(result, privacy.Curve.Params().N)
	return result
}
