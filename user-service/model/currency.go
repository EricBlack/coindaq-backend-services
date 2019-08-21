package model

import (
	"bx.com/user-service/bxgo"
	"time"
)

type Currency struct {
	CurrencyId                     string    `xorm:"not null pk comment('主键') VARCHAR(32)"`
	CurrencyName                   string    `xorm:"not null comment('币名称') VARCHAR(50)"`
	CurrencyTag                    string    `xorm:"not null comment('币标签') VARCHAR(50)"`
	CurrencyType                   int       `xorm:"comment('币种类型，A类或B类') INT(11)"`
	CurrencyTranstionCommission    int64     `xorm:"not null default 2 comment('单笔手续费（按千分数计） 默认千分之二    ') BIGINT(11)"`
	CurrencyTranstionCommissionMin int64     `xorm:"not null default 0 comment('最小手续费（非百分数）') BIGINT(11)"`
	CurrencyCreateTime             time.Time `xorm:"comment('币创建日期') DATETIME"`
	CurrencyStatus                 int       `xorm:"default 0 comment('币状态	0 前端不可见 1 前端可见') INT(11)"`
	CurrencyRechargeStatus         int       `xorm:"default 0 comment('币充值状态 0 不可充值 1可充值') INT(11)"`
	CurrencyWithdrawStatus         int       `xorm:"default 0 comment('币提现状态 0 不可提现 1 可提现') INT(11)"`
	CurrencyWithdrawCommission     string    `xorm:"not null default '2' comment('提现手续费(固定数值,每种代币具体值不一样)') VARCHAR(11)"`
	CurrencyWithdrawCommissionMin  int64     `xorm:"not null default 0 comment('提现最小手续费') BIGINT(11)"`
	CurrencyFundingPrice           string    `xorm:"comment('众筹价格  1ETH=1231231TNB') VARCHAR(50)"`
	CurrencyWhitePaperUrl          string    `xorm:"comment('白皮书地址') VARCHAR(255)"`
	CurrencyOfficialUrl            string    `xorm:"comment('官网地址') VARCHAR(255)"`
	CurrencyBlockCheckUrl          string    `xorm:"comment('区块查询地址') VARCHAR(255)"`
	CurrencyIssueTime              time.Time `xorm:"comment('发行时间') DATETIME"`
	CurrencyExtractCountMin        string    `xorm:"comment('提币最小数量') VARCHAR(20)"`
	CurrencyAvatar                 string    `xorm:"comment('币头像') VARCHAR(255)"`
	RatioToUsdt                    string    `xorm:"comment('当前币兑换USDT的比例(eg:value=100，则代表100个USDT=1个该币)') DECIMAL(50)"`
	EthereumUsableStatus           int       `xorm:"comment('是否以太坊 0:非以太坊1：以太坊 默认为0') INT(11)"`
	ContractAddress                string    `xorm:"comment('合约地址') VARCHAR(255)"`
	ContractDecimalNb              int       `xorm:"comment('合约小数点位数') INT(11)"`
}

func (cr Currency) TableName() string {
	return "currency"
}

func GetCurrency(currencyId string) (*Currency, error) {
	var currency Currency

	_, err := bxgo.OrmEngin.Where("currency_id=? ", currencyId).Get(&currency)

	return &currency, err
}

func GetAllCurrency() ([]*Currency, error) {
	var currencyList []*Currency

	err := bxgo.OrmEngin.Find(&currencyList)

	return currencyList, err
}
