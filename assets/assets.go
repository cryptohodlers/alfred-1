package assets

import (
	"fmt"
	"strings"

	"github.com/stellar/go/build"
	"github.com/stellar/go/clients/horizon"
)

type Asset struct {
	BuilderAsset build.Asset
	Domain       string
	Instructions string
	Type         string
}

var (
	Assets = []Asset{
		{
			BuilderAsset: build.NativeAsset(),
			Domain:       "stellar.org",
		},
		{
			BuilderAsset: build.CreditAsset(
				"CENTUS",
				"GAKMVPHBET4T7DPN32ODVSI4AA3YEZX2GHGNNSBGFNRQ6QEVKFO4MNDZ",
			),
			Domain: "centus.one",
			Type:   "token",
		},
		{
			BuilderAsset: build.CreditAsset(
				"SLT",
				"GCKA6K5PCQ6PNF5RQBF7PQDJWRHO6UOGFMRLK3DYHDOI244V47XKQ4GP",
			),
			Domain:       "smartlands.io",
			Instructions: "https://smartlands.io",
			Type:         "token",
		},
		{
			BuilderAsset: build.CreditAsset(
				"CNY",
				"GAREELUB43IRHWEASCFBLKHURCGMHE5IF6XSE7EXDLACYHGRHM43RFOX",
			),
			Domain:       "ripplefox.com",
			Instructions: "Leave your address in the message to seller when you order theIte:https://shop109149722.taobao.com",
		},
		{
			BuilderAsset: build.CreditAsset(
				"USD",
				"GDUKMGUGDZQK6YHYA5Z6AY2G4XDSZPSZ3SW5UN3ARVMO6QSRDWP5YLEX",
			),
			Domain:       "stablecoin.anchorusd.com",
			Instructions: "stablecoin.anchorusd.com",
			Type:         "token",
		},
		{
			BuilderAsset: build.CreditAsset(
				"KIN",
				"GBDEVU63Y6NTHJQQZIKVTC23NWLQVP3WJ2RI2OTSJTNYOIGICST6DUXR",
			),
			Domain:       "apay.io",
			Instructions: "https://apay.io/",
		},
		{
			BuilderAsset: build.CreditAsset(
				"USDT",
				"GCQTGZQQ5G4PTM2GL7CDIFKUBIPEC52BROAQIAPW53XBRJVN6ZJVTG6V",
			),
			Domain:       "apay.io",
			Instructions: "apay.io",
		},
	}
)

// indexes
var (
	CodeToAsset = map[string][]Asset{}
)

func init() {
	for _, a := range Assets {
		code := a.BuilderAsset.Code
		if a.BuilderAsset.Native {
			code = "XLM"
		}
		CodeToAsset[code] = append(CodeToAsset[code], a)
	}
}

func GetAssets(code string) []Asset {
	return CodeToAsset[strings.ToUpper(code)]
}

func GetByCodeIssuer(code, issuer string) *Asset {
	for _, a := range CodeToAsset[code] {
		if a.BuilderAsset.Issuer == issuer {
			return &a
		}
	}

	return nil
}

func (a Asset) ToHorizonAsset() horizon.Asset {
	if a.BuilderAsset.Native {
		return horizon.Asset{
			Type: "native",
		}
	}

	typ := "credit_alphanum4"
	if len(a.BuilderAsset.Code) > 4 {
		typ = "credit_alphanum12"
	}

	return horizon.Asset{
		Type:   typ,
		Code:   a.BuilderAsset.Code,
		Issuer: a.BuilderAsset.Issuer,
	}
}

func (a Asset) String() string {
	if a.BuilderAsset.Native {
		return "XLM stellar.org"
	}

	return fmt.Sprintf("%s %s (%s)",
		a.BuilderAsset.Code, a.Domain,
		a.BuilderAsset.Issuer)
}

func (a Asset) CodeString() string {
	if a.BuilderAsset.Native {
		return "XLM"
	}

	return a.BuilderAsset.Code
}
