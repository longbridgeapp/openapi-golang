package quote

import (
	"time"

	"github.com/longbridgeapp/openapi-go"
	"github.com/longbridgeapp/openapi-protobufs/gen/go/quote"
)

type TradeStatus int32
type TradeSession int32
type TradeSessionType int32

type PushEvent struct {
	Quote   *PushQuote
	Depth   *PushDepth
	Brokers *PushBrokers
	Trade   *PushTrades
}

type PushQuote struct {
	Symbol       string
	Sequence     int64
	LastDone     float64
	Open         float64
	High         float64
	Low          float64
	Timestamp    int64
	Volume       int64
	Turnover     float64
	TradeStatus  TradeStatus
	TradeSession TradeSessionType
}

type PushDepth struct {
	Symbol   string
	Sequence int64
	Ask      []*Depth
	Bid      []*Depth
}

type Depth struct {
	Position int32
	Price    string
	Volume   int64
	OrderNum int64
}

func toDepth(origin *quotev1.Depth) *Depth {
	return &Depth{
		Position: origin.GetPosition(),
		Price:    origin.GetPrice(),
		Volume:   origin.GetVolume(),
		OrderNum: origin.GetOrderNum(),
	}
}

func toDepths(origin []*quotev1.Depth) (depths []*Depth) {
	depths = make([]*Depth, 0, len(origin))
	for _, item := range origin {
		depths = append(depths, toDepth(item))
	}
	return
}

type PushBrokers struct {
	Symbol     string
	Sequence   int64
	AskBrokers []*Brokers
	BidBrokers []*Brokers
}

type Brokers struct {
	Position  int32
	BrokerIds []int32
}

func toBrokers(origin []*quotev1.Brokers) (brokers []*Brokers) {
	brokers = make([]*Brokers, len(origin))
	for _, item := range origin {
		brokers = append(brokers, &Brokers{
			Position:  item.GetPosition(),
			BrokerIds: item.GetBrokerIds(),
		})
	}
	return
}

type PushTrades struct {
	Symbol   string
	Sequence int64
	Trade    []*Trade
}

type Trade struct {
	Price        string
	Volume       int64
	Timestamp    int64
	TradeType    string
	Direction    int32
	TradeSession TradeSession
}

func toTrades(origin []*quotev1.Trade) (trades []*Trade) {
	trades = make([]*Trade, 0, len(origin))
	for _, item := range origin {
		trades = append(trades, &Trade{
			Price:        item.GetPrice(),
			Volume:       item.GetVolume(),
			Timestamp:    item.GetTimestamp(),
			TradeType:    item.GetTradeType(),
			Direction:    item.GetDirection(),
			TradeSession: TradeSession(item.GetTradeSession()),
		})
	}
	return
}

type SecurityStaticInfo struct {
	Symbol            string
	NameCn            string
	NameEn            string
	NameHk            string
	Exchange          string
	Currency          string
	LotSize           int32
	TotalShares       int64
	CirculatingShares int64
	HkShares          int64
	Eps               string
	EpsTtm            string
	Bps               string
	DividendYield     string
	StockDerivatives  []int32
}

type Issuer struct {
	ID     int32
	NameCn string
	NameEn string
	NameHk string
}

type RealtimeOptionQuote struct {
	Symbol       string
	LastDone     float64
	PrevClose    float64
	Open         float64
	High         float64
	Low          float64
	Timestamp    int64
	Volume       int64
	Turnover     float64
	OptionExtend *OptionExtend
}

type OptionQuote struct {
	Symbol       string
	LastDone     string
	PrevClose    string
	Open         string
	High         string
	Low          string
	Timestamp    int64
	Volume       int64
	Turnover     string
	TradeStatus  TradeStatus
	OptionExtend *OptionExtend
}

type OptionExtend struct {
	ImpliedVolatility    string
	OpenInterest         int64
	ExpiryDate           string // YYMMDD
	StrikePrice          string
	ContractMultiplier   string
	ContractType         string
	ContractSize         string
	Direction            string
	HistoricalVolatility string
	UnderlyingSymbol     string
}

type StrikePriceInfo struct {
	Price      string
	CallSymbol string
	PutSymbol  string
	Standard   bool
}

