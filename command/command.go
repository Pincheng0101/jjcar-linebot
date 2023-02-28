package command

type AskCommand string

const (
	RegistryAskCommand      AskCommand = "ASK_REGISTRY"
	InputNameAskCommand     AskCommand = "ASK_INPUT_NAME"
	InputPhoneAskCommand    AskCommand = "ASK_INPUT_PHONE"
	InputRegionAskCommand   AskCommand = "ASK_INPUT_REGION_COMMAND"
	InputBirthdayAskCommand AskCommand = "ASK_INPUT_BIRTHDAY_COMMAND"
	InputCarTypeAskCommand  AskCommand = "ASK_INPUT_CARTYPE"
)
