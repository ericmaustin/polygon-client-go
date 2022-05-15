package models

// ListTickersParams is the set of parameters for the ListTickers method.
type ListTickersParams struct {
	// Specify a ticker symbol. Defaults to empty string which queries all tickers.
	TickerEQ  *string `query:"ticker"`
	TickerLT  *string `query:"ticker.lt"`
	TickerLTE *string `query:"ticker.lte"`
	TickerGT  *string `query:"ticker.gt"`
	TickerGTE *string `query:"ticker.gte"`

	// Specify the type of the tickers. Find the types that we support via our Ticker Types API. Defaults to empty
	// string which queries all types.
	Type *string `query:"type"`

	// Filter by market type. By default all markets are included.
	Market *AssetClass `query:"market"`

	// Specify the primary exchange of the asset in the ISO code format. Find more information about the ISO codes at
	// the ISO org website. Defaults to empty string which queries all exchanges.
	Exchange *int `query:"exchange"`

	// Specify the CUSIP code of the asset you want to search for. Find more information about CUSIP codes at their
	// website. Defaults to empty string which queries all CUSIPs.
	//
	// Note: Although you can query by CUSIP, due to legal reasons we do not return the CUSIP in the response.
	CUSIP *int `query:"cusip"`

	// Specify the CIK of the asset you want to search for. Find more information about CIK codes at their website.
	// Defaults to empty string which queries all CIKs.
	CIK *int `query:"cik"`

	// Specify a point in time to retrieve tickers available on that date. Defaults to the most recent available date.
	Date *Date `query:"date"`

	// Specify if the tickers returned should be actively traded on the queried date. Default is true.
	Active *bool `query:"active"`

	// Search for terms within the ticker and/or company name.
	Search *string `query:"search"`

	// The field to sort the results on. Default is ticker. If the search query parameter is present, sort is ignored
	// and results are ordered by relevance.
	Sort *Sort `query:"sort"`

	// The order to sort the results on. Default is asc (ascending).
	Order *Order `query:"order"`

	// Limit the size of the response, default is 100 and max is 1000.
	Limit *int `query:"limit"`
}

func (p ListTickersParams) WithTicker(c Comparator, q string) *ListTickersParams {
	if c == EQ {
		p.TickerEQ = &q
	} else if c == LT {
		p.TickerLT = &q
	} else if c == LTE {
		p.TickerLTE = &q
	} else if c == GT {
		p.TickerGT = &q
	} else if c == GTE {
		p.TickerGTE = &q
	}
	return &p
}

func (p ListTickersParams) WithType(q string) *ListTickersParams {
	p.Type = &q
	return &p
}

func (p ListTickersParams) WithMarket(q AssetClass) *ListTickersParams {
	p.Market = &q
	return &p
}

func (p ListTickersParams) WithExchange(q int) *ListTickersParams {
	p.Exchange = &q
	return &p
}

func (p ListTickersParams) WithCUSIP(q int) *ListTickersParams {
	p.CUSIP = &q
	return &p
}

func (p ListTickersParams) WithCIK(q int) *ListTickersParams {
	p.CIK = &q
	return &p
}

func (p ListTickersParams) WithDate(q Date) *ListTickersParams {
	p.Date = &q
	return &p
}

func (p ListTickersParams) WithActive(q bool) *ListTickersParams {
	p.Active = &q
	return &p
}

func (p ListTickersParams) WithSearch(q string) *ListTickersParams {
	p.Search = &q
	return &p
}

func (p ListTickersParams) WithSort(q Sort) *ListTickersParams {
	p.Sort = &q
	return &p
}

func (p ListTickersParams) WithOrder(q Order) *ListTickersParams {
	p.Order = &q
	return &p
}

func (p ListTickersParams) WithLimit(q int) *ListTickersParams {
	p.Limit = &q
	return &p
}

// ListTickersResponse is the response returned by the ListTickers method.
type ListTickersResponse struct {
	BaseResponse

	// An array of tickers that match your query. Note: Although you can query by CUSIP, due to legal reasons we do not
	// return the CUSIP in the response.
	Results []Ticker `json:"results,omitempty"`
}

