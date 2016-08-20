package data

import (
	"regexp"
	"strconv"

	"github.com/jinzhu/gorm"
	"github.com/kleister/kleister-api/model"
)

// GetTeams retrieves all available teams from the database.
func (db *data) GetTeams() (*model.Teams, error) {
	records := &model.Teams{}

	err := db.Order(
		"name ASC",
	).Preload(
		"Users",
	).Preload(
		"Packs",
	).Preload(
		"Mods",
	).Find(
		records,
	).Error

	return records, err
}

// CreateTeam creates a new team.
func (db *data) CreateTeam(record *model.Team) error {
	return db.Create(
		record,
	).Error
}

// UpdateTeam updates a team.
func (db *data) UpdateTeam(record *model.Team) error {
	return db.Save(
		record,
	).Error
}

// DeleteTeam deletes a team.
func (db *data) DeleteTeam(record *model.Team) error {
	return db.Delete(
		record,
	).Error
}

// GetTeam retrieves a specific team from the database.
func (db *data) GetTeam(id string) (*model.Team, *gorm.DB) {
	var (
		record = &model.Team{}
		query  *gorm.DB
	)

	if match, _ := regexp.MatchString("^([0-9]+)$", id); match {
		val, _ := strconv.ParseInt(id, 10, 64)

		query = db.Where(
			"id = ?",
			val,
		)
	} else {
		query = db.Where(
			"slug = ?",
			id,
		)
	}

	res := query.Model(
		record,
	).Preload(
		"Users",
	).Preload(
		"Packs",
	).Preload(
		"Mods",
	).First(
		record,
	)

	return record, res
}

// GetTeamUsers retrieves users for a team.
func (db *data) GetTeamUsers(params *model.TeamUserParams) (*model.Users, error) {
	team, _ := db.GetTeam(params.Team)

	records := &model.Users{}

	err := db.Model(
		team,
	).Association(
		"Users",
	).Find(
		records,
	).Error

	return records, err
}

// GetTeamHasUser checks if a specific user is assigned to a team.
func (db *data) GetTeamHasUser(params *model.TeamUserParams) bool {
	team, _ := db.GetTeam(params.Team)
	user, _ := db.GetUser(params.User)

	res := db.Model(
		team,
	).Association(
		"Users",
	).Find(
		user,
	).Error

	return res == nil
}

func (db *data) CreateTeamUser(params *model.TeamUserParams) error {
	team, _ := db.GetTeam(params.Team)
	user, _ := db.GetUser(params.User)

	return db.Model(
		team,
	).Association(
		"Users",
	).Append(
		user,
	).Error
}

func (db *data) DeleteTeamUser(params *model.TeamUserParams) error {
	team, _ := db.GetTeam(params.Team)
	user, _ := db.GetUser(params.User)

	return db.Model(
		team,
	).Association(
		"Users",
	).Delete(
		user,
	).Error
}

// GetTeamPacks retrieves packs for a team.
func (db *data) GetTeamPacks(params *model.TeamPackParams) (*model.Packs, error) {
	team, _ := db.GetTeam(params.Team)

	records := &model.Packs{}

	err := db.Model(
		team,
	).Association(
		"Packs",
	).Find(
		records,
	).Error

	return records, err
}

// GetTeamHasPack checks if a specific pack is assigned to a team.
func (db *data) GetTeamHasPack(params *model.TeamPackParams) bool {
	team, _ := db.GetTeam(params.Team)
	pack, _ := db.GetPack(params.Pack)

	res := db.Model(
		team,
	).Association(
		"Packs",
	).Find(
		pack,
	).Error

	return res == nil
}

func (db *data) CreateTeamPack(params *model.TeamPackParams) error {
	team, _ := db.GetTeam(params.Team)
	pack, _ := db.GetPack(params.Pack)

	return db.Model(
		team,
	).Association(
		"Packs",
	).Append(
		pack,
	).Error
}

func (db *data) DeleteTeamPack(params *model.TeamPackParams) error {
	team, _ := db.GetTeam(params.Team)
	pack, _ := db.GetPack(params.Pack)

	return db.Model(
		team,
	).Association(
		"Packs",
	).Delete(
		pack,
	).Error
}

// GetTeamMods retrieves mods for a team.
func (db *data) GetTeamMods(params *model.TeamModParams) (*model.Mods, error) {
	team, _ := db.GetTeam(params.Team)

	records := &model.Mods{}

	err := db.Model(
		team,
	).Association(
		"Mods",
	).Find(
		records,
	).Error

	return records, err
}

// GetTeamHasMod checks if a specific mod is assigned to a team.
func (db *data) GetTeamHasMod(params *model.TeamModParams) bool {
	team, _ := db.GetTeam(params.Team)
	mod, _ := db.GetMod(params.Mod)

	res := db.Model(
		team,
	).Association(
		"Mods",
	).Find(
		mod,
	).Error

	return res == nil
}

func (db *data) CreateTeamMod(params *model.TeamModParams) error {
	team, _ := db.GetTeam(params.Team)
	mod, _ := db.GetMod(params.Mod)

	return db.Model(
		team,
	).Association(
		"Mods",
	).Append(
		mod,
	).Error
}

func (db *data) DeleteTeamMod(params *model.TeamModParams) error {
	team, _ := db.GetTeam(params.Team)
	mod, _ := db.GetMod(params.Mod)

	return db.Model(
		team,
	).Association(
		"Mods",
	).Delete(
		mod,
	).Error
}