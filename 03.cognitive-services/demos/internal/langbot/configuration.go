package langbot

import (
	"github.com/proge-software/lectio-csml-csbot/internal/tgconf"
	"github.com/proge-software/lectio-csml-csbot/pkg/wssface"
	"github.com/proge-software/lectio-csml-csbot/pkg/wssformrecognizer"
	"github.com/proge-software/lectio-csml-csbot/pkg/wsssentiment"
	"github.com/proge-software/lectio-csml-csbot/pkg/wsstranslator"
	"github.com/proge-software/lectio-csml-csbot/pkg/wssvision"
)

//Configuration Bot's Configuration
type Configuration struct {
	TelegramConf       tgconf.Configuration
	FaceConf           *wssface.Configuration
	VisionConf         *wssvision.Configuration
	TextAnalyticsConf  *wsssentiment.Configuration
	TranslatorConf     *wsstranslator.Configuration
	FormRecognizerConf *wssformrecognizer.Configuration
}
