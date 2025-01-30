package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	res := subdomainVisits([]string{
		"9001 discuss.leetcode.com",
	})
	for _, r := range res {
		fmt.Println(r)
	}
	fmt.Println("-----------------")

	res = subdomainVisits([]string{
		"900 google.mail.com", "50 yahoo.com", "1 intel.mail.com", "5 wiki.org",
	})
	for _, r := range res {
		fmt.Println(r)
	}
	fmt.Println("-----------------")
}

func subdomainVisits(cpdomains []string) []string {
	var res []string
	domainCount := make(map[string]int)

	for _, cp := range cpdomains {
		split := strings.Split(cp, " ")
		count, _ := strconv.Atoi(split[0])
		domains := strings.Split(split[1], ".")

		domain := ""
		for i := len(domains) - 1; i >= 0; i-- {
			domain = strings.TrimSuffix(domains[i]+"."+domain, ".")
			domainCount[domain] += count
		}
	}

	for domain, count := range domainCount {
		res = append(res, strconv.Itoa(count)+" "+domain)
	}

	return res
}
