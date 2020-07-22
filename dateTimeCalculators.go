/*
Author: Jasmin A. Smith
Date: 06/25/2020
*/

package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// Calculator 1
func timeDuration() {
	var intTime1, intTime2 []int
	var inputError error
	var errorCheck, resume1, resume2 bool
	var firstTime, secondTime, operation string
	in := bufio.NewReader(os.Stdin)

	// Gets the first time from user
Message1:
	fmt.Println("\nPlease give me the number of days, hours, minutes, and seconds seperated by spaces.")
	fmt.Println("Example: 3 5 15 0 represents 3 days, 5 hours, 15 minutes, and 0 seconds")
	fmt.Print("User's response: ")
	firstTime, err := in.ReadString('\n')
	firstTime = strings.TrimRight(firstTime, "\r\n")
	firstTime = strings.ToLower(firstTime)
	errorCheck = checkError(err)
	if errorCheck == true {
		goto Message1
	}
	if firstTime == "clear" {
		goto Message1
	}
	if firstTime == "exit" || firstTime == "quit" {
		os.Exit(1)
	}
	// spilts the input into an array
	r := regexp.MustCompile(` `)
	stringTime1 := r.Split(firstTime, -1)
	size1 := len(stringTime1)

	// checks every input has a value
	if size1 != 4 {
		if size1 > 4 {
			inputError = errors.New("You put in too many inputs. Try again.")
		} else {
			inputError = errors.New("You did not put in enough inputs. Try again.")
		}
		fmt.Println(inputError)
		goto Message1
	}

	//  Turns string array into integers
	intTime1, resume1 = convertingArray(stringTime1)
	if resume1 == false {
		fmt.Println("Your input needs to contain only integers! Try again.")
		goto Message1
	}

	// Prompts the second time
Message2:
	fmt.Println("\nPlease give me another set of days, hours, minutes, and seconds seperated by spaces.")
	fmt.Println("Example: 7 20 50 10 represents 7 days, 20 hours, 50 minutes, and 10 seconds")
	fmt.Print("User's response: ")
	secondTime, err = in.ReadString('\n')
	secondTime = strings.TrimRight(secondTime, "\r\n")
	secondTime = strings.ToLower(secondTime)
	errorCheck = checkError(err)
	if errorCheck == true {
		goto Message2
	}
	if secondTime == "clear" {
		goto Message1
	}
	if secondTime == "exit" || secondTime == "quit" {
		os.Exit(1)
	}

	// spilts input into array
	stringTime2 := r.Split(secondTime, -1)
	size2 := len(stringTime2)

	// making sure all inputs has a value
	if size2 != 4 {
		if size2 > 4 {
			inputError = errors.New("You put in too many inputs. Try again.")
		} else {
			inputError = errors.New("You did not put in enough inputs. Try again.")
		}
		fmt.Println(inputError)
		goto Message2
	}

	// turns strings into integers
	intTime2, resume2 = convertingArray(stringTime2)
	if resume2 == false {
		fmt.Println("Your input needs to contain only integers! Try again.")
		goto Message2
	}

	// gets opertaion
Message3:
	fmt.Println("\nWould you like to perform addition or subtraction?")
	fmt.Println("Enter addition for additon or subtraction for subtraction.")
	fmt.Print("User's Input: ")
	operation, err = in.ReadString('\n')
	operation = strings.TrimRight(operation, "\r\n")
	operation = strings.ToLower(operation)
	errorCheck = checkError(err)
	if errorCheck == true {
		goto Message2
	}

	if operation == "subtraction" {
		difference(intTime1, intTime2)
	} else if operation == "addition" {
		addition(intTime1, intTime2)
	} else if operation == "clear" {
		goto Message1
	} else if operation == "exit" || operation == "quit" {
		os.Exit(1)
	} else {
		fmt.Println("That was not a valid option. Try again!")
		goto Message3
	}
}

// calulates total seconds in a day
func totalSeconds(array []int) int {
	var total int

	size := len(array)
	if size == 3 {
		total = array[0]*86400 + array[1]*3600 + array[2]*31536000
	} else if size == 4 {
		total = array[0]*86400 + array[1]*3600 + array[2]*60 + array[3]
	}

	return total
}

// takes difference of times from first calculator
func difference(first []int, second []int) {
	time1 := totalSeconds(first)
	time2 := totalSeconds(second)
	seconds := time1 - time2
	mathSeconds := float64(seconds)
	minutes := mathSeconds / 60
	hours := mathSeconds / 3600
	days := mathSeconds / 86400

	formalDifference(seconds)

	fmt.Printf("%0.6f days\n", days)
	fmt.Printf("%0.5f hours\n", hours)
	fmt.Printf("%0.3f minutes\n", minutes)
	fmt.Println(seconds, "seconds")
}

