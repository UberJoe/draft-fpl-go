package fpl

import (
	"time"
)

type General struct {
	Elements []struct {
		ID                               int         `json:"id"`
		Assists                          int         `json:"assists"`
		Bonus                            int         `json:"bonus"`
		Bps                              int         `json:"bps"`
		CleanSheets                      int         `json:"clean_sheets"`
		Creativity                       string      `json:"creativity"`
		GoalsConceded                    int         `json:"goals_conceded"`
		GoalsScored                      int         `json:"goals_scored"`
		IctIndex                         string      `json:"ict_index"`
		Influence                        string      `json:"influence"`
		Minutes                          int         `json:"minutes"`
		OwnGoals                         int         `json:"own_goals"`
		PenaltiesMissed                  int         `json:"penalties_missed"`
		PenaltiesSaved                   int         `json:"penalties_saved"`
		RedCards                         int         `json:"red_cards"`
		Saves                            int         `json:"saves"`
		Threat                           string      `json:"threat"`
		YellowCards                      int         `json:"yellow_cards"`
		Starts                           int         `json:"starts"`
		ExpectedGoals                    string      `json:"expected_goals"`
		ExpectedAssists                  string      `json:"expected_assists"`
		ExpectedGoalInvolvements         string      `json:"expected_goal_involvements"`
		ExpectedGoalsConceded            string      `json:"expected_goals_conceded"`
		Added                            time.Time   `json:"added"`
		ChanceOfPlayingNextRound         int         `json:"chance_of_playing_next_round"`
		ChanceOfPlayingThisRound         int         `json:"chance_of_playing_this_round"`
		Code                             int         `json:"code"`
		DraftRank                        int         `json:"draft_rank"`
		DreamteamCount                   int         `json:"dreamteam_count"`
		EpNext                           interface{} `json:"ep_next"`
		EpThis                           interface{} `json:"ep_this"`
		EventPoints                      int         `json:"event_points"`
		FirstName                        string      `json:"first_name"`
		Form                             string      `json:"form"`
		InDreamteam                      bool        `json:"in_dreamteam"`
		News                             string      `json:"news"`
		NewsAdded                        time.Time   `json:"news_added"`
		NewsReturn                       interface{} `json:"news_return"`
		NewsUpdated                      interface{} `json:"news_updated"`
		PointsPerGame                    string      `json:"points_per_game"`
		SecondName                       string      `json:"second_name"`
		SquadNumber                      interface{} `json:"squad_number"`
		Status                           string      `json:"status"`
		TotalPoints                      int         `json:"total_points"`
		WebName                          string      `json:"web_name"`
		InfluenceRank                    int         `json:"influence_rank"`
		InfluenceRankType                int         `json:"influence_rank_type"`
		CreativityRank                   int         `json:"creativity_rank"`
		CreativityRankType               int         `json:"creativity_rank_type"`
		ThreatRank                       int         `json:"threat_rank"`
		ThreatRankType                   int         `json:"threat_rank_type"`
		IctIndexRank                     int         `json:"ict_index_rank"`
		IctIndexRankType                 int         `json:"ict_index_rank_type"`
		FormRank                         interface{} `json:"form_rank"`
		FormRankType                     interface{} `json:"form_rank_type"`
		PointsPerGameRank                interface{} `json:"points_per_game_rank"`
		PointsPerGameRankType            interface{} `json:"points_per_game_rank_type"`
		CornersAndIndirectFreekicksOrder interface{} `json:"corners_and_indirect_freekicks_order"`
		CornersAndIndirectFreekicksText  string      `json:"corners_and_indirect_freekicks_text"`
		DirectFreekicksOrder             interface{} `json:"direct_freekicks_order"`
		DirectFreekicksText              string      `json:"direct_freekicks_text"`
		PenaltiesOrder                   interface{} `json:"penalties_order"`
		PenaltiesText                    string      `json:"penalties_text"`
		ElementType                      int         `json:"element_type"`
		Team                             int         `json:"team"`
	} `json:"elements"`
	ElementTypes []struct {
		ID                int    `json:"id"`
		ElementCount      int    `json:"element_count"`
		SingularName      string `json:"singular_name"`
		SingularNameShort string `json:"singular_name_short"`
		PluralName        string `json:"plural_name"`
		PluralNameShort   string `json:"plural_name_short"`
	} `json:"element_types"`
	ElementStats []struct {
		Name           string      `json:"name"`
		Label          string      `json:"label"`
		Abbreviation   string      `json:"abbreviation"`
		IsMatchStat    bool        `json:"is_match_stat"`
		MatchStatOrder interface{} `json:"match_stat_order"`
		Sort           string      `json:"sort"`
	} `json:"element_stats"`
	Events struct {
		Current int `json:"current"`
		Data    []struct {
			AverageEntryScore   interface{} `json:"average_entry_score"`
			DeadlineTime        time.Time   `json:"deadline_time"`
			ID                  int         `json:"id"`
			Name                string      `json:"name"`
			Finished            bool        `json:"finished"`
			HighestScoringEntry interface{} `json:"highest_scoring_entry"`
			TradesTime          time.Time   `json:"trades_time"`
			WaiversTime         time.Time   `json:"waivers_time"`
		} `json:"data"`
		Next int `json:"next"`
	} `json:"events"`
	Fixtures map[int][]struct {
		ID                   int         `json:"id"`
		Started              bool        `json:"started"`
		Code                 int         `json:"code"`
		Finished             bool        `json:"finished"`
		FinishedProvisional  bool        `json:"finished_provisional"`
		KickoffTime          time.Time   `json:"kickoff_time"`
		Minutes              int         `json:"minutes"`
		ProvisionalStartTime bool        `json:"provisional_start_time"`
		TeamAScore           interface{} `json:"team_a_score"`
		TeamHScore           interface{} `json:"team_h_score"`
		PulseID              int         `json:"pulse_id"`
		Event                int         `json:"event"`
		TeamA                int         `json:"team_a"`
		TeamH                int         `json:"team_h"`
	} `json:"fixtures"`
	Settings SettingsResponse `json:"settings"`
	Teams    []struct {
		Code      int    `json:"code"`
		ID        int    `json:"id"`
		Name      string `json:"name"`
		PulseID   int    `json:"pulse_id"`
		ShortName string `json:"short_name"`
	} `json:"teams"`
}

