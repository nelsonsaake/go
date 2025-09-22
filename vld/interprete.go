package vld

import (
	"fmt"
	"reflect"

	"github.com/go-playground/validator/v10"
)

// interprete: organize error in a user-friendly way like in laravel
func interprete(err validator.FieldError) string {
	field := toWords(err.Field())

	tr := map[string]string{
		"email": fmt.Sprintf("%s provided is not a valid email.", field),

		// field comparison tags
		"eqfield":       fmt.Sprintf("%s must be equal to %s.", field, err.Param()),
		"eqcsfield":     fmt.Sprintf("%s must be equal to the relative field %s.", field, err.Param()),
		"fieldcontains": fmt.Sprintf("%s must contain the characters: %s.", field, err.Param()),
		"fieldexcludes": fmt.Sprintf("%s must not contain the characters: %s.", field, err.Param()),
		"gtcsfield":     fmt.Sprintf("%s must be greater than the relative field %s.", field, err.Param()),
		"gtecsfield":    fmt.Sprintf("%s must be greater than or equal to the relative field %s.", field, err.Param()),
		"gtefield":      fmt.Sprintf("%s must be greater than or equal to the field %s.", field, err.Param()),
		"gtfield":       fmt.Sprintf("%s must be greater than the field %s.", field, err.Param()),
		"ltcsfield":     fmt.Sprintf("%s must be less than the relative field %s.", field, err.Param()),
		"ltecsfield":    fmt.Sprintf("%s must be less than or equal to the relative field %s.", field, err.Param()),
		"ltefield":      fmt.Sprintf("%s must be less than or equal to the field %s.", field, err.Param()),
		"ltfield":       fmt.Sprintf("%s must be less than the field %s.", field, err.Param()),
		"necsfield":     fmt.Sprintf("%s must not be equal to the relative field %s.", field, err.Param()),
		"nefield":       fmt.Sprintf("%s must not be equal to the field %s.", field, err.Param()),

		// Comparison tags
		"eq": func() string {
			switch err.Kind() {
			case reflect.String:
				return fmt.Sprintf("%s must be the same as %s.", field, err.Param())
			default:
				return fmt.Sprintf("%s must be equal to %s.", field, err.Param())
			}
		}(),
		"eq_ignore_case": fmt.Sprintf("%s must be equal to %s (ignoring case).", field, err.Param()),
		"gt":             fmt.Sprintf("%s must be greater than %s.", field, err.Param()),
		"gte":            fmt.Sprintf("%s must be greater than or equal to %s.", field, err.Param()),
		"lt":             fmt.Sprintf("%s must be less than %s.", field, err.Param()),
		"lte":            fmt.Sprintf("%s must be less than or equal to %s.", field, err.Param()),
		"ne":             fmt.Sprintf("%s must not be equal to %s.", field, err.Param()),
		"ne_ignore_case": fmt.Sprintf("%s must not be equal to %s (ignoring case).", field, err.Param()),

		// Network and URI tags
		"cidr":             fmt.Sprintf("%s must be a valid CIDR notation.", field),
		"cidrv4":           fmt.Sprintf("%s must be a valid CIDR IPv4 notation.", field),
		"cidrv6":           fmt.Sprintf("%s must be a valid CIDR IPv6 notation.", field),
		"datauri":          fmt.Sprintf("%s must be a valid Data URL.", field),
		"fqdn":             fmt.Sprintf("%s must be a valid fully qualified domain name (FQDN).", field),
		"hostname":         fmt.Sprintf("%s must be a valid hostname (RFC 952).", field),
		"hostname_port":    fmt.Sprintf("%s must be a valid host:port combination.", field),
		"hostname_rfc1123": fmt.Sprintf("%s must be a valid hostname (RFC 1123).", field),
		"ip":               fmt.Sprintf("%s must be a valid IP address.", field),
		"ip4_addr":         fmt.Sprintf("%s must be a valid IPv4 address.", field),
		"ip6_addr":         fmt.Sprintf("%s must be a valid IPv6 address.", field),
		"ip_addr":          fmt.Sprintf("%s must be a valid IP address.", field),
		"ipv4":             fmt.Sprintf("%s must be a valid IPv4 address.", field),
		"ipv6":             fmt.Sprintf("%s must be a valid IPv6 address.", field),
		"mac":              fmt.Sprintf("%s must be a valid MAC address.", field),
		"tcp4_addr":        fmt.Sprintf("%s must be a valid TCPv4 address.", field),
		"tcp6_addr":        fmt.Sprintf("%s must be a valid TCPv6 address.", field),
		"tcp_addr":         fmt.Sprintf("%s must be a valid TCP address.", field),
		"udp4_addr":        fmt.Sprintf("%s must be a valid UDPv4 address.", field),
		"udp6_addr":        fmt.Sprintf("%s must be a valid UDPv6 address.", field),
		"udp_addr":         fmt.Sprintf("%s must be a valid UDP address.", field),
		"unix_addr":        fmt.Sprintf("%s must be a valid Unix domain socket address.", field),
		"uri":              fmt.Sprintf("%s must be a valid URI.", field),
		"url":              fmt.Sprintf("%s must be a valid URL.", field),
		"http_url":         fmt.Sprintf("%s must be a valid HTTP URL.", field),
		"url_encoded":      fmt.Sprintf("%s must be a valid URL encoded string.", field),
		"urn_rfc2141":      fmt.Sprintf("%s must be a valid URN (RFC 2141).", field),

		// String and character tags
		"alpha":           fmt.Sprintf("%s must contain only alphabetic characters.", field),
		"alphanum":        fmt.Sprintf("%s must contain only alphanumeric characters.", field),
		"alphanumunicode": fmt.Sprintf("%s must contain only alphanumeric Unicode characters.", field),
		"alphaunicode":    fmt.Sprintf("%s must contain only alphabetic Unicode characters.", field),
		"ascii":           fmt.Sprintf("%s must contain only ASCII characters.", field),
		"boolean":         fmt.Sprintf("%s must be a boolean value.", field),
		"contains":        fmt.Sprintf("%s must contain %s.", field, err.Param()),
		"containsany":     fmt.Sprintf("%s must contain at least one of the following: %s.", field, err.Param()),
		"containsrune":    fmt.Sprintf("%s must contain the rune: %s.", field, err.Param()),
		"endsnotwith":     fmt.Sprintf("%s must not end with %s.", field, err.Param()),
		"endswith":        fmt.Sprintf("%s must end with %s.", field, err.Param()),
		"excludes":        fmt.Sprintf("%s must not contain %s.", field, err.Param()),
		"excludesall":     fmt.Sprintf("%s must not contain any of the following: %s.", field, err.Param()),
		"excludesrune":    fmt.Sprintf("%s must not contain the rune: %s.", field, err.Param()),
		"lowercase":       fmt.Sprintf("%s must be lowercase.", field),
		"multibyte":       fmt.Sprintf("%s must contain multi-byte characters.", field),
		"number":          fmt.Sprintf("%s must be a number.", field),
		"numeric":         fmt.Sprintf("%s must be numeric.", field),
		"printascii":      fmt.Sprintf("%s must contain only printable ASCII characters.", field),
		"startsnotwith":   fmt.Sprintf("%s must not start with %s.", field, err.Param()),
		"startswith":      fmt.Sprintf("%s must start with %s.", field, err.Param()),
		"uppercase":       fmt.Sprintf("%s must be uppercase.", field),

		// formats
		"base64":                        fmt.Sprintf("%s must be a valid Base64 string.", field),
		"base64url":                     fmt.Sprintf("%s must be a valid Base64URL string.", field),
		"base64rawurl":                  fmt.Sprintf("%s must be a valid Base64RawURL string.", field),
		"bic":                           fmt.Sprintf("%s must be a valid Business Identifier Code (ISO 9362).", field),
		"bcp47_language_tag":            fmt.Sprintf("%s must be a valid BCP 47 language tag.", field),
		"btc_addr":                      fmt.Sprintf("%s must be a valid Bitcoin address.", field),
		"btc_addr_bech32":               fmt.Sprintf("%s must be a valid Bitcoin Bech32 address.", field),
		"credit_card":                   fmt.Sprintf("%s must be a valid credit card number.", field),
		"mongodb":                       fmt.Sprintf("%s must be a valid MongoDB ObjectID.", field),
		"cron":                          fmt.Sprintf("%s must be a valid cron expression.", field),
		"spicedb":                       fmt.Sprintf("%s must be a valid SpiceDb ObjectID, Permission, or Type.", field),
		"datetime":                      fmt.Sprintf("%s must be a valid datetime.", field),
		"e164":                          fmt.Sprintf("%s must be a valid E.164 formatted phone number.", field),
		"eth_addr":                      fmt.Sprintf("%s must be a valid Ethereum address.", field),
		"hexadecimal":                   fmt.Sprintf("%s must be a valid hexadecimal string.", field),
		"hexcolor":                      fmt.Sprintf("%s must be a valid hex color string.", field),
		"hsl":                           fmt.Sprintf("%s must be a valid HSL string.", field),
		"hsla":                          fmt.Sprintf("%s must be a valid HSLA string.", field),
		"html":                          fmt.Sprintf("%s must contain valid HTML tags.", field),
		"html_encoded":                  fmt.Sprintf("%s must be a valid HTML encoded string.", field),
		"isbn":                          fmt.Sprintf("%s must be a valid International Standard Book Number.", field),
		"isbn10":                        fmt.Sprintf("%s must be a valid ISBN-10.", field),
		"isbn13":                        fmt.Sprintf("%s must be a valid ISBN-13.", field),
		"issn":                          fmt.Sprintf("%s must be a valid International Standard Serial Number.", field),
		"iso3166_1_alpha2":              fmt.Sprintf("%s must be a valid two-letter country code (ISO 3166-1 alpha-2).", field),
		"iso3166_1_alpha3":              fmt.Sprintf("%s must be a valid three-letter country code (ISO 3166-1 alpha-3).", field),
		"iso3166_1_alpha_numeric":       fmt.Sprintf("%s must be a valid numeric country code (ISO 3166-1 numeric).", field),
		"iso3166_2":                     fmt.Sprintf("%s must be a valid country subdivision code (ISO 3166-2).", field),
		"iso4217":                       fmt.Sprintf("%s must be a valid currency code (ISO 4217).", field),
		"json":                          fmt.Sprintf("%s must be a valid JSON string.", field),
		"jwt":                           fmt.Sprintf("%s must be a valid JSON Web Token (JWT).", field),
		"latitude":                      fmt.Sprintf("%s must be a valid latitude.", field),
		"longitude":                     fmt.Sprintf("%s must be a valid longitude.", field),
		"luhn_checksum":                 fmt.Sprintf("%s must have a valid Luhn algorithm checksum.", field),
		"postcode_iso3166_alpha2":       fmt.Sprintf("%s must be a valid postcode.", field),
		"postcode_iso3166_alpha2_field": fmt.Sprintf("%s must be a valid postcode.", field),
		"rgb":                           fmt.Sprintf("%s must be a valid RGB string.", field),
		"rgba":                          fmt.Sprintf("%s must be a valid RGBA string.", field),
		"ssn":                           fmt.Sprintf("%s must be a valid Social Security Number (SSN).", field),
		"timezone":                      fmt.Sprintf("%s must be a valid timezone.", field),
		"uuid":                          fmt.Sprintf("%s must be a valid UUID.", field),
		"uuid3":                         fmt.Sprintf("%s must be a valid UUID v3.", field),
		"uuid3_rfc4122":                 fmt.Sprintf("%s must be a valid UUID v3 (RFC 4122).", field),
		"uuid4":                         fmt.Sprintf("%s must be a valid UUID v4.", field),
		"uuid4_rfc4122":                 fmt.Sprintf("%s must be a valid UUID v4 (RFC 4122).", field),
		"uuid5":                         fmt.Sprintf("%s must be a valid UUID v5.", field),
		"uuid5_rfc4122":                 fmt.Sprintf("%s must be a valid UUID v5 (RFC 4122).", field),
		"uuid_rfc4122":                  fmt.Sprintf("%s must be a valid UUID (RFC 4122).", field),
		"md4":                           fmt.Sprintf("%s must be a valid MD4 hash.", field),
		"md5":                           fmt.Sprintf("%s must be a valid MD5 hash.", field),
		"sha256":                        fmt.Sprintf("%s must be a valid SHA256 hash.", field),
		"sha384":                        fmt.Sprintf("%s must be a valid SHA384 hash.", field),
		"sha512":                        fmt.Sprintf("%s must be a valid SHA512 hash.", field),
		"ripemd128":                     fmt.Sprintf("%s must be a valid RIPEMD-128 hash.", field),
		"ripemd160":                     fmt.Sprintf("%s must be a valid RIPEMD-160 hash.", field),
		"tiger128":                      fmt.Sprintf("%s must be a valid TIGER128 hash.", field),
		"tiger160":                      fmt.Sprintf("%s must be a valid TIGER160 hash.", field),
		"tiger192":                      fmt.Sprintf("%s must be a valid TIGER192 hash.", field),
		"semver":                        fmt.Sprintf("%s must be a valid semantic version (SemVer 2.0.0).", field),
		"ulid":                          fmt.Sprintf("%s must be a valid ULID.", field),
		"cve":                           fmt.Sprintf("%s must be a valid CVE identifier.", field),

		// Other tags
		"dir":                  fmt.Sprintf("%s must be an existing directory.", field),
		"dirpath":              fmt.Sprintf("%s must be a valid directory path.", field),
		"file":                 fmt.Sprintf("%s must be an existing file.", field),
		"filepath":             fmt.Sprintf("%s must be a valid file path.", field),
		"image":                fmt.Sprintf("%s must be a valid image file.", field),
		"isdefault":            fmt.Sprintf("%s must be the default value.", field),
		"len":                  fmt.Sprintf("%s must be %s characters.", field, err.Param()),
		"max":                  fmt.Sprintf("%s must be at most %s characters.", field, err.Param()),
		"min":                  fmt.Sprintf("%s must be at least %s characters.", field, err.Param()),
		"oneof":                fmt.Sprintf("%s must be one of the following: %s.", field, err.Param()),
		"required":             fmt.Sprintf("%s is required.", field),
		"required_if":          fmt.Sprintf("%s is required when %s is present.", field, err.Param()),
		"required_unless":      fmt.Sprintf("%s is required unless %s is present.", field, err.Param()),
		"required_with":        fmt.Sprintf("%s is required when %s is present.", field, err.Param()),
		"required_with_all":    fmt.Sprintf("%s is required when all of %s are present.", field, err.Param()),
		"required_without":     fmt.Sprintf("%s is required when %s is not present.", field, err.Param()),
		"required_without_all": fmt.Sprintf("%s is required when none of %s are present.", field, err.Param()),
		"excluded_if":          fmt.Sprintf("%s must be excluded if %s is present.", field, err.Param()),
		"excluded_unless":      fmt.Sprintf("%s must be excluded unless %s is present.", field, err.Param()),
		"excluded_with":        fmt.Sprintf("%s must be excluded when %s is present.", field, err.Param()),
		"excluded_with_all":    fmt.Sprintf("%s must be excluded when all of %s are present.", field, err.Param()),
		"excluded_without":     fmt.Sprintf("%s must be excluded when %s is not present.", field, err.Param()),
		"excluded_without_all": fmt.Sprintf("%s must be excluded when none of %s are present.", field, err.Param()),
		"unique":               fmt.Sprintf("%s must be unique.", field),
		"iscolor":              fmt.Sprintf("%s must be a valid color (hex, rgb, rgba, hsl, or hsla).", field),
		"country_code":         fmt.Sprintf("%s must be a valid country code (ISO 3166-1 alpha-2, alpha-3, or numeric).", field),
		"strong":               fmt.Sprintf("%s must contain at least one capital letter, one number, and one special character.", field),
		"owasp_password":       fmt.Sprintf("%s must be at least 8 characters and contain at least one uppercase letter, one lowercase letter, one number, and one special character.", field),
	}

	msg, exists := tr[err.Tag()]

	if !exists {
		msg = fmt.Sprintf("%s is not valid.", field)
	}

	return toSentenceCase(msg)
}
