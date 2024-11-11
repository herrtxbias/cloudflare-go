// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workers

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"reflect"

	"github.com/cloudflare/cloudflare-go/v3/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v3/internal/param"
	"github.com/cloudflare/cloudflare-go/v3/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v3/option"
	"github.com/cloudflare/cloudflare-go/v3/shared"
	"github.com/tidwall/gjson"
)

// AIService contains methods and other services that help with interacting with
// the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAIService] method instead.
type AIService struct {
	Options []option.RequestOption
	Models  *AIModelService
}

// NewAIService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewAIService(opts ...option.RequestOption) (r *AIService) {
	r = &AIService{}
	r.Options = opts
	r.Models = NewAIModelService(opts...)
	return
}

// This endpoint provides users with the capability to run specific AI models
// on-demand.
//
// By submitting the required input data, users can receive real-time predictions
// or results generated by the chosen AI model. The endpoint supports various AI
// model types, ensuring flexibility and adaptability for diverse use cases.
//
// Model specific inputs available in
// [Cloudflare Docs](https://developers.cloudflare.com/workers-ai/models/).
func (r *AIService) Run(ctx context.Context, modelName string, params AIRunParams, opts ...option.RequestOption) (res *AIRunResponseUnion, err error) {
	var env AIRunResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if modelName == "" {
		err = errors.New("missing required model_name parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/ai/run/%s", params.AccountID, modelName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// An array of classification results for the input text
//
// Union satisfied by [workers.AIRunResponseTextClassification],
// [shared.UnionString], [workers.AIRunResponseAudio],
// [workers.AIRunResponseTextEmbeddings],
// [workers.AIRunResponseAutomaticSpeechRecognition],
// [workers.AIRunResponseImageClassification],
// [workers.AIRunResponseObjectDetection], [workers.AIRunResponseObject],
// [workers.AIRunResponseTranslation], [workers.AIRunResponseSummarization] or
// [workers.AIRunResponseImageToText].
type AIRunResponseUnion interface {
	ImplementsWorkersAIRunResponseUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AIRunResponseUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AIRunResponseTextClassification{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AIRunResponseAudio{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AIRunResponseTextEmbeddings{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AIRunResponseAutomaticSpeechRecognition{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AIRunResponseImageClassification{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AIRunResponseObjectDetection{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AIRunResponseObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AIRunResponseTranslation{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AIRunResponseSummarization{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AIRunResponseImageToText{}),
		},
	)
}

type AIRunResponseTextClassification []AIRunResponseTextClassificationItem

func (r AIRunResponseTextClassification) ImplementsWorkersAIRunResponseUnion() {}

type AIRunResponseTextClassificationItem struct {
	// The classification label assigned to the text (e.g., 'POSITIVE' or 'NEGATIVE')
	Label string `json:"label"`
	// Confidence score indicating the likelihood that the text belongs to the
	// specified label
	Score float64                                 `json:"score"`
	JSON  aiRunResponseTextClassificationItemJSON `json:"-"`
}

// aiRunResponseTextClassificationItemJSON contains the JSON metadata for the
// struct [AIRunResponseTextClassificationItem]
type aiRunResponseTextClassificationItemJSON struct {
	Label       apijson.Field
	Score       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AIRunResponseTextClassificationItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiRunResponseTextClassificationItemJSON) RawJSON() string {
	return r.raw
}

type AIRunResponseAudio struct {
	// The generated audio in MP3 format, base64-encoded
	Audio string                 `json:"audio"`
	JSON  aiRunResponseAudioJSON `json:"-"`
}

// aiRunResponseAudioJSON contains the JSON metadata for the struct
// [AIRunResponseAudio]
type aiRunResponseAudioJSON struct {
	Audio       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AIRunResponseAudio) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiRunResponseAudioJSON) RawJSON() string {
	return r.raw
}

func (r AIRunResponseAudio) ImplementsWorkersAIRunResponseUnion() {}

type AIRunResponseTextEmbeddings struct {
	// Embeddings of the requested text values
	Data  [][]float64                     `json:"data"`
	Shape []float64                       `json:"shape"`
	JSON  aiRunResponseTextEmbeddingsJSON `json:"-"`
}

// aiRunResponseTextEmbeddingsJSON contains the JSON metadata for the struct
// [AIRunResponseTextEmbeddings]
type aiRunResponseTextEmbeddingsJSON struct {
	Data        apijson.Field
	Shape       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AIRunResponseTextEmbeddings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiRunResponseTextEmbeddingsJSON) RawJSON() string {
	return r.raw
}

func (r AIRunResponseTextEmbeddings) ImplementsWorkersAIRunResponseUnion() {}

type AIRunResponseAutomaticSpeechRecognition struct {
	// The transcription
	Text      string                                        `json:"text,required"`
	Vtt       string                                        `json:"vtt"`
	WordCount float64                                       `json:"word_count"`
	Words     []AIRunResponseAutomaticSpeechRecognitionWord `json:"words"`
	JSON      aiRunResponseAutomaticSpeechRecognitionJSON   `json:"-"`
}

// aiRunResponseAutomaticSpeechRecognitionJSON contains the JSON metadata for the
// struct [AIRunResponseAutomaticSpeechRecognition]
type aiRunResponseAutomaticSpeechRecognitionJSON struct {
	Text        apijson.Field
	Vtt         apijson.Field
	WordCount   apijson.Field
	Words       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AIRunResponseAutomaticSpeechRecognition) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiRunResponseAutomaticSpeechRecognitionJSON) RawJSON() string {
	return r.raw
}

func (r AIRunResponseAutomaticSpeechRecognition) ImplementsWorkersAIRunResponseUnion() {}

type AIRunResponseAutomaticSpeechRecognitionWord struct {
	// The ending second when the word completes
	End float64 `json:"end"`
	// The second this word begins in the recording
	Start float64                                         `json:"start"`
	Word  string                                          `json:"word"`
	JSON  aiRunResponseAutomaticSpeechRecognitionWordJSON `json:"-"`
}

// aiRunResponseAutomaticSpeechRecognitionWordJSON contains the JSON metadata for
// the struct [AIRunResponseAutomaticSpeechRecognitionWord]
type aiRunResponseAutomaticSpeechRecognitionWordJSON struct {
	End         apijson.Field
	Start       apijson.Field
	Word        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AIRunResponseAutomaticSpeechRecognitionWord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiRunResponseAutomaticSpeechRecognitionWordJSON) RawJSON() string {
	return r.raw
}

type AIRunResponseImageClassification []AIRunResponseImageClassificationItem

func (r AIRunResponseImageClassification) ImplementsWorkersAIRunResponseUnion() {}

type AIRunResponseImageClassificationItem struct {
	// The predicted category or class for the input image based on analysis
	Label string `json:"label"`
	// A confidence value, between 0 and 1, indicating how certain the model is about
	// the predicted label
	Score float64                                  `json:"score"`
	JSON  aiRunResponseImageClassificationItemJSON `json:"-"`
}

// aiRunResponseImageClassificationItemJSON contains the JSON metadata for the
// struct [AIRunResponseImageClassificationItem]
type aiRunResponseImageClassificationItemJSON struct {
	Label       apijson.Field
	Score       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AIRunResponseImageClassificationItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiRunResponseImageClassificationItemJSON) RawJSON() string {
	return r.raw
}

type AIRunResponseObjectDetection []AIRunResponseObjectDetectionItem

func (r AIRunResponseObjectDetection) ImplementsWorkersAIRunResponseUnion() {}

type AIRunResponseObjectDetectionItem struct {
	// Coordinates defining the bounding box around the detected object
	Box AIRunResponseObjectDetectionBox `json:"box"`
	// The class label or name of the detected object
	Label string `json:"label"`
	// Confidence score indicating the likelihood that the detection is correct
	Score float64                              `json:"score"`
	JSON  aiRunResponseObjectDetectionItemJSON `json:"-"`
}

// aiRunResponseObjectDetectionItemJSON contains the JSON metadata for the struct
// [AIRunResponseObjectDetectionItem]
type aiRunResponseObjectDetectionItemJSON struct {
	Box         apijson.Field
	Label       apijson.Field
	Score       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AIRunResponseObjectDetectionItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiRunResponseObjectDetectionItemJSON) RawJSON() string {
	return r.raw
}

// Coordinates defining the bounding box around the detected object
type AIRunResponseObjectDetectionBox struct {
	// The x-coordinate of the bottom-right corner of the bounding box
	Xmax float64 `json:"xmax"`
	// The x-coordinate of the top-left corner of the bounding box
	Xmin float64 `json:"xmin"`
	// The y-coordinate of the bottom-right corner of the bounding box
	Ymax float64 `json:"ymax"`
	// The y-coordinate of the top-left corner of the bounding box
	Ymin float64                             `json:"ymin"`
	JSON aiRunResponseObjectDetectionBoxJSON `json:"-"`
}

// aiRunResponseObjectDetectionBoxJSON contains the JSON metadata for the struct
// [AIRunResponseObjectDetectionBox]
type aiRunResponseObjectDetectionBoxJSON struct {
	Xmax        apijson.Field
	Xmin        apijson.Field
	Ymax        apijson.Field
	Ymin        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AIRunResponseObjectDetectionBox) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiRunResponseObjectDetectionBoxJSON) RawJSON() string {
	return r.raw
}

type AIRunResponseObject struct {
	// The generated text response from the model
	Response string `json:"response"`
	// An array of tool calls requests made during the response generation
	ToolCalls []AIRunResponseObjectToolCall `json:"tool_calls"`
	JSON      aiRunResponseObjectJSON       `json:"-"`
}

// aiRunResponseObjectJSON contains the JSON metadata for the struct
// [AIRunResponseObject]
type aiRunResponseObjectJSON struct {
	Response    apijson.Field
	ToolCalls   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AIRunResponseObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiRunResponseObjectJSON) RawJSON() string {
	return r.raw
}

func (r AIRunResponseObject) ImplementsWorkersAIRunResponseUnion() {}

type AIRunResponseObjectToolCall struct {
	// The arguments passed to be passed to the tool call request
	Arguments interface{} `json:"arguments"`
	// The name of the tool to be called
	Name string                          `json:"name"`
	JSON aiRunResponseObjectToolCallJSON `json:"-"`
}

// aiRunResponseObjectToolCallJSON contains the JSON metadata for the struct
// [AIRunResponseObjectToolCall]
type aiRunResponseObjectToolCallJSON struct {
	Arguments   apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AIRunResponseObjectToolCall) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiRunResponseObjectToolCallJSON) RawJSON() string {
	return r.raw
}

type AIRunResponseTranslation struct {
	// The translated text in the target language
	TranslatedText string                       `json:"translated_text"`
	JSON           aiRunResponseTranslationJSON `json:"-"`
}

// aiRunResponseTranslationJSON contains the JSON metadata for the struct
// [AIRunResponseTranslation]
type aiRunResponseTranslationJSON struct {
	TranslatedText apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AIRunResponseTranslation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiRunResponseTranslationJSON) RawJSON() string {
	return r.raw
}

func (r AIRunResponseTranslation) ImplementsWorkersAIRunResponseUnion() {}

type AIRunResponseSummarization struct {
	// The summarized version of the input text
	Summary string                         `json:"summary"`
	JSON    aiRunResponseSummarizationJSON `json:"-"`
}

// aiRunResponseSummarizationJSON contains the JSON metadata for the struct
// [AIRunResponseSummarization]
type aiRunResponseSummarizationJSON struct {
	Summary     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AIRunResponseSummarization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiRunResponseSummarizationJSON) RawJSON() string {
	return r.raw
}

func (r AIRunResponseSummarization) ImplementsWorkersAIRunResponseUnion() {}

type AIRunResponseImageToText struct {
	Description string                       `json:"description"`
	JSON        aiRunResponseImageToTextJSON `json:"-"`
}

// aiRunResponseImageToTextJSON contains the JSON metadata for the struct
// [AIRunResponseImageToText]
type aiRunResponseImageToTextJSON struct {
	Description apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AIRunResponseImageToText) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiRunResponseImageToTextJSON) RawJSON() string {
	return r.raw
}

func (r AIRunResponseImageToText) ImplementsWorkersAIRunResponseUnion() {}

type AIRunParams struct {
	AccountID param.Field[string]  `path:"account_id,required"`
	Body      AIRunParamsBodyUnion `json:"body,required"`
}

func (r AIRunParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type AIRunParamsBody struct {
	Audio param.Field[interface{}] `json:"audio"`
	// Decreases the likelihood of the model repeating the same lines verbatim.
	FrequencyPenalty param.Field[float64]     `json:"frequency_penalty"`
	Functions        param.Field[interface{}] `json:"functions"`
	// Controls how closely the generated image should adhere to the prompt; higher
	// values make the image more aligned with the prompt
	Guidance param.Field[float64] `json:"guidance"`
	// The height of the generated image in pixels
	Height param.Field[int64]       `json:"height"`
	Image  param.Field[interface{}] `json:"image"`
	// For use with img2img tasks. A base64-encoded string of the input image
	ImageB64 param.Field[string] `json:"image_b64"`
	// The text that you want the model to summarize
	InputText param.Field[string] `json:"input_text"`
	// The speech language (e.g., 'en' for English, 'fr' for French). Defaults to 'en'
	// if not specified
	Lang param.Field[string] `json:"lang"`
	// Name of the LoRA (Low-Rank Adaptation) model to fine-tune the base model.
	Lora param.Field[string]      `json:"lora"`
	Mask param.Field[interface{}] `json:"mask"`
	// The maximum length of the generated summary in tokens
	MaxLength param.Field[int64] `json:"max_length"`
	// The maximum number of tokens to generate in the response.
	MaxTokens param.Field[int64]       `json:"max_tokens"`
	Messages  param.Field[interface{}] `json:"messages"`
	// Text describing elements to avoid in the generated image
	NegativePrompt param.Field[string] `json:"negative_prompt"`
	// The number of diffusion steps; higher values can improve quality but take longer
	NumSteps param.Field[int64] `json:"num_steps"`
	// Increases the likelihood of the model introducing new topics.
	PresencePenalty param.Field[float64] `json:"presence_penalty"`
	// A text description of the image you want to generate
	Prompt param.Field[string] `json:"prompt"`
	// If true, a chat template is not applied and you must adhere to the specific
	// model's expected formatting.
	Raw param.Field[bool] `json:"raw"`
	// Penalty for repeated tokens; higher values discourage repetition.
	RepetitionPenalty param.Field[float64] `json:"repetition_penalty"`
	// Random seed for reproducibility of the image generation
	Seed param.Field[int64] `json:"seed"`
	// The language of the recorded audio
	SourceLang param.Field[string] `json:"source_lang"`
	// If true, the response will be streamed back incrementally using SSE, Server Sent
	// Events.
	Stream param.Field[bool] `json:"stream"`
	// A value between 0 and 1 indicating how strongly to apply the transformation
	// during img2img tasks; lower values make the output closer to the input image
	Strength param.Field[float64] `json:"strength"`
	// The language to translate the transcription into. Currently only English is
	// supported.
	TargetLang param.Field[string] `json:"target_lang"`
	// Controls the randomness of the output; higher values produce more random
	// results.
	Temperature param.Field[float64]     `json:"temperature"`
	Text        param.Field[interface{}] `json:"text"`
	Tools       param.Field[interface{}] `json:"tools"`
	// Limits the AI to choose from the top 'k' most probable words. Lower values make
	// responses more focused; higher values introduce more variety and potential
	// surprises.
	TopK param.Field[int64] `json:"top_k"`
	// Adjusts the creativity of the AI's responses by controlling how many possible
	// words it considers. Lower values make outputs more predictable; higher values
	// allow for more varied and creative responses.
	TopP param.Field[float64] `json:"top_p"`
	// The width of the generated image in pixels
	Width param.Field[int64] `json:"width"`
}

func (r AIRunParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AIRunParamsBody) implementsWorkersAIRunParamsBodyUnion() {}

// Satisfied by [workers.AIRunParamsBodyTextClassification],
// [workers.AIRunParamsBodyTextToImage], [workers.AIRunParamsBodyTextToSpeech],
// [workers.AIRunParamsBodyTextEmbeddings],
// [workers.AIRunParamsBodyAutomaticSpeechRecognition],
// [workers.AIRunParamsBodyImageClassification],
// [workers.AIRunParamsBodyObjectDetection], [workers.AIRunParamsBodyPrompt],
// [workers.AIRunParamsBodyMessages], [workers.AIRunParamsBodyTranslation],
// [workers.AIRunParamsBodySummarization], [workers.AIRunParamsBodyImageToText],
// [AIRunParamsBody].
type AIRunParamsBodyUnion interface {
	implementsWorkersAIRunParamsBodyUnion()
}

type AIRunParamsBodyTextClassification struct {
	// The text that you want to classify
	Text param.Field[string] `json:"text,required"`
}

func (r AIRunParamsBodyTextClassification) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AIRunParamsBodyTextClassification) implementsWorkersAIRunParamsBodyUnion() {}

type AIRunParamsBodyTextToImage struct {
	// A text description of the image you want to generate
	Prompt param.Field[string] `json:"prompt,required"`
	// Controls how closely the generated image should adhere to the prompt; higher
	// values make the image more aligned with the prompt
	Guidance param.Field[float64] `json:"guidance"`
	// The height of the generated image in pixels
	Height param.Field[int64] `json:"height"`
	// For use with img2img tasks. An array of integers that represent the image data
	// constrained to 8-bit unsigned integer values
	Image param.Field[[]float64] `json:"image"`
	// For use with img2img tasks. A base64-encoded string of the input image
	ImageB64 param.Field[string] `json:"image_b64"`
	// An array representing An array of integers that represent mask image data for
	// inpainting constrained to 8-bit unsigned integer values
	Mask param.Field[[]float64] `json:"mask"`
	// Text describing elements to avoid in the generated image
	NegativePrompt param.Field[string] `json:"negative_prompt"`
	// The number of diffusion steps; higher values can improve quality but take longer
	NumSteps param.Field[int64] `json:"num_steps"`
	// Random seed for reproducibility of the image generation
	Seed param.Field[int64] `json:"seed"`
	// A value between 0 and 1 indicating how strongly to apply the transformation
	// during img2img tasks; lower values make the output closer to the input image
	Strength param.Field[float64] `json:"strength"`
	// The width of the generated image in pixels
	Width param.Field[int64] `json:"width"`
}

func (r AIRunParamsBodyTextToImage) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AIRunParamsBodyTextToImage) implementsWorkersAIRunParamsBodyUnion() {}

type AIRunParamsBodyTextToSpeech struct {
	// A text description of the image you want to generate
	Prompt param.Field[string] `json:"prompt,required"`
	// The speech language (e.g., 'en' for English, 'fr' for French). Defaults to 'en'
	// if not specified
	Lang param.Field[string] `json:"lang"`
}

func (r AIRunParamsBodyTextToSpeech) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AIRunParamsBodyTextToSpeech) implementsWorkersAIRunParamsBodyUnion() {}

type AIRunParamsBodyTextEmbeddings struct {
	// The text to embed
	Text param.Field[AIRunParamsBodyTextEmbeddingsTextUnion] `json:"text,required"`
}

func (r AIRunParamsBodyTextEmbeddings) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AIRunParamsBodyTextEmbeddings) implementsWorkersAIRunParamsBodyUnion() {}

