package common

import (
	"fmt"
	"strings"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	"github.com/go-playground/locales/zh_Hant"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zh_tw_translations "github.com/go-playground/validator/v10/translations/zh_tw"
)

var Validate *validator.Validate
var (
	ZH    ut.Translator
	ZH_TW ut.Translator
)

type Translator struct {
	tag         string
	translation string
	override    bool
}

func init() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		Validate = v
	} else {
		Validate = validator.New()
	}

	en := en.New()
	zh_Hant := zh_Hant.New()
	zh := zh.New()
	uni := ut.New(en, zh_Hant, zh)

	ZH, _ = uni.GetTranslator("zh")
	ZH_TW, _ = uni.GetTranslator("zh_Hant")

	AddTranslations([]Translator{
		{tag: "required", translation: "{0} 是必須的"},
		{tag: "max", translation: "{0} 長度必須是 {1} 或更短"},
		{tag: "min", translation: "{0} 長度必須至少為 {1} 個字元"},
		{tag: "alphanum", translation: "{0} 只能包含字母和數字"},
		{tag: "lowercase", translation: "{0} 只能包含小寫字母"},
	}, ZH_TW)

	zh_tw_translations.RegisterDefaultTranslations(Validate, ZH_TW)
}

func TranslateError(lang string, err error) map[string]string {
	translated, lang_map := map[string]string{}, GetLangTranslator(lang)
	if err == nil {
		return translated
	}
	if errs, ok := err.(validator.ValidationErrors); ok {
		for _, e := range errs {
			translated[strings.ToLower(e.Field())] = e.Translate(lang_map)
		}
	} else {
		translated["error"] = err.Error()
	}
	return translated
}

func GetLangTranslator(lang string) ut.Translator {
	switch strings.ToLower(lang) {
	case "zh_tw":
		return ZH_TW
	case "zh":
		return ZH
	default:
		return ZH_TW
	}
}

func AddTranslations(translations []Translator, trans ut.Translator) {
	for _, t := range translations {
		AddTranslation(t.tag, t.translation, trans, t.override)
	}
}

func AddTranslation(key, value string, trans ut.Translator, override bool) {
	Validate.RegisterTranslation(key, trans, func(ut ut.Translator) error {
		return ut.Add(key, value, override)
	}, func(ut ut.Translator, fe validator.FieldError) (t string) {
		t, _ = ut.T(key,
			strings.ToLower(fe.Field()),   // {0} Field
			fe.Param(),                    // {1} Param
			fe.Tag(),                      // {2} Tag
			fmt.Sprintf("%v", fe.Value()), // {3} Value
			fe.Kind().String(),            // {4} Kind
			fe.Type().String(),            // {5} Type
			fe.Namespace(),                // {6} Namespace
			fe.StructNamespace(),          // {7} StructNamespace
			fe.StructField(),              // {8} StructField
			fe.ActualTag(),                // {9} ActualTag
		)
		return
	})
}
