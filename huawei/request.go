package huawei

// Request defines data we expose to partners for an application.
type Request struct {
	// Attribute:
	//   ID
	// Type:
	//   string; required
	// Description:
	//   Unique ID of the bid request, provided by the exchange.
	ID string `json:"id"`

	// Attribute:
	//   imp
	// Type:
	//   object array; required
	// Description:
	//   Array of Imp objects (Section 3.2.4) representing the
	//   impressions offered. At least 1 Imp object is required.
	Imp []Imp `json:"imp"`

	// Attribute:
	//   app
	// Type:
	//   object; recommended
	// Description:
	//    Details via an App object (Section 3.2.14) about the publisher’s
	//    app (i.e., non-browser applications). Only applicable and
	//    recommended for apps.
	App *App `json:"app,omitempty"`

	// Attribute:
	//   device
	// Type:
	//   object; recommended
	// Description:
	//   Details via a Device object (Section 3.2.18) about the user’s
	//   device to which the impression will be delivered.
	Device *Device `json:"device,omitempty"`

	// Attribute:
	//   test
	// Type:
	//   integer; default 0
	// Description:
	//    Indicator of test mode in which auctions are not billable,
	//    where 0 = live mode, 1 = test mode.
	Test int8 `json:"test,omitempty"`

	// Attribute:
	//   at
	// Type:
	//   integer; default 2
	// Description:
	//    Auction type, where 1 = First Price, 2 = Second Price Plus.
	//    Exchange-specific auction types can be defined using values
	//    greater than 500.
	AT uint8 `json:"at,omitempty"`

	// Attribute:
	//   tmax
	// Type:
	//   integer
	// Description:
	//    Maximum time in milliseconds the exchange allows for bids to
	//    be received including Internet latency to avoid timeout. This
	//    value supersedes any a priori guidance from the exchange.
	TMax int64 `json:"tmax,omitempty"`

	// Attribute:
	//   wlang
	// Type:
	//   string array
	// Description:
	//   White list of languages for creatives using ISO-639-1-alpha-2.
	//   Omission implies no specific restrictions, but buyers would be
	//   advised to consider language attribute in the Device and/or
	//   Content objects if available.
	WLang []string `json:"wlang,omitempty"`

	WLocaleCountry []string `json:"wlocalecountry,omitempty"`
}

// App contains app details. This attribute is used for app traffic.
type App struct {

	// Attribute:
	//   ID
	// Type:
	//   string; recommended
	// Description:
	//   Exchange-specific app ID.
	ID string `json:"id,omitempty"`

	// Attribute:
	//   name
	// Type:
	//   string
	// Description:
	//   App name (may be aliased at the publisher’s request).
	Name string `json:"name,omitempty"`

	// Attribute:
	//   bundle
	// Type:
	//   string
	// Description:
	//   A platform-specific application identifier intended to be
	//   unique to the app and independent of the exchange. On
	//   Android, this should be a bundle or package name (e.g.,
	//   com.foo.mygame). On iOS, it is typically a numeric ID.
	Bundle string `json:"bundle,omitempty"`

	// Attribute:
	//   ver
	// Type:
	//   string
	// Description:
	//   Application version.
	Ver string `json:"ver,omitempty"`
}

// Asset - object array, which indicates the ad style supported by the Ad Exchange.
type Asset struct {
	TemplateID uint8  `json:"templateid"`
	W          uint16 `json:"w,omitempty"`
	H          uint16 `json:"h,omitempty"`
}

// Deal object that convey the specific deals applicable to this impression.
type Deal struct {
	ID          string   `json:"id"`
	BidFloor    float32  `json:"bidfloor,omitempty"`
	BidFloorCur string   `json:"bidfloorcur,omitempty"`
	At          uint8    `json:"at,omitempty"`
	Cur         []string `json:"cur,omitempty"`
}

