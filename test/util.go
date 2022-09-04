package test

import "github.com/jaswdr/faker"

var f faker.Faker

func init() {
	f = faker.New()
}

func InvalidRegion() string {
	return f.RandomLetter()
}

func ValidRegions() []string {
	return []string{
		"af-south-1",
		"eu-north-1",
		"ap-south-1",
		"eu-west-3",
		"eu-west-2",
		"eu-south-1",
		"eu-west-1",
		"ap-northeast-3",
		"ap-northeast-2",
		"me-south-1",
		"ap-northeast-1",
		"sa-east-1",
		"ca-central-1",
		"ap-east-1",
		"ap-southeast-1",
		"ap-southeast-2",
		"ap-southeast-3",
		"eu-central-1",
		"us-east-1",
		"us-east-2",
		"us-west-1",
		"us-west-2",
	}
}

func ValidRegion() string {
	regions := ValidRegions()

	return regions[f.IntBetween(0, len(regions)-1)]
}
