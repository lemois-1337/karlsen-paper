package model

type PaperAPI interface {
	GenerateWallet() (PaperWallet, error)
}

type PaperWallet interface {
	Mnemonic() *MnemonicString
	Address(index int) (string, error)
	AddressQR(index int) ([]byte, error)
	KPubKey() (string)
	KPubKeyQR() ([]byte, error)
}
