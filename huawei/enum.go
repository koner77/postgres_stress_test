package huawei

// List of default values
const (
	DefaultCur     = "USD"
	DefaultLang    = "en"
	DefaultCountry = "gb"
	MaxBrandLength = 28
	MaxCrIDLength  = 49
	// BuyerID this is ID of the impression buyer (advertiser or agent who participates in the bidding)
	BuyerID = "1234"
)

// List of possible templates
const (
	BannerType     uint8 = 18
	BannerImage    uint8 = 9
	VideoType      uint8 = 6
	SmallImageType uint8 = 4
	ThreeImageType uint8 = 3
	ImageType      uint8 = 1
)

// List of SQUID impression and click trackers
const (
	impTrackersLink   = "https://ads.squidapp.co/log/imp?id="
	clickTrackersLink = "https://ads.squidapp.co/log/click?id="
)

// CatType represents categories of Huawei content
type CatType string

// represent values of CatType
const (
	ArtsEntertainment    CatType = "IAB1"
	Automotive           CatType = "IAB2"
	Business             CatType = "IAB3"
	Careers              CatType = "IAB4"
	Education            CatType = "IAB5"
	FamilyParenting      CatType = "IAB6"
	FoodDrink            CatType = "IAB8"
	HealthFitness        CatType = "IAB7"
	HobbiesInterests     CatType = "IAB9"
	HomeGarden           CatType = "IAB10"
	LawGovtPolitics      CatType = "IAB11"
	News                 CatType = "IAB12"
	PersonalFinance      CatType = "IAB13"
	Pets                 CatType = "IAB16"
	RealEstate           CatType = "IAB21"
	ReligionSpirituality CatType = "IAB23"
	Science              CatType = "IAB15"
	Shopping             CatType = "IAB22"
	Society              CatType = "IAB14"
	Sports               CatType = "IAB17"
	StyleFashion         CatType = "IAB18"
	TechnologyComputing  CatType = "IAB19"
	Travel               CatType = "IAB20"
	Uncategorized        CatType = "IAB24"
)

// AvailableCountry list of countries in which we show ads
var AvailableCountry = map[string]bool{
	"fo": true,
	"mc": true,
	"lv": true,
	"me": true,
	"va": true,
	"md": true,
	"ee": true,

	"bo": true,
	"do": true,
	"ve": true,
	"sv": true,
	"gt": true,
	"ni": true,
	"hn": true,
	"pa": true,
	"py": true,
	"uy": true,

	"se": true,
	"de": true,
	"fr": true,
	"pl": true,
	"it": true,
	"es": true,
	"ua": true,
	"gb": true,
	"hu": true,
	"cz": true,
	"fi": true,
	"no": true,
	"dk": true,

	"ro": true, /* romania */
	"nl": true, /*netherlands*/
	"be": true, /*belgium*/
	"ie": true, /* ierland*/
	"pt": true, /* portugal*/
	"at": true, /* austria*/
	"gr": true, /* greece*/
	"bg": true, /*bulgaria*/
	"rs": true, /* serbia*/
	"sk": true, /*slovakia*/
	"hr": true, /*croatia*/
	"ch": true, /*swiss*/
	"lt": true,
	"ad": true,
	"si": true, /* sloveniia*/
	"mx": true,

	"co": true,
	"cl": true,
	"pe": true,
	"ar": true,
	"br": true,
	"sm": true,
	"ec": true,
	"cr": true,
}

// AvailableLang the list of supported languages in these countries we do not show our ads yet
var AvailableLang = map[string]bool{
	"ru": true,
}

// availableTemplates the list of supported templates
var availableTemplates = []uint8{ImageType, BannerImage, ThreeImageType, SmallImageType, VideoType, BannerType}

const (
	UsualPrice   = 4
	SpecialPrice = 12
)

const (
	Assistant        string = "assistant"
	Browser          string = "browser"
	Petal            string = "petal"
	OtherApps        string = "OtherApps"
	OtherAppsSpecial string = "OtherAppsSpecial"
)

var BundleMapSpecial = map[string]string{
	"com.snaptube.premium":   OtherAppsSpecial,
	"com.nemo.vidmate":       OtherAppsSpecial,
	"com.dwarfplanet.bundle": OtherAppsSpecial,
	"org.ytv.downsave":       OtherAppsSpecial,
}