// GetTickerDetailsParams is the set of parameters for the GetTickerDetails method.
type GetTickerDetailsParams struct {
	// The ticker symbol of the asset.
	Ticker string `validate:"required" path:"ticker"`

	// Specify a point in time to get information about the ticker available on that date. When retrieving information
	// from SEC filings, we compare this date with the period of report date on the SEC filing.
	//
	// For example, consider an SEC filing submitted by AAPL on 2019-07-31, with a period of report date ending on
	// 2019-06-29. That means that the filing was submitted on 2019-07-31, but the filing was created based on
	// information from 2019-06-29. If you were to query for AAPL details on 2019-06-29, the ticker details would
	// include information from the SEC filing.
	//
	// Defaults to the most recent available date.
	Date *Date `query:"date"`
}

func (p GetTickerDetailsParams) WithDate(q Date) *GetTickerDetailsParams {
	p.Date = &q
	return &p
}

// GetTickerDetailsResponse is the response returned by the GetTickerDetails method.
type GetTickerDetailsResponse struct {
	BaseResponse

	// Ticker with details.
	Results Ticker `json:"results,omitempty"`
}

// ListTickerNewsParams is the set of parameters for the ListTickerNews method.
type ListTickerNewsParams struct {
	// Return results that contain this ticker.
	TickerEQ  *string `query:"ticker"`
	TickerLT  *string `query:"ticker.lt"`
	TickerLTE *string `query:"ticker.lte"`
	TickerGT  *string `query:"ticker.gt"`
	TickerGTE *string `query:"ticker.gte"`

	// Return results published on, before, or after this date.
	PublishedUtcEQ  *Millis `query:"published_utc"`
	PublishedUtcLT  *Millis `query:"published_utc.lt"`
	PublishedUtcLTE *Millis `query:"published_utc.lte"`
	PublishedUtcGT  *Millis `query:"published_utc.gt"`
	PublishedUtcGTE *Millis `query:"published_utc.gte"`

	// Sort field used for ordering.
	Sort *Sort `query:"sort"`

	// Order results based on the sort field.
	Order *Order `query:"order"`

	// Limit the number of results returned, default is 10 and max is 1000.
	Limit *int `query:"limit"`
}

func (p ListTickerNewsParams) WithTicker(c Comparator, q string) *ListTickerNewsParams {
	if c == EQ {
		p.TickerEQ = &q
	} else if c == LT {
		p.TickerLT = &q
	} else if c == LTE {
		p.TickerLTE = &q
	} else if c == GT {
		p.TickerGT = &q
	} else if c == GTE {
		p.TickerGTE = &q
	}
	return &p
}

func (p ListTickerNewsParams) WithPublishedUTC(c Comparator, q Millis) *ListTickerNewsParams {
	if c == EQ {
		p.PublishedUtcEQ = &q
	} else if c == LT {
		p.PublishedUtcLT = &q
	} else if c == LTE {
		p.PublishedUtcLTE = &q
	} else if c == GT {
		p.PublishedUtcGT = &q
	} else if c == GTE {
		p.PublishedUtcGTE = &q
	}
	return &p
}

func (p ListTickerNewsParams) WithSort(q Sort) *ListTickerNewsParams {
	p.Sort = &q
	return &p
}

func (p ListTickerNewsParams) WithOrder(q Order) *ListTickerNewsParams {
	p.Order = &q
	return &p
}

func (p ListTickerNewsParams) WithLimit(q int) *ListTickerNewsParams {
	p.Limit = &q
	return &p
}

// ListTickerNewsResponse is the response returned by the ListTickerNews method.
type ListTickerNewsResponse struct {
	BaseResponse

	// Ticker news results.
	Results []TickerNews `json:"results,omitempty"`
}

// GetTickerTypesParams is the set of parameters for the GetTickerTypes method.
type GetTickerTypesParams struct {
	// Filter by asset class.
	AssetClass *AssetClass `query:"asset_class"`

	// Filter by locale.
	Locale *MarketLocale `query:"locale"`
}

func (p GetTickerTypesParams) WithAssetClass(q AssetClass) *GetTickerTypesParams {
	p.AssetClass = &q
	return &p
}

func (p GetTickerTypesParams) WithLocale(q MarketLocale) *GetTickerTypesParams {
	p.Locale = &q
	return &p
}

// GetTickerTypesResponse is the response returned by the GetTickerTypes method.
type GetTickerTypesResponse struct {
	BaseResponse

	// Ticker type results.
	Results []TickerType `json:"results,omitempty"`
}

