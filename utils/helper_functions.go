package utils

import (
	"net"
	"net/url"

	"github.com/Sumant-Dusane/nginx-deployer/dtos"
)

func IsDomainOrIP(address string) dtos.AddressType {
	ip := net.ParseIP(address)

	if ip != nil {
		return dtos.IP
	}

	domainUrl, domainErr := url.ParseRequestURI(address)

	if domainUrl != nil && domainErr == nil {
		return dtos.Domain
	}

	return dtos.Invalid
}
