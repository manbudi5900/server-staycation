package formatter

import (
	"staycation/domain"
)

type ProductDetailFormatter struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Slug string `json:"slug"`
	Price int `json:"price"`
	Country string `json:"country"`
	Province string `json:"province"`
	City string `json:"city"`
	Unit string `json:"unit"`
	Type string `json:"type"`
	Description string `json:"description"`
	IsPopuler bool `json:"is_popular"`
	ProductImage []ProductImageFormatter `json:"images"`
	Feature []FeatureFormatter `json:"features"`
	Activity []ActivityFormatter `json:"activities"`
	Category []CategoryLandingFormatter `json:"categories"`
	Testimonial []TestimonialFormatter `json:"testimonial"`
	
}


func FormatProductDetail(product domain.Product) ProductDetailFormatter {

	formatter := ProductDetailFormatter{}
	formatter.ID = product.ID
	formatter.Name = product.Name
	formatter.Slug = product.Slug
	formatter.Price = product.Price
	formatter.Country = product.Country
	formatter.Province = product.Province
	formatter.City = product.City
	formatter.Unit = product.Unit
	formatter.Type = product.Type
	formatter.Description = product.Description
	formatter.IsPopuler = product.IsPopuler

	var formatterItemImage []ProductImageFormatter


	for _,items := range product.ProductImage {
		formatterProductImage :=  FormatProductImage(items)
		formatterItemImage = append(formatterItemImage, formatterProductImage)

	}
	formatter.ProductImage = formatterItemImage

	var formatterItemFeature []FeatureFormatter


	for _,items := range product.Feature {
		formatterProductFeature :=  FormatFeature(items)
		formatterItemFeature = append(formatterItemFeature, formatterProductFeature)

	}
	formatter.Feature = formatterItemFeature

	var formatterItemActivity[]ActivityFormatter


	for _,items := range product.Activity {
		formatterProductActivity :=  FormatActivity(items)
		formatterItemActivity = append(formatterItemActivity, formatterProductActivity)

	}
	formatter.Activity = formatterItemActivity

	var formatterCategory []CategoryLandingFormatter
	var formatterItem []MostPickedFormatter
		for _,items := range product.Category.Product {
			formatterProduct :=  FormatProductLanding(items)
			formatterItem = append(formatterItem, formatterProduct)
	
		}
		formatter1 := CategoryLandingFormatter{
			ID : product.Category.ID,
			Name : product.Category.Name,
			Items : formatterItem,
		}
		formatterCategory = append(formatterCategory, formatter1)
	
		formatter.Category = formatterCategory
	var formatterTestimoni []TestimonialFormatter
	var testimonial domain.Testimonial
	testimonial.Name = "Happy Family"
	testimonial.ID = 1
	testimonial.ImageUrl = "/images/testimonial-landing-page.png"
	testimonial.Rate = 4.3
	testimonial.Content =  "What a great trip with my family and I should try again and again next time soon..."
	testimonial.FamilyName = "Budi"
	testimonial.FamilyOccupation = "Programmer"
	formatterTestimoni = append(formatterTestimoni, FormatTestimonial(testimonial))
	formatter.Testimonial = formatterTestimoni
	return formatter
}
