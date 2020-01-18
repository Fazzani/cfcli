package service

import (
	"encoding/json"
	"errors"

	. "github.com/fazzani/cfcli/domain"

	log "github.com/sirupsen/logrus"

	req "github.com/parnurzeal/gorequest"
)

type IZoneService interface {
	FindAll() ([]Zone, error)
	Get(zoneId int) (Zone, error)
	Create(z Zone) (bool, error)
	Update(z Zone) (bool, error)
	Delete(zone_id string) (bool, error)
	RessourceName() string
}

type HttpZoneService struct {
	Config Configuration
}

func (s HttpZoneService) RessourceName() string {
	return "zones"
}

// https://api.cloudflare.com/client/v4/zones?name=synker.ovh&status=active&page=1&per_page=20&order=status&direction=desc&match=all
func (s HttpZoneService) FindAll() ([]Zone, error) {
	url := s.Config.API.BaseURL + s.RessourceName() + "?name=synker.ovh&status=active&page=1&per_page=20&order=status&direction=desc&match=any"
	res, errs := makeRequest(url, "GET", s.Config.Authentication.AccountID, s.Config.Authentication.Email, s.Config.Authentication.Key)
	if errs != nil {
		log.Errorln(errs)
		return nil, errors.New("HttpRequest error")
	}
	zones := []Zone{}
	err := json.Unmarshal([]byte(res), &zones)
	return zones, err
}

//https://api.cloudflare.com/client/v4/zones/{{zone_id}}
func (s HttpZoneService) Get(zoneId string) (Zone, error) {
	url := s.Config.API.BaseURL + s.RessourceName() + "/" + zoneId
	res, errs := makeRequest(url, "GET", s.Config.Authentication.AccountID, s.Config.Authentication.Email, s.Config.Authentication.Key)
	zone := Zone{}

	if errs != nil {
		log.Errorln(errs)
		return zone, errors.New("HttpRequest error")
	}

	err := json.Unmarshal([]byte(res), &zone)
	return zone, err
}

//https://api.cloudflare.com/client/v4/zones/{{zone_id}}
/*
	{
		"name": "example.com",

		"jump_start": true,false,
		"organization":{
			"id":"{{org_id}}"
		}
	}
*/
func (s HttpZoneService) Create(z Zone) (bool, error) {
	url := s.Config.API.BaseURL + s.RessourceName() + "/" + z.ID
	res, errs := makeRequest(url, "POST", s.Config.Authentication.AccountID, s.Config.Authentication.Email, s.Config.Authentication.Key)
	if errs != nil {
		log.Errorln(errs)
		return false, errors.New("HttpRequest error")
	}
	err := json.Unmarshal([]byte(res), &z)
	return err == nil, err
}

//https://api.cloudflare.com/client/v4/zones/{{zone_id}}
/*
		{
	"paused":true,false,
	"vanity_name_servers":[
	  "ns1.example.com","ns2.example.com"],
	"plan":{
	  "id":{{plan_id}}
	}
	}
*/
func (s HttpZoneService) Update(z Zone) (bool, error) {
	url := s.Config.API.BaseURL + s.RessourceName() + "/" + z.ID
	res, errs := makeRequest(url, "PATCH", s.Config.Authentication.AccountID, s.Config.Authentication.Email, s.Config.Authentication.Key)
	if errs != nil {
		log.Errorln(errs)
		return false, errors.New("HttpRequest error")
	}
	err := json.Unmarshal([]byte(res), &z)
	return err == nil, err
}

//https://api.cloudflare.com/client/v4/zones/{{zone_id}}
func (s HttpZoneService) Delete(zoneId string) (bool, error) {
	url := s.Config.API.BaseURL + s.RessourceName() + "/" + zoneId
	res, errs := makeRequest(url, "DELETE", s.Config.Authentication.AccountID, s.Config.Authentication.Email, s.Config.Authentication.Key)
	if errs != nil {
		log.Errorln(errs)
		return false, errors.New("HttpRequest error")
	}
	zone := Zone{}
	err := json.Unmarshal([]byte(res), &zone)
	return err == nil, err
}

func makeRequest(url string, method string, accountId string, email string, key string) (string, []error) {
	if method == "" {
		method = "GET"
	}

	request := req.New()
	_, body, errs := request.CustomMethod(method, url).
		Set("Content-Type", "application/json").
		Set("account_id", accountId).
		Set("x-auth-email", email).
		Set("x-auth-key", key).
		End()
	return body, errs
}