// formally puts all information into one line
func formalDifference(diff int) {
	var days, hours, minutes, seconds int
	days = diff / 86400
	daysRem := diff % 86400
	hours = daysRem / 3600
	hoursRem := daysRem % 3600
	minutes = hoursRem / 60
	seconds = minutes % 60

	fmt.Printf("\n%v days %v hours %v minutes %v seconds\n", days, hours, minutes, seconds)
}

// takes adds times from the first calculator
func addition(first []int, second []int) {
	time1 := totalSeconds(first)
	time2 := totalSeconds(second)
	seconds := time2 + time1
	mathSeconds := float64(seconds)
	minutes := mathSeconds / 60
	hours := mathSeconds / 3600
	days := mathSeconds / 86400

	formalAddition(seconds)

	fmt.Printf("%0.6f days\n", days)
	fmt.Printf("%0.5f hours\n", hours)
	fmt.Printf("%0.3f minutes\n", minutes)
	fmt.Println(seconds, "seconds")
}

func formalAddition(add int) {
	var days, hours, minutes, seconds int
	days = add / 86400
	daysRem := add % 86400
	hours = daysRem / 3600
	hoursRem := daysRem % 3600
	minutes = hoursRem / 60
	seconds = hoursRem % 60

	fmt.Printf("\n%v days %v hours %v minutes %v seconds\n", days, hours, minutes, seconds)
}