// The text to embed
//
// Satisfied by [shared.UnionString],
// [workers.AIRunParamsBodyTextEmbeddingsTextArray].
type AIRunParamsBodyTextEmbeddingsTextUnion interface {
	ImplementsWorkersAIRunParamsBodyTextEmbeddingsTextUnion()
}

type AIRunParamsBodyTextEmbeddingsTextArray []string

func (r AIRunParamsBodyTextEmbeddingsTextArray) ImplementsWorkersAIRunParamsBodyTextEmbeddingsTextUnion() {
}

type AIRunParamsBodyAutomaticSpeechRecognition struct {
	// An array of integers that represent the audio data constrained to 8-bit unsigned
	// integer values
	Audio param.Field[[]float64] `json:"audio,required"`
	// The language of the recorded audio
	SourceLang param.Field[string] `json:"source_lang"`
	// The language to translate the transcription into. Currently only English is
	// supported.
	TargetLang param.Field[string] `json:"target_lang"`
}

func (r AIRunParamsBodyAutomaticSpeechRecognition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AIRunParamsBodyAutomaticSpeechRecognition) implementsWorkersAIRunParamsBodyUnion() {}

type AIRunParamsBodyImageClassification struct {
	// An array of integers that represent the image data constrained to 8-bit unsigned
	// integer values
	Image param.Field[[]float64] `json:"image,required"`
}

