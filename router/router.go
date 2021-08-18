package router

import (
	"PoPic/database"
	"PoPic/model"
	"PoPic/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/url"
	"os"
	"path"
	"strings"
	"time"
)

func Routers(r *gin.Engine) *gin.Engine {
	// database
	db := database.InitMysql()
	db.Debug()

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"Title": utils.GetConfig("common.name"),
		})
	})

	// api router
	v1 := r.Group("/api/v1")
	v1.POST("/config", func(c *gin.Context) {
		response := model.NewResponse()
		key := c.DefaultPostForm("key", "")
		response.Message = "success"
		response.Result = utils.GetConfig(key)
		c.JSON(http.StatusOK, response)
	})

	v1.POST("/auth", func(c *gin.Context) {
		response := model.NewResponse()
		password := c.DefaultPostForm("password", "")
		isEncrypt := c.DefaultPostForm("isEncrypt", "0")
		if isEncrypt == "1" {
			password, _ = url.QueryUnescape(password) //url解码
			password = utils.AesDecrypt(password, utils.GetConfig("encrypter.key"))
		}
		if password == utils.GetConfig("admin.password") {
			response.Message = "success"
			response.Result = utils.AesEncrypt(password, utils.GetConfig("encrypter.key"))
		} else {
			response.Code = 500
			response.Message = "the password is wrong"
		}
		c.JSON(http.StatusOK, response)
	})

	//// 新建/修改计算类型
	v1.POST("/file", func(c *gin.Context) {
		response := model.NewResponse()
		f, err := c.FormFile("file")
		if err != nil {
			response.Code = 500
			response.Message = "获取文件失败"
			fmt.Println(err)
			c.JSON(http.StatusOK, response)
			return
		}
		fileExt := strings.ToLower(path.Ext(f.Filename))
		if fileExt != ".png" && fileExt != ".jpg" && fileExt != ".gif" && fileExt != ".jpeg" {
			response.Code = 500
			response.Message = "上传失败!只允许png,jpg,gif,jpeg文件"
			c.JSON(http.StatusOK, response)
			return
		}
		fileName := time.Now().Format("20060102150405")
		corrPath, _ := os.Getwd() //获取项目的执行路径
		fileDir := corrPath + "/web" + utils.GetConfig("upload.cover")
		filepath := fmt.Sprintf("%s%s%s", fileDir, fileName, fileExt)
		if err := c.SaveUploadedFile(f, filepath); err != nil {
			fmt.Println(err)
			response.Code = 500
			response.Message = "图片上传失败"
		} else {
			response.Message = "图片上传成功"
			response.Result = fmt.Sprintf("%s%s", fileName, fileExt)
		}
		c.JSON(http.StatusOK, response)
	})

	// 计算方式
	//v1.POST("/types", func(c *gin.Context) {
	//	response := model.NewResponse()
	//	var mTypes []model.Type
	//	db.Find(&mTypes)
	//	jsonType2, err := json.Marshal(mTypes)
	//	if err != nil {
	//		fmt.Println(err)
	//	}
	//	response.Message = "success"
	//	response.Result = string(jsonType2)
	//	c.JSON(http.StatusOK, response)
	//})
	//
	//// 关联字段
	//v1.POST("/fields", func(c *gin.Context) {
	//	response := model.NewResponse()
	//	typeId := c.DefaultPostForm("typeId", "")
	//	typeIdInt, err := strconv.Atoi(typeId)
	//	if err != nil {
	//		fmt.Println(err)
	//	}
	//	var mFields []model.Field
	//	db.Where(&model.Field{TId: typeIdInt}).Find(&mFields)
	//
	//	jsonFields, err := json.Marshal(mFields)
	//	if err != nil {
	//		fmt.Println(err)
	//	}
	//	response.Message = "success"
	//	response.Result = string(jsonFields)
	//	c.JSON(http.StatusOK, response)
	//})
	//
	//// 计算记录
	//v1.POST("/records", func(c *gin.Context) {
	//	response := model.NewResponse()
	//	dateForm := c.DefaultPostForm("dateForm", "")
	//	dateTo := c.DefaultPostForm("dateTo", "")
	//	typeId := c.DefaultPostForm("typeId", "")
	//	typeIdInt, err := strconv.Atoi(typeId)
	//	if err != nil {
	//		fmt.Println(err)
	//	}
	//
	//	var mRecords []model.Record
	//	if dateForm != "" && dateTo != "" {
	//		db.Where(&model.Record{TId: typeIdInt}).Where("created_at BETWEEN ? AND ?", dateForm, dateTo).Find(&mRecords)
	//	} else {
	//		db.Where(&model.Record{TId: typeIdInt}).Find(&mRecords)
	//	}
	//
	//	jsonRecords, err := json.Marshal(mRecords)
	//	if err != nil {
	//		fmt.Println(err)
	//	}
	//	response.Message = "success"
	//	response.Result = string(jsonRecords)
	//	c.JSON(http.StatusOK, response)
	//})
	//
	//// 新建/修改计算类型
	//v1.POST("/type/update", func(c *gin.Context) {
	//	response := model.NewResponse()
	//	code := c.DefaultPostForm("code", "")
	//	desc := c.DefaultPostForm("desc", "")
	//	cover := c.DefaultPostForm("cover", "")
	//	remarks := c.DefaultPostForm("remarks", "")
	//	status := c.DefaultPostForm("status", "")
	//
	//	mType := model.Type{Code: code, Desc: desc, Cover: cover, Remarks: remarks}
	//	if status != "" {
	//		mType.Status = status
	//	}
	//	typeId := c.DefaultPostForm("typeId", "")
	//	result := db
	//	if typeId != "" {
	//		typeIdInt, err := strconv.Atoi(typeId)
	//		if err != nil {
	//			fmt.Println(err)
	//		}
	//		//struct转map，否则0值字段更新时会被忽略
	//		mTypeMap := utils.ModelStructToMap(mType)
	//		result = db.Model(&model.Type{}).Where("ID = ?", typeIdInt).Updates(mTypeMap)
	//	} else {
	//		result = db.Create(&mType)
	//	}
	//	if result.RowsAffected > 0 {
	//		response.Message = "success"
	//	} else {
	//		response.Message = "insert type failed"
	//	}
	//	c.JSON(http.StatusOK, response)
	//})
	//
	//// 新建/修改计算类型
	//v1.POST("/type/cover/update", func(c *gin.Context) {
	//	response := model.NewResponse()
	//	f, err := c.FormFile("file")
	//	if err != nil {
	//		response.Code = 500
	//		response.Message = "获取文件失败"
	//		fmt.Println(err)
	//		c.JSON(http.StatusOK, response)
	//		return
	//	}
	//	fileExt := strings.ToLower(path.Ext(f.Filename))
	//	if fileExt != ".png" && fileExt != ".jpg" && fileExt != ".gif" && fileExt != ".jpeg" {
	//		response.Code = 500
	//		response.Message = "上传失败!只允许png,jpg,gif,jpeg文件"
	//		c.JSON(http.StatusOK, response)
	//		return
	//	}
	//	fileName := time.Now().Format("20060102150405")
	//	corrPath, _ := os.Getwd() //获取项目的执行路径
	//	fileDir := corrPath + "/web" + utils.GetConfig("upload.cover")
	//	filepath := fmt.Sprintf("%s%s%s", fileDir, fileName, fileExt)
	//	if err := c.SaveUploadedFile(f, filepath); err != nil {
	//		fmt.Println(err)
	//		response.Code = 500
	//		response.Message = "图片上传失败"
	//	} else {
	//		response.Message = "图片上传成功"
	//		response.Result = fmt.Sprintf("%s%s", fileName, fileExt)
	//	}
	//	c.JSON(http.StatusOK, response)
	//})
	//
	//// 新建/修改关联字段
	//v1.POST("/field/update", func(c *gin.Context) {
	//	response := model.NewResponse()
	//	typeId := c.DefaultPostForm("typeId", "")
	//	typeIdInt, err := strconv.Atoi(typeId)
	//	if err != nil {
	//		fmt.Println(err)
	//	}
	//
	//	code := c.DefaultPostForm("code", "")
	//	desc := c.DefaultPostForm("desc", "")
	//	fieldType := c.DefaultPostForm("fieldType", "0") //是否为计算字段
	//	formula := c.DefaultPostForm("formula", "")      //公式
	//	decimal := c.DefaultPostForm("decimal", "")      //保留小数位数
	//	orderNum := c.DefaultPostForm("orderNum", "")    //字段排序
	//	status := c.DefaultPostForm("status", "")
	//
	//	fieldTypeInt, err := strconv.Atoi(fieldType)
	//	if err != nil {
	//		fmt.Println(err)
	//	}
	//	decimalInt, err := strconv.Atoi(decimal)
	//	if err != nil {
	//		fmt.Println(err)
	//	}
	//	orderNumInt, err := strconv.Atoi(orderNum)
	//	if err != nil {
	//		fmt.Println(err)
	//	}
	//
	//	mField := model.Field{TId: typeIdInt, Code: code, Desc: desc, FieldType: fieldTypeInt, Formula: formula, Decimal: decimalInt, OrderNum: orderNumInt}
	//	if status != "" {
	//		mField.Status = status
	//	}
	//	fieldId := c.DefaultPostForm("fieldId", "")
	//	result := db
	//	if fieldId != "" {
	//		fieldIdInt, err := strconv.Atoi(fieldId)
	//		if err != nil {
	//			fmt.Println(err)
	//		}
	//		//struct转map，否则0值字段更新时会被忽略
	//		mFieldMap := utils.ModelStructToMap(mField)
	//		result = db.Model(&model.Field{}).Where("ID = ?", fieldIdInt).Updates(mFieldMap)
	//	} else {
	//		result = db.Create(&mField)
	//	}
	//	if result.RowsAffected > 0 {
	//		response.Message = "success"
	//	} else {
	//		response.Code = 500
	//		response.Message = "insert field failed"
	//	}
	//	c.JSON(http.StatusOK, response)
	//})
	//// 删除类型
	//v1.POST("/type/delete", func(c *gin.Context) {
	//	response := model.NewResponse()
	//	typeId := c.DefaultPostForm("typeId", "")
	//	result := db
	//	if typeId != "" {
	//		typeIdInt, err := strconv.Atoi(typeId)
	//		if err != nil {
	//			fmt.Println(err)
	//		}
	//		result = db.Delete(&model.Type{}, typeIdInt)
	//	}
	//	if result.RowsAffected > 0 {
	//		response.Message = "success"
	//	} else {
	//		response.Code = 500
	//		response.Message = "delete field failed"
	//	}
	//	c.JSON(http.StatusOK, response)
	//})
	//// 删除字段
	//v1.POST("/field/delete", func(c *gin.Context) {
	//	response := model.NewResponse()
	//	fieldId := c.DefaultPostForm("fieldId", "")
	//	result := db
	//	if fieldId != "" {
	//		fieldIdInt, err := strconv.Atoi(fieldId)
	//		if err != nil {
	//			fmt.Println(err)
	//		}
	//		result = db.Delete(&model.Field{}, fieldIdInt)
	//	}
	//	if result.RowsAffected > 0 {
	//		response.Message = "success"
	//	} else {
	//		response.Code = 500
	//		response.Message = "delete field failed"
	//	}
	//	c.JSON(http.StatusOK, response)
	//})
	//
	//// 详细记录
	//v1.POST("/records/fields", func(c *gin.Context) {
	//	type AllRecordFields struct {
	//		Record model.Record
	//		Fields []model.RecordField
	//	}
	//	response := model.NewResponse()
	//	typeId := c.DefaultPostForm("typeId", "1")
	//	var mRecords []model.Record
	//	if typeId != "" {
	//		typeIdInt, err := strconv.Atoi(typeId)
	//		if err != nil {
	//			fmt.Println(err)
	//		}
	//		db.Where(&model.Record{TId: typeIdInt}).Find(&mRecords)
	//	}
	//	var recordIDArr []int //切片
	//	for _, Record := range mRecords {
	//		recordIDArr = append(recordIDArr, int(Record.Model.ID))
	//	}
	//	var mRecordFields []model.RecordField
	//	db.Where("r_id in ?", recordIDArr).Find(&mRecordFields)
	//	jsonRecords, err := json.Marshal(mRecordFields)
	//	if err != nil {
	//		fmt.Println(err)
	//	}
	//	response.Message = "success"
	//	response.Result = string(jsonRecords)
	//	c.JSON(http.StatusOK, response)
	//})
	//// 新建/修改记录
	//v1.POST("/record/update", func(c *gin.Context) {
	//	response := model.NewResponse()
	//	typeId := c.DefaultPostForm("typeId", "")
	//	typeIdInt, err := strconv.Atoi(typeId)
	//	if err != nil {
	//		fmt.Println(err)
	//	}
	//	partDesc := c.DefaultPostForm("partDesc", "")
	//
	//	record := model.Record{TId: typeIdInt, PartDesc: partDesc}
	//	fieldId := c.DefaultPostForm("fieldId", "")
	//	result := db
	//	if fieldId != "" {
	//		fieldIdInt, err := strconv.Atoi(fieldId)
	//		if err != nil {
	//			fmt.Println(err)
	//		}
	//		//struct转map，否则0值字段更新时会被忽略
	//		recordMap := utils.ModelStructToMap(record)
	//		result = db.Model(&model.Record{}).Where("ID = ?", fieldIdInt).Updates(recordMap)
	//	} else {
	//		result = db.Create(&record)
	//	}
	//	if result.RowsAffected > 0 {
	//		response.Message = "success"
	//		response.Result = strconv.Itoa(int(record.Model.ID))
	//	} else {
	//		response.Code = 500
	//		response.Message = "insert field failed"
	//	}
	//	c.JSON(http.StatusOK, response)
	//})
	//
	//// 新建/修改记录字段
	//v1.POST("/record/field/update", func(c *gin.Context) {
	//	response := model.NewResponse()
	//	recordId := c.DefaultPostForm("recordId", "")
	//	recordIdInt, err := strconv.Atoi(recordId)
	//	if err != nil {
	//		fmt.Println(err)
	//	}
	//	fieldId := c.DefaultPostForm("fieldId", "")
	//	fieldIdInt, err := strconv.Atoi(fieldId)
	//	if err != nil {
	//		fmt.Println(err)
	//	}
	//	numValue := c.DefaultPostForm("numValue", "")
	//
	//	recordField := &model.RecordField{RId: recordIdInt, FId: fieldIdInt, NumValue: numValue}
	//	result := db
	//	recordFieldID := c.DefaultPostForm("recordFieldID", "")
	//	if recordFieldID != "" {
	//		recordFieldIDInt, err := strconv.Atoi(recordFieldID)
	//		if err != nil {
	//			fmt.Println(err)
	//		}
	//		//struct转map，否则0值字段更新时会被忽略
	//		recordFieldMap := utils.ModelStructToMap(recordField)
	//		result = db.Model(&model.RecordField{}).Where("ID = ?", recordFieldIDInt).Updates(recordFieldMap)
	//	} else {
	//		result = db.Create(&recordField)
	//	}
	//	if result.RowsAffected > 0 {
	//		response.Message = "success"
	//		response.Result = strconv.Itoa(int(recordField.Model.ID))
	//	} else {
	//		response.Code = 500
	//		response.Message = "insert failed"
	//	}
	//	c.JSON(http.StatusOK, response)
	//})
	//
	//// 新建/修改记录字段
	//v1.POST("/record/delete", func(c *gin.Context) {
	//	response := model.NewResponse()
	//	recordId := c.DefaultPostForm("recordId", "")
	//	recordIdInt, err := strconv.Atoi(recordId)
	//	if err != nil {
	//		fmt.Println(err)
	//	}
	//	result := db
	//	record := &model.Record{}
	//	result = db.Delete(&record, recordIdInt)
	//	result = db.Where("r_id = ?", recordIdInt).Delete(&model.RecordField{})
	//	if result.RowsAffected > 0 {
	//		response.Message = "success"
	//		response.Result = strconv.Itoa(int(record.Model.ID))
	//	} else {
	//		response.Code = 500
	//		response.Message = "insert failed"
	//	}
	//	c.JSON(http.StatusOK, response)
	//})
	return r
}
