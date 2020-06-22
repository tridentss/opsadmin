package service

import (
	"errors"
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)

// @title    GetEnvList
// @description   get env list by pagination, 分页获取数据
// @auth                      （2020/04/05  20:22）
// @param     info             request.PageInfo
// @return    err              error
// @return    list             interface{}
// @return    total            int

func GetEnvList(info request.PageInfo) (err error, list interface{}, total int) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.Model(&model.DepolyEnv{})
	var envList []model.DepolyEnv
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&envList).Error
	return err, envList, total
}

// @title    EnvCreate
// @description   create base env, 新增环境
// @auth                     （2020/06/22  10:28）
// @param     api             model.DepolyEnv
// @return                    error

func EnvCreate(env model.DepolyEnv) (err error) {
	findOne := global.GVA_DB.Where("name = ?", env.Name).Find(&model.DepolyEnv{}).Error
	if findOne == nil {
		return errors.New("存在相同环境")
	} else {
		err = global.GVA_DB.Create(&env).Error
	}
	return err
}

// @title    EnvUpdate
// @description   更新环境
// @auth                     （2020/04/05  20:22）
// @param     env             model.DepolyEnv
// @return                    error

func EnvUpdate(env model.DepolyEnv) (err error) {
	err = global.GVA_DB.Where("id = ?", env.ID).First(&model.DepolyEnv{}).Updates(&env).Error
	return err
}

// @title    EnvDelete
// @description    删除环境
// @auth                     （2020/04/05  20:22）
// @param     id
// @return    err             error

func EnvDelete(id float64) (err error) {
	var env model.DepolyEnv
	err = global.GVA_DB.Where("id = ?", id).Delete(&env).Error
	return err
}