type TeamResponse struct {
	Code      int    `json:"code"`
	ID        int    `json:"id"`
	Name      string `json:"name"`
	PulseID   int    `json:"pulse_id"`
	ShortName string `json:"short_name"`
}

type EventsResponse struct {
	Current int `json:"current"`
	Data    []struct {
		AverageEntryScore   interface{} `json:"average_entry_score"`
		DeadlineTime        time.Time   `json:"deadline_time"`
		ID                  int         `json:"id"`
		Name                string      `json:"name"`
		Finished            bool        `json:"finished"`
		HighestScoringEntry interface{} `json:"highest_scoring_entry"`
		TradesTime          time.Time   `json:"trades_time"`
		WaiversTime         time.Time   `json:"waivers_time"`
	} `json:"data"`
	Next int `json:"next"`
}

type PlayerDetailedInfo struct {
	ChanceOfPlayingNextRound         int         `json:"chance_of_playing_next_round"`
	ChanceOfPlayingThisRound         int         `json:"chance_of_playing_this_round"`
	Code                             int         `json:"code"`
	CostChangeEvent                  int         `json:"cost_change_event"`
	CostChangeEventFall              int         `json:"cost_change_event_fall"`
	CostChangeStart                  int         `json:"cost_change_start"`
	CostChangeStartFall              int         `json:"cost_change_start_fall"`
	DreamteamCount                   int         `json:"dreamteam_count"`
	ElementType                      int         `json:"element_type"`
	EpNext                           string      `json:"ep_next"`
	EpThis                           string      `json:"ep_this"`
	EventPoints                      int         `json:"event_points"`
	FirstName                        string      `json:"first_name"`
	Form                             string      `json:"form"`
	ID                               int         `json:"id"`
	InDreamteam                      bool        `json:"in_dreamteam"`
	News                             string      `json:"news"`
	NewsAdded                        time.Time   `json:"news_added"`
	NowCost                          int         `json:"now_cost"`
	Photo                            string      `json:"photo"`
	PointsPerGame                    string      `json:"points_per_game"`
	SecondName                       string      `json:"second_name"`
	SelectedByPercent                string      `json:"selected_by_percent"`
	Special                          bool        `json:"special"`
	SquadNumber                      interface{} `json:"squad_number"`
	Status                           string      `json:"status"`
	Team                             int         `json:"team"`
	TeamCode                         int         `json:"team_code"`
	TotalPoints                      int         `json:"total_points"`
	TransfersIn                      int         `json:"transfers_in"`
	TransfersInEvent                 int         `json:"transfers_in_event"`
	TransfersOut                     int         `json:"transfers_out"`
	TransfersOutEvent                int         `json:"transfers_out_event"`
	ValueForm                        string      `json:"value_form"`
	ValueSeason                      string      `json:"value_season"`
	WebName                          string      `json:"web_name"`
	Minutes                          int         `json:"minutes"`
	GoalsScored                      int         `json:"goals_scored"`
	Assists                          int         `json:"assists"`
	CleanSheets                      int         `json:"clean_sheets"`
	GoalsConceded                    int         `json:"goals_conceded"`
	OwnGoals                         int         `json:"own_goals"`
	PenaltiesSaved                   int         `json:"penalties_saved"`
	PenaltiesMissed                  int         `json:"penalties_missed"`
	YellowCards                      int         `json:"yellow_cards"`
	RedCards                         int         `json:"red_cards"`
	Saves                            int         `json:"saves"`
	Bonus                            int         `json:"bonus"`
	Bps                              int         `json:"bps"`
	Influence                        string      `json:"influence"`
	Creativity                       string      `json:"creativity"`
	Threat                           string      `json:"threat"`
	IctIndex                         string      `json:"ict_index"`
	InfluenceRank                    int         `json:"influence_rank"`
	InfluenceRankType                int         `json:"influence_rank_type"`
	CreativityRank                   int         `json:"creativity_rank"`
	CreativityRankType               int         `json:"creativity_rank_type"`
	ThreatRank                       int         `json:"threat_rank"`
	ThreatRankType                   int         `json:"threat_rank_type"`
	IctIndexRank                     int         `json:"ict_index_rank"`
	IctIndexRankType                 int         `json:"ict_index_rank_type"`
	CornersAndIndirectFreekicksOrder interface{} `json:"corners_and_indirect_freekicks_order"`
	CornersAndIndirectFreekicksText  string      `json:"corners_and_indirect_freekicks_text"`
	DirectFreekicksOrder             interface{} `json:"direct_freekicks_order"`
	DirectFreekicksText              string      `json:"direct_freekicks_text"`
	PenaltiesOrder                   interface{} `json:"penalties_order"`
	PenaltiesText                    string      `json:"penalties_text"`
}