func toStrikePriceInfos(origin []*quotev1.StrikePriceInfo) (priceInfos []*StrikePriceInfo) {
	priceInfos = make([]*StrikePriceInfo, 0, len(origin))
	// TODO use copier
	for _, item := range origin {
		priceInfos = append(priceInfos, &StrikePriceInfo{
			Price: item.GetPrice(),
			CallSymbol: item.GetCallSymbol(),
			PutSymbol: item.GetPutSymbol(),
			Standard: item.GetStandard(),
		})
	}
	return priceInfos
}

type WarrantExtend struct {
	ImpliedVolatility string
	ExpiryDate        string
	LastTradeDate     string
	OutstandingRatio  string
	OutstandingQty    int64
	ConversionRatio   string
	Category          string
	StrikePrice       string
	UpperStrikePrice  string
	LowerStrikePrice  string
	CallPrice         string
	UnderlyingSymbol  string
}

type RealtimeWarrantQuote struct {
	Symbol        string
	LastDone      float64
	PrevClose     float64
	Open          float64
	High          float64
	Low           float64
	Timestamp     int64
	Volume        int64
	Turnover      float64
	WarrantExtend *WarrantExtend
}

func toWarrantExtend(origin *quotev1.WarrantExtend) *WarrantExtend {
	return &WarrantExtend{
		ImpliedVolatility: origin.GetImpliedVolatility(),
		ExpiryDate:        origin.GetExpiryDate(),
		LastTradeDate:     origin.GetLastTradeDate(),
		OutstandingRatio:  origin.GetOutstandingRatio(),
		OutstandingQty:    origin.GetOutstandingQty(),
		ConversionRatio:   origin.GetConversionRatio(),
		Category:          origin.GetCategory(),
		StrikePrice:       origin.GetStrikePrice(),
		UpperStrikePrice:  origin.GetUpperStrikePrice(),
		LowerStrikePrice:  origin.GetLowerStrikePrice(),
		CallPrice:         origin.GetCallPrice(),
		UnderlyingSymbol:  origin.GetUnderlyingSymbol(),
	}
}

type WarrantQuote struct {
	Symbol        string
	LastDone      string
	PrevClose     string
	Open          string
	High          string
	Low           string
	Timestamp     int64
	Volume        int64
	Turnover      string
	TradeStatus   TradeStatus
	WarrantExtend *WarrantExtend
}

func toWarrantQuotes(origin []*quotev1.WarrantQuote) (warrantQuotes []*WarrantQuote) {
	warrantQuotes = make([]*WarrantQuote, 0, len(origin))
	// TODO use copier
	for _, item := range origin {
		warrantQuotes = append(warrantQuotes, &WarrantQuote{
			Symbol:        item.GetSymbol(),
			LastDone:      item.GetLastDone(),
			PrevClose:     item.GetPrevClose(),
			Open:          item.GetOpen(),
			High:          item.GetHigh(),
			Low:           item.GetLow(),
			Timestamp:     item.GetTimestamp(),
			Volume:        item.GetVolume(),
			Turnover:      item.GetTurnover(),
			TradeStatus:   TradeStatus(item.GetTradeStatus()),
			WarrantExtend: toWarrantExtend(item.GetWarrantExtend()),
		})
	}
	return
}

type WarrantFilter struct {
	Symbol   string
	Language string

	SortBy     int32
	SortOrder  int32 // 0 Ascending 1 Desending
	SortOffset int32
	PageSize   int32 // Up to 500

	// The following are optional

	Type      []int32 // optional values: 0 - Call	1 - Put 2 - Bull 3 - Bear 4 - Inline
	IssuerIDs []int32

	// ExpiryDateType can have the following values.
	// 1 - Less than 3 months
	// 2 - 3 - 6 months
	// 3 - 6 - 12 months
	// 4 - greater than 12 months
	ExpiryDateType []int32

	// Optional values for PriceType
	// 1 - In bounds
	// 2 - Out bounds
	PriceType []int32

	// Optional values for Status:
	// 2 - Suspend trading
	// 3 - Papare List
	// 4 - Normal
	Status []int32
}

type Warrant struct {
	Symbol            string
	Name              string
	LastDone          float64
	ChangeRate        float64
	ChangeVal         float64
	Turnover          float64
	ExpiryDate        string // YYYYMMDD
	StrikePrice       float64
	UpperStrikePrice  float64
	LowerStrikePrice  float64
	OutstandingQty    float64
	OutstandingRatio  float64
	Premium           float64
	ItmOtm            float64
	ImpliedVolatility float64
	Delta             float64
	CallPrice         float64
	EffectiveLeverage float64
	LeverageRatio     float64
	ConversionRatio   float64
	BalancePoint      float64
	State             string
}

