package service

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"go.mongodb.org/mongo-driver/bson"

	"parking/repository"
)

type CommandInput struct {
	LengthOfParam int
	Action        string
	StParam       string
	NdParam       string
}

func (impl *implementation) Stdin(input CommandInput) string {
	var result string

	switch input.Action {
	case "create_parking_lot":
		result = impl.CreateParking(input.LengthOfParam, input.StParam)
	case "parking_lot":
		result = impl.CreateParking(input.LengthOfParam, input.StParam)
	case "park":
		result = impl.Park(input.LengthOfParam, input.StParam, input.NdParam)
	case "status":
		impl.Status(input.LengthOfParam)
	case "leave":
		result = impl.Leave(input.LengthOfParam, input.StParam)
	case "registration_numbers_for_cars_with_colour":
		result = impl.RegistrationNumbersWithColour(input.LengthOfParam, input.StParam)
	case "slot_numbers_for_cars_with_colour":
		result = impl.SlotNumbersWithColour(input.LengthOfParam, input.StParam)
	case "slot_number_for_registration_number":
		result = impl.SlotNumberForRegistrationNumber(input.LengthOfParam, input.StParam)
	case "exit":
		result = "exit"
	}

	return result
}

func (impl *implementation) CreateParking(lengthOfParam int, stParam string) string {
	slot, err := strconv.ParseUint(stParam, 10, 16)
	if lengthOfParam != 2 || err != nil {
		return ""
	}

	err = impl.Create(int(slot))
	if err != nil {
		return err.Error()
	}

	result := fmt.Sprintf("Created a parking lot with %v slots", slot)
	return result
}

func (impl *implementation) Park(lengthOfParam int, stParam, ndParam string) string {
	if lengthOfParam != 3 {
		return ""
	}
	id, err := impl.Parking(&repository.Park{
		Car:    stParam,
		Colour: ndParam,
	})
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			return "Sorry, parking lot is full"
		}
		return err.Error()
	}

	result := fmt.Sprintf("Allocated slot number: %s", id)
	return result
}

func (impl *implementation) Status(lengthOfParam int) {
	if lengthOfParam != 1 {
		return
	}
	park, err := impl.List()
	if err != nil {
		return
	}

	header := fmt.Sprintf("%s  %s  %s", "Slot No.", "Registration No", "Colour")
	fmt.Println(header)
	for _, item := range park {
		if item.Car == "" {
			continue
		}
		fmt.Println(fmt.Sprintf("%d         %s    %s", item.Index, item.Car, item.Colour))
	}

}

func (impl *implementation) Leave(lengthOfParam int, stParam string) string {
	if lengthOfParam != 2 {
		return ""
	}

	slot, err := strconv.ParseUint(stParam, 10, 16)
	if lengthOfParam != 2 || err != nil {
		return ""
	}

	err = impl.Delete(int(slot))
	if err != nil {
		return err.Error()
	}
	result := fmt.Sprintf("Slot number %v is free", slot)

	return result
}

func (impl *implementation) RegistrationNumbersWithColour(lengthOfParam int, stParam string) string {
	if lengthOfParam != 2 {
		return ""
	}

	if len(stParam) < 1 {
		return ""
	}

	filter := bson.M{"$and": bson.A{
		bson.M{"colour": stParam},
		bson.M{"deleted": bson.M{"$ne": "true"}},
	}}

	items, err := impl.repo.List(context.Background(), filter)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			return "Not found"
		}
		return err.Error()
	}

	result := make([]string, len(items))
	for i, item := range items {
		result[i] = item.Car
	}

	return strings.Join(result, ", ")
}

func (impl *implementation) SlotNumbersWithColour(lengthOfParam int, stParam string) string {
	if lengthOfParam != 2 {
		return ""
	}

	if len(stParam) < 1 {
		return ""
	}

	filter := bson.M{"$and": bson.A{
		bson.M{"colour": stParam},
		bson.M{"deleted": bson.M{"$ne": "true"}},
	}}
	items, err := impl.repo.List(context.Background(), filter)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			return "Not found"
		}
		return err.Error()
	}

	result := make([]string, len(items))
	for i, item := range items {
		result[i] = strconv.Itoa(item.Index)
	}

	return strings.Join(result, ", ")
}

func (impl *implementation) SlotNumberForRegistrationNumber(lengthOfParam int, stParam string) string {
	if lengthOfParam != 2 {
		return ""
	}

	if len(stParam) < 1 {
		return ""
	}

	car := &repository.Park{}
	filter := bson.M{"$and": bson.A{
		bson.M{"car": stParam},
		bson.M{"deleted": bson.M{"$ne": "true"}},
	}}
	err := impl.repo.Read(context.Background(), filter, car)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			return "Not found"
		}
		return err.Error()
	}

	return strconv.Itoa(car.Index)
}
