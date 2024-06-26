package main

import "fmt"

type Employee struct {
	Id        int
	FirstName string
	LastName  string
}

type Developer struct {
	Individual Employee
	hourlyRate int
	WorkWeek   [7]int
}

type Weekday int

const (
	Sunday Weekday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

func (d *Developer) LogHours(day Weekday, hours int) {
	d.WorkWeek[day] = hours
}

func (d *Developer) HoursWorked() int {
	total := 0
	for _, v := range d.WorkWeek {
		total += v
	}
	return total
}

func nonLoggedHours() func(int) int {
	total := 0
	f := func(t int) int {
		total += t
		return total
	}
	return f
}

func (d *Developer) PayDay() (int, bool) {

	totalHoursWorked := d.HoursWorked()
	hasOvertime := false
	if totalHoursWorked > 40 {
		hasOvertime = true
		overTime := totalHoursWorked - 40
		overTimePay := overTime*2*d.hourlyRate + 40*d.hourlyRate
		return overTimePay, hasOvertime
	}
	return totalHoursWorked * d.hourlyRate, hasOvertime
}

func (d *Developer) PayDetails() {
	for i, v := range d.WorkWeek {
		fmt.Printf("Number of hours worked on %v: %d\n", i, v)
	}
}

func main() {
	ivieEdebiri := Developer{Individual: Employee{Id: 1, FirstName: "Ivie", LastName: "Edebiri"}, hourlyRate: 35}
	x := nonLoggedHours()
	fmt.Println("Tracking hours worked so far today: ", x(5))
	fmt.Println("Tracking hours worked so far today: ", x(3))
	fmt.Println("Tracking hours worked so far today: ", x(5))
	ivieEdebiri.LogHours(Monday, 12)
	ivieEdebiri.LogHours(Tuesday, 8)
	ivieEdebiri.LogHours(Wednesday, 10)
	ivieEdebiri.LogHours(Thursday, 12)
	pay, workedOvertime := ivieEdebiri.PayDay()
	if !workedOvertime {
		fmt.Println("Did not work overtime")
	} else {
		fmt.Println("Worked overtime")
	}
	ivieEdebiri.PayDetails()
	fmt.Printf("Total pay for the week: %d\n", pay)
	fmt.Printf("Hours worked on Monday: %d\n", ivieEdebiri.WorkWeek[Monday])
	fmt.Printf("Hours worked on Tuesday: %d\n", ivieEdebiri.WorkWeek[Tuesday])
	fmt.Printf("Hours worked this week %d\n", ivieEdebiri.HoursWorked())
}