type ElementStatsResponse struct {
	Name           string      `json:"name"`
	Label          string      `json:"label"`
	Abbreviation   string      `json:"abbreviation"`
	IsMatchStat    bool        `json:"is_match_stat"`
	MatchStatOrder interface{} `json:"match_stat_order"`
	Sort           string      `json:"sort"`
}

type SettingsResponse struct {
	League struct {
		DefaultEntries          int    `json:"default_entries"`
		DraftReminderHours      []int  `json:"draft_reminder_hours"`
		DraftPostponeHours      int    `json:"draft_postpone_hours"`
		DraftPushbackTimes      int    `json:"draft_pushback_times"`
		H2HDraw                 int    `json:"h2h_draw"`
		H2HLose                 int    `json:"h2h_lose"`
		H2HWin                  int    `json:"h2h_win"`
		MaxEntries              int    `json:"max_entries"`
		MinEntries              int    `json:"min_entries"`
		PrivateMax              int    `json:"private_max"`
		PublicDraftDelayMinutes int    `json:"public_draft_delay_minutes"`
		PublicDraftTzDefault    string `json:"public_draft_tz_default"`
		PublicEntrySizes        []int  `json:"public_entry_sizes"`
		PublicMax               int    `json:"public_max"`
	} `json:"league"`
	Scoring struct {
		LongPlayLimit    int `json:"long_play_limit"`
		ShortPlay        int `json:"short_play"`
		LongPlay         int `json:"long_play"`
		ConcedeLimit     int `json:"concede_limit"`
		GoalsConcededGKP int `json:"goals_conceded_GKP"`
		GoalsConcededDEF int `json:"goals_conceded_DEF"`
		GoalsConcededMID int `json:"goals_conceded_MID"`
		GoalsConcededFWD int `json:"goals_conceded_FWD"`
		SavesLimit       int `json:"saves_limit"`
		Saves            int `json:"saves"`
		GoalsScoredGKP   int `json:"goals_scored_GKP"`
		GoalsScoredDEF   int `json:"goals_scored_DEF"`
		GoalsScoredMID   int `json:"goals_scored_MID"`
		GoalsScoredFWD   int `json:"goals_scored_FWD"`
		Assists          int `json:"assists"`
		CleanSheetsGKP   int `json:"clean_sheets_GKP"`
		CleanSheetsDEF   int `json:"clean_sheets_DEF"`
		CleanSheetsMID   int `json:"clean_sheets_MID"`
		CleanSheetsFWD   int `json:"clean_sheets_FWD"`
		PenaltiesSaved   int `json:"penalties_saved"`
		PenaltiesMissed  int `json:"penalties_missed"`
		YellowCards      int `json:"yellow_cards"`
		RedCards         int `json:"red_cards"`
		OwnGoals         int `json:"own_goals"`
		Bonus            int `json:"bonus"`
	} `json:"scoring"`
	Squad struct {
		Size              int `json:"size"`
		SelectGKP         int `json:"select_GKP"`
		SelectDEF         int `json:"select_DEF"`
		SelectMID         int `json:"select_MID"`
		SelectFWD         int `json:"select_FWD"`
		Play              int `json:"play"`
		MinPlayGKP        int `json:"min_play_GKP"`
		MaxPlayGKP        int `json:"max_play_GKP"`
		MinPlayDEF        int `json:"min_play_DEF"`
		MaxPlayDEF        int `json:"max_play_DEF"`
		MinPlayMID        int `json:"min_play_MID"`
		MaxPlayMID        int `json:"max_play_MID"`
		MinPlayFWD        int `json:"min_play_FWD"`
		MaxPlayFWD        int `json:"max_play_FWD"`
		PositionTypeLocks struct {
			Num12 string `json:"12"`
		} `json:"position_type_locks"`
		CaptainsDisabled bool `json:"captains_disabled"`
	} `json:"squad"`
	Transactions struct {
		NewElementLockedHours           int `json:"new_element_locked_hours"`
		TradeVetoMinimum                int `json:"trade_veto_minimum"`
		TradeVetoHours                  int `json:"trade_veto_hours"`
		WaiversBeforeStartMinHours      int `json:"waivers_before_start_min_hours"`
		WaiversBeforeDeadlineHours      int `json:"waivers_before_deadline_hours"`
		WaiversBeforeDeadlineHoursEvent struct {
			Num36 int `json:"36"`
		} `json:"waivers_before_deadline_hours_event"`
	} `json:"transactions"`
	UI struct {
		SpecialShirtExclusions []interface{} `json:"special_shirt_exclusions"`
		UseSpecialShirts       bool          `json:"use_special_shirts"`
	} `json:"ui"`
}

