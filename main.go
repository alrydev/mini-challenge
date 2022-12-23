package main

import "fmt"

type Profile struct {
	Name   string
	Health int
	Power  int
	Exp    int
}

func MakeProfile(name string) Profile {
	var profile Profile
	if name == "Sasuke" {
		profile = Profile{
			Name:   name,
			Health: 200,
			Power:  100,
			Exp:    0,
		}
	} else if name == "Goku" {
		profile = Profile{
			Name:   name,
			Health: 400,
			Power:  300,
			Exp:    100,
		}
	} else if name == "Naruto" {
		profile = Profile{
			Name:   name,
			Health: 150,
			Power:  200,
			Exp:    50,
		}
	}

	return profile
}

func PowerUp(profile Profile, multiplier int) Profile {

	return Profile{
		Name:   profile.Name,
		Health: profile.Health + (profile.Health * multiplier),
		Power:  profile.Power + (profile.Power * multiplier),
		Exp:    profile.Exp + (profile.Exp * multiplier),
	}
}

func main() {
	profile := MakeProfile("Goku")
	fmt.Println("Name:", profile.Name)
	fmt.Println("Health:", profile.Health)
	fmt.Println("Power:", profile.Power)
	fmt.Println("Exp:", profile.Exp)
	fmt.Println("Heroes Power Up")
	powerUp := PowerUp(profile, 2)
	fmt.Println("Name : ", powerUp.Name)
	fmt.Println("Health : ", powerUp.Health)
	fmt.Println("Power : ", powerUp.Power)
	fmt.Println("Exp : ", powerUp.Exp)
}
