package request

type UserCreate struct {
	FullName string `json:"full_name"`
	Email    string `json:"email"`
}

type VerifyOtpRequest struct {
	Email        string `json:"email"`
	Otp          string `json:"otp"`
	CompanyName  string `json:"company_name"`
	EmployeeSize int    `json:"number_of_employees"`
	Password     string `json:"password"`
}
