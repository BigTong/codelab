package rpcbm

import (
	"time"

	"codelab/rpcbentchmark/proto"
)

type Tag struct {
	Value    string `json:"value"`
	Category string `json:"category"`
}

type Answer struct {
	QuestionId string `json:"id"`
	Answer     string `json:"answer"`
}

type Account struct {
	Value   string `json:"v"`
	Network string `json:"n"`
}

type Region struct {
	Country  Tlation `json:"c"`
	Province Tlation `json:"p"`
	City     Tlation `json:"city"`
	District Tlation `json:"district"`
}

type Tlation struct {
	Id string `json:"id"`
	Zh string `json:"zh"`
	En string `json:"en"`
	Ko string `json:"ko"`
	Ja string `json:"ja"`
}

type Media struct {
	Id          string    `json:"id"`
	UserId      string    `json:"uid"`
	CreatedTime time.Time `json:"c_time"`
	UpdatedTime time.Time `json:"u_time"`
	Index       int       `json:"index"`
	Identifier  string    `json:"identifier"`
	Width       int       `json:"width"`
	Height      int       `json:"height"`
	DHash       uint64    `json:"d_hash"`
	VIdentifier string    `json:"v_i"`
	VWidth      int       `json:"v_width"`
	VHeight     int       `json:"v_hegiht"`
	VDuration   float64   `json:"v_duration"`
	AIdentifier string    `json:"a_i"`
	ADuration   float64   `json:"a_duration"`
	Status      string    `json:"status"`
	Rel         string    `json:"rel"`
	First       bool      `json:"first"`
	Other       bool      `json:"other"`
}

type Pass struct {
	Id          string    `json:"id"`
	UserId      string    `json:"uid"`
	OtherUserId string    `json:"oid"`
	Count       int       `json:"cnt"`
	CTime       time.Time `json:"ctime"`
	LTime       time.Time `json:"ltime"`
	Longitude   float64   `json:"longitude"`
	Latitude    float64   `json:"latitude"`
}

type User struct {
	Id              string     `json:"id"`
	Name            string     `json:"name"`
	Des             string     `json:"des"`
	Industry        string     `json:"industry"`
	Company         string     `json:"company"`
	Department      string     `json:"department"`
	School          string     `json:"school"`
	Major           string     `json:"major"`
	Hometown        string     `json:"hometown"`
	Hangouts        string     `json:"hangouts"`
	Tags            []Tag      `json:"tags"`
	Accounts        []Account  `json:"accounts"`
	Answers         []Answer   `json:"answers"`
	Gender          string     `json:"gender"`
	Looking         string     `json:"looking"`
	Intent          string     `json:"intent"`
	Longitude       float64    `json:"longitude"`
	Latitude        float64    `json:"latitude"`
	Distance        float64    `json:"distance"`
	Region          Region     `json:"region"`
	Radius          int        `json:"radius"`
	MaxAge          int        `json:"max_age"`
	MinAge          int        `json:"min_age"`
	Contacts        bool       `json:"contacts"`
	MutualContacts  bool       `json:"mutual_contacts"`
	Email           string     `json:"email"`
	CountryCode     int        `json:"country_code"`
	MobileNumber    string     `json:"mobile_number"`
	Password        string     `json:"password"`
	Birthdate       time.Time  `json:"birthdate"`
	Pictures        []Media    `json:"pictures"`
	LastActivity    time.Time  `json:"last_activity"`
	UpdatedTime     time.Time  `json:"update_time"`
	CreatedTime     time.Time  `json:"create_time"`
	Passby          *Pass      `json:"passby"`
	Popularity      float64    `json:"popularity"`
	PreviewMessage  bool       `json:"preview_message"`
	ShowMomentLikes bool       `json:"show_moment_likes"`
	WorkActive      bool       `json:"work_activi"`
	StudyActive     bool       `json:"study_active"`
	ScenarioIds     []string   `json:"scenario_ids"`
	Category        *string    `json:"category"`
	ExpiresTime     *time.Time `json:"expires_time"`
	Status          string     `json:"status"`
	Type            string     `json:"type"`
}