func (r AIRunParamsBodyImageClassification) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AIRunParamsBodyImageClassification) implementsWorkersAIRunParamsBodyUnion() {}

type AIRunParamsBodyObjectDetection struct {
	// An array of integers that represent the image data constrained to 8-bit unsigned
	// integer values
	Image param.Field[[]float64] `json:"image"`
}

func (r AIRunParamsBodyObjectDetection) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AIRunParamsBodyObjectDetection) implementsWorkersAIRunParamsBodyUnion() {}

type AIRunParamsBodyPrompt struct {
	// The input text prompt for the model to generate a response.
	Prompt param.Field[string] `json:"prompt,required"`
	// Decreases the likelihood of the model repeating the same lines verbatim.
	FrequencyPenalty param.Field[float64] `json:"frequency_penalty"`
	// Name of the LoRA (Low-Rank Adaptation) model to fine-tune the base model.
	Lora param.Field[string] `json:"lora"`
	// The maximum number of tokens to generate in the response.
	MaxTokens param.Field[int64] `json:"max_tokens"`
	// Increases the likelihood of the model introducing new topics.
	PresencePenalty param.Field[float64] `json:"presence_penalty"`
	// If true, a chat template is not applied and you must adhere to the specific
	// model's expected formatting.
	Raw param.Field[bool] `json:"raw"`
	// Penalty for repeated tokens; higher values discourage repetition.
	RepetitionPenalty param.Field[float64] `json:"repetition_penalty"`
	// Random seed for reproducibility of the generation.
	Seed param.Field[int64] `json:"seed"`
	// If true, the response will be streamed back incrementally using SSE, Server Sent
	// Events.
	Stream param.Field[bool] `json:"stream"`
	// Controls the randomness of the output; higher values produce more random
	// results.
	Temperature param.Field[float64] `json:"temperature"`
	// Limits the AI to choose from the top 'k' most probable words. Lower values make
	// responses more focused; higher values introduce more variety and potential
	// surprises.
	TopK param.Field[int64] `json:"top_k"`
	// Adjusts the creativity of the AI's responses by controlling how many possible
	// words it considers. Lower values make outputs more predictable; higher values
	// allow for more varied and creative responses.
	TopP param.Field[float64] `json:"top_p"`
}