func timeFromDate() {
	var date, time1, time2, timeOfDay, operator string
	var errorCheck bool
	var err error
	in := bufio.NewReader(os.Stdin)

Message1:
	fmt.Println("\nPlease enter your date in the form of mm/dd/yyyy")
	fmt.Println("For example, 12/1/2021 represents December 1, 2021")
	fmt.Print("User's response: ")
	date, err = in.ReadString('\n')
	date = strings.TrimRight(date, "\r\n")
	date = strings.ToLower(date)
	errorCheck = checkError(err)
	if errorCheck == true {
		goto Message1
	}
	if date == "clear" {
		goto Message1
	}
	if date == "exit" || date == "quit" {
		os.Exit(1)
	}
	r := regexp.MustCompile(`/`)
	stringDate := r.Split(date, -1)
	dateSize := len(stringDate)

	if dateSize != 3 {
		if dateSize > 4 {
			inputError1 := errors.New("You put in too many inputs. Try again.")
			fmt.Println(inputError1)
			goto Message1
		} else if dateSize == 1 {
			inputError1 := errors.New("You did not use any backslashs '/'. Try again.")
			fmt.Println(inputError1)
			goto Message1
		} else {
			inputError1 := errors.New("You did not put in enough inputs. Try again.")
			fmt.Println(inputError1)
			goto Message1
		}
	}
	intDate, resume1 := convertingArray(stringDate)
	if resume1 == false {
		fmt.Println("Your input needs to contain only integers! Try again.")
		goto Message1
	}
Message2:
	fmt.Println("\nPlease enter the time in the form of hh:mm:ss")
	fmt.Println("For example, 12:04:00")
	fmt.Print("User's input: ")
	time1, err = in.ReadString('\n')
	time1 = strings.TrimRight(time1, "\r\n")
	time1 = strings.ToLower(time1)
	errorCheck = checkError(err)
	if errorCheck == true {
		goto Message2
	}
	if time1 == "clear" {
		goto Message1
	}
	if time1 == "exit" || time1 == "quit" {
		os.Exit(1)
	}
	re := regexp.MustCompile(`:`)
	stringTime := re.Split(time1, -1)
	timeSize := len(stringTime)

	if timeSize != 3 {
		if timeSize > 3 {
			inputError2 := errors.New("You put in too many inputs. Try again.")
			fmt.Println(inputError2)
			goto Message2
		} else if timeSize == 1 {
			inputError2 := errors.New("You did not use any semicolons ':'. Try again.")
			fmt.Println(inputError2)
			goto Message2
		} else {
			inputError2 := errors.New("You did not put in enough inputs. Try again.")
			fmt.Println(inputError2)
			goto Message2
		}
	}

	intTime, resume2 := convertingArray(stringTime)
	if resume2 == false {
		fmt.Println("Your input needs to contain only integers! Try again.")
		goto Message2
	}

Message3:
	fmt.Print("\nWhat time of day AM or PM? ")
	timeOfDay, err = in.ReadString('\n')
	timeOfDay = strings.TrimRight(timeOfDay, "\r\n")
	timeOfDay = strings.ToLower(timeOfDay)
	errorCheck = checkError(err)
	if errorCheck == true {
		goto Message2
	}

	if timeOfDay == "pm" {
		if intTime[0] != 12 {
			intTime[0] += 12
		} else {
			intTime[0] = intTime[0]
		}

	} else if timeOfDay == "am" {
		if intTime[0] == 12 {
			intTime[0] -= 12
		} else {
			intTime[0] = intTime[0]
		}
	} else if timeOfDay == "clear" {
		goto Message1
	} else if timeOfDay == "exit" || timeOfDay == "quit" {
		os.Exit(1)
	} else {
		fmt.Println("You did not put in a valid input")
		goto Message3
	}

	fullDate := time.Date(intDate[2], time.Month(intDate[0]), intDate[1], intTime[0], intTime[1], intTime[2], int(0), time.Local)

Message4:
	fmt.Println("\nPlease enter desired time to be taken off with spaces inbetween")
	fmt.Println("For example, 5 3 30 represents 5 days 3 hours and 30 minutes")
	fmt.Print("User input: ")
	time2, err = in.ReadString('\n')
	time2 = strings.TrimRight(time2, "\r\n")
	time2 = strings.ToLower(time2)
	if errorCheck == true {
		goto Message4
	}
	if time2 == "clear" {
		goto Message1
	}
	if time2 == "exit" || time2 == "quit" {
		os.Exit(1)
	}
	reg := regexp.MustCompile(` `)
	stringTime2 := reg.Split(time2, -1)
	timeSize2 := len(stringTime2)
	if timeSize2 != 3 {
		if timeSize2 > 3 {
			inputError3 := errors.New("You put in too many inputs. Try again.")
			fmt.Println(inputError3)
			goto Message4
		} else if timeSize2 == 1 {
			inputError3 := errors.New("You did not use any spaces' '. Try again.")
			fmt.Println(inputError3)
			goto Message4
		} else {
			inputError3 := errors.New("You did not put in enough inputs. Try again.")
			fmt.Println(inputError3)
			goto Message4
		}
	}
	intTime2, resume3 := convertingArray(stringTime2)
	if resume3 == false {
		fmt.Println("Your input needs to contain only integers! Try again.")
		goto Message4
	}

Message5:
	fmt.Println("\nWould you like to do addition or subtraction? ")
	fmt.Println("Enter addition for addition or subtraction for subtraction.")
	fmt.Print("User's response: ")
	operator, err = in.ReadString('\n')
	operator = strings.TrimRight(operator, "\r\n")
	operator = strings.ToLower(operator)
	errorCheck = checkError(err)
	if errorCheck == true {
		goto Message2
	}
	if operator == "addition" {
		addTime(fullDate, intTime2)
	} else if operator == "subtraction" {
		subtractTime(fullDate, intTime2)
	} else if operator == "clear" {
		goto Message1
	} else if operator == "exit" || operator == "quit" {
		os.Exit(1)
	} else {
		fmt.Println("You did not put in a valid option. Try again!")
		goto Message5
	}

}
func addTime(date time.Time, timeArray []int) {
	days := timeArray[0]
	hours := timeArray[1]
	minutes := timeArray[2]

	date = date.AddDate(0, 0, days)
	date = date.Add(time.Hour * time.Duration(hours))
	date = date.Add(time.Minute * time.Duration(minutes))

	fmt.Println("\nYour new date is:", date)
}
func subtractTime(date time.Time, timeArray []int) {
	days := timeArray[0] + 1
	hours := timeArray[1]
	minutes := timeArray[2]

	date = date.AddDate(0, 0, -days)
	date = date.Add(-time.Hour * time.Duration(hours))
	date = date.Add(-time.Minute * time.Duration(minutes))

	fmt.Println("\nYour new date is:", date)
}
func ageCalculator() {
	var date1, date2 string
	var err error
	var errorCheck bool
	in := bufio.NewReader(os.Stdin)

Message1:
	fmt.Println("\nPlease enter your date in the form of mm/dd/yyyy")
	fmt.Println("For example, 09/1/1994 represents September 1, 1994")
	fmt.Print("User's response: ")
	date1, err = in.ReadString('\n')
	date1 = strings.TrimRight(date1, "\r\n")
	date1 = strings.ToLower(date1)
	errorCheck = checkError(err)
	if errorCheck == true {
		goto Message1
	}
	if date1 == "clear" {
		goto Message1
	}
	if date1 == "exit" || date1 == "quit" {
		os.Exit(1)
	}
	r := regexp.MustCompile(`/`)
	stringDate1 := r.Split(date1, -1)
	dateSize1 := len(stringDate1)

	if dateSize1 != 3 {
		if dateSize1 > 3 {
			inputError1 := errors.New("You put in too many inputs. Try again.")
			fmt.Println(inputError1)
			goto Message1
		} else if dateSize1 == 1 {
			inputError1 := errors.New("You did not use any backslashs '/'. Try again.")
			fmt.Println(inputError1)
			goto Message1
		} else {
			inputError1 := errors.New("You did not put in enough inputs. Try again.")
			fmt.Println(inputError1)
			goto Message1
		}
	}
	intDate1, resume1 := convertingArray(stringDate1)
	if resume1 == false {
		fmt.Println("Your input needs to contain only integers! Try again.")
		goto Message1
	}

Message2:
	fmt.Println("\nPlease enter your date in the form of mm/dd/yyyy")
	fmt.Println("For example, 12/1/2021 represents December 1, 2021")
	fmt.Print("User's response: ")
	date2, err = in.ReadString('\n')
	date2 = strings.TrimRight(date2, "\r\n")
	date2 = strings.ToLower(date2)
	errorCheck = checkError(err)
	if errorCheck == true {
		goto Message2
	}
	if date2 == "clear" {
		goto Message1
	}
	if date1 == "exit" || date1 == "quit" {
		os.Exit(1)
	}
	stringDate2 := r.Split(date2, -1)
	dateSize2 := len(stringDate2)

	if dateSize2 != 3 {
		if dateSize2 > 3 {
			inputError2 := errors.New("You put in too many inputs. Try again.")
			fmt.Println(inputError2)
			goto Message2
		} else if dateSize2 == 1 {
			inputError2 := errors.New("You did not use any backslashs '/'. Try again.")
			fmt.Println(inputError2)
			goto Message2
		} else {
			inputError2 := errors.New("You did not put in enough inputs. Try again.")
			fmt.Println(inputError2)
			goto Message2
		}
	}
	intDate2, resume2 := convertingArray(stringDate2)
	if resume2 == false {
		fmt.Println("Your input needs to contain only integers! Try again.")
		goto Message2
	}

	fullDate1 := time.Date(intDate1[2], time.Month(intDate1[0]), intDate1[1], int(0), int(0), int(0), int(0), time.Local)
	fullDate2 := time.Date(intDate2[2], time.Month(intDate2[0]), intDate2[1], int(0), int(0), int(0), int(0), time.Local)
	getAge(fullDate1, fullDate2)
}
func getAge(date1 time.Time, date2 time.Time) {
	var diff time.Duration

	if date1.After(date2) {
		diff = date1.Sub(date2)
	} else {
		diff = date2.Sub(date1)
	}
	years := int(diff.Hours()) / 24 / 365
	monthsRem := int(diff.Hours()) / 24 / 365 % 12
	daysRem := (int(diff.Hours()) - (years * 365 * 24)) / 12 % 7
	fmt.Println(years, "years", monthsRem, "months", daysRem, "days")
	months := (years * 12) + monthsRem
	rem := (int(diff.Hours()) - (years * 365 * 24)) / 12 % 7
	fmt.Println(months, "months", rem, "days")
	weeks := int(diff.Hours()) / 24 / 7
	daysLeft := int(diff.Hours()) / 24 % 7
	fmt.Println(weeks, "weeks", daysLeft, "days")
	numDays := int(diff.Hours() / 24)
	fmt.Println(numDays, "days")
	fmt.Println(diff.Hours(), "hours")
	fmt.Println(diff.Minutes(), "minutes")
	fmt.Println(diff.Seconds(), "seconds")
}
func convertingArray(array []string) ([]int, bool) {
	var newArray []int
	for _, i := range array {
		j, err := strconv.Atoi(i)
		errorCheck := checkError(err)
		if errorCheck == true {
			return newArray, false
		}
		newArray = append(newArray, j)
	}
	return newArray, true

}
func checkError(err error) bool {
	errorMessage := errors.New("You did not put in a valid input. Try again!")
	if err != nil {
		fmt.Println(errorMessage)
		return true
	}
	return false

}

func Menu() {
	var calculate bool = true
	var input string
	var errorCheck bool
	var err error
	in := bufio.NewReader(os.Stdin)

	for calculate != false {
	Welcome:
		fmt.Println("Here are your calculators!")
		fmt.Println("1. Time Duration Calculator")
		fmt.Println("2. Time From Date Calculator")
		fmt.Println("3. Age Calculator")
		fmt.Println("4. Quit")
		fmt.Print("Choose a number 1-3 to use our calculators or choose 4 to exit! ")
		input, err = in.ReadString('\n')
		input = strings.TrimRight(input, "\r\n")
		errorCheck = checkError(err)
		if errorCheck == true {
			goto Welcome
		}

		if input == "1" {
			timeDuration()
		} else if input == "2" {
			timeFromDate()
		} else if input == "3" {
			ageCalculator()
		} else if input == "4" {
			fmt.Println("Thank you for calculating with us! Hope to see you soon!")
			calculate = false
		} else {
			fmt.Println("Invalid input! Try again.")
			goto Welcome
		}

	}
}

func main() {
	Menu()
}
