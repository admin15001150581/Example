package until

import (
	"github.com/dgrijalva/jwt-go"
	"example/pkg/setting"
	"time"

)
var jwtSecret = []byte(setting.Viper.GetString("JwtSecret"))

type Claims struct {
	username string`json:"username"`
	password string `json:"password"`
	jwt.StandardClaims
}

//生成token
func GenerateToken(username, password string) (string, error) {
	nowTime:=time.Now() //当前时间
	expireTime :=nowTime.Add(3 * time.Hour) //有效时间
	//结构体赋值
	claims :=Claims{
		 username,
	     password,
	     jwt.StandardClaims{
		 	ExpiresAt: expireTime.Unix(),
		 	Issuer: "example",
		 },
	}

	tokenClaims :=jwt.NewWithClaims(jwt.SigningMethodHS256,claims) //创建新的token
	token, err := tokenClaims.SignedString(jwtSecret) //获取完整的token
	return token,err
}
//验证token
func ParseToken(token string)(*Claims,error){
	tokenClaims,err:=jwt.ParseWithClaims(token,&Claims{},func(token *jwt.Token)(interface{},error){
		return jwtSecret,nil
	})
	if tokenClaims!=nil{
		if claims,ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid{
			return claims,nil
		}
	}
	return nil,err
}