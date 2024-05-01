package makexendgenesis_test

import (
	"math/big"
	"testing"

	futils "github.com/Fantom-foundation/go-opera/utils"
	"github.com/stretchr/testify/suite"
)

type MakeGenesisTestSuite struct {
	suite.Suite
}

func TestMakeGenesis(t *testing.T) {
	suite.Run(t, new(MakeGenesisTestSuite))
}

func (s *MakeGenesisTestSuite) TestCalc() {
	validators := []int{1, 2, 3, 4, 5}

	totalSupplyTarget := futils.ToFtm(200_000_000)
	eachValidatorStake := futils.ToFtm(1_000_000)

	freeRWA := totalSupplyTarget.Sub(totalSupplyTarget, eachValidatorStake.Mul(eachValidatorStake, new(big.Int).SetInt64(int64(len(validators)))))

	s.T().Logf("free rwa %v", freeRWA.String())

	s.EqualValues(0, futils.ToFtm(195_000_000).Cmp(freeRWA))
}
