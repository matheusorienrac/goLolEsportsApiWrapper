package enums

// Enum for the hl parameter
type HlType int

const (
	EnUS HlType = iota
	EnGB
	EnAU
	CsCZ
	DeDE
	ElGR
	EsEs
	EsMX
	FrFR
	HuHU
	ItIT
	PlPL
	PtBR
	RoRO
	RuRU
	TrTR
	JaJP
	KoKR
)

func (hl HlType) String() string {
	return [...]string{"en-US", "en-GB", "en-AU", "cs-CZ", "de-DE", "el-GR", "es-ES", "es-MX", "fr-FR", "hu-HU", "it-IT", "pl-PL", "pt-BR", "ro-RO", "ru-RU", "tr-TR", "ja-JP", "ko-KR"}[hl]
}
