package response_define

type ResponseCreateWallet struct {
	Sign      string                   `json:"sign"`
	Timestamp string                   `json:"timestamp"`
	Data      ResponseCreateWalletData `json:"data"`
	Msg       string                   `json:"msg"`
	Code      int                      `json:"code"`
}

type ResponseCreateWalletData struct {
	Address   string `json:"address"`
	UserID    string `json:"UserId"`
	PartnerID string `json:"PartnerId"`
	ChainID   int    `json:"ChainID"`
	OpenID    string `json:"OpenId"`
}
