**go-register-gmail-otp**
--------------------------------
A mini project that integrates with the Google API, allowing users to register using their Google email and a one-time password (OTP).

**Env Files**
---
```
Change the name of the .env.example file to .env and update the values within the .env file.
```

**Pull / Download in ur Local**
---
Clone this repository into ur Local
```
git clone https://github.com/VincentLimarus/go-register-gmail-otps.git
```
Go to the Directory
```
cd go-register-gmail-otps
```
Install Dependency
```
go get .
```
Run Server 
```
go run .
```

**Endpoints**
---
```
http://localhost:3000/api/v1/auth/register

{
    "email" : "user@gmail.com"
}
```
```
http://localhost:3000/api/v1/auth/register/verify
{
    "email" : "user@gmail.com",
    "otp" : "466538"
}

make sure the otp is correct (same with the OTP in Gmail)
```

**Key Takeaways**
---
A user who registers in the system using Gmail will be initially added to the database with the `Is_Active status set to False`. However, once the user completes registration with an OTP, the `Is_Active status will be updated to True`.
