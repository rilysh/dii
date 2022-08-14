package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"os"
	"strconv"
)

type Response struct {
	Data struct {
		Geo struct {
			Host          string  `json:"host"`
			Asn           int32   `json:"asn"`
			Isp           string  `json:"isp"`
			Country       string  `json:"country_name"`
			CountryCode   string  `json:"country_code"`
			RegionName    string  `json:"region_name"`
			RegionCode    string  `json:"region_code"`
			City          string  `json:"city"`
			PostalCode    string  `json:"postal_code"`
			ContinentName string  `json:"continent_name"`
			ContinentCode string  `json:"continent_code"`
			Latitude      float32 `json:"latitude"`
			Longitude     float32 `json:"longitude"`
			MetroCode     string  `json:"metro_code"`
			Timezone      string  `json:"timezone"`
			Datetime      string  `json:"datetime"`
		} `json:"geo"`
	} `json:"data"`
}

func vaprt(msg string) {
	fmt.Printf("%s", msg)
}

func help() {
	vaprt(`
dii - Portable IP and DNS resolver
	
Usage
-----
dii -ip [IP ADDRESS], resolve IP query
dii -dns [WEBSITE URL], resolve DNS query
dii -h, shows this help
`)
}

func main() {
	none := "Unknown"
	color_start := "\033[0;38m"
	color_end := "\033[0m"
	arg := os.Args

	if len(arg) == 1 {
		help()
		return
	}
	switch arg[1] {
	case "-ip":
		if len(arg) == 2 {
			help()
			return
		}
		client := http.Client{}
		request, _ := http.NewRequest("GET", "https://tools.keycdn.com/geo.json?host="+arg[2], nil)
		request.Header = http.Header{
			"User-Agent": {"keycdn-tools:https://github.com/kiwimoe/dii"},
		}
		res, err := client.Do(request)
		if err != nil {
			vaprt("An error occurred, log: " + err.Error() + "\n")
			return
		}
		if res.StatusCode != 200 {
			vaprt("Failed to get results, status code: " + strconv.Itoa(res.StatusCode) + "\n")
			return
		}
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			vaprt("An error occurred, log: " + err.Error() + "\n")
			return
		}
		var resp Response
		json.Unmarshal(body, &resp)
		// Gp doesn't have ternary operation feature, therefore I need to use "if-else" as an eqavalent
		host := resp.Data.Geo.Host
		asn := resp.Data.Geo.Asn
		isp := resp.Data.Geo.Isp
		country := resp.Data.Geo.Country
		country_code := resp.Data.Geo.CountryCode
		region_name := resp.Data.Geo.RegionName
		region_code := resp.Data.Geo.RegionCode
		city := resp.Data.Geo.City
		postal_code := resp.Data.Geo.PostalCode
		continent_name := resp.Data.Geo.ContinentName
		continent_code := resp.Data.Geo.ContinentCode
		latitude := resp.Data.Geo.Latitude
		longitude := resp.Data.Geo.Longitude

		if asn == 0 {
			vaprt("No data found from the provided IP\n")
			return
		}
		if len(isp) == 0 {
			isp = none
		}
		if len(country) == 0 {
			country = none
			country_code = none
		}
		if len(region_name) == 0 {
			region_name = none
			region_code = none
		}
		if len(city) == 0 {
			city = none
		}
		if len(postal_code) == 0 {
			postal_code = none
		}
		if len(continent_name) == 0 {
			continent_name = none
			continent_code = none
		}
		fmt.Printf(`
		%s
---------Info------------
 Host: %s
 ASN: %d
 ISP: %s
 Country: %s
 Country Code: %s
 Region Name: %s
 Region Code: %s
 City: %s
 Postal Code: %s
 Continent Name: %s
 Continent Code: %s
 Latitude: %f
 Longitude: %f
-------------------------
%s
`, color_start,
			host,
			asn,
			isp,
			country,
			country_code,
			region_name,
			region_code,
			city,
			postal_code,
			continent_name,
			continent_code,
			latitude,
			longitude,
			color_end)
		break

	case "-dns":
		if len(arg) == 2 {
			help()
			return
		}
		// Ignore all errors
		ns, _ := net.LookupNS(arg[2])
		ip, _ := net.LookupIP(arg[2])
		txt, _ := net.LookupTXT(arg[2])
		mx, _ := net.LookupMX(arg[2])
		host_str := ""
		txt_str := ""
		mx_str := ""
		mx_uint := 0
		ip_list := ""

		if len(ns) == 0 {
			vaprt("Failed to resolve DNS for provided query\n")
			return
		}
		for _, name := range ns {
			host_str += name.Host + " "
		}
		for _, txtid := range txt {
			txt_str += txtid + " "
		}
		for _, mxid := range mx {
			mx_str += mxid.Host + " "
			mx_uint += int(mxid.Pref)
		}
		for _, ip_loc := range ip {
			ip_list += ip_loc.String() + " "
		}

		fmt.Printf(`
		%s
------------- Info -----------------
 NS Record(s): %s
 IPv4: %s
 IPv6: %s
 IPs: %s
 MX: %s
 MX Pref: %d
 TXT: %s
------------------------------------
%s
`,
			color_start,
			host_str,
			ip[0],
			ip[len(ip)-1],
			ip_list,
			mx_str,
			mx_uint,
			txt_str,
			color_end)
		break

	case "-h", "-help":
		help()
		break

	default:
		help()
		break
	}
}
