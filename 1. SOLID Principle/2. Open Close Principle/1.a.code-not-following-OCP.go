package main

import "fmt"

// This code is not following OCP as if any new type of customer profile will come in future
// then InsurancePremiumDiscountCalculator code will be changed so we are violating the
// principle i.e closed for modification and open for extensibility

type InsurancePremiumDiscountCalculator struct {
}

func (ipdc InsurancePremiumDiscountCalculator) calPremiumDiscountPercent1(customer HealthInsuranceCustomerProfile) int {
	if customer.isLoyalCustomer() {
		return 20
	}
	return 0
}

func (ipdc InsurancePremiumDiscountCalculator) calPremiumDiscountPercent2(customer VehicleInsuranceCustomerProfile) int {
	if customer.isLoyalCustomer() {
		return 20
	}
	return 0
}

type HealthInsuranceCustomerProfile struct {
}

func (hicp HealthInsuranceCustomerProfile) isLoyalCustomer() bool {
	return true
}

type VehicleInsuranceCustomerProfile struct {
}

func (vicp VehicleInsuranceCustomerProfile) isLoyalCustomer() bool {
	return false
}

func main() {
	insPremDiscCal := InsurancePremiumDiscountCalculator{}
	healthInsCustomer := HealthInsuranceCustomerProfile{}
	vehicleInsCustomer := VehicleInsuranceCustomerProfile{}

	dis := insPremDiscCal.calPremiumDiscountPercent1(healthInsCustomer)
	fmt.Printf("Discount for health insurance customer %d", dis)

	dis = insPremDiscCal.calPremiumDiscountPercent2(vehicleInsCustomer)
	fmt.Printf("\nDiscount for health insurance customer %d", dis)
}
