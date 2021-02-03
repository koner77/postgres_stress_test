package huawei

// Response represents struct that we send to Huawei
type Response struct {
	ID      string    `json:"id"`
	BidID   *string   `json:"bidid,omitempty"`
	SeatBid []SeatBid `json:"seatbid"`
}

// SeatBid represents a bidding seat, that is, each SeatBid represents a bidding response of an advertiser (or agent).
type SeatBid struct {
	Bid  []Bid  `json:"bid"`
	Seat string `json:"seat"`
}

// Bid - Add one bid for each native object and duplicate ImpID
type Bid struct {
	ID         string      `json:"id,omitempty"`
	ImpID      string      `json:"impid"`
	Price      float32     `json:"price"`          // hard coded
	CID        string      `json:"cid"`            // id from post
	CrID       string      `json:"crid"`           // hard code it
	Cur        string      `json:"cur"`            //SEK
	Lurl       *string     `json:"lurl,omitempty"` //SEK
	Nurl       *string     `json:"nurl,omitempty"` //SEK
	NativeResp HNativeResp `json:"nativersp,omitempty"`
	Cat        []CatType   `json:"cat"`
	Language   string      `json:"language,omitempty"`
	ImpExs     []ImpEx     `json:"impEXs,omitempty"`
}

// HNativeResp contains creative information.
// This attribute is mandatory when the ad position type is native, native in-feed, native focus image, HiFolder,
// splash, mobile magazine unlock, rewarded video, banner or interstitial.
type HNativeResp struct {
	TemplateID  uint8    `json:"templateid"`
	Title       string   `json:"title,omitempty"`
	Brand       string   `json:"brand,omitempty"`
	Images      []Image  `json:"images,omitempty"`
	ImpTrackers []string `json:"imptrackers,omitempty"`
	ExpireTime  int64    `json:"expiretime,omitempty"`
	Link        *Link    `json:"link,omitempty"`
	Video       *Video   `json:"video,omitempty"`
}

// Video contains Video file information.
type Video struct {
	URL      string `json:"url"`
	Duration uint16 `json:"duration"`
	FileSize uint32 `json:"filesize"`
	Mime     string `json:"mime"`
	W        uint16 `json:"w,omitempty"`
	H        uint16 `json:"h,omitempty"`
	Sha256   string `json:"sha256,omitempty"`
}

// Link represents target link.
// This attribute is optional for the pure display promotion type and mandatory for other promotion types.
type Link struct {
	InteractionType uint8    `json:"interactiontype"`
	URL             string   `json:"url,omitempty"`
	ClickTrackers   []string `json:"clicktrackers,omitempty"`
}

// Image represents Huawei struct images

type Image struct {
	URL  string `json:"url"`
	W    uint16 `json:"w"`
	H    uint16 `json:"h"`
	Mime string `json:"mime"`
}

type TaskResponse struct {
	Response        Response
	GeneralProvider string
	RealProvider    string
	Error           error
}