type ElementTypesResponse struct {
	ID                 int    `json:"id"`
	PluralName         string `json:"plural_name"`
	PluralNameShort    string `json:"plural_name_short"`
	SingularName       string `json:"singular_name"`
	SingularNameShort  string `json:"singular_name_short"`
	SquadSelect        int    `json:"squad_select"`
	SquadMinPlay       int    `json:"squad_min_play"`
	SquadMaxPlay       int    `json:"squad_max_play"`
	UIShirtSpecific    bool   `json:"ui_shirt_specific"`
	SubPositionsLocked []int  `json:"sub_positions_locked"`
	ElementCount       int    `json:"element_count"`
}
type Fixture []struct {
	Code                 int           `json:"code"`
	Event                interface{}   `json:"event"`
	Finished             bool          `json:"finished"`
	FinishedProvisional  bool          `json:"finished_provisional"`
	ID                   int           `json:"id"`
	KickoffTime          interface{}   `json:"kickoff_time"`
	Minutes              int           `json:"minutes"`
	ProvisionalStartTime bool          `json:"provisional_start_time"`
	Started              interface{}   `json:"started"`
	TeamA                int           `json:"team_a"`
	TeamAScore           interface{}   `json:"team_a_score"`
	TeamH                int           `json:"team_h"`
	TeamHScore           interface{}   `json:"team_h_score"`
	Stats                []interface{} `json:"stats"`
	TeamHDifficulty      int           `json:"team_h_difficulty"`
	TeamADifficulty      int           `json:"team_a_difficulty"`
	PulseID              int           `json:"pulse_id"`
}

type WeeklyFixture []struct {
	Code                 int    `json:"code"`
	Event                int    `json:"event"`
	Finished             bool   `json:"finished"`
	FinishedProvisional  bool   `json:"finished_provisional"`
	ID                   int    `json:"id"`
	KickoffTime          string `json:"kickoff_time"`
	Minutes              int    `json:"minutes"`
	ProvisionalStartTime bool   `json:"provisional_start_time"`
	Started              bool   `json:"started"`
	TeamA                int    `json:"team_a"`
	TeamAScore           int    `json:"team_a_score"`
	TeamH                int    `json:"team_h"`
	TeamHScore           int    `json:"team_h_score"`
	Stats                []struct {
		Identifier string `json:"identifier"`
		A          []struct {
			Value   int `json:"value"`
			Element int `json:"element"`
		} `json:"a"`
		H []struct {
			Value   int `json:"value"`
			Element int `json:"element"`
		} `json:"h"`
	} `json:"stats"`
	TeamHDifficulty int `json:"team_h_difficulty"`
	TeamADifficulty int `json:"team_a_difficulty"`
	PulseID         int `json:"pulse_id"`
}

