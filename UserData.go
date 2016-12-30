package main

var Account, Password string

func GetAccount() string  {
    return Account;
}

func GetPassWord() string {
    return Password;
}

func SetAccountInformation(account string, password string)  {
    Account     = account;
    Password    = password;
}