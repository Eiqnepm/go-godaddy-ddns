package main

import (
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/Eiqnepm/go-godaddy-ddns/internal/api"
	"github.com/Eiqnepm/go-godaddy-ddns/internal/mullvad"
)

func envDefault(key string, defaultValue string) (value string) {
	value = os.Getenv(key)
	if value == "" {
		value = defaultValue
	}

	return
}

func envRequired(key string) (value string) {
	value = os.Getenv(key)
	if value == "" {
		log.Fatalf("%v environment variable required", key)
	}

	return
}

func main() {
	api := api.Credentials{
		Key:    envRequired("API_KEY"),
		Secret: envRequired("API_SECRET"),
	}
	domain := envRequired("DOMAIN")
	name := envDefault("NAME", "@")
	ipv4 := envDefault("IPV4", "true")
	ipv6 := envDefault("IPV6", "false")
	t, err := strconv.Atoi(envDefault("TIMEOUT", "300"))
	if err != nil {
		log.Fatal(err)
	}

	timeout := time.Duration(t) * time.Second

	var previous struct {
		ipv4 string
		ipv6 string
	}
	first := true
	for {
		if !first {
			time.Sleep(timeout)
		}
		first = false

		if strings.ToLower(ipv4) == "true" {
			ip, err := mullvad.GetIP("https://ipv4.am.i.mullvad.net/json")
			if err != nil {
				log.Println(err)
				continue
			}

			if ip != previous.ipv4 {
				err = api.PutRecord(domain, "A", name, ip)
				if err != nil {
					log.Println(err)
					continue
				}

				previous.ipv4 = ip
			}
		}

		if strings.ToLower(ipv6) == "true" {
			ip, err := mullvad.GetIP("https://ipv6.am.i.mullvad.net/json")
			if err != nil {
				log.Println(err)
				continue
			}

			if ip != previous.ipv6 {
				err = api.PutRecord(domain, "AAAA", name, ip)
				if err != nil {
					log.Println(err)
					continue
				}

				previous.ipv6 = ip
			}
		}
	}
}
