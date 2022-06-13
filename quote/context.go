package quote

import (
	"context"
	"time"

	"github.com/longbridgeapp/openapi-go/config"
	"github.com/longbridgeapp/openapi-go"
)

const DateLayout = "20160102"

type QuoteContext struct {
	opts Options
	core   *Core
}

func (c *QuoteContext) Subscribe(ctx context.Context, symbols []string, subFlags []SubFlag, isFirstPush bool) (err error) {
	return c.core.Subscribe(ctx, symbols, subFlags, isFirstPush)
}

func (c *QuoteContext) Unsubscribe(ctx context.Context, unSubAll bool, symbols []string) (err error) {
	return c.core.Unsubscribe(ctx, unSubAll, symbols)
}

func (c *QuoteContext) Subscriptions(ctx context.Context) (subscriptions map[string][]SubFlag, err error) {
	return c.core.Subscriptions(ctx)
}

func (c *QuoteContext) StaticInfo(ctx context.Context, symbols []string) (staticInfos []*SecurityStaticInfo, err error) {
	return c.core.StaticInfo(ctx, symbols)
}

func (c *QuoteContext) Quote(ctx context.Context, symbols []string) (quotes []*SecurityQuote, err error) {
	return c.core.Quote(ctx, symbols)
}

func (c *QuoteContext) OptionQuote(ctx context.Context, symbols []string) (optionQuotes []*OptionQuote, err error) {
	return c.core.OptionQuote(ctx, symbols)
}

func (c *QuoteContext) WarrantQuote(ctx context.Context, symbols []string) (warrantQuotes []*WarrantQuote, err error) {
	return c.core.WarrantQuote(ctx, symbols)
}

func (c *QuoteContext) Depth(ctx context.Context, symbol string) (securityDepth *SecurityDepth, err error) {
	return c.core.Depth(ctx, symbol)
}

func (c *QuoteContext) Brokers(ctx context.Context, symbol string) (securityBrokers *SecurityBrokers, err error) {
	return c.core.Brokers(ctx, symbol)
}

func (c *QuoteContext) Participants(ctx context.Context) (infos []*ParticipantInfo, err error) {
	return c.core.Participants(ctx)
}

func (c *QuoteContext) Trades(ctx context.Context, symbol string, count int32) (trades []*Trade, err error){
	return c.core.Trades(ctx, symbol, count)
}

func (c *QuoteContext) Intraday(ctx context.Context, symbol string) (lines []*IntradayLine, err error){
	return c.core.Intraday(ctx, symbol)
}


func (c *QuoteContext) Candlesticks(ctx context.Context, symbol string, period Period, count int32, adjustType AdjustType) (sticks []*Candlestick, err error){
	return c.core.Candlesticks(ctx, symbol, period, count, adjustType)
}


func (c *QuoteContext) OptionChainExpiryDateList(ctx context.Context, symbol string) (times []*time.Time, err error){
	return c.core.OptionChainExpiryDateList(ctx, symbol)
}

func (c *QuoteContext) OptionChainInfoByDate(ctx context.Context, symbol string, expiryDate *time.Time) (priceInfos []*StrikePriceInfo, err error){
	return c.core.OptionChainInfoByDate(ctx, symbol, expiryDate)
}

func (c *QuoteContext) WarrantIssuers(ctx context.Context) (infos []*IssuerInfo, err error){
	return c.core.WarrantIssuers(ctx)
}

func (c *QuoteContext) TradingSession(ctx context.Context) (sessions []*MarketTradingSession, err error){
	return c.core.TradingSession(ctx)
}

func (c *QuoteContext) TradingDays(ctx context.Context, market openapi.Market, begin *time.Time, end *time.Time) (days *MarketTradingDay, err error){
	return c.core.TradingDays(ctx, market, begin, end)
}

func (c *QuoteContext) RealtimeQuote(ctx context.Context, symbols []string) (RealtimeQuote, error) {
	return nil, nil
}

func (c *QuoteContext) RealtimeDepth(ctx context.Context, symbol string) error {
	return nil
}

func (c *QuoteContext) RealtimeTrades(ctx context.Context, symbol string, count int) error {
	return nil
}

func (c *QuoteContext) RealtimeBrokers(ctx context.Context, symbol string) error {
	return nil
}


func NewFromCfg(config.Config) *QuoteContext {
	return nil
}

func New(opts ...Option) *QuoteContext {
	return nil
}