// Device contain information about the user’s device to which the impression will be delivered.
type Device struct {
	// Attribute:
	//   ua
	// Type:
	//   string; recommended
	// Description:
	//   Browser user agent string.
	UA string `json:"ua,omitempty"`

	// Attribute:
	//   geo
	// Type:
	//   object; recommended
	// Description:
	//   Location of the device assumed to be the user’s current
	//   location defined by a Geo object (Section 3.2.19).
	Geo *Geo `json:"geo,omitempty"`

	// Attribute:
	//   lmt
	// Type:
	//   integer; recommended
	// Description:
	//   “Limit Ad Tracking” signal commercially endorsed (e.g., iOS,
	//   Android), where 0 = tracking is unrestricted, 1 = tracking must
	//   be limited per commercial guidelines.
	Lmt string `json:"lmt,omitempty"`

	// Attribute:
	//   ip
	// Type:
	//   string; recommended
	// Description:
	//   IPv4 address closest to device.
	IP string `json:"ip,omitempty"`

	// Attribute:
	//   devicetype
	// Type:
	//   integer
	// Description:
	//   The general type of device. Refer to List 5.21.
	DeviceType *int8 `json:"devicetype,omitempty"`

	// Attribute:
	//   make
	// Type:
	//   string
	// Description:
	//   Device make (e.g., “Apple”).
	Make string `json:"make,omitempty"`

	// Attribute:
	//   model
	// Type:
	//   string
	// Description:
	//   Device model (e.g., “iPhone”).
	Model string `json:"model,omitempty"`

	// Attribute:
	//   os
	// Type:
	//   string
	// Description:
	//   Device operating system (e.g., “iOS”).
	OS string `json:"os,omitempty"`

	// Attribute:
	//   osv
	// Type:
	//   string
	// Description:
	//   Device operating system version (e.g., “3.1.2”).
	OSV string `json:"osv,omitempty"`

	// Attribute:
	//   h
	// Type:
	//   integer
	// Description:
	//   Physical height of the screen in pixels.
	H uint16 `json:"h,omitempty"`

	// Attribute:
	//   w
	// Type:
	//   integer
	// Description:
	//   Physical width of the screen in pixels.
	W uint16 `json:"w,omitempty"`

	// Attribute:
	//   ppi
	// Type:
	//   integer
	// Description:
	//   Screen size as pixels per linear inch.
	PPI uint16 `json:"ppi,omitempty"`

	// Attribute:
	//   pxratio
	// Type:
	//   float
	// Description:
	//   The ratio of physical pixels to device independent pixels.
	PxRatio float32 `json:"pxratio,omitempty"`

	// Attribute:
	//   language
	// Type:
	//   string
	// Description:
	//   Browser language using ISO-639-1-alpha-2.
	Language string `json:"language,omitempty"`

	// Attribute:
	//   carrier
	// Type:
	//   string
	// Description:
	//   Carrier or ISP (e.g., “VERIZON”) using exchange curated string
	//   names which should be published to bidders a priori.
	Carrier string `json:"carrier,omitempty"`

	// Attribute:
	//   connectiontype
	// Type:
	//   integer
	// Description:
	//   Network connection type. Refer to List 5.22.
	ConnectionType *int8 `json:"connectiontype,omitempty"`

	// Attribute:
	//   didmd5
	// Type:
	//   string
	// Description:
	//  Hardware device ID (e.g., IMEI); hashed via MD5.
	DIDMD5 string `json:"didmd5,omitempty"`

	// Attribute:
	//   dpidmd5
	// Type:
	//   string
	// Description:
	//   Platform device ID (e.g., Android ID); hashed via MD5.
	DPIDMD5 string `json:"dpidmd5,omitempty"`

	// TODO not mapped
	DIDSHA256 string `json:"didsha256,omitempty"`

	MACSHA256 string `json:"macsha256,omitempty"`

	Oaid string `json:"oaid,omitempty"`

	Imt string `json:"imt,omitempty"`

	Gaid string `json:"gaid,omitempty"`

	Mvgeoid string `json:"mvgeoid,omitempty"`

	LocaleCountry string `json:"localeCountry,omitempty"`
}

// Geo contains geographic location information.
type Geo struct {
	Country string `json:"country"`
}

// Imp representing the impressions offered. At least 1 Imp object is required.
type Imp struct {

	// Attribute:
	//   ID
	// Type:
	//   string; required
	// Description:
	//   A unique identifier for this impression within the context of
	//   the bid request (typically, starts with 1 and increments.
	ID string `json:"id"`

	// Attribute:
	//   native
	// Type:
	//   object
	// Description:
	//   A Native object (Section 3.2.9); required if this impression is
	//   offered as a native ad opportunity.
	Native *Native `json:"native,omitempty"`

	// Attribute:
	//   pmp
	// Type:
	//   object
	// Description:
	//   A Pmp object (Section 3.2.11) containing any private
	//   marketplace deals in effect for this impression.
	PMP *PMP `json:"pmp,omitempty"`

	// Attribute:
	//   bidfloor
	// Type:
	//   float; default 0
	// Description:
	//   Minimum bid for this impression expressed in CPM.
	BidFloor float32 `json:"bidfloor,omitempty"`

	// Attribute:
	//   bidfloorcur
	// Type:
	//   string; default “USD”
	// Description:
	//   Currency specified using ISO-4217 alpha codes. This may be
	//   different from bid currency returned by bidder if this is
	//   allowed by the exchange.
	BidFloorCur string `json:"bidfloorcur,omitempty"`

	// No in prebir
	ImpExs []ImpEx `json:"impEXs,omitempty"`

	// No in prebir
	AdType uint8 `json:"adtype"`

	Bcat []string `json:"bcat,omitempty"`

	Wcat []string `json:"wcat,omitempty"`

	WinterActionType []uint8 `json:"winteractiontype,omitempty"`

	BpkgName []string `json:"bpkgname,omitempty"`

	Cur []string `json:"cur,omitempty"`

	MaxCount uint8 `json:"maxCount,omitempty"`

	TrafficDistributeMode uint8 `json:"trafficDistributeMode,omitempty"`
}

// ImpEx contains extended information array. The DSP returns the original attribute without change in the bidding response.
// This attribute is used by the current browser media.
type ImpEx struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

// Native This attribute is mandatory when the ad position type is native, native in-feed, native focus image,
// HiFolder promotion, ,splash, mobile magazine unlock, rewarded video, banner or interstitial.
type Native struct {
	Assets []Asset `json:"assets"`
}

// PMP containing any private marketplace deals in effect for this impression.
type PMP struct {

	// Attribute:
	//   private_auction
	// Type:
	//   integer; default 0
	// Description:
	//   Indicator of auction eligibility to seats named in the Direct
	//   Deals object, where 0 = all bids are accepted, 1 = bids are
	//   restricted to the deals specified and the terms thereof.
	PrivateAuction int8 `json:"private_auction,omitempty"`

	// Attribute:
	//   deals
	// Type:
	//   object array
	// Description:
	//   Array of Deal (Section 3.2.12) objects that convey the specific
	//   deals applicable to this impression.
	Deals []Deal `json:"deals"`
}
