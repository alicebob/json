package json

// rtb is rather verbose. Test and benchmark!

import (
	"encoding/json"
	"reflect"
	"testing"
)

type RTB struct {
	Id      string       `json:"id"`
	Imp     []Impression `json:"imp,omitempty"`
	Site    *Site        `json:"site,omitempty"`
	App     *App         `json:"app,omitempty"`
	Device  *Device      `json:"device,omitempty"`
	User    *User        `json:"user,omitempty"`
	Test    int          `json:"test,omitempty"`
	At      int          `json:"at,omitempty"`
	Tmax    int          `json:"tmax,omitempty"`
	Wseat   []string     `json:"wseat,omitempty"`
	AllImps int          `json:"allimps,omitempty"`
	Cur     []string     `json:"cur,omitempty"`
	BCat    []string     `json:"bcat,omitempty"`
	BAdv    []string     `json:"badv,omitempty"`
	// Regs    *Regs           `json:"regs,omitempty"`
	Ext RawMessage `json:"ext,omitempty"`
}
type App struct {
	Id            string     `json:"id,omitempty"`
	Name          string     `json:"name,omitempty"`
	Bundle        string     `json:"bundle,omitempty"`
	Domain        string     `json:"domain,omitempty"`
	StoreURL      string     `json:"storeurl,omitempty"`
	Cat           []string   `json:"cat,omitempty"`
	SectionCat    []string   `json:"sectioncat,omitempty"`
	PageCat       []string   `json:"pagecat,omitempty"`
	Ver           string     `json:"ver,omitempty"`
	PrivacyPolicy int        `json:"privacypolicy,omitempty"`
	Paid          int        `json:"paid,omitempty"`
	Publisher     *Publisher `json:"publisher,omitempty"`
	Content       *Content   `json:"content,omitempty"`
	Keywords      string     `json:"keywords,omitempty"`
	Ext           RawMessage `json:"ext,omitempty"`
}
type Publisher struct {
	Id     string     `json:"id,omitempty"`
	Name   string     `json:"name,omitempty"`
	Cat    []string   `json:"cat,omitempty"`
	Domain string     `json:"domain,omitempty"`
	Ext    RawMessage `json:"ext,omitempty"`
}
type Device struct {
	Ua             string     `json:"ua,omitempty"`
	Geo            *Geo       `json:"geo,omitempty"`
	Dnt            int        `json:"dnt,omitempty"`
	LMT            int        `json:"lmt,omitempty"`
	Ip             string     `json:"ip,omitempty"`
	Ipv6           string     `json:"ipv6,omitempty"`
	DeviceType     int        `json:"devicetype,omitempty"`
	Make           string     `json:"make,omitempty"`
	Model          string     `json:"model,omitempty"`
	Os             string     `json:"os,omitempty"`
	Osv            string     `json:"osv,omitempty"`
	Hwv            string     `json:"hwv,omitempty"`
	H              int        `json:"h,omitempty"`
	W              int        `json:"w,omitempty"`
	PPI            int        `json:"ppi,omitempty"`
	PXRatio        float64    `json:"pxration,omitempty"`
	Js             int        `json:"js,omitempty"`
	FlashVer       string     `json:"flashver,omitempty"`
	Language       string     `json:"language,omitempty"`
	Carrier        string     `json:"carrier,omitempty"`
	ConnectionType int        `json:"connectiontype,omitempty"`
	Ifa            string     `json:"ifa,omitempty"`
	DidSHA1        string     `json:"didsha1,omitempty"`
	DidMD5         string     `json:"didmd5,omitempty"`
	DpidSHA1       string     `json:"dpidsha1,omitempty"`
	DpidMD5        string     `json:"dpidmd5,omitempty"`
	MacSHA1        string     `json:"macsha1,omitempty"`
	MacMD5         string     `json:"macmd5,omitempty"`
	Ext            RawMessage `json:"ext,omitempty"`
}
type Geo struct {
	Lat           float64    `json:"lat,omitempty"`
	Lon           float64    `json:"lon,omitempty"`
	Type          int        `json:"type,omitempty"`
	Country       string     `json:"country,omitempty"`
	Region        string     `json:"region,omitempty"`
	RegionFips104 string     `json:"regionfips104,omitempty"`
	Metro         string     `json:"metro,omitempty"`
	City          string     `json:"city,omitempty"`
	Zip           string     `json:"zip,omitempty"`
	UTCOffset     int        `json:"utcoffset,omitempty"`
	Ext           RawMessage `json:"ext,omitempty"`
}
type User struct {
	Id         string     `json:"id,omitempty"`
	BuyerUid   string     `json:"buyeruid,omitempty"`
	Yob        int        `json:"yob,omitempty"`
	Gender     string     `json:"gender,omitempty"`
	Keywords   string     `json:"keywords,omitempty"`
	CustomData string     `json:"customdata,omitempty"`
	Geo        *Geo       `json:"geo,omitempty"`
	Data       []Data     `json:"data,omitempty"`
	Ext        RawMessage `json:"ext,omitempty"`
}
type Data struct {
	Id      string     `json:"id,omitempty"`
	Name    string     `json:"name,omitempty"`
	Segment []Segment  `json:"segment,omitempty"`
	Ext     RawMessage `json:"ext,omitempty"`
}

