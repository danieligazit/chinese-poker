

type MessageHandler interface{
	UnmarshalMessage(message []byte) (clientMessage Message, err error)
	HandleMessage(clientId uint64, message []byte)
	SendMessage(clientId uint64, clientMessage Message)
	SendConnectionStatus()
}