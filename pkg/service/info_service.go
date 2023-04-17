package services

import (
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"net/http"
	"sync"
	"time"
	"zoro/pkg/model"
)

func (sv *MyServer) GetDb(ctx *gin.Context) {
	keyCache := "chien"
	data, err := sv.R.Get(keyCache).Result()
	var dataInfo []model.Info

	if err != nil {
		var wgScan sync.WaitGroup

		wgScan.Add(1) // 2 job scan data
		var infoChannel = make(chan []model.Info, 1)
		go AsyncRawDataInfo(sv, &wgScan, infoChannel)
		wgScan.Wait()

		dataInfo = <-infoChannel

		dataCache, err := json.Marshal(dataInfo)
		if err == nil {
			err := sv.R.Set(keyCache, string(dataCache), time.Duration(1)*time.Hour).Err()
			if err != nil {
				log.Printf("Cache event time scan erorr [%v]", err)
			}
		}
	} else {
		log.Printf("Get Cache")
		err := json.Unmarshal([]byte(data), &dataInfo)
		if err != nil {
			ctx.JSON(http.StatusOK, "")
			return
		}
	}

	ctx.JSON(http.StatusOK, dataInfo)
}

func AsyncRawDataInfo(sv *MyServer, wgrp *sync.WaitGroup, result chan<- []model.Info) {
	defer func() {
		close(result)
		wgrp.Done()
	}()

	var infosModel []model.Info
	sqlResponse := sv.H.DB.Where("(? = ? )", 1, 1).Find(&infosModel)
	if errors.Is(sqlResponse.Error, gorm.ErrRecordNotFound) {
		log.Printf("RawInfo err [%v]", sqlResponse.Error)
		result <- []model.Info{}
		return
	}
	result <- infosModel
}

func (sv *MyServer) GetAllDb(ctx *gin.Context) {
	var wgScan sync.WaitGroup
	wgScan.Add(2) // 2 job scan data
	var infoChannel = make(chan []model.Info, 1)
	var purchaseChannel = make(chan []model.Purchase, 1)
	go AsyncRawDataInfo(sv, &wgScan, infoChannel)
	go AsyncRawDataPurchase(sv, &wgScan, purchaseChannel)
	wgScan.Wait()

	dataInfo := <-infoChannel
	dataPurchase := <-purchaseChannel
	ctx.JSON(http.StatusOK, gin.H{
		"info":     dataInfo,
		"purchase": dataPurchase,
	})
}

func AsyncRawDataPurchase(sv *MyServer, wgrp *sync.WaitGroup, result chan<- []model.Purchase) {
	defer func() {
		close(result)
		wgrp.Done()
	}()

	var infosModel []model.Purchase
	sqlResponse := sv.H.DB.Where("(? = ? )", 1, 1).Find(&infosModel)
	if errors.Is(sqlResponse.Error, gorm.ErrRecordNotFound) {
		log.Printf("RawInfo err [%v]", sqlResponse.Error)
		result <- []model.Purchase{}
		return
	}
	result <- infosModel
}
