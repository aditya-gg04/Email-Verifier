package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)



func main(){
	fmt.Printf("enter email to verify: \n")
	fmt.Printf("domain, hasMX, hasSPF, hasDMARC, dmarcRecords, spfRecords\n")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		email := scanner.Text()
		isValidEmail(email)
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}

func isValidEmail(email string)  {
	var hasMX, hasSPF, hasDMARC bool
	var dmarcRecords, spfRecords string
	//mxrecords
	mxRecords, err := net.LookupMX(email)
	if err == nil && len(mxRecords) > 0 {
		hasMX = true
	}

	//spfrecords 
	spfRecordsList, err := net.LookupTXT(email)
	if err ==nil{
		for _ , text:= range spfRecordsList{
			if strings.HasPrefix(text, "v=spf1"){
				hasSPF = true
				spfRecords = text
				break
			}
		}
	}


	//dmarcrecords
	dmarcRecordsList, err := net.LookupTXT(email)
	if err ==nil{
		for _ , text:= range dmarcRecordsList{
			if strings.HasPrefix(text, "v=dmarc1"){
				hasDMARC = true
				dmarcRecords = text
				break
			}
		}
	}


	fmt.Printf("Email: %s, hasMX: %t, hasSPF: %t, hasDMARC: %t, dmarcRecords: %s, spfRecords: %s\n", email, hasMX, hasSPF, hasDMARC, dmarcRecords, spfRecords)

}