// Ticker contains detailed information on a specified ticker symbol.
type Ticker struct {
	Active                      bool           `json:"active"`
	Address                     CompanyAddress `json:"address,omitempty"`
	Branding                    Branding       `json:"branding,omitempty"`
	CIK                         string         `json:"cik,omitempty"`
	CompositeFIGI               string         `json:"composite_figi,omitempty"`
	CurrencyName                string         `json:"currency_name,omitempty"`
	DelistedUTC                 Time           `json:"delisted_utc,omitempty"`
	Description                 string         `json:"description,omitempty"`
	HomepageURL                 string         `json:"homepage_url,omitempty"`
	LastUpdatedUTC              Time           `json:"last_updated_utc,omitempty"`
	ListDate                    Date           `json:"list_date,omitempty"`
	Locale                      string         `json:"locale,omitempty"`
	Market                      string         `json:"market,omitempty"`
	MarketCap                   float64        `json:"market_cap,omitempty"`
	Name                        string         `json:"name,omitempty"`
	PhoneNumber                 string         `json:"phone_number,omitempty"`
	PrimaryExchange             string         `json:"primary_exchange,omitempty"`
	ShareClassFIGI              string         `json:"share_class_figi,omitempty"`
	ShareClassSharesOutstanding int64          `json:"share_class_shares_outstanding,omitempty"`
	SICCode                     string         `json:"sic_code,omitempty"`
	SICDescription              string         `json:"sic_description,omitempty"`
	Ticker                      string         `json:"ticker,omitempty"`
	TickerRoot                  string         `json:"ticker_root,omitempty"`
	TickerSuffix                string         `json:"ticker_suffix,omitempty"`
	TotalEmployees              int32          `json:"total_employees,omitempty"`
	Type                        string         `json:"type,omitempty"`
	WeightedSharesOutstanding   int64          `json:"weighted_shares_outstanding,omitempty"`
}

// CompanyAddress contains information on the physical address of a company.
type CompanyAddress struct {
	Address1   string `json:"address1,omitempty"`
	Address2   string `json:"address2,omitempty"` // todo: add this to the spec
	City       string `json:"city,omitempty"`
	PostalCode string `json:"postal_code,omitempty"`
	State      string `json:"state,omitempty"`
}

// Branding contains information related to a company's brand.
type Branding struct {
	LogoURL string `json:"logo_url,omitempty"`
	IconURL string `json:"icon_url,omitempty"`
}

// TickerNews contains information on a ticker news article.
type TickerNews struct {
	AMPURL       string    `json:"amp_url,omitempty"`
	ArticleURL   string    `json:"article_url,omitempty"`
	Author       string    `json:"author,omitempty"`
	Description  string    `json:"description,omitempty"`
	ID           string    `json:"id,omitempty"`
	ImageURL     string    `json:"image_url,omitempty"`
	Keywords     []string  `json:"keywords,omitempty"`
	PublishedUTC Time      `json:"published_utc,omitempty"`
	Publisher    Publisher `json:"publisher,omitempty"`
	Tickers      []string  `json:"tickers,omitempty"`
	Title        string    `json:"title,omitempty"`
}

// Publisher contains information on a new article publisher.
type Publisher struct {
	FaviconURL  string `json:"favicon_url,omitempty"`
	HomepageURL string `json:"homepage_url,omitempty"`
	LogoURL     string `json:"logo_url,omitempty"`
	Name        string `json:"name,omitempty"`
}

// TickerType represents a type of ticker with a code that the API understands.
type TickerType struct {
	AssetClass  string `json:"asset_class,omitempty"`
	Code        string `json:"code,omitempty"`
	Description string `json:"description,omitempty"`
	Locale      string `json:"locale,omitempty"`
}

// ListOptionContractsParams is the set of parameters for the ListContracts method.
type ListOptionContractsParams struct {
	TickerEQ  *string `query:"underlying_ticker"`
	TickerLT  *string `query:"underlying_ticker.lt"`
	TickerLTE *string `query:"underlying_ticker.lte"`
	TickerGT  *string `query:"underlying_ticker.gt"`
	TickerGTE *string `query:"underlying_ticker.gte"`

	ContractType *string `query:"contract_type"`

	ExpirationDateEQ  *Date `query:"expiration_date"`
	ExpirationDateLT  *Date `query:"expiration_date.lt"`
	ExpirationDateLTE *Date `query:"expiration_date.lte"`
	ExpirationDateGT  *Date `query:"expiration_date.gt"`
	ExpirationDateGTE *Date `query:"expiration_date.gte"`

	StrikePriceEQ  *float64 `query:"strike_price"`
	StrikePriceLT  *float64 `query:"strike_price.lt"`
	StrikePriceLTE *float64 `query:"strike_price.lte"`
	StrikePriceGT  *float64 `query:"strike_price.gt"`
	StrikePriceGTE *float64 `query:"strike_price.gte"`

	AsOf    *Date `query:"as_of"`
	Expired *bool `query:"expired"`
}

