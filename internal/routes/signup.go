package routes

type Signup struct{
	Username string `json:username`
	Email string `json:email`
	Password string `json:password`
}

var input Signup

func SignupHandler(con *fiber.ctx)error{
	if err:=con.BodyParser(&input);err!=nil{
		return con.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":"Invalid JSON",
		})
	}
	if input.Username==""||input.Email==""||input.Password==""{
		return con.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":"All fields are required",
		})
	}

	hash,err:=bcrypt.GenerateFromPassword([]byte(input.Password),12)
	if err!=nil{
		return con.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":"Failed to process password",
		})
	}
	return con.Redirect("/login")
}