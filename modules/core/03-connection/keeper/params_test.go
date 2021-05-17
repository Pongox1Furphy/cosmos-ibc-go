package keeper_test

import (
	"github.com/cosmos/ibc-go/modules/core/03-connection/types"
)

func (suite *KeeperTestSuite) TestParams() {
	expParams := types.DefaultParams()

	params := suite.chainA.App.GetIBCKeeper().ConnectionKeeper.GetParams(suite.chainA.GetContext())
	suite.Require().Equal(expParams, params)

	expParams.MaxTimePerBlock = 10
	suite.chainA.App.GetIBCKeeper().ConnectionKeeper.SetParams(suite.chainA.GetContext(), expParams)
	params = suite.chainA.App.GetIBCKeeper().ConnectionKeeper.GetParams(suite.chainA.GetContext())
	suite.Require().Equal(uint64(10), expParams.MaxTimePerBlock)
}