type TradeDate struct {
	Date          string
	TradeDateType int32 // 0 full day, 1 morning only, 2 afternoon only(not happened before)
}

type SubscriptionType = quotev1.SubType

const (
	SubscriptionRealtimeQuote = SubscriptionType(quotev1.SubType_QUOTE)
	SubscriptionOrderBook     = SubscriptionType(quotev1.SubType_DEPTH)
	SubscriptionBrokerQueue   = SubscriptionType(quotev1.SubType_BROKERS)
	SubscriptionTicker        = SubscriptionType(quotev1.SubType_TRADE)
)

type QotSubscription struct {
	Symbol        string
	Subscriptions []SubscriptionType
}

type Period int32

const (
	Period_ONE_MINUTE     = Period(quotev1.Period_ONE_MINUTE)
	Period_FIVE_MINUTE    = Period(quotev1.Period_FIVE_MINUTE)
	Period_FIFTEEN_MINUTE = Period(quotev1.Period_FIFTEEN_MINUTE)
	Period_THIRTY_MINUTE  = Period(quotev1.Period_THIRTY_MINUTE)
	Period_SIXTY_MINUTE   = Period(quotev1.Period_SIXTY_MINUTE)
	Period_DAY            = Period(quotev1.Period_DAY)
	Period_WEEK           = Period(quotev1.Period_WEEK)
	Period_MONTH          = Period(quotev1.Period_MONTH)
	Period_YEAR           = Period(quotev1.Period_YEAR)
)

type AdjustType int32

const (
	AdjustType_NO_ADJUST       = AdjustType(quotev1.AdjustType_NO_ADJUST)
	AdjustType_FORWARD_ADJUST  = AdjustType(quotev1.AdjustType_FORWARD_ADJUST)
)

type Candlestick struct {
	Close     string
	Open      string
	Low       string
	High      string
	Volume    int64
	Turnover  string
	Timestamp int64
}

func toCandlesticks(origin []*quotev1.Candlestick) (sticks []*Candlestick) {
	sticks = make([]*Candlestick, 0, len(origin))
	// TODO use copier
	for _, item := range origin {
		sticks = append(sticks, &Candlestick{
			Close: item.GetClose(),
			Open: item.GetOpen(),
			Low: item.GetLow(),
			High: item.GetHigh(),
			Volume: item.GetVolume(),
			Turnover: item.GetTurnover(),
			Timestamp: item.GetTimestamp(),
		})
	}
	return
}

type SecurityQuote struct {
	Symbol          string
	LastDone        string
	PrevClose       string
	Open            string
	High            string
	Low             string
	Timestamp       int64
	Volume          int64
	Turnover        string
	TradeStatus     TradeStatus
	PreMarketQuote  *PrePostQuote
	PostMarketQuote *PrePostQuote
}

func toSecurityQuotes(origin []*quotev1.SecurityQuote) (quotes []*SecurityQuote) {
	quotes = make([]*SecurityQuote, 0, len(origin))
	// TODO use copier
	for _, item := range origin {
		quotes = append(quotes, &SecurityQuote{
			Symbol:      item.GetSymbol(),
			LastDone:    item.GetLastDone(),
			PrevClose:   item.GetPrevClose(),
			Open:        item.GetOpen(),
			High:        item.GetHigh(),
			Low:         item.GetLow(),
			Timestamp:   item.GetTimestamp(),
			Volume:      item.GetVolume(),
			Turnover:    item.GetTurnover(),
			TradeStatus: TradeStatus(item.GetTradeStatus()),
			PreMarketQuote: &PrePostQuote{
				LastDone:  item.GetPreMarketQuote().GetLastDone(),
				Timestamp: item.GetPreMarketQuote().GetTimestamp(),
				Volume:    item.GetPreMarketQuote().GetVolume(),
				Turnover:  item.GetPreMarketQuote().GetTurnover(),
				High:      item.GetPreMarketQuote().GetHigh(),
				Low:       item.GetPreMarketQuote().GetHigh(),
				PrevClose: item.GetPreMarketQuote().GetPrevClose(),
			},
			PostMarketQuote: &PrePostQuote{
				LastDone:  item.GetPostMarketQuote().GetLastDone(),
				Timestamp: item.GetPostMarketQuote().GetTimestamp(),
				Volume:    item.GetPostMarketQuote().GetVolume(),
				Turnover:  item.GetPostMarketQuote().GetTurnover(),
				High:      item.GetPostMarketQuote().GetHigh(),
				Low:       item.GetPostMarketQuote().GetHigh(),
				PrevClose: item.GetPostMarketQuote().GetPrevClose(),
			},
		})
	}
	return
}