type Player struct {
	Fixtures []struct {
		ID                   int         `json:"id"`
		Code                 int         `json:"code"`
		TeamH                int         `json:"team_h"`
		TeamHScore           interface{} `json:"team_h_score"`
		TeamA                int         `json:"team_a"`
		TeamAScore           interface{} `json:"team_a_score"`
		Event                int         `json:"event"`
		Finished             bool        `json:"finished"`
		Minutes              int         `json:"minutes"`
		ProvisionalStartTime bool        `json:"provisional_start_time"`
		KickoffTime          string      `json:"kickoff_time"`
		EventName            string      `json:"event_name"`
		IsHome               bool        `json:"is_home"`
		Difficulty           int         `json:"difficulty"`
	} `json:"fixtures"`
	History []struct {
		Element          int       `json:"element"`
		Fixture          int       `json:"fixture"`
		OpponentTeam     int       `json:"opponent_team"`
		TotalPoints      int       `json:"total_points"`
		WasHome          bool      `json:"was_home"`
		KickoffTime      time.Time `json:"kickoff_time"`
		TeamHScore       int       `json:"team_h_score"`
		TeamAScore       int       `json:"team_a_score"`
		Round            int       `json:"round"`
		Minutes          int       `json:"minutes"`
		GoalsScored      int       `json:"goals_scored"`
		Assists          int       `json:"assists"`
		CleanSheets      int       `json:"clean_sheets"`
		GoalsConceded    int       `json:"goals_conceded"`
		OwnGoals         int       `json:"own_goals"`
		PenaltiesSaved   int       `json:"penalties_saved"`
		PenaltiesMissed  int       `json:"penalties_missed"`
		YellowCards      int       `json:"yellow_cards"`
		RedCards         int       `json:"red_cards"`
		Saves            int       `json:"saves"`
		Bonus            int       `json:"bonus"`
		Bps              int       `json:"bps"`
		Influence        string    `json:"influence"`
		Creativity       string    `json:"creativity"`
		Threat           string    `json:"threat"`
		IctIndex         string    `json:"ict_index"`
		Value            int       `json:"value"`
		TransfersBalance int       `json:"transfers_balance"`
		Selected         int       `json:"selected"`
		TransfersIn      int       `json:"transfers_in"`
		TransfersOut     int       `json:"transfers_out"`
	} `json:"history"`
	HistoryPast []struct {
		SeasonName      string `json:"season_name"`
		ElementCode     int    `json:"element_code"`
		StartCost       int    `json:"start_cost"`
		EndCost         int    `json:"end_cost"`
		TotalPoints     int    `json:"total_points"`
		Minutes         int    `json:"minutes"`
		GoalsScored     int    `json:"goals_scored"`
		Assists         int    `json:"assists"`
		CleanSheets     int    `json:"clean_sheets"`
		GoalsConceded   int    `json:"goals_conceded"`
		OwnGoals        int    `json:"own_goals"`
		PenaltiesSaved  int    `json:"penalties_saved"`
		PenaltiesMissed int    `json:"penalties_missed"`
		YellowCards     int    `json:"yellow_cards"`
		RedCards        int    `json:"red_cards"`
		Saves           int    `json:"saves"`
		Bonus           int    `json:"bonus"`
		Bps             int    `json:"bps"`
		Influence       string `json:"influence"`
		Creativity      string `json:"creativity"`
		Threat          string `json:"threat"`
		IctIndex        string `json:"ict_index"`
	} `json:"history_past"`
}
type PlayerFixture struct {
	ID                   int         `json:"id"`
	Code                 int         `json:"code"`
	TeamH                int         `json:"team_h"`
	TeamHScore           interface{} `json:"team_h_score"`
	TeamA                int         `json:"team_a"`
	TeamAScore           interface{} `json:"team_a_score"`
	Event                int         `json:"event"`
	Finished             bool        `json:"finished"`
	Minutes              int         `json:"minutes"`
	ProvisionalStartTime bool        `json:"provisional_start_time"`
	KickoffTime          string      `json:"kickoff_time"`
	EventName            string      `json:"event_name"`
	IsHome               bool        `json:"is_home"`
	Difficulty           int         `json:"difficulty"`
}

type PlayerHistory struct {
	Element          int       `json:"element"`
	Fixture          int       `json:"fixture"`
	OpponentTeam     int       `json:"opponent_team"`
	TotalPoints      int       `json:"total_points"`
	WasHome          bool      `json:"was_home"`
	KickoffTime      time.Time `json:"kickoff_time"`
	TeamHScore       int       `json:"team_h_score"`
	TeamAScore       int       `json:"team_a_score"`
	Round            int       `json:"round"`
	Minutes          int       `json:"minutes"`
	GoalsScored      int       `json:"goals_scored"`
	Assists          int       `json:"assists"`
	CleanSheets      int       `json:"clean_sheets"`
	GoalsConceded    int       `json:"goals_conceded"`
	OwnGoals         int       `json:"own_goals"`
	PenaltiesSaved   int       `json:"penalties_saved"`
	PenaltiesMissed  int       `json:"penalties_missed"`
	YellowCards      int       `json:"yellow_cards"`
	RedCards         int       `json:"red_cards"`
	Saves            int       `json:"saves"`
	Bonus            int       `json:"bonus"`
	Bps              int       `json:"bps"`
	Influence        string    `json:"influence"`
	Creativity       string    `json:"creativity"`
	Threat           string    `json:"threat"`
	IctIndex         string    `json:"ict_index"`
	Value            int       `json:"value"`
	TransfersBalance int       `json:"transfers_balance"`
	Selected         int       `json:"selected"`
	TransfersIn      int       `json:"transfers_in"`
	TransfersOut     int       `json:"transfers_out"`
}