func (r AIRunParamsBodyPrompt) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AIRunParamsBodyPrompt) implementsWorkersAIRunParamsBodyUnion() {}

type AIRunParamsBodyMessages struct {
	// An array of message objects representing the conversation history.
	Messages param.Field[[]AIRunParamsBodyMessagesMessage] `json:"messages,required"`
	// Decreases the likelihood of the model repeating the same lines verbatim.
	FrequencyPenalty param.Field[float64]                           `json:"frequency_penalty"`
	Functions        param.Field[[]AIRunParamsBodyMessagesFunction] `json:"functions"`
	// The maximum number of tokens to generate in the response.
	MaxTokens param.Field[int64] `json:"max_tokens"`
	// Increases the likelihood of the model introducing new topics.
	PresencePenalty param.Field[float64] `json:"presence_penalty"`
	// Penalty for repeated tokens; higher values discourage repetition.
	RepetitionPenalty param.Field[float64] `json:"repetition_penalty"`
	// Random seed for reproducibility of the generation.
	Seed param.Field[int64] `json:"seed"`
	// If true, the response will be streamed back incrementally.
	Stream param.Field[bool] `json:"stream"`
	// Controls the randomness of the output; higher values produce more random
	// results.
	Temperature param.Field[float64] `json:"temperature"`
	// A list of tools available for the assistant to use.
	Tools param.Field[[]AIRunParamsBodyMessagesToolUnion] `json:"tools"`
	// Limits the AI to choose from the top 'k' most probable words. Lower values make
	// responses more focused; higher values introduce more variety and potential
	// surprises.
	TopK param.Field[int64] `json:"top_k"`
	// Controls the creativity of the AI's responses by adjusting how many possible
	// words it considers. Lower values make outputs more predictable; higher values
	// allow for more varied and creative responses.
	TopP param.Field[float64] `json:"top_p"`
}

