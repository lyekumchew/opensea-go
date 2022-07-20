package model

import (
	"github.com/shopspring/decimal"
)

type Fee struct {
	BasisPoints string     `opensea:"basis_points" json:"basisPoints"`
	Account     *AccountV2 `opensea:"account" json:"account"`
}

type AccountV2 struct {
	User          int    `opensea:"user" json:"user"`
	ProfileImgURL string `opensea:"profile_img_url" json:"profileImgURL"`
	Address       string `opensea:"address" json:"address"`
	Config        string `opensea:"config" json:"config"`
}

type ProtocolData struct {
	Parameters *Parameters `opensea:"parameters" json:"parameters"`
	Signature  string      `opensea:"signature" json:"signature"`
}

type Parameters struct {
	Offerer                         string           `json:"offerer"`
	Offer                           []*Offer         `json:"offer"`
	Consideration                   []*Consideration `json:"consideration"`
	StartTime                       string           `json:"startTime"`
	EndTime                         string           `json:"endTime"`
	OrderType                       int              `json:"orderType"`
	Zone                            string           `json:"zone"`
	ZoneHash                        string           `json:"zoneHash"`
	Salt                            string           `json:"salt"`
	ConduitKey                      string           `json:"conduitKey"`
	TotalOriginalConsiderationItems int              `json:"totalOriginalConsiderationItems"`
	Counter                         int              `json:"counter"`
}

type Offer struct {
	ItemType             int    `json:"itemType"`
	Token                string `json:"token"`
	IdentifierOrCriteria string `json:"identifierOrCriteria"`
	StartAmount          string `json:"startAmount"`
	EndAmount            string `json:"endAmount"`
}

type Consideration struct {
	ItemType             int    `json:"itemType"`
	Token                string `json:"token"`
	IdentifierOrCriteria string `json:"identifierOrCriteria"`
	StartAmount          string `json:"startAmount"`
	EndAmount            string `json:"endAmount"`
	Recipient            string `json:"recipient"`
}

type SeaportOrder struct {
	CreatedDate     string           `opensea:"created_date" json:"createdDate"`
	ClosingDate     string           `opensea:"closing_date" json:"closingDate"`
	ListingTime     int              `opensea:"listing_time" json:"listingTime"`
	ExpirationTime  int              `opensea:"expiration_time" json:"expirationTime"`
	OrderHash       string           `opensea:"order_hash" json:"orderHash"`
	ProtocolData    *ProtocolData    `opensea:"protocol_data" json:"protocolData"`
	ProtocolAddress string           `opensea:"protocol_address" json:"protocolAddress"`
	Maker           *AccountV2       `opensea:"maker" json:"maker"`
	Taker           *AccountV2       `opensea:"taker" json:"taker"`
	CurrentPrice    *decimal.Decimal `opensea:"current_price" json:"currentPrice"`
	MakerFees       []*Fee           `opensea:"maker_fees" json:"makerFees"`
	TakerFees       []*Fee           `opensea:"taker_fees" json:"takerFees"`
	Side            string           `opensea:"side" json:"side"`
	OrderType       string           `opensea:"order_type" json:"orderType"`
	Canceled        bool             `opensea:"canceled" json:"canceled"`
	Finalized       bool             `opensea:"finalized" json:"finalized"`
	MarkedInvalid   bool             `opensea:"marked_invalid" json:"markedInvalid"`
	ClientSignature string           `opensea:"client_signature" json:"clientSignature"`
}