type PlayerHistoryPast struct {
	SeasonName      string `json:"season_name"`
	ElementCode     int    `json:"element_code"`
	StartCost       int    `json:"start_cost"`
	EndCost         int    `json:"end_cost"`
	TotalPoints     int    `json:"total_points"`
	Minutes         int    `json:"minutes"`
	GoalsScored     int    `json:"goals_scored"`
	Assists         int    `json:"assists"`
	CleanSheets     int    `json:"clean_sheets"`
	GoalsConceded   int    `json:"goals_conceded"`
	OwnGoals        int    `json:"own_goals"`
	PenaltiesSaved  int    `json:"penalties_saved"`
	PenaltiesMissed int    `json:"penalties_missed"`
	YellowCards     int    `json:"yellow_cards"`
	RedCards        int    `json:"red_cards"`
	Saves           int    `json:"saves"`
	Bonus           int    `json:"bonus"`
	Bps             int    `json:"bps"`
	Influence       string `json:"influence"`
	Creativity      string `json:"creativity"`
	Threat          string `json:"threat"`
	IctIndex        string `json:"ict_index"`
}
type GameWeek struct {
	Elements []struct {
		ID    int `json:"id"`
		Stats struct {
			Minutes         int    `json:"minutes"`
			GoalsScored     int    `json:"goals_scored"`
			Assists         int    `json:"assists"`
			CleanSheets     int    `json:"clean_sheets"`
			GoalsConceded   int    `json:"goals_conceded"`
			OwnGoals        int    `json:"own_goals"`
			PenaltiesSaved  int    `json:"penalties_saved"`
			PenaltiesMissed int    `json:"penalties_missed"`
			YellowCards     int    `json:"yellow_cards"`
			RedCards        int    `json:"red_cards"`
			Saves           int    `json:"saves"`
			Bonus           int    `json:"bonus"`
			Bps             int    `json:"bps"`
			Influence       string `json:"influence"`
			Creativity      string `json:"creativity"`
			Threat          string `json:"threat"`
			IctIndex        string `json:"ict_index"`
			TotalPoints     int    `json:"total_points"`
			InDreamteam     bool   `json:"in_dreamteam"`
		} `json:"stats"`
		Explain []struct {
			Fixture int `json:"fixture"`
			Stats   []struct {
				Identifier string `json:"identifier"`
				Points     int    `json:"points"`
				Value      int    `json:"value"`
			} `json:"stats"`
		} `json:"explain"`
	} `json:"elements"`
}

type Manager struct {
	ID                       int       `json:"id"`
	JoinedTime               time.Time `json:"joined_time"`
	StartedEvent             int       `json:"started_event"`
	FavouriteTeam            int       `json:"favourite_team"`
	PlayerFirstName          string    `json:"player_first_name"`
	PlayerLastName           string    `json:"player_last_name"`
	PlayerRegionID           int       `json:"player_region_id"`
	PlayerRegionName         string    `json:"player_region_name"`
	PlayerRegionIsoCodeShort string    `json:"player_region_iso_code_short"`
	PlayerRegionIsoCodeLong  string    `json:"player_region_iso_code_long"`
	SummaryOverallPoints     int       `json:"summary_overall_points"`
	SummaryOverallRank       int       `json:"summary_overall_rank"`
	SummaryEventPoints       int       `json:"summary_event_points"`
	SummaryEventRank         int       `json:"summary_event_rank"`
	CurrentEvent             int       `json:"current_event"`
	Leagues                  struct {
		Classic []struct {
			ID             int         `json:"id"`
			Name           string      `json:"name"`
			ShortName      string      `json:"short_name"`
			Created        time.Time   `json:"created"`
			Closed         bool        `json:"closed"`
			Rank           interface{} `json:"rank"`
			MaxEntries     interface{} `json:"max_entries"`
			LeagueType     string      `json:"league_type"`
			Scoring        string      `json:"scoring"`
			AdminEntry     interface{} `json:"admin_entry"`
			StartEvent     int         `json:"start_event"`
			EntryCanLeave  bool        `json:"entry_can_leave"`
			EntryCanAdmin  bool        `json:"entry_can_admin"`
			EntryCanInvite bool        `json:"entry_can_invite"`
			HasCup         bool        `json:"has_cup"`
			CupLeague      interface{} `json:"cup_league"`
			CupQualified   interface{} `json:"cup_qualified"`
			EntryRank      int         `json:"entry_rank"`
			EntryLastRank  int         `json:"entry_last_rank"`
		} `json:"classic"`
		H2H []interface{} `json:"h2h"`
		Cup struct {
			Matches []interface{} `json:"matches"`
			Status  struct {
				QualificationEvent   int    `json:"qualification_event"`
				QualificationNumbers int    `json:"qualification_numbers"`
				QualificationRank    int    `json:"qualification_rank"`
				QualificationState   string `json:"qualification_state"`
			} `json:"status"`
			CupLeague int `json:"cup_league"`
		} `json:"cup"`
	} `json:"leagues"`
	Name                       string      `json:"name"`
	Kit                        interface{} `json:"kit"`
	LastDeadlineBank           int         `json:"last_deadline_bank"`
	LastDeadlineValue          int         `json:"last_deadline_value"`
	LastDeadlineTotalTransfers int         `json:"last_deadline_total_transfers"`
}
type ManagerResponse struct {
}

