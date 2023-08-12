package formatter

import (
	"staycation/domain"
)

type LandingPageFormatter struct {
	Hero HeroFormatter `json:hero`
	MostPicked []MostPickedFormatter `json:mostPicked`
	Categories []CategoryLandingFormatter `json:categories`
}
type MostPickedFormatter struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Slug string `json:"slug"`
	Price int `json:"price"`
	Country string `json:"country"`
	Province string `json:"province"`
	City string `json:"city"`
	Unit string `json:"unit"`
	Type string `json:"type"`
	ImageUrl string `json:"url"`
	IsPopuler bool `json:"is_popular"`
}
func FormatProductLanding(product domain.Product) MostPickedFormatter {

	formatter := MostPickedFormatter{}
	formatter.ID = product.ID
	formatter.Name = product.Name
	formatter.Slug = product.Slug
	formatter.Price = product.Price
	formatter.Country = product.Country
	formatter.Province = product.Province
	formatter.City = product.City
	formatter.Unit = product.Unit
	formatter.Type = product.Type
	if len(product.ProductImage) > 0 {
		formatter.ImageUrl = product.ProductImage[0].Url
	}
	formatter.IsPopuler = product.IsPopuler
	return formatter
}
func FormatMostPicked(product []domain.Product) []MostPickedFormatter {
	var formatter []MostPickedFormatter
	for _,items := range product {
		formatterProduct :=  FormatProductLanding(items)
		formatter = append(formatter, formatterProduct)

	}

	return formatter
}
type HeroFormatter struct {
	Travelers int
	Treasures int
	Cities int
}
type CategoryLandingFormatter struct {
	ID       int `json:"id"`
	Name     string `json:"name"`
	Items []MostPickedFormatter `json:"product"`
}
func FormatCategoryLanding(categorys []domain.Category) []CategoryLandingFormatter {
	var formatter []CategoryLandingFormatter
	var formatterItem []MostPickedFormatter
	for _,category := range categorys{
		for _,items := range category.Product {
			formatterProduct :=  FormatProductLanding(items)
			formatterItem = append(formatterItem, formatterProduct)
	
		}
		formatter1 := CategoryLandingFormatter{
			ID : category.ID,
			Name : category.Name,
			Items : formatterItem,
	
		}
		formatter = append(formatter, formatter1)
	}
	
	

	return formatter
}
func FormatHero(hero domain.Hero) HeroFormatter {
	formatter := HeroFormatter{
		Travelers: hero.Travelers,
		Treasures: hero.Treasures,
		Cities: hero.Cities,
	}

	return formatter
}

	



	

	



func FormatLandingPage(hero domain.Hero, category []domain.Category, product []domain.Product) LandingPageFormatter {
	formatter := LandingPageFormatter{
		Hero: FormatHero(hero),
		Categories: FormatCategoryLanding(category),
		MostPicked: FormatMostPicked(product),


	}

	return formatter
}