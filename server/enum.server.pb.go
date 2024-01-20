// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.4.0
// source: enum.server.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Gender int32

const (
	Gender_GenderNone  Gender = 0
	Gender_GenderMan   Gender = 1
	Gender_GenderWoman Gender = 2
)

// Enum value maps for Gender.
var (
	Gender_name = map[int32]string{
		0: "GenderNone",
		1: "GenderMan",
		2: "GenderWoman",
	}
	Gender_value = map[string]int32{
		"GenderNone":  0,
		"GenderMan":   1,
		"GenderWoman": 2,
	}
)

func (x Gender) Enum() *Gender {
	p := new(Gender)
	*p = x
	return p
}

func (x Gender) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Gender) Descriptor() protoreflect.EnumDescriptor {
	return file_enum_server_proto_enumTypes[0].Descriptor()
}

func (Gender) Type() protoreflect.EnumType {
	return &file_enum_server_proto_enumTypes[0]
}

func (x Gender) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Gender.Descriptor instead.
func (Gender) EnumDescriptor() ([]byte, []int) {
	return file_enum_server_proto_rawDescGZIP(), []int{0}
}

type HeroBasicType int32

const (
	HeroBasicType_None        HeroBasicType = 0
	HeroBasicType_BoyWarrior  HeroBasicType = 8001
	HeroBasicType_GirlWarrior HeroBasicType = 8002
	HeroBasicType_BoyKnight   HeroBasicType = 8003
	HeroBasicType_GirlKnight  HeroBasicType = 8004
	HeroBasicType_BoyRogue    HeroBasicType = 8005
	HeroBasicType_GirlRogue   HeroBasicType = 8006
	HeroBasicType_BoyMage     HeroBasicType = 8007
	HeroBasicType_GirlMage    HeroBasicType = 8008
	HeroBasicType_BoyShaman   HeroBasicType = 8009
	HeroBasicType_GirlShaman  HeroBasicType = 8010
	HeroBasicType_BoyWarlock  HeroBasicType = 8011
	HeroBasicType_GirlWarlock HeroBasicType = 8012
	HeroBasicType_BoyPriest   HeroBasicType = 8013
	HeroBasicType_GirlPriest  HeroBasicType = 8014
)

// Enum value maps for HeroBasicType.
var (
	HeroBasicType_name = map[int32]string{
		0:    "None",
		8001: "BoyWarrior",
		8002: "GirlWarrior",
		8003: "BoyKnight",
		8004: "GirlKnight",
		8005: "BoyRogue",
		8006: "GirlRogue",
		8007: "BoyMage",
		8008: "GirlMage",
		8009: "BoyShaman",
		8010: "GirlShaman",
		8011: "BoyWarlock",
		8012: "GirlWarlock",
		8013: "BoyPriest",
		8014: "GirlPriest",
	}
	HeroBasicType_value = map[string]int32{
		"None":        0,
		"BoyWarrior":  8001,
		"GirlWarrior": 8002,
		"BoyKnight":   8003,
		"GirlKnight":  8004,
		"BoyRogue":    8005,
		"GirlRogue":   8006,
		"BoyMage":     8007,
		"GirlMage":    8008,
		"BoyShaman":   8009,
		"GirlShaman":  8010,
		"BoyWarlock":  8011,
		"GirlWarlock": 8012,
		"BoyPriest":   8013,
		"GirlPriest":  8014,
	}
)

func (x HeroBasicType) Enum() *HeroBasicType {
	p := new(HeroBasicType)
	*p = x
	return p
}

