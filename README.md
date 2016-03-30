Golang library to use the [Watson Developer Cloud][wdc] services, a collection of REST
APIs and SDKs that use cognitive computing to solve complex problems.

## Table of Contents
    * Preamble (#preamble)
    * Installation (#installation)
    * Usage (#usage)
    * [IBM Watson Services](#ibm-watson-services)
        * [Speech to Text](#speech-to-text)
    * Contributing (#contributing)
    * License (#license)

## Preamble
The ultimate target of this library is to support all of IBM's Watson APIs. However at the current time we only support access to the following endpoints:
    * Speech-To-Text (no streaming)

I'll try to add more functions in the future. Your patience is highly appreciated. 
If you can't wait for a specific function, feel free to help me out and contribute!

Also, most of the parameter documentation for the currently supported APIs was not written by me, but by the brave men and woman of the IBM developer team. 
All credit for excellent documentation goes to them and I'm not going to take that away from them.

## Installation
```sh
$ go get github.com/maxl28/watson-developer-cloud
```

## Usage
Acquire your credentials via IBM's Bluemix service and put them in a config file of your choice. You'll need them to initialize the Watson client like this:
```go
import "github.com/maxl28/watson-developer-cloud/watson"

w := watson.New(watson.WatsonOptions{
		Username: "YOUR_USERNAME",
		Password: "YOUR_PASSWORD",
	})
```

## IBM Watson Services

### Speech to Text
Assuming that you've already created your Watson client, you may now initialize the Speech-To-Text client like this:
```go
import "github.com/maxl28/watson-developer-cloud/speech-to-text"

stt := speech_to_text.New(w, speech_to_text.SpeechToTextOptions{})
```

Now, we'll make a request to the recognize endpoint with minimal parameters:
```go
file, _ := os.Open("YOUR_TEST_FILE.ogg")

defer file.Close()

data, _ := stt.Recognize(speech_to_text.RecognizeParameters{
	Audio: file,
	Content_type: "audio/ogg;codecs=opus",
	Model: "en-US_BroadbandModel",
	Max_alternatives: 3,
})

result, _ := ioutil.ReadAll(data.Body)
```

## Contributing
Feel free to help me expand this little library and help the Go community use IBM's awesome API services.

## License
This library is licensed under the MIT license. Full license text is available in
[COPYING][license].