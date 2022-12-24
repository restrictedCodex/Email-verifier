package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("domain, hasMX, hasSPF, spfRecords, hasDMARCH, dmarcRecords \n")

	for scanner.Scan() {
		checkDomain(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal("Could Not read from the Input: \n", err)
	}
}

func checkDomain(domain string) {
	var hasMX, hasSPF, hasDMARCH bool
	var spfRecords, dmarcRecords string

	mxRecord, err := net.LookupMX(domain)
	if err != nil {
		log.Println(err)
	}

	if len(mxRecord) > 0 {
		hasMX = true
	}

	txtRecord, err := net.LookupTXT(domain)
	if err != nil {
		log.Println(err)
	}

	for _, record := range txtRecord {
		if strings.HasPrefix(record, "v=spf1") {
			hasSPF = true
			spfRecords = record
			break
		}
	}

	dmarcRecord, err := net.LookupTXT("_dmarc." + domain)
	if err != nil {
		log.Println(err)
	}

	for _, record := range dmarcRecord {
		if strings.HasPrefix(record, "v=DMARC1") {
			hasDMARCH = true
			dmarcRecords = record
			break
		}
	}

	fmt.Printf("%v, %v, %v, %v", hasMX, hasSPF, spfRecords, "\n")
	fmt.Printf("%v, %v, %v", hasDMARCH, dmarcRecords, "\n")
}
