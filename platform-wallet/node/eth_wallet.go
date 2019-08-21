package node

import (
	"math"
	"math/big"

	web3 "github.com/johnnycx127/go-web3"
	"github.com/johnnycx127/go-web3/dto"
	"github.com/johnnycx127/go-web3/eth/block"
	"github.com/johnnycx127/go-web3/providers"
	"github.com/shopspring/decimal"
)

type EthWallet struct {
	name   string
	client *web3.Web3
	// coin      CoinInfo
	contracts map[string]*ERCToken
}

func NewEthWallet(name string, config map[string]interface{}) EthWallet {
	connection := web3.NewWeb3(providers.NewHTTPProvider(config["host"].(string), 10, false))
	contract, err := connection.Eth.Contract("")
	if err != nil {
		panic(err)
	}
	wallet := EthWallet{name, connection, make(map[string]*ERCToken)}
	if config["contracts"] == nil {
		return wallet
	}
	contractConfig := config["contracts"].([]map[string]interface{})
	for _, cc := range contractConfig {
		wallet.contracts[cc["symbol"].(string)] = &ERCToken{
			symbol:           cc["symbol"].(string),
			name:             cc["name"].(string),
			decimals:         cc["decimals"].(int),
			constractAddress: cc["constract_address"].(string),
			constractCli:     contract,
		}
	}
	return wallet
}

func (ethw *EthWallet) GetNewAddress(account string, password string) (string, error) {
	return ethw.client.Personal.NewAccount(password)
}

func (ethw *EthWallet) GetBlockNum() (int64, error) {
	num, err := ethw.client.Eth.GetBlockNumber()
	if err != nil {
		return 0, err
	}
	return num.ToInt64(), nil
}

func (ethw *EthWallet) GetBalanceByAccount(address string, minConf int) (string, error) {
	amount, err := ethw.client.Eth.GetBalance(address, block.LATEST)
	if err != nil {
		return "", err
	}
	return decimal.NewFromFloat(math.Pow(10, (-18))).Mul(decimal.NewFromBigInt(amount.ToBigInt(), 0)).Round(8).String(), nil
}

func (ethw *EthWallet) UnlockAccount(address string, passphrase string, duration int) error {
	_, err := ethw.client.Personal.UnlockAccount(address, passphrase, uint64(duration))
	return err
}

func (ethw *EthWallet) SendTransaction(fromAccount string, toAddress string, amount string) (string, error) {
	amountDecimal, err := decimal.NewFromString(amount)
	amountWei, _ := big.NewInt(0).SetString(amountDecimal.Mul(decimal.NewFromFloat(math.Pow(10, 18))).String(), 10)
	txId, err := ethw.client.Eth.SendTransaction(&dto.TransactionParameters{
		From:     fromAccount,
		To:       toAddress,
		Gas:      big.NewInt(3000000),
		GasPrice: big.NewInt(18000000000),
		Value:    amountWei,
	})
	return txId, err
}
