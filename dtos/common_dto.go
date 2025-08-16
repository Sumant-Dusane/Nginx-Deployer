package dtos

type AddressType string

const (
	Domain  AddressType = "domain"
	IP      AddressType = "ip"
	Invalid AddressType = "invalid"
)
