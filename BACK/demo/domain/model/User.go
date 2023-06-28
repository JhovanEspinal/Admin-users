package model

type User struct {
	Id 		    string  `json:"_id" bson:"_id"`
	Name   	    string  `json:"name" bson:"name"`
	Cc 		    int     `json:"cc" bson:"cc"`
	Age         int     `json:"age" bson:"age"`
	Gender      string  `json:"gender" bson:"gender"`
	Job         string  `json:"job" bson:"job"`
	Description string  `json:"description" bson:"description"`
	Email       string  `json:"email" bson:"email"`
	Password    string  `json:"password" bson:"password"`
	Img         string  `json:"img" bson:"img"`


}

type RegisteredUser struct {
	Email 	 string `json:"email" bson:"email"`
	Password string	`json:"password" bson:"password"`
}