func (x HeroBasicType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (HeroBasicType) Descriptor() protoreflect.EnumDescriptor {
	return file_enum_server_proto_enumTypes[1].Descriptor()
}

func (HeroBasicType) Type() protoreflect.EnumType {
	return &file_enum_server_proto_enumTypes[1]
}

func (x HeroBasicType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use HeroBasicType.Descriptor instead.
func (HeroBasicType) EnumDescriptor() ([]byte, []int) {
	return file_enum_server_proto_rawDescGZIP(), []int{1}
}

type AvatarType int32

const (
	AvatarType_AVATAR_TYPE_NONE   AvatarType = 0
	AvatarType_AVATAR_TRIAL_TYPE  AvatarType = 1
	AvatarType_AVATAR_LIMIT_TYPE  AvatarType = 2
	AvatarType_AVATAR_FORMAL_TYPE AvatarType = 3
	AvatarType_AVATAR_ASSIST_TYPE AvatarType = 4
)

// Enum value maps for AvatarType.
var (
	AvatarType_name = map[int32]string{
		0: "AVATAR_TYPE_NONE",
		1: "AVATAR_TRIAL_TYPE",
		2: "AVATAR_LIMIT_TYPE",
		3: "AVATAR_FORMAL_TYPE",
		4: "AVATAR_ASSIST_TYPE",
	}
	AvatarType_value = map[string]int32{
		"AVATAR_TYPE_NONE":   0,
		"AVATAR_TRIAL_TYPE":  1,
		"AVATAR_LIMIT_TYPE":  2,
		"AVATAR_FORMAL_TYPE": 3,
		"AVATAR_ASSIST_TYPE": 4,
	}
)

func (x AvatarType) Enum() *AvatarType {
	p := new(AvatarType)
	*p = x
	return p
}

func (x AvatarType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AvatarType) Descriptor() protoreflect.EnumDescriptor {
	return file_enum_server_proto_enumTypes[2].Descriptor()
}

func (AvatarType) Type() protoreflect.EnumType {
	return &file_enum_server_proto_enumTypes[2]
}

func (x AvatarType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AvatarType.Descriptor instead.
func (AvatarType) EnumDescriptor() ([]byte, []int) {
	return file_enum_server_proto_rawDescGZIP(), []int{2}
}

type ExtraLineupType int32

const (
	ExtraLineupType_LINEUP_NONE             ExtraLineupType = 0
	ExtraLineupType_LINEUP_CHALLENGE        ExtraLineupType = 1
	ExtraLineupType_LINEUP_ROGUE            ExtraLineupType = 2
	ExtraLineupType_LINEUP_CHALLENGE_2      ExtraLineupType = 3
	ExtraLineupType_LINEUP_CHALLENGE_3      ExtraLineupType = 4
	ExtraLineupType_LINEUP_ROGUE_CHALLENGE  ExtraLineupType = 5
	ExtraLineupType_LINEUP_STAGE_TRIAL      ExtraLineupType = 6
	ExtraLineupType_LINEUP_ROGUE_TRIAL      ExtraLineupType = 7
	ExtraLineupType_LINEUP_ACTIVITY         ExtraLineupType = 8
	ExtraLineupType_LINEUP_BOXING_CLUB      ExtraLineupType = 9
	ExtraLineupType_LINEUP_TREASURE_DUNGEON ExtraLineupType = 11
	ExtraLineupType_LINEUP_CHESS_ROGUE      ExtraLineupType = 12
	ExtraLineupType_LINEUP_HELIOBUS         ExtraLineupType = 13
)

// Enum value maps for ExtraLineupType.
var (
	ExtraLineupType_name = map[int32]string{
		0:  "LINEUP_NONE",
		1:  "LINEUP_CHALLENGE",
		2:  "LINEUP_ROGUE",
		3:  "LINEUP_CHALLENGE_2",
		4:  "LINEUP_CHALLENGE_3",
		5:  "LINEUP_ROGUE_CHALLENGE",
		6:  "LINEUP_STAGE_TRIAL",
		7:  "LINEUP_ROGUE_TRIAL",
		8:  "LINEUP_ACTIVITY",
		9:  "LINEUP_BOXING_CLUB",
		11: "LINEUP_TREASURE_DUNGEON",
		12: "LINEUP_CHESS_ROGUE",
		13: "LINEUP_HELIOBUS",
	}
	ExtraLineupType_value = map[string]int32{
		"LINEUP_NONE":             0,
		"LINEUP_CHALLENGE":        1,
		"LINEUP_ROGUE":            2,
		"LINEUP_CHALLENGE_2":      3,
		"LINEUP_CHALLENGE_3":      4,
		"LINEUP_ROGUE_CHALLENGE":  5,
		"LINEUP_STAGE_TRIAL":      6,
		"LINEUP_ROGUE_TRIAL":      7,
		"LINEUP_ACTIVITY":         8,
		"LINEUP_BOXING_CLUB":      9,
		"LINEUP_TREASURE_DUNGEON": 11,
		"LINEUP_CHESS_ROGUE":      12,
		"LINEUP_HELIOBUS":         13,
	}
)

func (x ExtraLineupType) Enum() *ExtraLineupType {
	p := new(ExtraLineupType)
	*p = x
	return p
}

func (x ExtraLineupType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ExtraLineupType) Descriptor() protoreflect.EnumDescriptor {
	return file_enum_server_proto_enumTypes[3].Descriptor()
}

func (ExtraLineupType) Type() protoreflect.EnumType {
	return &file_enum_server_proto_enumTypes[3]
}

func (x ExtraLineupType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ExtraLineupType.Descriptor instead.
func (ExtraLineupType) EnumDescriptor() ([]byte, []int) {
	return file_enum_server_proto_rawDescGZIP(), []int{3}
}

type BattleType int32

const (
	BattleType_Battle_NONE            BattleType = 0
	BattleType_Battle_CHALLENGE       BattleType = 1
	BattleType_Battle_CHALLENGE_Story BattleType = 2
	BattleType_Battle_ROGUE           BattleType = 3
	BattleType_Battle_TrialActivity   BattleType = 4
)

// Enum value maps for BattleType.
var (
	BattleType_name = map[int32]string{
		0: "Battle_NONE",
		1: "Battle_CHALLENGE",
		2: "Battle_CHALLENGE_Story",
		3: "Battle_ROGUE",
		4: "Battle_TrialActivity",
	}
	BattleType_value = map[string]int32{
		"Battle_NONE":            0,
		"Battle_CHALLENGE":       1,
		"Battle_CHALLENGE_Story": 2,
		"Battle_ROGUE":           3,
		"Battle_TrialActivity":   4,
	}
)

func (x BattleType) Enum() *BattleType {
	p := new(BattleType)
	*p = x
	return p
}

func (x BattleType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (BattleType) Descriptor() protoreflect.EnumDescriptor {
	return file_enum_server_proto_enumTypes[4].Descriptor()
}

func (BattleType) Type() protoreflect.EnumType {
	return &file_enum_server_proto_enumTypes[4]
}

func (x BattleType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use BattleType.Descriptor instead.
func (BattleType) EnumDescriptor() ([]byte, []int) {
	return file_enum_server_proto_rawDescGZIP(), []int{4}
}

type LanguageType int32

const (
	LanguageType_LANGUAGE_NONE LanguageType = 0
	LanguageType_LANGUAGE_SC   LanguageType = 1
	LanguageType_LANGUAGE_TC   LanguageType = 2
	LanguageType_LANGUAGE_EN   LanguageType = 3
	LanguageType_LANGUAGE_KR   LanguageType = 4
	LanguageType_LANGUAGE_JP   LanguageType = 5
	LanguageType_LANGUAGE_FR   LanguageType = 6
	LanguageType_LANGUAGE_DE   LanguageType = 7
	LanguageType_LANGUAGE_ES   LanguageType = 8
	LanguageType_LANGUAGE_PT   LanguageType = 9
	LanguageType_LANGUAGE_RU   LanguageType = 10
	LanguageType_LANGUAGE_TH   LanguageType = 11
	LanguageType_LANGUAGE_VI   LanguageType = 12
	LanguageType_LANGUAGE_ID   LanguageType = 13
)

// Enum value maps for LanguageType.
var (
	LanguageType_name = map[int32]string{
		0:  "LANGUAGE_NONE",
		1:  "LANGUAGE_SC",
		2:  "LANGUAGE_TC",
		3:  "LANGUAGE_EN",
		4:  "LANGUAGE_KR",
		5:  "LANGUAGE_JP",
		6:  "LANGUAGE_FR",
		7:  "LANGUAGE_DE",
		8:  "LANGUAGE_ES",
		9:  "LANGUAGE_PT",
		10: "LANGUAGE_RU",
		11: "LANGUAGE_TH",
		12: "LANGUAGE_VI",
		13: "LANGUAGE_ID",
	}
	LanguageType_value = map[string]int32{
		"LANGUAGE_NONE": 0,
		"LANGUAGE_SC":   1,
		"LANGUAGE_TC":   2,
		"LANGUAGE_EN":   3,
		"LANGUAGE_KR":   4,
		"LANGUAGE_JP":   5,
		"LANGUAGE_FR":   6,
		"LANGUAGE_DE":   7,
		"LANGUAGE_ES":   8,
		"LANGUAGE_PT":   9,
		"LANGUAGE_RU":   10,
		"LANGUAGE_TH":   11,
		"LANGUAGE_VI":   12,
		"LANGUAGE_ID":   13,
	}
)

func (x LanguageType) Enum() *LanguageType {
	p := new(LanguageType)
	*p = x
	return p
}

func (x LanguageType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (LanguageType) Descriptor() protoreflect.EnumDescriptor {
	return file_enum_server_proto_enumTypes[5].Descriptor()
}

func (LanguageType) Type() protoreflect.EnumType {
	return &file_enum_server_proto_enumTypes[5]
}

func (x LanguageType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use LanguageType.Descriptor instead.
func (LanguageType) EnumDescriptor() ([]byte, []int) {
	return file_enum_server_proto_rawDescGZIP(), []int{5}
}

type PlatformType int32

const (
	PlatformType_EDITOR             PlatformType = 0
	PlatformType_IOS                PlatformType = 1
	PlatformType_ANDROID            PlatformType = 2
	PlatformType_PC                 PlatformType = 3
	PlatformType_WEB                PlatformType = 4
	PlatformType_WAP                PlatformType = 5
	PlatformType_PS4                PlatformType = 6
	PlatformType_NINTENDO           PlatformType = 7
	PlatformType_CLOUD_ANDROID      PlatformType = 8
	PlatformType_CLOUD_PC           PlatformType = 9
	PlatformType_CLOUD_IOS          PlatformType = 10
	PlatformType_PS5                PlatformType = 11
	PlatformType_MAC                PlatformType = 12
	PlatformType_CLOUD_MAC          PlatformType = 13
	PlatformType_CLOUD_WEB_ANDROID  PlatformType = 20
	PlatformType_CLOUD_WEB_IOS      PlatformType = 21
	PlatformType_CLOUD_WEB_PC       PlatformType = 22
	PlatformType_CLOUD_WEB_MAC      PlatformType = 23
	PlatformType_CLOUD_WEB_TOUCH    PlatformType = 24
	PlatformType_CLOUD_WEB_KEYBOARD PlatformType = 25
)

// Enum value maps for PlatformType.
var (
	PlatformType_name = map[int32]string{
		0:  "EDITOR",
		1:  "IOS",
		2:  "ANDROID",
		3:  "PC",
		4:  "WEB",
		5:  "WAP",
		6:  "PS4",
		7:  "NINTENDO",
		8:  "CLOUD_ANDROID",
		9:  "CLOUD_PC",
		10: "CLOUD_IOS",
		11: "PS5",
		12: "MAC",
		13: "CLOUD_MAC",
		20: "CLOUD_WEB_ANDROID",
		21: "CLOUD_WEB_IOS",
		22: "CLOUD_WEB_PC",
		23: "CLOUD_WEB_MAC",
		24: "CLOUD_WEB_TOUCH",
		25: "CLOUD_WEB_KEYBOARD",
	}
	PlatformType_value = map[string]int32{
		"EDITOR":             0,
		"IOS":                1,
		"ANDROID":            2,
		"PC":                 3,
		"WEB":                4,
		"WAP":                5,
		"PS4":                6,
		"NINTENDO":           7,
		"CLOUD_ANDROID":      8,
		"CLOUD_PC":           9,
		"CLOUD_IOS":          10,
		"PS5":                11,
		"MAC":                12,
		"CLOUD_MAC":          13,
		"CLOUD_WEB_ANDROID":  20,
		"CLOUD_WEB_IOS":      21,
		"CLOUD_WEB_PC":       22,
		"CLOUD_WEB_MAC":      23,
		"CLOUD_WEB_TOUCH":    24,
		"CLOUD_WEB_KEYBOARD": 25,
	}
)

func (x PlatformType) Enum() *PlatformType {
	p := new(PlatformType)
	*p = x
	return p
}

func (x PlatformType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PlatformType) Descriptor() protoreflect.EnumDescriptor {
	return file_enum_server_proto_enumTypes[6].Descriptor()
}

func (PlatformType) Type() protoreflect.EnumType {
	return &file_enum_server_proto_enumTypes[6]
}

func (x PlatformType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PlatformType.Descriptor instead.
func (PlatformType) EnumDescriptor() ([]byte, []int) {
	return file_enum_server_proto_rawDescGZIP(), []int{6}
}

type ServerType int32

const (
	ServerType_SERVICE_NONE     ServerType = 0
	ServerType_SERVICE_NODE     ServerType = 1
	ServerType_SERVICE_GAME     ServerType = 2
	ServerType_SERVICE_GATE     ServerType = 3
	ServerType_SERVICE_DISPATCH ServerType = 4
	ServerType_SERVICE_MULTI    ServerType = 5
	ServerType_SERVICE_GM       ServerType = 6
)

// Enum value maps for ServerType.
var (
	ServerType_name = map[int32]string{
		0: "SERVICE_NONE",
		1: "SERVICE_NODE",
		2: "SERVICE_GAME",
		3: "SERVICE_GATE",
		4: "SERVICE_DISPATCH",
		5: "SERVICE_MULTI",
		6: "SERVICE_GM",
	}
	ServerType_value = map[string]int32{
		"SERVICE_NONE":     0,
		"SERVICE_NODE":     1,
		"SERVICE_GAME":     2,
		"SERVICE_GATE":     3,
		"SERVICE_DISPATCH": 4,
		"SERVICE_MULTI":    5,
		"SERVICE_GM":       6,
	}
)

func (x ServerType) Enum() *ServerType {
	p := new(ServerType)
	*p = x
	return p
}

func (x ServerType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ServerType) Descriptor() protoreflect.EnumDescriptor {
	return file_enum_server_proto_enumTypes[7].Descriptor()
}

func (ServerType) Type() protoreflect.EnumType {
	return &file_enum_server_proto_enumTypes[7]
}

func (x ServerType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ServerType.Descriptor instead.
func (ServerType) EnumDescriptor() ([]byte, []int) {
	return file_enum_server_proto_rawDescGZIP(), []int{7}
}

var File_enum_server_proto protoreflect.FileDescriptor

var file_enum_server_proto_rawDesc = []byte{
	0x0a, 0x11, 0x65, 0x6e, 0x75, 0x6d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2a, 0x38, 0x0a, 0x06, 0x47, 0x65,
	0x6e, 0x64, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x0a, 0x47, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x4e, 0x6f,
	0x6e, 0x65, 0x10, 0x00, 0x12, 0x0d, 0x0a, 0x09, 0x47, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x4d, 0x61,
	0x6e, 0x10, 0x01, 0x12, 0x0f, 0x0a, 0x0b, 0x47, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x57, 0x6f, 0x6d,
	0x61, 0x6e, 0x10, 0x02, 0x2a, 0xfe, 0x01, 0x0a, 0x0d, 0x48, 0x65, 0x72, 0x6f, 0x42, 0x61, 0x73,
	0x69, 0x63, 0x54, 0x79, 0x70, 0x65, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x6f, 0x6e, 0x65, 0x10, 0x00,
	0x12, 0x0f, 0x0a, 0x0a, 0x42, 0x6f, 0x79, 0x57, 0x61, 0x72, 0x72, 0x69, 0x6f, 0x72, 0x10, 0xc1,
	0x3e, 0x12, 0x10, 0x0a, 0x0b, 0x47, 0x69, 0x72, 0x6c, 0x57, 0x61, 0x72, 0x72, 0x69, 0x6f, 0x72,
	0x10, 0xc2, 0x3e, 0x12, 0x0e, 0x0a, 0x09, 0x42, 0x6f, 0x79, 0x4b, 0x6e, 0x69, 0x67, 0x68, 0x74,
	0x10, 0xc3, 0x3e, 0x12, 0x0f, 0x0a, 0x0a, 0x47, 0x69, 0x72, 0x6c, 0x4b, 0x6e, 0x69, 0x67, 0x68,
	0x74, 0x10, 0xc4, 0x3e, 0x12, 0x0d, 0x0a, 0x08, 0x42, 0x6f, 0x79, 0x52, 0x6f, 0x67, 0x75, 0x65,
	0x10, 0xc5, 0x3e, 0x12, 0x0e, 0x0a, 0x09, 0x47, 0x69, 0x72, 0x6c, 0x52, 0x6f, 0x67, 0x75, 0x65,
	0x10, 0xc6, 0x3e, 0x12, 0x0c, 0x0a, 0x07, 0x42, 0x6f, 0x79, 0x4d, 0x61, 0x67, 0x65, 0x10, 0xc7,
	0x3e, 0x12, 0x0d, 0x0a, 0x08, 0x47, 0x69, 0x72, 0x6c, 0x4d, 0x61, 0x67, 0x65, 0x10, 0xc8, 0x3e,
	0x12, 0x0e, 0x0a, 0x09, 0x42, 0x6f, 0x79, 0x53, 0x68, 0x61, 0x6d, 0x61, 0x6e, 0x10, 0xc9, 0x3e,
	0x12, 0x0f, 0x0a, 0x0a, 0x47, 0x69, 0x72, 0x6c, 0x53, 0x68, 0x61, 0x6d, 0x61, 0x6e, 0x10, 0xca,
	0x3e, 0x12, 0x0f, 0x0a, 0x0a, 0x42, 0x6f, 0x79, 0x57, 0x61, 0x72, 0x6c, 0x6f, 0x63, 0x6b, 0x10,
	0xcb, 0x3e, 0x12, 0x10, 0x0a, 0x0b, 0x47, 0x69, 0x72, 0x6c, 0x57, 0x61, 0x72, 0x6c, 0x6f, 0x63,
	0x6b, 0x10, 0xcc, 0x3e, 0x12, 0x0e, 0x0a, 0x09, 0x42, 0x6f, 0x79, 0x50, 0x72, 0x69, 0x65, 0x73,
	0x74, 0x10, 0xcd, 0x3e, 0x12, 0x0f, 0x0a, 0x0a, 0x47, 0x69, 0x72, 0x6c, 0x50, 0x72, 0x69, 0x65,
	0x73, 0x74, 0x10, 0xce, 0x3e, 0x2a, 0x80, 0x01, 0x0a, 0x0a, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x10, 0x41, 0x56, 0x41, 0x54, 0x41, 0x52, 0x5f, 0x54,
	0x59, 0x50, 0x45, 0x5f, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x15, 0x0a, 0x11, 0x41, 0x56,
	0x41, 0x54, 0x41, 0x52, 0x5f, 0x54, 0x52, 0x49, 0x41, 0x4c, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x10,
	0x01, 0x12, 0x15, 0x0a, 0x11, 0x41, 0x56, 0x41, 0x54, 0x41, 0x52, 0x5f, 0x4c, 0x49, 0x4d, 0x49,
	0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x10, 0x02, 0x12, 0x16, 0x0a, 0x12, 0x41, 0x56, 0x41, 0x54,
	0x41, 0x52, 0x5f, 0x46, 0x4f, 0x52, 0x4d, 0x41, 0x4c, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x10, 0x03,
	0x12, 0x16, 0x0a, 0x12, 0x41, 0x56, 0x41, 0x54, 0x41, 0x52, 0x5f, 0x41, 0x53, 0x53, 0x49, 0x53,
	0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x10, 0x04, 0x2a, 0xbd, 0x02, 0x0a, 0x0f, 0x45, 0x78, 0x74,
	0x72, 0x61, 0x4c, 0x69, 0x6e, 0x65, 0x75, 0x70, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0f, 0x0a, 0x0b,
	0x4c, 0x49, 0x4e, 0x45, 0x55, 0x50, 0x5f, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x14, 0x0a,
	0x10, 0x4c, 0x49, 0x4e, 0x45, 0x55, 0x50, 0x5f, 0x43, 0x48, 0x41, 0x4c, 0x4c, 0x45, 0x4e, 0x47,
	0x45, 0x10, 0x01, 0x12, 0x10, 0x0a, 0x0c, 0x4c, 0x49, 0x4e, 0x45, 0x55, 0x50, 0x5f, 0x52, 0x4f,
	0x47, 0x55, 0x45, 0x10, 0x02, 0x12, 0x16, 0x0a, 0x12, 0x4c, 0x49, 0x4e, 0x45, 0x55, 0x50, 0x5f,
	0x43, 0x48, 0x41, 0x4c, 0x4c, 0x45, 0x4e, 0x47, 0x45, 0x5f, 0x32, 0x10, 0x03, 0x12, 0x16, 0x0a,
	0x12, 0x4c, 0x49, 0x4e, 0x45, 0x55, 0x50, 0x5f, 0x43, 0x48, 0x41, 0x4c, 0x4c, 0x45, 0x4e, 0x47,
	0x45, 0x5f, 0x33, 0x10, 0x04, 0x12, 0x1a, 0x0a, 0x16, 0x4c, 0x49, 0x4e, 0x45, 0x55, 0x50, 0x5f,
	0x52, 0x4f, 0x47, 0x55, 0x45, 0x5f, 0x43, 0x48, 0x41, 0x4c, 0x4c, 0x45, 0x4e, 0x47, 0x45, 0x10,
	0x05, 0x12, 0x16, 0x0a, 0x12, 0x4c, 0x49, 0x4e, 0x45, 0x55, 0x50, 0x5f, 0x53, 0x54, 0x41, 0x47,
	0x45, 0x5f, 0x54, 0x52, 0x49, 0x41, 0x4c, 0x10, 0x06, 0x12, 0x16, 0x0a, 0x12, 0x4c, 0x49, 0x4e,
	0x45, 0x55, 0x50, 0x5f, 0x52, 0x4f, 0x47, 0x55, 0x45, 0x5f, 0x54, 0x52, 0x49, 0x41, 0x4c, 0x10,
	0x07, 0x12, 0x13, 0x0a, 0x0f, 0x4c, 0x49, 0x4e, 0x45, 0x55, 0x50, 0x5f, 0x41, 0x43, 0x54, 0x49,
	0x56, 0x49, 0x54, 0x59, 0x10, 0x08, 0x12, 0x16, 0x0a, 0x12, 0x4c, 0x49, 0x4e, 0x45, 0x55, 0x50,
	0x5f, 0x42, 0x4f, 0x58, 0x49, 0x4e, 0x47, 0x5f, 0x43, 0x4c, 0x55, 0x42, 0x10, 0x09, 0x12, 0x1b,
	0x0a, 0x17, 0x4c, 0x49, 0x4e, 0x45, 0x55, 0x50, 0x5f, 0x54, 0x52, 0x45, 0x41, 0x53, 0x55, 0x52,
	0x45, 0x5f, 0x44, 0x55, 0x4e, 0x47, 0x45, 0x4f, 0x4e, 0x10, 0x0b, 0x12, 0x16, 0x0a, 0x12, 0x4c,
	0x49, 0x4e, 0x45, 0x55, 0x50, 0x5f, 0x43, 0x48, 0x45, 0x53, 0x53, 0x5f, 0x52, 0x4f, 0x47, 0x55,
	0x45, 0x10, 0x0c, 0x12, 0x13, 0x0a, 0x0f, 0x4c, 0x49, 0x4e, 0x45, 0x55, 0x50, 0x5f, 0x48, 0x45,
	0x4c, 0x49, 0x4f, 0x42, 0x55, 0x53, 0x10, 0x0d, 0x2a, 0x7b, 0x0a, 0x0a, 0x42, 0x61, 0x74, 0x74,
	0x6c, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0f, 0x0a, 0x0b, 0x42, 0x61, 0x74, 0x74, 0x6c, 0x65,
	0x5f, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x14, 0x0a, 0x10, 0x42, 0x61, 0x74, 0x74, 0x6c,
	0x65, 0x5f, 0x43, 0x48, 0x41, 0x4c, 0x4c, 0x45, 0x4e, 0x47, 0x45, 0x10, 0x01, 0x12, 0x1a, 0x0a,
	0x16, 0x42, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x5f, 0x43, 0x48, 0x41, 0x4c, 0x4c, 0x45, 0x4e, 0x47,
	0x45, 0x5f, 0x53, 0x74, 0x6f, 0x72, 0x79, 0x10, 0x02, 0x12, 0x10, 0x0a, 0x0c, 0x42, 0x61, 0x74,
	0x74, 0x6c, 0x65, 0x5f, 0x52, 0x4f, 0x47, 0x55, 0x45, 0x10, 0x03, 0x12, 0x18, 0x0a, 0x14, 0x42,
	0x61, 0x74, 0x74, 0x6c, 0x65, 0x5f, 0x54, 0x72, 0x69, 0x61, 0x6c, 0x41, 0x63, 0x74, 0x69, 0x76,
	0x69, 0x74, 0x79, 0x10, 0x04, 0x2a, 0xfe, 0x01, 0x0a, 0x0c, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61,
	0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x11, 0x0a, 0x0d, 0x4c, 0x41, 0x4e, 0x47, 0x55, 0x41,
	0x47, 0x45, 0x5f, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x0f, 0x0a, 0x0b, 0x4c, 0x41, 0x4e,
	0x47, 0x55, 0x41, 0x47, 0x45, 0x5f, 0x53, 0x43, 0x10, 0x01, 0x12, 0x0f, 0x0a, 0x0b, 0x4c, 0x41,
	0x4e, 0x47, 0x55, 0x41, 0x47, 0x45, 0x5f, 0x54, 0x43, 0x10, 0x02, 0x12, 0x0f, 0x0a, 0x0b, 0x4c,
	0x41, 0x4e, 0x47, 0x55, 0x41, 0x47, 0x45, 0x5f, 0x45, 0x4e, 0x10, 0x03, 0x12, 0x0f, 0x0a, 0x0b,
	0x4c, 0x41, 0x4e, 0x47, 0x55, 0x41, 0x47, 0x45, 0x5f, 0x4b, 0x52, 0x10, 0x04, 0x12, 0x0f, 0x0a,
	0x0b, 0x4c, 0x41, 0x4e, 0x47, 0x55, 0x41, 0x47, 0x45, 0x5f, 0x4a, 0x50, 0x10, 0x05, 0x12, 0x0f,
	0x0a, 0x0b, 0x4c, 0x41, 0x4e, 0x47, 0x55, 0x41, 0x47, 0x45, 0x5f, 0x46, 0x52, 0x10, 0x06, 0x12,
	0x0f, 0x0a, 0x0b, 0x4c, 0x41, 0x4e, 0x47, 0x55, 0x41, 0x47, 0x45, 0x5f, 0x44, 0x45, 0x10, 0x07,
	0x12, 0x0f, 0x0a, 0x0b, 0x4c, 0x41, 0x4e, 0x47, 0x55, 0x41, 0x47, 0x45, 0x5f, 0x45, 0x53, 0x10,
	0x08, 0x12, 0x0f, 0x0a, 0x0b, 0x4c, 0x41, 0x4e, 0x47, 0x55, 0x41, 0x47, 0x45, 0x5f, 0x50, 0x54,
	0x10, 0x09, 0x12, 0x0f, 0x0a, 0x0b, 0x4c, 0x41, 0x4e, 0x47, 0x55, 0x41, 0x47, 0x45, 0x5f, 0x52,
	0x55, 0x10, 0x0a, 0x12, 0x0f, 0x0a, 0x0b, 0x4c, 0x41, 0x4e, 0x47, 0x55, 0x41, 0x47, 0x45, 0x5f,
	0x54, 0x48, 0x10, 0x0b, 0x12, 0x0f, 0x0a, 0x0b, 0x4c, 0x41, 0x4e, 0x47, 0x55, 0x41, 0x47, 0x45,
	0x5f, 0x56, 0x49, 0x10, 0x0c, 0x12, 0x0f, 0x0a, 0x0b, 0x4c, 0x41, 0x4e, 0x47, 0x55, 0x41, 0x47,
	0x45, 0x5f, 0x49, 0x44, 0x10, 0x0d, 0x2a, 0xae, 0x02, 0x0a, 0x0c, 0x50, 0x6c, 0x61, 0x74, 0x66,
	0x6f, 0x72, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0a, 0x0a, 0x06, 0x45, 0x44, 0x49, 0x54, 0x4f,
	0x52, 0x10, 0x00, 0x12, 0x07, 0x0a, 0x03, 0x49, 0x4f, 0x53, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07,
	0x41, 0x4e, 0x44, 0x52, 0x4f, 0x49, 0x44, 0x10, 0x02, 0x12, 0x06, 0x0a, 0x02, 0x50, 0x43, 0x10,
	0x03, 0x12, 0x07, 0x0a, 0x03, 0x57, 0x45, 0x42, 0x10, 0x04, 0x12, 0x07, 0x0a, 0x03, 0x57, 0x41,
	0x50, 0x10, 0x05, 0x12, 0x07, 0x0a, 0x03, 0x50, 0x53, 0x34, 0x10, 0x06, 0x12, 0x0c, 0x0a, 0x08,
	0x4e, 0x49, 0x4e, 0x54, 0x45, 0x4e, 0x44, 0x4f, 0x10, 0x07, 0x12, 0x11, 0x0a, 0x0d, 0x43, 0x4c,
	0x4f, 0x55, 0x44, 0x5f, 0x41, 0x4e, 0x44, 0x52, 0x4f, 0x49, 0x44, 0x10, 0x08, 0x12, 0x0c, 0x0a,
	0x08, 0x43, 0x4c, 0x4f, 0x55, 0x44, 0x5f, 0x50, 0x43, 0x10, 0x09, 0x12, 0x0d, 0x0a, 0x09, 0x43,
	0x4c, 0x4f, 0x55, 0x44, 0x5f, 0x49, 0x4f, 0x53, 0x10, 0x0a, 0x12, 0x07, 0x0a, 0x03, 0x50, 0x53,
	0x35, 0x10, 0x0b, 0x12, 0x07, 0x0a, 0x03, 0x4d, 0x41, 0x43, 0x10, 0x0c, 0x12, 0x0d, 0x0a, 0x09,
	0x43, 0x4c, 0x4f, 0x55, 0x44, 0x5f, 0x4d, 0x41, 0x43, 0x10, 0x0d, 0x12, 0x15, 0x0a, 0x11, 0x43,
	0x4c, 0x4f, 0x55, 0x44, 0x5f, 0x57, 0x45, 0x42, 0x5f, 0x41, 0x4e, 0x44, 0x52, 0x4f, 0x49, 0x44,
	0x10, 0x14, 0x12, 0x11, 0x0a, 0x0d, 0x43, 0x4c, 0x4f, 0x55, 0x44, 0x5f, 0x57, 0x45, 0x42, 0x5f,
	0x49, 0x4f, 0x53, 0x10, 0x15, 0x12, 0x10, 0x0a, 0x0c, 0x43, 0x4c, 0x4f, 0x55, 0x44, 0x5f, 0x57,
	0x45, 0x42, 0x5f, 0x50, 0x43, 0x10, 0x16, 0x12, 0x11, 0x0a, 0x0d, 0x43, 0x4c, 0x4f, 0x55, 0x44,
	0x5f, 0x57, 0x45, 0x42, 0x5f, 0x4d, 0x41, 0x43, 0x10, 0x17, 0x12, 0x13, 0x0a, 0x0f, 0x43, 0x4c,
	0x4f, 0x55, 0x44, 0x5f, 0x57, 0x45, 0x42, 0x5f, 0x54, 0x4f, 0x55, 0x43, 0x48, 0x10, 0x18, 0x12,
	0x16, 0x0a, 0x12, 0x43, 0x4c, 0x4f, 0x55, 0x44, 0x5f, 0x57, 0x45, 0x42, 0x5f, 0x4b, 0x45, 0x59,
	0x42, 0x4f, 0x41, 0x52, 0x44, 0x10, 0x19, 0x2a, 0x8d, 0x01, 0x0a, 0x0a, 0x53, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x12, 0x10, 0x0a, 0x0c, 0x53, 0x45, 0x52, 0x56, 0x49, 0x43,
	0x45, 0x5f, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x10, 0x0a, 0x0c, 0x53, 0x45, 0x52, 0x56,
	0x49, 0x43, 0x45, 0x5f, 0x4e, 0x4f, 0x44, 0x45, 0x10, 0x01, 0x12, 0x10, 0x0a, 0x0c, 0x53, 0x45,
	0x52, 0x56, 0x49, 0x43, 0x45, 0x5f, 0x47, 0x41, 0x4d, 0x45, 0x10, 0x02, 0x12, 0x10, 0x0a, 0x0c,
	0x53, 0x45, 0x52, 0x56, 0x49, 0x43, 0x45, 0x5f, 0x47, 0x41, 0x54, 0x45, 0x10, 0x03, 0x12, 0x14,
	0x0a, 0x10, 0x53, 0x45, 0x52, 0x56, 0x49, 0x43, 0x45, 0x5f, 0x44, 0x49, 0x53, 0x50, 0x41, 0x54,
	0x43, 0x48, 0x10, 0x04, 0x12, 0x11, 0x0a, 0x0d, 0x53, 0x45, 0x52, 0x56, 0x49, 0x43, 0x45, 0x5f,
	0x4d, 0x55, 0x4c, 0x54, 0x49, 0x10, 0x05, 0x12, 0x0e, 0x0a, 0x0a, 0x53, 0x45, 0x52, 0x56, 0x49,
	0x43, 0x45, 0x5f, 0x47, 0x4d, 0x10, 0x06, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_enum_server_proto_rawDescOnce sync.Once
	file_enum_server_proto_rawDescData = file_enum_server_proto_rawDesc
)

func file_enum_server_proto_rawDescGZIP() []byte {
	file_enum_server_proto_rawDescOnce.Do(func() {
		file_enum_server_proto_rawDescData = protoimpl.X.CompressGZIP(file_enum_server_proto_rawDescData)
	})
	return file_enum_server_proto_rawDescData
}

var file_enum_server_proto_enumTypes = make([]protoimpl.EnumInfo, 8)
var file_enum_server_proto_goTypes = []interface{}{
	(Gender)(0),          // 0: proto.Gender
	(HeroBasicType)(0),   // 1: proto.HeroBasicType
	(AvatarType)(0),      // 2: proto.AvatarType
	(ExtraLineupType)(0), // 3: proto.ExtraLineupType
	(BattleType)(0),      // 4: proto.BattleType
	(LanguageType)(0),    // 5: proto.LanguageType
	(PlatformType)(0),    // 6: proto.PlatformType
	(ServerType)(0),      // 7: proto.ServerType
}
var file_enum_server_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_enum_server_proto_init() }
func file_enum_server_proto_init() {
	if File_enum_server_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_enum_server_proto_rawDesc,
			NumEnums:      8,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_enum_server_proto_goTypes,
		DependencyIndexes: file_enum_server_proto_depIdxs,
		EnumInfos:         file_enum_server_proto_enumTypes,
	}.Build()
	File_enum_server_proto = out.File
	file_enum_server_proto_rawDesc = nil
	file_enum_server_proto_goTypes = nil
	file_enum_server_proto_depIdxs = nil
}
