package main

import (
	"time"
	"fmt"

	"github.com/hashgraph/hedera-sdk-go"
)

const server string = "testnet.hedera.com:51009"
const secret string = "302e020100300506032b657004220420d4d575f0a33cc860bcd205d1cb5efd2775614954fcc0fe484e81ba6117b0eeda"

func calCost(secs float64, price float64) float64 {
	return 6.66
}
//only use Account: 1001 for operator
func transferAmount(operatorAccountID hedera.AccountID, targetAccountID hedera.AccountID, amount int64){
	operatorSecret, err := hedera.SecretKeyFromString(secret)
	if err != nil{
		panic(err)
	}
	client, err := hedera.Dial(server)
	if err != nil{
		panic(err)
	}
	client.SetNode(hedera.AccountID{Account: 3})
	client.SetOperator(operatorAccountID, func() hedera.SecretKey{
		return operatorSecret
	})
	defer client.Close()
	opBalance, err := client.Account(operatorAccountID).Balance().Get()
	if err != nil{
		panic(err)
	}
	targBalance, err := client.Account(targetAccountID).Balance().Get()
	if err != nil{
		panic(err)
	}
	fmt.Printf("Target Balance: %v\nOp Balance %v\n",targBalance, opBalance)

	nodeAccountID := hedera.AccountID{Account: 3}
	response, err := client.TransferCrypto().Transfer(operatorAccountID, (0-amount) ).Transfer(targetAccountID, amount).Operator(operatorAccountID).Node(nodeAccountID).Memo("Test transfer 1 to 2").Sign(operatorSecret).Sign(operatorSecret).Execute()
	if err != nil{
		panic(err)
	}

	transId := response.ID
	fmt.Printf("transferred; transaction = %v\n", transId)

	time.Sleep(2* time.Second)

	receipt, err := client.Transaction(*transId).Receipt().Get()
	if err != nil{
		panic(err)
	}

	if receipt.Status != hedera.StatusSuccess{
		panic(fmt.Errorf("Transaction was not successful: %v", receipt.Status.String()))
	}
	time.Sleep(2*time.Second)

	targBalance, err = client.Account(targetAccountID).Balance().Get()
	if err != nil {
		panic(err)
	}
	opBalance,err = client.Account(operatorAccountID).Balance().Get()
	if err != nil{
		panic(err)
	}
	fmt.Printf("Target Balance: %v\nOp Balance %v\n",targBalance, opBalance)


}

func makeAccount(){
	operatorSecret, err := hedera.SecretKeyFromString(secret)
	if err != nil{
		panic(err)
	}

	secretKey, _ := hedera.GenerateSecretKey()
	public := secretKey.Public()

	fmt.Printf("secret = %v\n", secretKey)
	fmt.Printf("public = %v\n", public)

	client, err := hedera.Dial(server)
	if err !=nil{
		panic(err)
	}
	defer client.Close()

	nodeAccountID := hedera.AccountID{Account: 3}
	operatorAccountID := hedera.AccountID{Account: 1001}
	time.Sleep(2* time.Second)
	response, err := client.CreateAccount().Key(public).InitialBalance(0).Operator(operatorAccountID).Node(nodeAccountID).Memo("Test make Account").Sign(operatorSecret).Execute()
	if err != nil{
		panic(err)
	}

	transactionID := response.ID
	fmt.Printf("Created account; transaction = %v\n", transactionID)
	time.Sleep(2* time.Second)
    
	receipt,err := client.Transaction(*transactionID).Receipt().Get()
	if err != nil{
		panic(err)
	}
	fmt.Printf("Account = %v\n", *receipt.AccountID)

}



func getAccountBal(num hedera.AccountID) float64{
	accountID := num
	client, err := hedera.Dial(server)
	if err != nil{
		panic(err)
	}

	operatorAccountID := hedera.AccountID{Account: 1001}

	operatorSecret,err := hedera.SecretKeyFromString(secret)
	if err != nil{
		panic(err)
	}
	client.SetNode(hedera.AccountID{Account: 3})
	client.SetOperator(operatorAccountID, func() hedera.SecretKey {
		return operatorSecret
	})
	
	defer client.Close()

	balance, err := client.Account(accountID).Balance().Get()
	if err != nil{
		panic(err)
	}
	hbars := float64(balance)/100000000
	return hbars
}

// func main(){
// 	transferAmount(hedera.AccountID{Account: 1001}, hedera.AccountID{Account: 1002}, 100)
// 	bal := getAccontBal(hedera.AccountID{Account: 1001})
// 	bal2 := getAccontBal(hedera.AccountID{Account: 1002})
// 	fmt.Printf("Balance1: %.5f\nBalance2: %.5f\n\n", bal,bal2)
// }
