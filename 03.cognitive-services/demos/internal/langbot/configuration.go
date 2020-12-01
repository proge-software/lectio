package langbot

import (
	"github.com/proge-software/lectio/internal/tgconf"
	"github.com/proge-software/lectio/pkg/wssface"
	"github.com/proge-software/lectio/pkg/wssformrecognizer"
	"github.com/proge-software/lectio/pkg/wsssentiment"
	"github.com/proge-software/lectio/pkg/wsstranslator"
	"github.com/proge-software/lectio/pkg/wssvision"
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
