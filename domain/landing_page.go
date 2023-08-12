package domain

type LandingPage struct {
	Hero Hero
	Product []Product
	Category []Category
}
type Hero struct {
	Travelers int
	Treasures int
	Cities int
}
type MostPicked struct {
	Product []Product
}


