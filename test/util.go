package test

import "github.com/jaswdr/faker"

var f faker.Faker

func init() {
	f = faker.New()
}

func Region() string {
	return f.RandomLetter()
}