func NewFakeUser() *User {
	t := time.Now()
	unknown := "unkown"
	ret := &User{
		Id:         "000001",
		Name:       "zhangsan",
		Des:        "I like runing",
		Industry:   "IT",
		Company:    "baidu.inc",
		Department: "tech",
		School:     "peking",
		Major:      "computer",
		Hometown:   "hunan",
		Hangouts:   "xxx",
		Tags: []Tag{
			Tag{"aaaa", "bbbbb"},
			Tag{"ccccc", "ddddd"}},
		Accounts: []Account{
			Account{"aaaa", "tcp"},
			Account{"bbbb", "udp"}},
		Answers: []Answer{
			Answer{"are you ok?", "yes"},
			Answer{"are you fine?", "fine"}},
		Gender:         "male",
		Looking:        "femal",
		Intent:         "unkown",
		Longitude:      123.1,
		Latitude:       180.4,
		Distance:       123.5,
		Radius:         100,
		MaxAge:         28,
		MinAge:         18,
		Contacts:       false,
		MutualContacts: true,
		Email:          "xxx@gmail.com",
		CountryCode:    122,
		MobileNumber:   "1334567899",
		Password:       "xxxxxxxx",
		Birthdate:      time.Now(),
		Pictures: []Media{
			Media{Id: "id",
				UserId:      "uid",
				CreatedTime: time.Now(),
				UpdatedTime: time.Now(),
				Index:       1,
				Identifier:  "what",
				Width:       10,
				Height:      10,
				DHash:       121424324234,
				VIdentifier: "vid",
				VWidth:      12,
				VHeight:     0,
			},
		},
		LastActivity:    time.Now(),
		UpdatedTime:     time.Now(),
		CreatedTime:     time.Now(),
		Popularity:      1.0,
		PreviewMessage:  false,
		ShowMomentLikes: true,
		WorkActive:      true,
		StudyActive:     false,
		ScenarioIds:     []string{"beijing", "nanjing"},
		Category:        &unknown,
		ExpiresTime:     &t,
		Status:          "ok",
		Type:            "unkown",
	}

	region := Region{
		Country: Tlation{
			Id: "ii",
			Zh: "zh",
			En: "en",
			Ko: "ko",
			Ja: "ja",
		},
		Province: Tlation{
			Id: "ii",
			Zh: "zh",
			En: "en",
			Ko: "ko",
			Ja: "ja",
		},
		City: Tlation{
			Id: "ii",
			Zh: "zh",
			En: "en",
			Ko: "ko",
			Ja: "ja",
		},
		District: Tlation{
			Id: "ii",
			Zh: "zh",
			En: "en",
			Ko: "ko",
			Ja: "ja",
		},
	}
	ret.Region = region
	return ret
}

func NewFakeProtoUser() *proto.User {
	ret := &proto.User{
		Id:         "000001",
		Name:       "zhangsan",
		Des:        "I like runing",
		Industry:   "IT",
		Company:    "baidu.inc",
		Department: "tech",
		School:     "peking",
		Major:      "computer",
		Hometown:   "hunan",
		Hangouts:   "xxx",
		Tags: []*proto.Tag{
			&proto.Tag{"aaaa", "bbbbb"},
			&proto.Tag{"ccccc", "ddddd"}},
		Accounts: []*proto.Account{
			&proto.Account{"aaaa", "tcp"},
			&proto.Account{"bbbb", "udp"}},
		Answers: []*proto.Answer{
			&proto.Answer{"are you ok?", "yes"},
			&proto.Answer{"are you fine?", "fine"}},
		Gender:         "male",
		Looking:        "femal",
		Intent:         "unkown",
		Longitude:      123.1,
		Latitude:       180.4,
		Distance:       123.5,
		Radius:         100,
		MaxAge:         28,
		MinAge:         18,
		Contacts:       false,
		MutualContacts: true,
		Email:          "xxx@gmail.com",
		CountryCode:    122,
		MobileNumber:   "1334567899",
		Password:       "xxxxxxxx",
		Birthdate:      time.Now().Unix(),
		Pictures: []*proto.Media{
			&proto.Media{Id: "id",
				UserId:      "uid",
				CreatedTime: time.Now().Unix(),
				UpdatedTime: time.Now().Unix(),
				Index:       1,
				Identifier:  "what",
				Width:       10,
				Height:      10,
				DHash:       121424324234,
				VIdentifier: "vid",
				VWidth:      12,
				VHeight:     0,
			},
		},
		LastActivity:    time.Now().Unix(),
		UpdatedTime:     time.Now().Unix(),
		CreatedTime:     time.Now().Unix(),
		Popularity:      1.0,
		PreviewMessage:  false,
		ShowMomentLikes: true,
		WorkActive:      true,
		StudyActive:     false,
		ScenarioIds:     []string{"beijing", "nanjing"},
		Category:        "unknown",
		ExpiresTime:     time.Now().Unix(),
		Status:          "ok",
		Type:            "unkown",
	}

	region := &proto.Region{
		Country: &proto.Tlation{
			Id: "ii",
			Zh: "zh",
			En: "en",
			Ko: "ko",
			Ja: "ja",
		},
		Province: &proto.Tlation{
			Id: "ii",
			Zh: "zh",
			En: "en",
			Ko: "ko",
			Ja: "ja",
		},
		City: &proto.Tlation{
			Id: "ii",
			Zh: "zh",
			En: "en",
			Ko: "ko",
			Ja: "ja",
		},
		District: &proto.Tlation{
			Id: "ii",
			Zh: "zh",
			En: "en",
			Ko: "ko",
			Ja: "ja",
		},
	}
	ret.Region = region
	return ret
}
