package db

import (
	"fmt"
	"log"

	"../dataStructure"
	"../function"
	"../timeData"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/guregu/dynamo"
)

//DBから単語を取得
func GetWordData(keyword string) string {
	log.Print(keyword)
	cred := credentials.NewStaticCredentials(ACCESS_KET_ID, SECRET_ACCESS_KEY, "") // 最後の引数は[セッショントークン]

	db := dynamodb.New(session.New(), &aws.Config{
		Credentials: cred,
		Region:      aws.String("ap-northeast-1"), // "ap-northeast-1"等
	})

	getParams := &dynamodb.ScanInput{
		TableName: aws.String("word"),

		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":word": { // :を付けるのがセオリーのようです
				S: aws.String(keyword),
			},
		},

		FilterExpression: aws.String("contains(word, :word)"),
	}

	getItem, getErr := db.Scan(getParams)

	if getErr != nil {
		panic(getErr)
	}
	fmt.Println(getItem)

	return function.ResWord(getItem, keyword)

}

// get Max Id
// func GetMaxID() {
// 	cred := credentials.NewStaticCredentials(ACCESS_KET_ID, SECRET_ACCESS_KEY, "") // 最後の引数は[セッショントークン]

// 	db := dynamo.New(session.New(), &aws.Config{
// 		Credentials: cred,
// 		Region:      aws.String("ap-northeast-1"), // "ap-northeast-1"等
// 	})

// }

//単語をデータベースに登録
func PutWord(word string, flag int) {
	cred := credentials.NewStaticCredentials(ACCESS_KET_ID, SECRET_ACCESS_KEY, "") // 最後の引数は[セッショントークン]

	db := dynamo.New(session.New(), &aws.Config{
		Credentials: cred,
		Region:      aws.String("ap-northeast-1"), // "ap-northeast-1"等
	})

	table := db.Table("history")

	history := dataStructure.History{HistoryID: 1, Time: timeData.GetNowTimeFormat(), Flag: flag}
	//u := User{User_ID: "lambda test"}
	fmt.Println(history)

	if err := table.Put(history).Run(); err != nil {
		fmt.Println("err")
		panic(err.Error())
	}

}

// func Insert() {

// 	table := db.Table("history")

// 	history := dataStructure.History{HistoryID: 1, Time: timeData.GetNowTimeFormat(), Flag: 3}
// 	//u := User{User_ID: "lambda test"}
// 	fmt.Println(history)

// 	if err := table.Put(history).Run(); err != nil {
// 		fmt.Println("err")
// 		panic(err.Error())
// 	}
// }
// func checkWord(value string, keyword string) bool {
// 	//文字列の先頭部分は末尾文字と一致しているかどうか
// 	//fmt.Println(strings.HasPrefix("ナツ", "ツナミ"))
// 	fmt.Println(strings.HasPrefix(value, keyword))
// 	return strings.HasPrefix(value, keyword)
// }

// type WordDB struct {
// 	WordID int    `json:"word_id" dynamodbav:"word_id"`
// 	Word   string `json:"word" dynamodbav:"word"`
// }

// type Words struct {
// 	WordID int    `json:"word_id"`
// 	Word   string `json:"word"`
// }

// func resWord(output *dynamodb.ScanOutput, keyword string) string {
// 	// DBから取得したデータのJSONの形を変換
// 	words := make([]*WordDB, 0)
// 	unMarshaListOfMapErr := dynamodbattribute.UnmarshalListOfMaps(output.Items, &words)
// 	if unMarshaListOfMapErr != nil {
// 		panic(fmt.Sprintf("failed to unmarshal Dynamodb Scan Items, %v", unMarshaListOfMapErr))
// 	}

// 	bytes, _ := json.Marshal(words)

// 	//変換されたデータ形をパースし、取得
// 	var data []Words
// 	unMarshaErr := json.Unmarshal(bytes, &data)
// 	if unMarshaErr != nil {
// 		fmt.Println("error:", unMarshaErr)
// 	}

// 	for _, word := range data {
// 		fmt.Printf("word_id: %v, word: %v\n", word.WordID, word.Word)
// 		isWord := checkWord(word.Word, keyword)
// 		if isWord == true {
// 			return word.Word
// 		}
// 	}
// 	return ""
// }