// WithUnderlyingTicker add underlying ticker filter
func (p ListOptionContractsParams) WithUnderlyingTicker(c Comparator, q string) *ListOptionContractsParams {
	if c == EQ {
		p.TickerEQ = &q
	} else if c == LT {
		p.TickerLT = &q
	} else if c == LTE {
		p.TickerLTE = &q
	} else if c == GT {
		p.TickerGT = &q
	} else if c == GTE {
		p.TickerGTE = &q
	}
	return &p
}

// WithContractType add contract type filter
func (p ListOptionContractsParams) WithContractType(q string) *ListOptionContractsParams {
	p.ContractType = &q
	return &p
}

// WithExpirationDate add expiration date filter
func (p ListOptionContractsParams) WithExpirationDate(c Comparator, q Date) *ListOptionContractsParams {
	if c == EQ {
		p.ExpirationDateEQ = &q
	} else if c == LT {
		p.ExpirationDateLT = &q
	} else if c == LTE {
		p.ExpirationDateLTE = &q
	} else if c == GT {
		p.ExpirationDateGT = &q
	} else if c == GTE {
		p.ExpirationDateGTE = &q
	}
	return &p
}

// WithStrikePrice add a strike price filter
func (p ListOptionContractsParams) WithStrikePrice(c Comparator, q float64) *ListOptionContractsParams {
	if c == EQ {
		p.StrikePriceEQ = &q
	} else if c == LT {
		p.StrikePriceLT = &q
	} else if c == LTE {
		p.StrikePriceLTE = &q
	} else if c == GT {
		p.StrikePriceGT = &q
	} else if c == GTE {
		p.StrikePriceGTE = &q
	}
	return &p
}

// WithAsOf add as of filter
func (p ListOptionContractsParams) WithAsOf(q Date) *ListOptionContractsParams {
	p.AsOf = &q
	return &p
}

// WithExpired add expiration filter
func (p ListOptionContractsParams) WithExpired(q bool) *ListOptionContractsParams {
	p.Expired = &q
	return &p
}

type OptionContractAdditionalUnderlying struct {
	Amount     float64 `json:"amount,omitempty"`
	Type       string  `json:"type,omitempty"`
	Underlying string  `json:"underlying,omitempty"`
}

// OptionContract contains an options contract details
type OptionContract struct {
	CFI                   string                               `json:"cfi,omitempty"`
	ContractType          string                               `json:"contract_type,omitempty"`
	ExerciseStyle         string                               `json:"exercise_style,omitempty"`
	ExpirationDate        Date                                 `json:"expiration_date,omitempty"`
	PrimaryExchange       string                               `json:"primary_exchange,omitempty"`
	SharesPerContract     int                                  `json:"shares_per_contract,omitempty"`
	StrikePrice           float64                              `json:"strike_price,omitempty"`
	Ticker                string                               `json:"ticker,omitempty"`
	UnderLyingTicker      string                               `json:"underlying_ticker,omitempty"`
	AdditionalUnderlyings []OptionContractAdditionalUnderlying `json:"additional_underlyings,omitempty"`
}

// ListOptionContractsResponse contains the result of the ListOptionsContracts API call
type ListOptionContractsResponse struct {
	BaseResponse
	Results []OptionContract `json:"results,omitempty"`
}

// GetOptionContractDetailsParams is the set of parameters for the GetOptionContractDetails method.
type GetOptionContractDetailsParams struct {
	// The OptionTicker symbol of the asset.
	OptionTicker string `validate:"required" path:"options_ticker"`

	// Specify a point in time for the contract as of this date with format YYYY-MM-DD. Defaults to today's date.
	AsOf *Date `query:"as_of"`
}

// WithAsOf sets the as_of parameter to the request
func (p GetOptionContractDetailsParams) WithAsOf(q Date) *GetOptionContractDetailsParams {
	p.AsOf = &q
	return &p
}

// GetOptionContractDetailsResponse contains the result of the GetOptionContractDetails API call
type GetOptionContractDetailsResponse struct {
	BaseResponse
	Results OptionContract `json:"results,omitempty"`
}