var availableBundles = []string{Assistant, Browser, Petal, OtherApps}

var bundleMap = map[string]string{
	"com.huawei.browser":     Browser,
	"com.huawei.hiboard":     Assistant,
	"com.huawei.intelligent": Assistant,
	"com.huawei.petal":       Petal,
	"com.huawei.hwsearch":    Petal,
}

// BundleURLMap maps bundle name with URL to huawei repository
var BundleURLMap = map[string]string{
	"com.huawei.petal":                                        "https://appgallery.cloud.huawei.com/ag/n/app/C100995735",
	"com.huawei.hwsearch":                                     "https://appgallery.cloud.huawei.com/ag/n/app/C100995735",
	"com.android.mediacenter":                                 "https://appgallery.cloud.huawei.com/ag/n/app/C10021449",
	"com.huawei.music":                                        "https://appgallery.cloud.huawei.com/ag/n/app/C10021449",
	"com.huawei.himovie.overseas":                             "https://appgallery.cloud.huawei.com/ag/n/app/C100065895",
	"com.snaptube.premium":                                    "https://appgallery.cloud.huawei.com/ag/n/app/C100180977",
	"com.popularapp.periodcalendar":                           "https://appgallery.cloud.huawei.com/ag/n/app/C10322075",
	"com.nemo.vidmate":                                        "https://appgallery.cloud.huawei.com/ag/n/app/C100854761",
	"com.appgeneration.itunerfree":                            "https://appgallery.cloud.huawei.com/ag/n/app/C100172409",
	"com.dwarfplanet.bundle":                                  "https://appgallery.cloud.huawei.com/ag/n/app/C102507039",
	"steptracker.healthandfitness.walkingtracker.pedometer":   "https://appgallery.cloud.huawei.com/ag/n/app/C101586767",
	"com.imo.android.imoim":                                   "https://appgallery.cloud.huawei.com/ag/n/app/C100924809",
	"pedometer.steptracker.calorieburner.stepcounter":         "https://appgallery.cloud.huawei.com/ag/n/app/C100450293",
	"com.xian.beatmusic.huawei":                               "https://appgallery.cloud.huawei.com/ag/n/app/C103160303",
	"pedometer.stepcounter.calorieburner.pedometerforwalking": "https://appgallery.cloud.huawei.com/ag/n/app/C100961647",
	"com.download.free.videodownload":                         "https://appgallery.cloud.huawei.com/ag/n/app/C101454149",
	"video.downloader.videodownloader":                        "https://appgallery.cloud.huawei.com/ag/n/app/C101505615",
}

var templateForBundleMap = map[string][]uint8{
	Assistant: {ImageType, ThreeImageType, SmallImageType, BannerType},
	Browser:   {ImageType, ThreeImageType, SmallImageType, VideoType, BannerType},
	Petal:     {ImageType, ThreeImageType, SmallImageType, BannerType},
	OtherApps: {ImageType, ThreeImageType, SmallImageType, BannerImage, BannerType},
}

// set priority order of templates
// the lower item has higher priority
// if assets not exists in map, then weight = 0
// for example:
// given array 1 2 3 4 18
// given weights {3:1, 4:2, 1:3, 6:4}
// expected result is 6 > 1 > 4 > 3 > 18
var assetsPriorityMap = map[uint8]uint8{
	ThreeImageType: 1,
	SmallImageType: 2,
	ImageType:      3,
	VideoType:      4,
}

var domainAllowList = []string{
	"adform.net",
	"adnxs.com",
	"adsafeprotected.com",
	"de17a.com",
	"doubleclick.net",
	"gemius.pl",
	"mgid.com",
	"outbrain.com",
	"outbrainimg.com",
	"plista.com",
	"s3.eu-central-1.amazonaws.com",
	"serving-sys.com",
	"smartadserver.com",
	"spklw.com",
	"sprinklecontent.com",
	"squidapp.co",
}

var domainBlockList = []string{}

/* looking for a complete match of strings */
var titleBlockList = []string{}

var keyWordBlockList = []string{}

var brandBlockList = []string{}
