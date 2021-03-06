package entity

import (
	"github.com/photoprism/photoprism/pkg/s2"
)

type LocationMap map[string]Location

func (m LocationMap) Get(name string) Location {
	if result, ok := m[name]; ok {
		return result
	}

	return UnknownLocation
}

func (m LocationMap) Pointer(name string) *Location {
	if result, ok := m[name]; ok {
		return &result
	}

	return &UnknownLocation
}

var LocationFixtures = LocationMap{
	"mexico": {
		ID:          s2.TokenPrefix + "85d1ea7d382c",
		PlaceID:     PlaceFixtures.Get("mexico").ID,
		LocName:     "Adosada Platform",
		LocCategory: "botanical garden",
		Place:       PlaceFixtures.Pointer("mexico"),
		LocSource:   "places",
		CreatedAt:   Timestamp(),
		UpdatedAt:   Timestamp(),
	},
	"caravan park": {
		ID:      s2.TokenPrefix + "1ef75a71a36c",
		PlaceID: s2.TokenPrefix + "1ef75a71a36c",
		Place: &Place{
			ID:         s2.TokenPrefix + "1ef75a71a36",
			LocLabel:   "Mandeni, KwaZulu-Natal, South Africa",
			LocCity:    "Mandeni",
			LocState:   "KwaZulu-Natal",
			LocCountry: "za",
			CreatedAt:  Timestamp(),
			UpdatedAt:  Timestamp(),
		},
		LocName:     "Lobotes Caravan Park",
		LocCategory: "camping",
		LocSource:   "manual",
		CreatedAt:   Timestamp(),
		UpdatedAt:   Timestamp(),
	},
	"zinkwazi": {
		ID:          s2.TokenPrefix + "1ef744d1e28c",
		PlaceID:     PlaceFixtures.Get("zinkwazi").ID,
		Place:       PlaceFixtures.Pointer("zinkwazi"),
		LocName:     "Zinkwazi Beach",
		LocCategory: "beach",
		LocSource:   "places",
		CreatedAt:   Timestamp(),
		UpdatedAt:   Timestamp(),
	},
	"hassloch": {
		ID:          s2.TokenPrefix + "1ef744d1e280",
		PlaceID:     PlaceFixtures.Get("holidaypark").ID,
		Place:       PlaceFixtures.Pointer("holidaypark"),
		LocName:     "Holiday Park",
		LocCategory: "park",
		LocSource:   "places",
		CreatedAt:   Timestamp(),
		UpdatedAt:   Timestamp(),
	},
	"emptyNameLongCity": {
		ID:          s2.TokenPrefix + "1ef744d1e281",
		PlaceID:     PlaceFixtures.Get("emptyNameLongCity").ID,
		Place:       PlaceFixtures.Pointer("emptyNameLongCity"),
		LocName:     "",
		LocCategory: "botanical garden",
		LocSource:   "places",
		CreatedAt:   Timestamp(),
		UpdatedAt:   Timestamp(),
	},
	"emptyNameShortCity": {
		ID:          s2.TokenPrefix + "1ef744d1e282",
		PlaceID:     PlaceFixtures.Get("emptyNameShortCity").ID,
		Place:       PlaceFixtures.Pointer("emptyNameShortCity"),
		LocName:     "",
		LocCategory: "botanical garden",
		LocSource:   "places",
		CreatedAt:   Timestamp(),
		UpdatedAt:   Timestamp(),
	},
	"veryLongLocName": {
		ID:          s2.TokenPrefix + "1ef744d1e283",
		PlaceID:     PlaceFixtures.Get("veryLongLocName").ID,
		Place:       PlaceFixtures.Pointer("veryLongLocName"),
		LocName:     "longlonglonglonglonglonglonglonglonglonglonglonglongName",
		LocCategory: "cape",
		LocSource:   "places",
		CreatedAt:   Timestamp(),
		UpdatedAt:   Timestamp(),
	},
	"mediumLongLocName": {
		ID:          s2.TokenPrefix + "1ef744d1e283",
		PlaceID:     PlaceFixtures.Get("mediumLongLocName").ID,
		Place:       PlaceFixtures.Pointer("mediumLongLocName"),
		LocName:     "longlonglonglonglonglongName",
		LocCategory: "botanical garden",
		LocSource:   "places",
		CreatedAt:   Timestamp(),
		UpdatedAt:   Timestamp(),
	},
}

// CreateLocationFixtures inserts known entities into the database for testing.
func CreateLocationFixtures() {
	for _, entity := range LocationFixtures {
		Db().Create(&entity)
	}
}
