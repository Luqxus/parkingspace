package service

type SpaceService interface {}

type spaceService struct {}


func NewSpaceService() *spaceService {
	return &spaceService{}
}

func GetNearestSpaces() {
	// TODO: parking spaces in my radius
	
}

func GetSpace() {
//	TODO: get parking space by id
}


func GetSpaceInLocation() {
//	TODO: get parking space in specific location
}

func GetFilteredSpaces() {
//	TODO: get parking spaces that fit a filter
}