func (r AIRunParamsBodyMessages) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AIRunParamsBodyMessages) implementsWorkersAIRunParamsBodyUnion() {}

type AIRunParamsBodyMessagesMessage struct {
	// The content of the message as a string.
	Content param.Field[string] `json:"content,required"`
	// The role of the message sender (e.g., 'user', 'assistant', 'system', 'tool').
	Role param.Field[string] `json:"role,required"`
}

func (r AIRunParamsBodyMessagesMessage) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AIRunParamsBodyMessagesFunction struct {
	Code param.Field[string] `json:"code,required"`
	Name param.Field[string] `json:"name,required"`
}

func (r AIRunParamsBodyMessagesFunction) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AIRunParamsBodyMessagesTool struct {
	// A brief description of what the tool does.
	Description param.Field[string]      `json:"description"`
	Function    param.Field[interface{}] `json:"function"`
	// The name of the tool. More descriptive the better.
	Name       param.Field[string]      `json:"name"`
	Parameters param.Field[interface{}] `json:"parameters"`
	// Specifies the type of tool (e.g., 'function').
	Type param.Field[string] `json:"type"`
}

func (r AIRunParamsBodyMessagesTool) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AIRunParamsBodyMessagesTool) implementsWorkersAIRunParamsBodyMessagesToolUnion() {}