type ManagerLeaguesClassic struct {
	ID             int         `json:"id"`
	Name           string      `json:"name"`
	ShortName      string      `json:"short_name"`
	Created        time.Time   `json:"created"`
	Closed         bool        `json:"closed"`
	Rank           interface{} `json:"rank"`
	MaxEntries     interface{} `json:"max_entries"`
	LeagueType     string      `json:"league_type"`
	Scoring        string      `json:"scoring"`
	AdminEntry     interface{} `json:"admin_entry"`
	StartEvent     int         `json:"start_event"`
	EntryCanLeave  bool        `json:"entry_can_leave"`
	EntryCanAdmin  bool        `json:"entry_can_admin"`
	EntryCanInvite bool        `json:"entry_can_invite"`
	HasCup         bool        `json:"has_cup"`
	CupLeague      interface{} `json:"cup_league"`
	CupQualified   interface{} `json:"cup_qualified"`
	EntryRank      int         `json:"entry_rank"`
	EntryLastRank  int         `json:"entry_last_rank"`
}

type ManagerLeaguesCup struct {
	Matches []interface{} `json:"matches"`
	Status  struct {
		QualificationEvent   int    `json:"qualification_event"`
		QualificationNumbers int    `json:"qualification_numbers"`
		QualificationRank    int    `json:"qualification_rank"`
		QualificationState   string `json:"qualification_state"`
	} `json:"status"`
	CupLeague int `json:"cup_league"`
}

type ManagerCup struct{}

type EntryHistory struct {
	History []EntryHistoryWeek `json:"history"`
	Entry   struct {
		EventPoints       int    `json:"event_points"`
		FavouriteTeam     int    `json:"favourite_team"`
		ID                int    `json:"id"`
		LeagueSet         []int  `json:"league_set"`
		Name              string `json:"name"`
		OverallPoints     int    `json:"overall_points"`
		PlayerFirstName   string `json:"player_first_name"`
		PlayerLastName    string `json:"player_last_name"`
		RegionName        string `json:"region_name"`
		RegionCodeShort   string `json:"region_code_short"`
		RegionCodeLong    string `json:"region_code_long"`
		StartedEvent      int    `json:"started_event"`
		TransactionsEvent int    `json:"transactions_event"`
		TransactionsTotal int    `json:"transactions_total"`
	} `json:"entry"`
}

