package main

import "fmt"

// This code is following OCP as if any new type of customer profile will come in future
// then InsurancePremiumDiscountCalculator code will not be changed so we are following the
// principle i.e closed for modification and open for extensibility

type InsurancePremiumDiscountCalculator struct {
}

type CustomerProfile interface {
	isLoyalCustomer() bool
}

func (ipdc InsurancePremiumDiscountCalculator) calPremiumDiscountPercent(customer CustomerProfile) int {
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

	dis := insPremDiscCal.calPremiumDiscountPercent(healthInsCustomer)
	fmt.Printf("Discount for health insurance customer %d", dis)

	dis = insPremDiscCal.calPremiumDiscountPercent(vehicleInsCustomer)
	fmt.Printf("\nDiscount for health insurance customer %d", dis)
}