// Satisfied by [workers.AIRunParamsBodyMessagesToolsObject],
// [workers.AIRunParamsBodyMessagesToolsObject], [AIRunParamsBodyMessagesTool].
type AIRunParamsBodyMessagesToolUnion interface {
	implementsWorkersAIRunParamsBodyMessagesToolUnion()
}

type AIRunParamsBodyMessagesToolsObject struct {
	// A brief description of what the tool does.
	Description param.Field[string] `json:"description,required"`
	// The name of the tool. More descriptive the better.
	Name param.Field[string] `json:"name,required"`
	// Schema defining the parameters accepted by the tool.
	Parameters param.Field[AIRunParamsBodyMessagesToolsObjectParameters] `json:"parameters,required"`
}

func (r AIRunParamsBodyMessagesToolsObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AIRunParamsBodyMessagesToolsObject) implementsWorkersAIRunParamsBodyMessagesToolUnion() {}

// Schema defining the parameters accepted by the tool.
type AIRunParamsBodyMessagesToolsObjectParameters struct {
	// Definitions of each parameter.
	Properties param.Field[map[string]AIRunParamsBodyMessagesToolsObjectParametersProperties] `json:"properties,required"`
	// The type of the parameters object (usually 'object').
	Type param.Field[string] `json:"type,required"`
	// List of required parameter names.
	Required param.Field[[]string] `json:"required"`
}