type Segment struct {
	Id    string     `json:"id,omitempty"`
	Name  string     `json:"name,omitempty"`
	Value string     `json:"value,omitempty"`
	Ext   RawMessage `json:"ext,omitempty"`
}
type Impression struct {
	Id                string     `json:"id,omitempty"`
	Banner            *Banner    `json:"banner,omitempty"`
	Video             *Video     `json:"video,omitempty"`
	Native            *Native    `json:"native,omitempty"`
	DisplayManager    string     `json:"displaymanager,omitempty"`
	DisplayManagerVer string     `json:"displaymanagerver,omitempty"`
	Instl             int        `json:"instl,omitempty"`
	TagId             string     `json:"tagid,omitempty"`
	BidFloor          float64    `json:"bidfloor,omitempty"`
	BidFloorCur       string     `json:"bidfloorcur,omitempty"`
	Secure            int        `json:"secure,omitempty"`
	IFrameBuster      []string   `json:"iframebuster,omitempty"`
	Pmp               *Pmp       `json:"pmp,omitempty"`
	Ext               RawMessage `json:"ext,omitempty"`
}
type Banner struct {
	W        int        `json:"w,omitempty"`
	H        int        `json:"h,omitempty"`
	Wmax     int        `json:"wmax,omitempty"`
	Hmax     int        `json:"hmax,omitempty"`
	Wmin     int        `json:"wmin,omitempty"`
	Hmin     int        `json:"hmin,omitempty"`
	Id       string     `json:"id,omitempty"`
	BType    []int      `json:"btype,omitempty"`
	BAttr    []int      `json:"battr,omitempty"`
	Pos      int        `json:"pos,omitempty"`
	Mimes    []string   `json:"mimes,omitempty"`
	TopFrame int        `json:"topframe,omitempty"`
	ExpDir   []int      `json:"expdir,omitempty"`
	Api      []int      `json:"api,omitempty"`
	Ext      RawMessage `json:"ext,omitempty"`
}
type Native struct {
	Request string `json:"request"`
	Version string `json:"version,omitempty"`
	BAttr   []int  `json:"battr,omitempty"`
}
type Video struct {
	Mimes          []string        `json:"mimes,omitempty"`
	MinDuration    int             `json:"minduration,omitempty"`
	MaxDuration    int             `json:"maxduration,omitempty"`
	Protocol       int             `json:"protocol,omitempty"`
	Protocols      []int           `json:"protocols,omitempty"`
	W              int             `json:"w,omitempty"`
	H              int             `json:"h,omitempty"`
	StartDelay     int             `json:"startdelay,omitempty"`
	Linearity      int             `json:"linearity,omitempty"`
	Sequence       int             `json:"sequence,omitempty"`
	BAttr          []int           `json:"battr,omitempty"`
	MaxExtended    int             `json:"maxextended,omitempty"`
	MinBitRate     int             `json:"minbitrate,omitempty"`
	MaxBitRate     int             `json:"maxbitrate,omitempty"`
	BoxingAllowed  int             `json:"boxingallowed,omitempty"`
	PlaybackMethod []int           `json:"playbackmethod,omitempty"`
	Delivery       []int           `json:"delivery,omitempty"`
	Pos            int             `json:"pos,omitempty"`
	CompanionAd    []Banner        `json:"companionad,omitempty"`
	Api            []int           `json:"api,omitempty"`
	CompanionType  []int           `json:"companiontype,omitempty"`
	Ext            json.RawMessage `json:"ext,omitempty"`
}
type Pmp struct {
	PrivateAuction int        `json:"private_auction,omitempty"`
	Deals          []Deal     `json:"deals,omitempty"`
	Ext            RawMessage `json:"ext,omitempty"`
}
type Deal struct {
	Id          string     `json:"id,omitempty"`
	BidFloor    float64    `json:"bidfloor,omitempty"`
	BidFloorCur string     `json:"bidfloorcur,omitempty"`
	At          int        `json:"at,omitempty"`
	WSeat       []string   `json:"wseat,omitempty"`
	WAdomain    []string   `json:"wadomain,omitempty"`
	Ext         RawMessage `json:"ext,omitempty"`
}
type Site struct {
	Id            string     `json:"id,omitempty"`
	Name          string     `json:"name,omitempty"`
	Domain        string     `json:"domain,omitempty"`
	Cat           []string   `json:"cat,omitempty"`
	SectionCat    []string   `json:"sectioncat,omitempty"`
	PageCat       []string   `json:"pagecat,omitempty"`
	Page          string     `json:"page,omitempty"`
	Ref           string     `json:"ref,omitempty"`
	Search        string     `json:"search,omitempty"`
	Mobile        int        `json:"mobile,omitempty"`
	PrivacyPolicy int        `json:"privacypolicy,omitempty"`
	Publisher     *Publisher `json:"publisher,omitempty"`
	Content       *Content   `json:"content,omitempty"`
	Keywords      string     `json:"keywords,omitempty"`
	Ext           RawMessage `json:"ext,omitempty"`
}
type Content struct {
	Id                 string     `json:"id,omitempty"`
	Episode            int        `json:"episode,omitempty"`
	Title              string     `json:"title,omitempty"`
	Series             string     `json:"series,omitempty"`
	Season             string     `json:"season,omitempty"`
	Producer           *Producer  `json:"producer,omitempty"`
	URL                string     `json:"url,omitempty"`
	Cat                []string   `json:"cat,omitempty"`
	VideoQuality       int        `json:"videoquality,omitempty"`
	ContentRating      string     `json:"contentrating,omitempty"`
	UserRating         string     `json:"userrating,omitempty"`
	QagMediaRating     int        `json:"qagmediarating,omitempty"`
	Keywords           string     `json:"keywords,omitempty"`
	Livestream         int        `json:"livestream,omitempty"`
	SourceRelationship int        `json:"sourcerelationship,omitempty"`
	Len                int        `json:"len,omitempty"`
	Language           string     `json:"language,omitempty"`
	Embeddable         int        `json:"embeddable,omitempty"`
	Context            int        `json:"context,omitempty"`
	Ext                RawMessage `json:"ext,omitempty"`
}
type Producer struct {
	Id     string          `json:"id,omitempty"`
	Name   string          `json:"name,omitempty"`
	Cat    []string        `json:"cat,omitempty"`
	Domain string          `json:"domain,omitempty"`
	Ext    json.RawMessage `json:"ext,omitempty"`
}

