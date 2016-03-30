package speech_to_text

import (
	"strings"
	"github.com/maxl28/watson-developer-cloud/watson"
	"github.com/fatih/structs"
	"strconv"
	"net/http"
	"errors"
	"io"
)

// Watson Speech-To-Text default endpoint URL.
const SpeechToTextAPI string = "https://stream.watsonplatform.net/speech-to-text/api"

// Available options for this Watson module
type SpeechToTextOptions struct {
	// Endpoint URL to use
	Url string
}

// ToDo: Write documentation
type RecognizeParameters struct {
	// The identifier of the session to be used.
	Session_id string

	// The audio to be transcribed in the format specified by the content_type parameter.
	Audio io.Reader

	// The MIME type of the audio:
	// - audio/flac
	// - audio/l16 (Also specify the rate and channels; for example, audio/l16; rate=48000; channels=2. Ensure that the rate matches the rate at which the audio is captured.)
	// - audio/wav
	// - audio/ogg;codecs=opus
	Content_type string

	// The identifier of the model to be used for the recognition request:
	// - ar-AR_BroadbandModel
	// - en-UK_BroadbandModel
	// - en-UK_NarrowbandModel
	// - en-US_BroadbandModel (the default)
	// - en-US_NarrowbandModel
	// - es-ES_BroadbandModel
	// - es-ES_NarrowbandModel
	// - ja-JP_BroadbandModel
	// - ja-JP_NarrowbandModel
	// - pt-BR_BroadbandModel
	// - pt-BR_NarrowbandModel
	// - zh-CN_BroadbandModel
	// - zh-CN_NarrowbandModel
	Model string

	// Indicates whether multiple final results that represent consecutive phrases separated by long pauses are returned.
	// If true, such phrases are returned; if false (the default), recognition ends after the first "end of speech" incident is detected.
	Continuous bool

	// The time in seconds after which, if only silence (no speech) is detected in submitted audio,
	// the connection is closed with a 400 response code and with session_closed set to true.
	// The default is 30 seconds. Useful for stopping audio submission from a live microphone when a user simply walks away. Use -1 for infinity.
	Inactivity_timeout int

	// A list of keywords to spot in the audio. Each keyword string can include one or more tokens.
	// Omit the parameter or specify an empty array if you do not need to spot keywords.
	Keywords []string

	// A confidence value that is the lower bound for spotting a keyword. A word is considered to match a keyword if its confidence is greater than or equal to the threshold.
	// Specify a probability between 0 and 1 inclusive. No keyword spotting is performed if you omit the parameter or specify the default value (null).
	// If you specify a threshold, you must also specify one or more keywords.
	Keywords_threshold float64

	// The maximum number of alternative transcripts to be returned. By default, a single transcription is returned.
	Max_alternatives int

	// A confidence value that is the lower bound for identifying a hypothesis as a possible word alternative (also known as "Confusion Networks").
	// An alternative word is considered if its confidence is greater than or equal to the threshold. Specify a probability between 0 and 1 inclusive.
	// No alternative words are computed if you omit the parameter or specify the default value (null).
	Word_alternatives_threshold float64

	// Indicates whether a confidence measure in the range of 0 to 1 is to be returned for each word. The default is false.
	Word_confidence bool

	// Indicates whether time alignment is returned for each word. The default is false.
	Timestamps bool
}

type SpeechRecognitionResult struct{
	Results map[string] interface{}
	Result_index int
}

// A client for the IBM Watson Speech-To-Text API endpoint.
type SpeechToText struct {
	// The Watson client used to authenticate our request.
	Watson watson.Watson
	Options SpeechToTextOptions
}

// Create a new Speech-To-Text client instance.
func New(watson watson.Watson, options SpeechToTextOptions) SpeechToText {
	// Check if a custom URL is set and use default if not.
	if (len(options.Url) < 1) {
		options.Url = SpeechToTextAPI
	}

	return SpeechToText{watson, options}
}

// Perform a call to the Speech-To-Text API's recognize method.
func (stt *SpeechToText) Recognize(parameters RecognizeParameters) (*http.Response, error) {
	if (structs.IsZero(parameters)) {
		return nil, errors.New("Given parameter object was empty")
	}

	r, _ := http.NewRequest("POST", stt.getEndpointUrl("recognize", parameters.Session_id), parameters.Audio)

	values := r.URL.Query()

	values.Add("model", parameters.Model)
	values.Add("max_alternatives", strconv.Itoa(parameters.Max_alternatives))
	values.Add("word_confidence", strconv.FormatBool(parameters.Word_confidence))
	values.Add("keywords", `"` + strings.Join(parameters.Keywords, ",") + `"`)
	values.Add("keywords_threshold", strconv.FormatFloat(parameters.Keywords_threshold, 'f', 1, 64))

	r.URL.RawQuery = values.Encode()

	r.Header.Add("content-type", parameters.Content_type)

	return http.DefaultClient.Do(stt.Watson.SignRequest(r))
}

// Generate the endpoint URL we want to use.
func (stt *SpeechToText) getEndpointUrl(action string, session_id string) string {
	var params []string

	// Prepare the URL components with or without a session ID.
	if (!stt.Watson.Options.Use_unauthenticated && len(session_id) > 0) {
		params = []string{
			stt.Options.Url,
			stt.Watson.Options.Version,
			"sessions",
			session_id,
			action,
		}
	} else {
		params = []string{
			stt.Options.Url,
			stt.Watson.Options.Version,
			action,
		}
	}

	return strings.Join(params, "/")
}