func (r AIRunParamsBodyMessagesToolsObjectParameters) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AIRunParamsBodyMessagesToolsObjectParametersProperties struct {
	// A description of the expected parameter.
	Description param.Field[string] `json:"description,required"`
	// The data type of the parameter.
	Type param.Field[string] `json:"type,required"`
}

func (r AIRunParamsBodyMessagesToolsObjectParametersProperties) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AIRunParamsBodyTranslation struct {
	// The language code to translate the text into (e.g., 'es' for Spanish)
	TargetLang param.Field[string] `json:"target_lang,required"`
	// The text to be translated
	Text param.Field[string] `json:"text,required"`
	// The language code of the source text (e.g., 'en' for English). Defaults to 'en'
	// if not specified
	SourceLang param.Field[string] `json:"source_lang"`
}

func (r AIRunParamsBodyTranslation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AIRunParamsBodyTranslation) implementsWorkersAIRunParamsBodyUnion() {}

type AIRunParamsBodySummarization struct {
	// The text that you want the model to summarize
	InputText param.Field[string] `json:"input_text,required"`
	// The maximum length of the generated summary in tokens
	MaxLength param.Field[int64] `json:"max_length"`
}

func (r AIRunParamsBodySummarization) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AIRunParamsBodySummarization) implementsWorkersAIRunParamsBodyUnion() {}

type AIRunParamsBodyImageToText struct {
	// An array of integers that represent the image data constrained to 8-bit unsigned
	// integer values
	Image param.Field[[]float64] `json:"image,required"`
	// The maximum number of tokens to generate in the response.
	MaxTokens param.Field[int64] `json:"max_tokens"`
	// The input text prompt for the model to generate a response.
	Prompt param.Field[string] `json:"prompt"`
	// If true, a chat template is not applied and you must adhere to the specific
	// model's expected formatting.
	Raw param.Field[bool] `json:"raw"`
	// Controls the randomness of the output; higher values produce more random
	// results.
	Temperature param.Field[float64] `json:"temperature"`
}

func (r AIRunParamsBodyImageToText) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AIRunParamsBodyImageToText) implementsWorkersAIRunParamsBodyUnion() {}

type AIRunResponseEnvelope struct {
	// An array of classification results for the input text
	Result AIRunResponseUnion        `json:"result" format:"binary"`
	JSON   aiRunResponseEnvelopeJSON `json:"-"`
}

// aiRunResponseEnvelopeJSON contains the JSON metadata for the struct
// [AIRunResponseEnvelope]
type aiRunResponseEnvelopeJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AIRunResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiRunResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
