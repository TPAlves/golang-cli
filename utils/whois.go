package utils

import (
	"fmt"
	"log"
	"regexp"
	"strings"

	"github.com/likexian/whois"
)

func GetWhois(domain string) string {
	result, err := whois.Whois(domain)
	if err != nil {
		log.Fatal("Não foi possível exibir dados do domínio")
	}

	name := getNameServer(result)
	admin := getAdmin(result)
	date := getDateExpired(result)

	content := fmt.Sprintf("Domínio: %s\nTitular: %s\nExpiração: %s\n",
		name, admin, date)

	return content
}

func getNameServer(content string) string {
	name := keywordContent(content, "domain")
	return name
}

func getAdmin(content string) string {
	admin := keywordContent(content, "owner")
	return admin
}

func getDateExpired(content string) string {
	dateString := strings.Split(keywordContent(content, "expires"), "")
	date := ""
	for i, d := range dateString {
		if i == 4 || i == 6 {
			date += fmt.Sprintf("-%s", d)
		} else {
			date += d
		}
	}
	dateFormart := strings.Split(date, "-")

	if len(dateFormart) != 3 {
		log.Fatal("Não foi possível converter a data de expiração\n")
	}
	dateLocation := fmt.Sprintf("%s/%s/%s", dateFormart[2], dateFormart[1], dateFormart[0])
	return dateLocation
}

func keywordContent(content string, keyword string) string {
	re := regexp.MustCompile(fmt.Sprintf(`(?m)^%s:\s+.*`, keyword))
	match := re.FindStringSubmatch(content)
	content = strings.TrimSpace(
		strings.Replace(
			match[0],
			fmt.Sprintf("%s:", keyword),
			"",
			1))
	return content
}
