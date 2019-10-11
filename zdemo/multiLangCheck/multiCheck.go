package main

/**
    多语言校验
 */
import (
	"github.com/gin-gonic/gin"
	en2 "github.com/go-playground/locales/en"
	zh2 "github.com/go-playground/locales/zh"
	"github.com/go-playground/universal-translator"
	"gopkg.in/go-playground/validator.v9"
	enTranslations "gopkg.in/go-playground/validator.v9/translations/en"
	zhTranslations "gopkg.in/go-playground/validator.v9/translations/zh"
	"net/http"
	"time"
)
var (
	Uni *ut.UniversalTranslator
	Validate *validator.Validate
	)
type PersonMultiLang struct {
	// 注意V9用的是validate 取代binding.其他属性一致
	Age int     `form:"age" validate:"required,gt=10"`
	Name string `form:"name" validate:"required"`
	Address string `form:"address"`
	Birthday time.Time `form:"birthday" time_format:"2006-01-02"`
}
func main(){
	// 升级版验证多语言错误信息 1.需要引入V9的包，该包才支持多语言验证
	multiLangChecker:=gin.Default()

	// 定义验证器
	Validate=validator.New()
	// 定义翻译器
	zh:=zh2.New()
	en:=en2.New()
	Uni=ut.New(zh,en)
	// 自定义验证规则
	multiLangChecker.GET("maltiLangCheck",func(gin *gin.Context){
		// 1.获取前端当前指定的语言类型标示 默认为中文环境
		localLanguage:=gin.DefaultQuery("localLanguge","zh")
		// 2.通过语言标示获取语言翻译器
		languageTrans,_:=Uni.GetTranslator(localLanguage)
		// 3.注册对应语言环境的验证器
		switch localLanguage{
		case "zh":
			zhTranslations.RegisterDefaultTranslations(Validate,languageTrans)
		case "en":
			enTranslations.RegisterDefaultTranslations(Validate,languageTrans)
		default:
		}
       // 4.绑定数据到结构体
		var personMultiLang PersonMultiLang
		if err:=gin.ShouldBind(&personMultiLang);err!=nil{
			gin.String(500,"%v",err)
			gin.Abort()
			return
		}
		// 5.用得到的验证器执行验证
        if err:=Validate.Struct(personMultiLang);err!=nil{
        	errs:=err.(validator.ValidationErrors)
        	errSlices:=[]string{}
        	// 将所有验证错误的信息封装到切片数组中返回
        	for _,e:=range errs{
        		errSlices=append(errSlices,e.Translate(languageTrans))
			}
        	gin.String(500,"%v",errSlices)
        	gin.Abort()
        	return
		}else{
			gin.String(http.StatusOK,"%v",personMultiLang)
		}
	})

	multiLangChecker.Run()
}
