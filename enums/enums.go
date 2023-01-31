package enums

// Enum for the hl parameter
type HlType int

const (
	enUS HlType = iota
	enGB
	enAU
	csCZ
	deDE
	elGR
	esEs
	esMX
	frFR
	huHU
	itIT
	plPL
	ptBR
	roRO
	ruRU
	trTR
	jaJP
	koKR
)

func (hl HlType) String() string {
	return [...]string{"en-US", "en-GB", "en-AU", "cs-CZ", "de-DE", "el-GR", "es-ES", "es-MX", "fr-FR", "hu-HU", "it-IT", "pl-PL", "pt-BR", "ro-RO", "ru-RU", "tr-TR", "ja-JP", "ko-KR"}[hl]
}
