package balance

import (
	"get-block/domain/balance"
	"math/big"
	"strconv"
)

const blockCount = 100

type Service struct {
	getBlockClient balance.GetBlockClientInterface
}

type MostChangedBalanceDto struct {
	Address string
}

func NewService(getBlockClient balance.GetBlockClientInterface) *Service {
	return &Service{getBlockClient: getBlockClient}
}

func (service *Service) GetMostChangedBalanceAddress() (MostChangedBalanceDto, error) {
	mostChangedBalanceDto := MostChangedBalanceDto{}
	addressMap := map[string]*big.Int{}
	lastBlock, err := service.getBlockClient.GetBlockNumber()
	lastBlockInt := convertHexToInt(lastBlock)
	if err != nil {
		return mostChangedBalanceDto, err
	}

	for i := lastBlockInt.Int64(); i >= lastBlockInt.Int64()-blockCount; i-- {
		currentBlockNumber := convertIntBlockNumberToHex(i)
		block, err := service.getBlockClient.GetBlockByNumber(currentBlockNumber)

		if err != nil {
			return mostChangedBalanceDto, err
		}
		for _, transaction := range block.Result.Transactions {
			from := transaction.From
			to := transaction.To
			value := convertHexToInt(transaction.Value)

			if _, ok := addressMap[from]; ok {
				addressMap[from] = big.NewInt(0).Sub(addressMap[from], value)
			} else {
				addressMap[from] = big.NewInt(0).Sub(big.NewInt(0), value)
			}

			if _, ok := addressMap[to]; ok {
				addressMap[to] = big.NewInt(0).Add(addressMap[to], value)
			} else {
				addressMap[to] = value
			}
		}
	}

	mostChangedBalanceDto = service.GetMostChangedBalanceDto(addressMap)
	return mostChangedBalanceDto, nil
}

func (service *Service) GetMostChangedBalanceDto(addressMap map[string]*big.Int) MostChangedBalanceDto {
	mostChangedBalanceDto := MostChangedBalanceDto{}
	maxValue := big.NewInt(0)

	for address, value := range addressMap {
		if value.CmpAbs(maxValue) == +1 { // value > maxValue
			maxValue = value
			mostChangedBalanceDto = MostChangedBalanceDto{Address: address}
		}
	}
	return mostChangedBalanceDto
}

func convertHexToInt(number string) *big.Int {
	bigNum, _ := new(big.Int).SetString(number, 0)
	return bigNum
}

func convertIntBlockNumberToHex(number int64) string {
	return "0x" + strconv.FormatInt(number, 16)
}