type EntryHistoryWeek struct {
	ID             int         `json:"id"`
	Points         int         `json:"points"`
	TotalPoints    int         `json:"total_points"`
	Rank           interface{} `json:"rank"`
	RankSort       interface{} `json:"rank_sort"`
	EventTransfers int         `json:"event_transfers"`
	PointsOnBench  int         `json:"points_on_bench"`
	Entry          int         `json:"entry"`
	Event          int         `json:"event"`
}
type LeagueInfo struct {
	League struct {
		AdminEntry         int       `json:"admin_entry"`
		Closed             bool      `json:"closed"`
		DraftDT            time.Time `json:"draft_dt"`
		DraftPickTimeLimit int       `json:"draft_pick_time_limit"`
		DraftStatus        string    `json:"draft_status"`
		DraftTZShow        string    `json:"draft_tz_show"`
		ID                 int       `json:"id"`
		KORounds           int       `json:"ko_rounds"`
		MakeCodePublic     bool      `json:"make_code_public"`
		MaxEntries         int       `json:"max_entries"`
		MinEntries         int       `json:"min_entries"`
		Name               string    `json:"name"`
		Scoring            string    `json:"scoring"`
		StartEvent         int       `json:"start_event"`
		StopEvent          int       `json:"stop_event"`
		Trades             string    `json:"trades"`
		TransactionMode    string    `json:"transaction_mode"`
		Variety            string    `json:"variety"`
	} `json:"league"`
	LeagueEntries []struct {
		EntryID         int       `json:"entry_id"`
		EntryName       string    `json:"entry_name"`
		ID              int       `json:"id"`
		JoinedTime      time.Time `json:"joined_time"`
		PlayerFirstName string    `json:"player_first_name"`
		PlayerLastName  string    `json:"player_last_name"`
		ShortName       string    `json:"short_name"`
		WaiverPick      int       `json:"waiver_pick"`
	} `json:"league_entries"`
	Standings []struct {
		EventTotal  int `json:"event_total"`
		LastRank    int `json:"last_rank"`
		LeagueEntry int `json:"league_entry"`
		Rank        int `json:"rank"`
		RankSort    int `json:"rank_sort"`
		Total       int `json:"total"`
	} `json:"standings"`
}
type LeagueResponse struct {
	AdminEntry         int       `json:"admin_entry"`
	Closed             bool      `json:"closed"`
	DraftDT            time.Time `json:"draft_dt"`
	DraftPickTimeLimit int       `json:"draft_pick_time_limit"`
	DraftStatus        string    `json:"draft_status"`
	DraftTZShow        string    `json:"draft_tz_show"`
	ID                 int       `json:"id"`
	KORounds           int       `json:"ko_rounds"`
	MakeCodePublic     bool      `json:"make_code_public"`
	MaxEntries         int       `json:"max_entries"`
	MinEntries         int       `json:"min_entries"`
	Name               string    `json:"name"`
	Scoring            string    `json:"scoring"`
	StartEvent         int       `json:"start_event"`
	StopEvent          int       `json:"stop_event"`
	Trades             string    `json:"trades"`
	TransactionMode    string    `json:"transaction_mode"`
	Variety            string    `json:"variety"`
}
type LeagueEntriesResponse struct {
	EntryID         int       `json:"entry_id"`
	EntryName       string    `json:"entry_name"`
	ID              int       `json:"id"`
	JoinedTime      time.Time `json:"joined_time"`
	PlayerFirstName string    `json:"player_first_name"`
	PlayerLastName  string    `json:"player_last_name"`
	ShortName       string    `json:"short_name"`
	WaiverPick      int       `json:"waiver_pick"`
}
type StandingsResponse struct {
	EventTotal  int `json:"event_total"`
	LastRank    int `json:"last_rank"`
	LeagueEntry int `json:"league_entry"`
	Rank        int `json:"rank"`
	RankSort    int `json:"rank_sort"`
	Total       int `json:"total"`
}
type MyTeam struct {
	Picks []struct {
		Element       int  `json:"element"`
		Position      int  `json:"position"`
		SellingPrice  int  `json:"selling_price"`
		Multiplier    int  `json:"multiplier"`
		PurchasePrice int  `json:"purchase_price"`
		IsCaptain     bool `json:"is_captain"`
		IsViceCaptain bool `json:"is_vice_captain"`
	} `json:"picks"`
	Chips []struct {
		StatusForEntry string        `json:"status_for_entry"`
		PlayedByEntry  []interface{} `json:"played_by_entry"`
		Name           string        `json:"name"`
		Number         int           `json:"number"`
		StartEvent     int           `json:"start_event"`
		StopEvent      int           `json:"stop_event"`
		ChipType       string        `json:"chip_type"`
	} `json:"chips"`
	Transfers struct {
		Cost   int    `json:"cost"`
		Status string `json:"status"`
		Limit  int    `json:"limit"`
		Made   int    `json:"made"`
		Bank   int    `json:"bank"`
		Value  int    `json:"value"`
	} `json:"transfers"`
}

type TransferHistory struct {
	ElementIn      int       `json:"element_in"`
	ElementInCost  int       `json:"element_in_cost"`
	ElementOut     int       `json:"element_out"`
	ElementOutCost int       `json:"element_out_cost"`
	Entry          int       `json:"entry"`
	Evnt           int       `json:"evnt,omitempty"`
	Time           time.Time `json:"time"`
	Event          int       `json:"event,omitempty"`
}

type TeamWeekly struct {
	Picks []struct {
		Element       int  `json:"element"`
		Position      int  `json:"position"`
		IsCaptain     bool `json:"is_captain"`
		IsViceCaptain bool `json:"is_vice_captain"`
		Multiplier    int  `json:"multiplier"`
	} `json:"picks"`
	EntryHistory struct {
	} `json:"entry_history"`
	Subs []interface{} `json:"subs"`
}
