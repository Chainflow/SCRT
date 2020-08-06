package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/Chainflow/SCRT/config"
)

//ValidatorStatusAlert which sends the status(voting, jailed) of validator alert to telegram
func ValidatorStatusAlert(cfg *config.Config) error {
	log.Println("Coming inside validator status alerting")
	ops := HTTPOptions{
		Endpoint: cfg.LCDEndpoint + "/staking/validators/" + cfg.ValOperatorAddress,
		Method:   http.MethodGet,
	}

	resp, err := HitHTTPTarget(ops)
	if err != nil {
		log.Printf("Error: %v", err)
		return err
	}

	var validatorResp ValidatorResp
	err = json.Unmarshal(resp.Body, &validatorResp)
	if err != nil {
		log.Printf("Error: %v", err)
		return err
	}

	alertTime1 := cfg.AlertTime1
	alertTime2 := cfg.AlertTime2

	t1, _ := time.Parse(time.Kitchen, alertTime1)
	t2, _ := time.Parse(time.Kitchen, alertTime2)

	now := time.Now().UTC()
	t := now.Format(time.Kitchen)

	a1 := t1.Format(time.Kitchen)
	a2 := t2.Format(time.Kitchen)

	log.Println("a1, a2 and present time : ", a1, a2, t)

	if t == a1 || t == a2 {
		validatorStatus := validatorResp.Result.Jailed
		if !validatorStatus {
			_ = SendTelegramAlert(fmt.Sprintf("Your SCRT validator %s is currently voting", cfg.ValidatorName), cfg)
			log.Println("Sent validator status alert")
		} else {
			_ = SendTelegramAlert(fmt.Sprintf("Your SCRT validator %s is in jailed status", cfg.ValidatorName), cfg)
			log.Println("Sent validator status alert")
		}
	}
	return nil
}

//CheckValidatorJailed to send transaction alert to telegram
// when the validator will be jailed
func CheckValidatorJailed(cfg *config.Config) error {
	log.Println("Coming inside jailed alerting")
	ops := HTTPOptions{
		Endpoint: cfg.LCDEndpoint + "/staking/validators/" + cfg.ValOperatorAddress,
		Method:   http.MethodGet,
	}

	resp, err := HitHTTPTarget(ops)
	if err != nil {
		log.Printf("Error: %v", err)
		return err
	}

	var validatorResp ValidatorResp
	err = json.Unmarshal(resp.Body, &validatorResp)
	if err != nil {
		log.Printf("Error: %v", err)
		return err
	}

	validatorStatus := validatorResp.Result.Jailed
	if validatorStatus {
		_ = SendTelegramAlert(fmt.Sprintf("Your SCRT validator %s is in jailed status", cfg.ValidatorName), cfg)
		log.Println("Sent validator jailed status alert")
	}
	return nil
}