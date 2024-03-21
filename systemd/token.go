package systemd
import (
	"github.com/golang-jwt/jwt/v5"
)
type R4UCustomerClaims struct {
	UserID string
	UserName string
	jwt.RegisteredClaims
}
const sign_key = "Hbn23J&jmo37JH9JKn0h390HKHNfvXsKjkJejKJd#pla@JLK2JHJJHmLkjdsjI8dhjhhjkdsj3232"
func GenerateToken(claim R4UCustomerClaims)(string,error){
	token,err := jwt.NewWithClaims(jwt.SigningMethodHS256, claim).SignedString([]byte(sign_key))
	return token,err
}