type PrePostQuote struct {
	LastDone  string
	Timestamp int64
	Volume    int64
	Turnover  string
	High      string
	Low       string
	PrevClose string
}

type SecurityDepth struct {
	Symbol string
	Ask    []*Depth
	Bid    []*Depth
}

func toSecurityDepth(origin *quotev1.SecurityDepthResponse) *SecurityDepth {
	return &SecurityDepth{
		Symbol: origin.GetSymbol(),
		Ask:    toDepths(origin.GetAsk()),
		Bid:    toDepths(origin.GetBid()),
	}
}

type SecurityBrokers struct {
	Symbol     string
	AskBrokers []*Brokers
	BidBrokers []*Brokers
}

func toSecurityBrokers(origin *quotev1.SecurityBrokersResponse) *SecurityBrokers {
	return &SecurityBrokers{
		Symbol:     origin.GetSymbol(),
		AskBrokers: toBrokers(origin.GetAskBrokers()),
		BidBrokers: toBrokers(origin.GetBidBrokers()),
	}
}

type ParticipantInfo struct {
	BrokerIds         []int32
	ParticipantNameCn string
	ParticipantNameEn string
	ParticipantNameHk string
}

func toParticipantInfos(origin []*quotev1.ParticipantInfo) (participantInfos []*ParticipantInfo) {
	participantInfos = make([]*ParticipantInfo, 0, len(origin))
	// TODO use copier
	for _, item := range origin {
		participantInfos = append(participantInfos, &ParticipantInfo{
			BrokerIds:         item.GetBrokerIds(),
			ParticipantNameCn: item.GetParticipantNameCn(),
			ParticipantNameEn: item.GetParticipantNameEn(),
			ParticipantNameHk: item.GetParticipantNameHk(),
		})
	}
	return
}

type IntradayLine struct {
	Price     string
	Timestamp int64
	Volume    int64
	Turnover  string
	AvgPrice  string
}

func toIntradayLines(origin []*quotev1.Line) (lines []*IntradayLine) {
	lines = make([]*IntradayLine, len(origin))
	for _, item := range origin {
		lines = append(lines, &IntradayLine{
			Price:     item.GetPrice(),
			Timestamp: item.GetTimestamp(),
			Volume:    item.GetVolume(),
			AvgPrice:  item.GetAvgPrice(),
		})
	}
	return
}

type IssuerInfo struct {
	Id     int32
	NameCn string
	NameEn string
	NameHk string
}

func toIssueInfos(origin []*quotev1.IssuerInfo) (infos []*IssuerInfo) {
	infos = make([]*IssuerInfo, len(origin))
	for _, item := range origin {
		infos = append(infos, &IssuerInfo{
			Id: item.GetId(),
			NameCn: item.GetNameCn(),
			NameEn: item.GetNameEn(),
			NameHk: item.GetNameHk(),
		})
	}
	return
}

type MarketTradingSession struct {
	Market       openapi.Market
	TradeSession []*TradePeriod
}

func toMarketTradingSessions(origin []*quotev1.MarketTradePeriod) (sessions []*MarketTradingSession) {
	sessions = make([]*MarketTradingSession, len(origin))
	for _, item := range origin {
		sessions = append(sessions, &MarketTradingSession{
			Market: openapi.Market(item.GetMarket()),
			TradeSession: toTradePeriods(item.GetTradeSession()),
		})
	}
	return
}

type TradePeriod struct {
	BegTime      int32
	EndTime      int32
	TradeSession TradeSession
}

func toTradePeriods(origin []*quotev1.TradePeriod) (periods []*TradePeriod) {
	periods = make([]*TradePeriod, len(origin))
	for _, item := range origin {
		periods = append(periods, &TradePeriod{
			BegTime: item.GetBegTime(),
			EndTime: item.GetEndTime(),
			TradeSession: TradeSession(item.GetTradeSession()),
		})
	}
	return
}

type MarketTradingDay struct {
	TradeDay     []*time.Time
	HalfTradeDay []*time.Time
}