var (
	rawRTB = `
{ 
   "app":{ 
      "bundle":"553834731", 
      "cat":[ 
         "IAB3", 
         "business" 
      ], 
      "id":"XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX", 
      "name":"App_Name", 
      "publisher":{ 
         "id":"01234567890abcdef01234567890abcd", 
         "name":"Publisher_Name" 
      }, 
      "storeurl":"https://itunes.apple.com/us/app/candy­crush­saga/id553834731?mt=8", 
      "ver":"1.0" 
   }, 
   "at":2, 
   "bcat":[ 
      "IAB7­39", 
      "IAB8­5", 
      "IAB8­18", 
      "IAB9­9", 
      "IAB25", 
      "IAB26", 
      "IAB3­7" 
   ], 
   "device":{ 
      "carrier":"310-260", 
      "connectiontype":2, 
      "devicetype":4, 
      "dnt":0, 
      "geo":{ 
         "country":"USA", 
         "lat":10.738701, 
         "lon":76.0037 
      }, 
      "h":1920, 
      "hwv":"iPhone 6+", 
      "ifa":"12345678-90ab-cdef-0123-456789abcdef", 
      "ip":"1.23.123.12", 
      "js":1, 
      "language":"en", 
      "make":"Apple", 
      "model":"iPhone", 
      "os":"iOS", 
      "osv":"8.1", 
      "ua":"Mozilla/5.0 (iPhone; CPU iPhone OS 8_1 like Mac OS X) AppleWebKit/600.1.4 (KHTML, like Gecko) Mobile/12B411", 
      "w":1080 
   }, 
   "id":"12345678-90ab-cdef-0123-456789abcdef", 
   "imp":[ 
      { 
         "banner":{ 
            "api":[ 
               3, 
               5 
            ], 
            "battr":[ 
               3, 
               8, 
               9, 
               10, 
               14, 
               6 
            ], 
            "btype":[ 
               4 
            ], 
            "h":50, 
            "pos":1, 
            "w":320 
         }, 
         "bidfloor":0.02, 
         "displaymanager":"mopub", 
         "displaymanagerver":"3.10.0", 
         "ext":{ 
            "brsrclk":1, 
            "dlp":1 
         }, 
         "id":"1", 
         "instl":0, 
         "secure":0, 
         "tagid":"abcdef0123456789abcdef0123456789" 
      } 
   ] 
}
`
)

func TestRTB(t *testing.T) {
	vj := RTB{}
	if err := json.Unmarshal([]byte(rawRTB), &vj); err != nil {
		t.Fatal(err)
	}

	v := RTB{}
	if err := Decode(rawRTB, &v); err != nil {
		t.Fatal(err)
	}

	if have, want := v, vj; !reflect.DeepEqual(have, want) {
		t.Errorf("have %+v, want %+v", have, want)
	}
}

func BenchmarkRTB(b *testing.B) {
	for i := 0; i < b.N; i++ {
		tv := RTB{}
		if err := Decode(rawRTB, &tv); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkRTBThem(b *testing.B) {
	raw := []byte(rawRTB)
	for i := 0; i < b.N; i++ {
		tv := RTB{}
		if err := json.Unmarshal(raw, &tv); err != nil {
			b.Fatal(err)
		}
	}
}
