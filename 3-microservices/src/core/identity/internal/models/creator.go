package models

import "time"

// CreatorType creator type
type CreatorType string

const (
	GameProject        CreatorType = "Game Project"
	CharityNGO         CreatorType = "Charity/NGO"
	Brand              CreatorType = "Brand"
	Influencer         CreatorType = "Influencer"
	FBVenue            CreatorType = "F&BVenue"
	Retail             CreatorType = "Retail"
	EventOrganizer     CreatorType = "Event Organizer"
	Agency             CreatorType = "Agency"
	SoftwareCompany    CreatorType = "Software company"
	IndependentCreator CreatorType = "Independent creator"
	Others             CreatorType = "Others"
)

func (creatorType CreatorType) String() string {
	switch creatorType {
	case GameProject:
		return "Game Project"
	case CharityNGO:
		return "Charity/NGO"
	case Brand:
		return "Brand"
	case Influencer:
		return "Influencer"
	case FBVenue:
		return "F&BVenue"
	case Retail:
		return "Retail"
	case EventOrganizer:
		return "Event Organizer"
	case Agency:
		return "Agency"
	case SoftwareCompany:
		return "Software company"
	case IndependentCreator:
		return "Independent creator"
	case Others:
		return "Others"
	}
	return "UNKNOWN"
}

func (creatorType CreatorType) isValid() bool {
	if creatorType.String() != "UNKNOWN" {
		return true
	}
	return false
}

type Creator struct {
	ID        string // uuid
	UserID    string
	Type      string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}
