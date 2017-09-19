/*
* Go Library (C) 2017 Inc.
*
* @project     Ukkbox
* @package     main
* @author      @jeffotoni
* @size        16/07/2017
*
 */

package main

func test_dynamodb_() {

	// input := &dynamodb.QueryInput{

	// 	ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{

	// 		":name": {
	// 			S: aws.String(upload_name),
	// 		},
	// 	},

	// 	//FilterExpression: aws.String("upload_uid = :name"),
	// 	KeyConditionExpression: aws.String("upload_name = :name"),

	// 	//ProjectionExpression: aws.String("upload_uid"),

	// 	TableName: aws.String("table"),
	// }

	// result, err := svc.Query(input)

	// // var UkkUpLoad []UkkoboxUpload2

	// oneItem := []UkkoboxUpload2{}

	// if err == nil {

	// 	errdy := dynamodbattribute.UnmarshalListOfMaps(result.Items, &oneItem)

	// 	fmt.Println("fol err: ", errdy)
	// 	fmt.Println("fol: ", oneItem[0].Upload_uid)
	// }

	// var queryInput = &dynamodb.QueryInput{

	// 	TableName: aws.String(DyTable),

	// 	KeyConditions: map[string]*dynamodb.Condition{

	// 		"upload_uid": {

	// 			// IN, NULL, BETWEEN, LT, NOT_CONTAINS, EQ, GT, NOT_NULL, NE, LE, BEGINS_WITH, GE, CONTAINS
	// 			ComparisonOperator: aws.String("EQ"),

	// 			AttributeValueList: []*dynamodb.AttributeValue{
	// 				{
	// 					S: aws.String("NULL"),
	// 				},
	// 			},
	// 		},

	// 		"upload_name": {

	// 			ComparisonOperator: aws.String("EQ"),

	// 			AttributeValueList: []*dynamodb.AttributeValue{
	// 				{
	// 					S: aws.String(upload_name),
	// 				},
	// 			},
	// 		},
	// 	},
	// }
	//result, err := svc.Query(queryInput)

	// params_get := &dynamodb.GetItemInput{

	// 	Key: map[string]*dynamodb.AttributeValue{

	// 		"upload_name": {
	// 			S: aws.String(upload_name),
	// 		},
	// 	},

	// 	TableName: aws.String(DyTable), // Required
	// }

	// result, err := svc.GetItem(params_get)

	// input := &dynamodb.ScanInput{

	// 	ExpressionAttributeNames: map[string]*string{

	// 		"upload_uid": aws.String("upload_uid"),
	// 	},
	// 	ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
	// 		":v_name": {
	// 			S: aws.String(upload_name),
	// 		},
	// 	},

	// 	FilterExpression:     aws.String("upload_uid = :v_name"),
	// 	ProjectionExpression: aws.String("#upload_uid"),
	// 	TableName:            aws.String(DyTable),
	// }

	// result, err := svc.Scan(input)

	//fmt.Println("result: ", result)
	//fmt.Println("result error: ", err)

	//return ""
	// params := &dynamodb.QueryInput{

	// 	TableName: aws.String(DyTable), // Required
	// 	Limit:     aws.Int64(3),

	// 	// IndexName: aws.String("localSecondaryIndex"),
	// 	ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{

	// 		":v_name": { // Required
	// 			S: aws.String(upload_name),
	// 		},
	// 		":v_id": {
	// 			S: aws.String("NULL"),
	// 		},
	// 	},

	// 	FilterExpression: aws.String("upload_name >= :v_name"),

	// 	// KeyConditionExpression: aws.String("upload_id = :v_id"),

	// 	KeyConditions: map[string]*dynamodb.Condition{

	// 		"upload_name": { // Required

	// 			ComparisonOperator: aws.String("GT"), // Required
	// 			AttributeValueList: []*dynamodb.AttributeValue{
	// 				{ // Required
	// 					S: aws.String(upload_name),
	// 				},
	// 				// More values...
	// 			},
	// 		},
	// 		"upload_uid": { // Required

	// 			ComparisonOperator: aws.String("EQ"), // Required

	// 			// AttributeValueList: []*dynamodb.AttributeValue{
	// 			//  S: aws.String("NOT_NULL"),
	// 			// },

	// 		},

	// 		// More values...
	// 	},

	// 	//ProjectionExpression: aws.String("#v_name, #v_id"),

	// 	Select:           aws.String("ALL_ATTRIBUTES"),
	// 	ScanIndexForward: aws.Bool(true),
	// }

	// //Get the response and print it out.
	// resp, err := svc.Query(params)

	// fmt.Println("result: ", resp)
	// fmt.Println("result error: ", err)

	// return ""
	// var queryInput = &dynamodb.QueryInput{

	// 	TableName: aws.String(DyTable),

	// 	KeyConditions: map[string]*dynamodb.Condition{

	// 		"upload_name": {

	// 			ComparisonOperator: aws.String("EQ"),

	// 			AttributeValueList: []*dynamodb.AttributeValue{
	// 				{
	// 					S: aws.String(upload_name),
	// 				},
	// 			},
	// 		},
	// 	},
	// }

	// var resp, err = svc.Query(queryInput)

	// if err != nil {

	// 	fmt.Println("erro gery : ", err)
	// 	//return nil, err
	// }

	// fmt.Println("resp query: ", resp)

	// return ""

	// input := &dynamodb.GetItemInput{

	// 	Key: map[string]*dynamodb.AttributeValue{
	// 		"upload_name": {
	// 			S: aws.String(upload_name),
	// 		},
	// 	},

	// 	TableName: aws.String(DyTable),
	// }

	// result, err_get := svc.GetItem(input)

	// fmt.Println("Name file: ", upload_name)

	// fmt.Println("getItem result: ", result)

	// fmt.Println("getItem err: ", err_get)

	// oneItem := UkkoboxUpload{}

	// if err_get == nil {

	// 	err := dynamodbattribute.UnmarshalMap(result.Item, &oneItem)

	// 	fmt.Println("error get file: ", err)

	// 	if err == nil {

	// 		if oneItem.Upload_uid != "" {

	// 			return oneItem.Upload_uid

	// 		} else {

	// 			return ""
	// 		}

	// 	} else {

	// 		return ""
	// 	}

	// } else {

	// 	//fmt.Println("GetItem err is: ", err_get)

	// 	return ""
	// }

	// return ""
}
