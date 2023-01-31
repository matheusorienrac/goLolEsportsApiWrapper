package models

import "time"

// Enum for the hl parameter
type hlType int

const (
	enUS hlType = iota
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

func (hl hlType) String() string {
	switch hl {
	case enUS:
		return "en-US"
	case enGB:
		return "en-GB"
	case enAU:
		return "en-AU"
	case csCZ:
		return "cs-CZ"
	case deDE:
		return "de-DE"
	case elGR:
		return "el-GR"
	case esEs:
		return "es-ES"
	case esMX:
		return "es-MX"
	case frFR:
		return "fr-FR"
	case huHU:
		return "hu-HU"
	case itIT:
		return "it-IT"
	case plPL:
		return "pl-PL"
	case ptBR:
		return "pt-BR"
	case roRO:
		return "ro-RO"
	case ruRU:
		return "ru-RU"
	case trTR:
		return "tr-TR"
	case jaJP:
		return "ja-JP"
	case koKR:
		return "ko-KR"
	default:
		return "en-US"
	}
	return [...]string{"en-US", "en-GB", "en-AU", "cs-CZ", "de-DE", "el-GR", "es-ES", "es-MX", "fr-FR", "hu-HU", "it-IT", "pl-PL", "pt-BR", "ro-RO", "ru-RU", "tr-TR", "ja-JP", "ko-KR"}[hl]
}

type Live struct {
	Data struct {
		Schedule struct {
			Events []struct {
				ID        string    `json:"id"`
				StartTime time.Time `json:"startTime"`
				State     string    `json:"state"`
				Type      string    `json:"type"`
				BlockName string    `json:"blockName"`
				League    struct {
					ID              string `json:"id"`
					Slug            string `json:"slug"`
					Name            string `json:"name"`
					Image           string `json:"image"`
					Priority        int    `json:"priority"`
					DisplayPriority struct {
						Position int    `json:"position"`
						Status   string `json:"status"`
					} `json:"displayPriority"`
				} `json:"league"`
				Tournament struct {
					ID string `json:"id"`
				} `json:"tournament"`
				Match struct {
					ID    string `json:"id"`
					Teams []struct {
						ID     string `json:"id"`
						Name   string `json:"name"`
						Slug   string `json:"slug"`
						Code   string `json:"code"`
						Image  string `json:"image"`
						Result struct {
							Outcome  interface{} `json:"outcome"`
							GameWins int         `json:"gameWins"`
						} `json:"result"`
						Record struct {
							Wins   int `json:"wins"`
							Losses int `json:"losses"`
						} `json:"record"`
					} `json:"teams"`
					Strategy struct {
						Type  string `json:"type"`
						Count int    `json:"count"`
					} `json:"strategy"`
					Games []struct {
						Number int    `json:"number"`
						ID     string `json:"id"`
						State  string `json:"state"`
						Teams  []struct {
							ID   string `json:"id"`
							Side string `json:"side"`
						} `json:"teams"`
						Vods []interface{} `json:"vods"`
					} `json:"games"`
				} `json:"match"`
				Streams []struct {
					Parameter   string `json:"parameter"`
					Locale      string `json:"locale"`
					MediaLocale struct {
						Locale         string `json:"locale"`
						EnglishName    string `json:"englishName"`
						TranslatedName string `json:"translatedName"`
					} `json:"mediaLocale"`
					Provider    string        `json:"provider"`
					Countries   []interface{} `json:"countries"`
					Offset      int           `json:"offset"`
					StatsStatus string        `json:"statsStatus"`
				} `json:"streams"`
			} `json:"events"`
		} `json:"schedule"`
	} `json:"data"`
}

type Window struct {
	EsportsGameID  string `json:"esportsGameId"`
	EsportsMatchID string `json:"esportsMatchId"`
	GameMetadata   struct {
		PatchVersion     string `json:"patchVersion"`
		BlueTeamMetadata struct {
			EsportsTeamID       string `json:"esportsTeamId"`
			ParticipantMetadata []struct {
				ParticipantID int    `json:"participantId"`
				SummonerName  string `json:"summonerName"`
				ChampionID    string `json:"championId"`
				Role          string `json:"role"`
			} `json:"participantMetadata"`
		} `json:"blueTeamMetadata"`
		RedTeamMetadata struct {
			EsportsTeamID       string `json:"esportsTeamId"`
			ParticipantMetadata []struct {
				ParticipantID int    `json:"participantId"`
				SummonerName  string `json:"summonerName"`
				ChampionID    string `json:"championId"`
				Role          string `json:"role"`
			} `json:"participantMetadata"`
		} `json:"redTeamMetadata"`
	} `json:"gameMetadata"`
	Frames []struct {
		Rfc460Timestamp time.Time `json:"rfc460Timestamp"`
		GameState       string    `json:"gameState"`
		BlueTeam        struct {
			TotalGold    int      `json:"totalGold"`
			Inhibitors   int      `json:"inhibitors"`
			Towers       int      `json:"towers"`
			Barons       int      `json:"barons"`
			TotalKills   int      `json:"totalKills"`
			Dragons      []string `json:"dragons"`
			Participants []struct {
				ParticipantID int `json:"participantId"`
				Level         int `json:"level"`
				Kills         int `json:"kills"`
				Deaths        int `json:"deaths"`
				Assists       int `json:"assists"`
				CreepScore    int `json:"creepScore"`
				TotalGold     int `json:"totalGold"`
				CurrentHealth int `json:"currentHealth"`
				MaxHealth     int `json:"maxHealth"`
			} `json:"participants"`
		} `json:"blueTeam"`
		RedTeam struct {
			TotalGold    int      `json:"totalGold"`
			Inhibitors   int      `json:"inhibitors"`
			Towers       int      `json:"towers"`
			Barons       int      `json:"barons"`
			TotalKills   int      `json:"totalKills"`
			Dragons      []string `json:"dragons"`
			Participants []struct {
				ParticipantID int `json:"participantId"`
				Level         int `json:"level"`
				Kills         int `json:"kills"`
				Deaths        int `json:"deaths"`
				Assists       int `json:"assists"`
				CreepScore    int `json:"creepScore"`
				TotalGold     int `json:"totalGold"`
				CurrentHealth int `json:"currentHealth"`
				MaxHealth     int `json:"maxHealth"`
			} `json:"participants"`
		} `json:"redTeam"`
	} `json:"frames"`
}
type Details struct {
	Frames []struct {
		Rfc460Timestamp time.Time `json:"rfc460Timestamp"`
		Participants    []struct {
			ParticipantID       int   `json:"participantId"`
			Level               int   `json:"level"`
			Kills               int   `json:"kills"`
			Deaths              int   `json:"deaths"`
			Assists             int   `json:"assists"`
			CreepScore          int   `json:"creepScore"`
			TotalGold           int   `json:"totalGold"`
			CurrentHealth       int   `json:"currentHealth"`
			MaxHealth           int   `json:"maxHealth"`
			TotalGoldEarned     int   `json:"totalGoldEarned"`
			KillParticipation   int   `json:"killParticipation"`
			ChampionDamageShare int   `json:"championDamageShare"`
			WardsPlaced         int   `json:"wardsPlaced"`
			WardsDestroyed      int   `json:"wardsDestroyed"`
			AttackDamage        int   `json:"attackDamage"`
			AbilityPower        int   `json:"abilityPower"`
			CriticalChance      int   `json:"criticalChance"`
			AttackSpeed         int   `json:"attackSpeed"`
			LifeSteal           int   `json:"lifeSteal"`
			Armor               int   `json:"armor"`
			MagicResistance     int   `json:"magicResistance"`
			Tenacity            int   `json:"tenacity"`
			Items               []int `json:"items"`
			PerkMetadata        struct {
				StyleID    int   `json:"styleId"`
				SubStyleID int   `json:"subStyleId"`
				Perks      []int `json:"perks"`
			} `json:"perkMetadata"`
			Abilities string `json:"abilities"`
		} `json:"participants"`
	} `json:"frames"`
}
