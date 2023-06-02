package validation

import (
	"fmt"

	"github.com/christian-gama/nutrai-api/pkg/errutil/errors"
)

// errorMsgs is a map of errors messages that are more user friendly.
var errorMsgs = map[string]func(field string, param string) error{
	"eqcsfield": func(field string, param string) error {
		return errors.Invalid(field, "must be equal to %s", param)
	},

	"eqfield": func(field string, param string) error {
		return errors.Invalid(field, "must be equal to %s", param)
	},

	"fieldcontains": func(field string, param string) error {
		return errors.Invalid(field, "must contain %s", param)
	},

	"fieldexcludes": func(field string, param string) error {
		return errors.Invalid(field, "must not contain %s", param)
	},

	"cidr": func(field string, param string) error {
		return errors.Invalid(field, "is not a valid CIDR")
	},

	"cidrv4": func(field string, param string) error {
		return errors.Invalid(field, "is not a valid IPv4 CIDR")
	},

	"cidrv6": func(field string, param string) error {
		return errors.Invalid(field, "is not a valid IPv6 CIDR")
	},

	"datauri": func(field string, param string) error {
		return errors.Invalid(field, "is not a valid data URI")
	},

	"fqdn": func(field string, param string) error {
		return errors.Invalid(field, "is not a valid FQDN")
	},

	"hostname": func(field string, param string) error {
		return errors.Invalid(field, "is not a valid hostname (RFC 952)")
	},

	"hostname_port": func(field string, param string) error {
		return errors.Invalid(field, "is not a valid HostPort")
	},

	"hostname_rfc1123": func(field string, param string) error {
		return errors.Invalid(field, "is not a valid hostname (RFC 1123)")
	},

	"ip": func(field string, param string) error {
		return errors.Invalid(field, "is not a valid IP address")
	},

	"ip4_addr": func(field string, param string) error {
		return errors.Invalid(field, "is not a valid IPv4 address")
	},

	"ip6_addr": func(field string, param string) error {
		return errors.Invalid(field, "is not a valid IPv6 address")
	},

	"ip_addr": func(field string, param string) error {
		return errors.Invalid(field, "is not a valid IP address")
	},

	"ipv4": func(field string, param string) error {
		return errors.Invalid(field, "is not a valid IPv4 address")
	},

	"ipv6": func(field string, param string) error {
		return errors.Invalid(field, "is not a valid IPv6 address")
	},

	"mac": func(field string, param string) error {
		return errors.Invalid(field, "is not a valid MAC address")
	},

	"tcp4_addr": func(field string, param string) error {
		return errors.Invalid(field, "is not a valid TCPv4 address")
	},

	"tcp6_addr": func(field string, param string) error {
		return errors.Invalid(field, "is not a valid TCPv6 address")
	},

	"tcp_addr": func(field string, param string) error {
		return errors.Invalid(field, "is not a valid TCP address")
	},

	"udp4_addr": func(field string, param string) error {
		return errors.Invalid(field, "is not a valid UDPv4 address")
	},

	"udp6_addr": func(field string, param string) error {
		return errors.Invalid(field, "is not a valid UDPv6 address")
	},

	"udp_addr": func(field string, param string) error {
		return errors.Invalid(field, "is not a valid UDP address")
	},

	"unix_addr": func(field string, param string) error {
		return errors.Invalid(field, "is not a valid UNIX address")
	},

	"uri": func(field string, param string) error {
		return errors.Invalid(field, "is not a valid URI")
	},

	"url_encoded": func(field string, param string) error {
		return errors.Invalid(field, "is not a valid URL-encoded string")
	},

	"urn_rfc2141": func(field string, param string) error {
		return errors.Invalid(field, "is not a valid URN (RFC 2141)")
	},

	"alpha": func(field string, param string) error {
		return errors.Invalid(field, "must contain only alphabetic characters")
	},

	"alphanum": func(field string, param string) error {
		return errors.Invalid(field, "must contain only alphanumeric characters")
	},

	"alphanumunicode": func(field string, param string) error {
		return errors.Invalid(field, "must contain only alphanumeric Unicode characters")
	},

	"alphaunicode": func(field string, param string) error {
		return errors.Invalid(field, "must contain only alphabetic Unicode characters")
	},

	"ascii": func(field string, param string) error {
		return errors.Invalid(field, "must contain only ASCII characters")
	},

	"boolean": func(field string, param string) error {
		return errors.Invalid(field, "must be a boolean value")
	},

	"contains": func(field string, param string) error {
		return errors.Invalid(field, "must contain %s", param)
	},

	"containsany": func(field string, param string) error {
		return errors.Invalid(
			field,
			fmt.Sprintf("must contain at least one of the following characters: %s", param),
		)
	},

	"containsrune": func(field string, param string) error {
		return errors.Invalid(field, "must contain the character %s", param)
	},

	"endsnotwith": func(field string, param string) error {
		return errors.Invalid(field, "must not end with %s", param)
	},

	"endswith": func(field string, param string) error {
		return errors.Invalid(field, "must end with %s", param)
	},

	"excludes": func(field string, param string) error {
		return errors.Invalid(field, "must not contain %s", param)
	},

	"excludesall": func(field string, param string) error {
		return errors.Invalid(
			field,
			fmt.Sprintf("must not contain any of the following characters: %s", param),
		)
	},

	"excludesrune": func(field string, param string) error {
		return errors.Invalid(field, "must not contain the character %s", param)
	},

	"lowercase": func(field string, param string) error {
		return errors.Invalid(field, "must contain only lowercase characters")
	},

	"multibyte": func(field string, param string) error {
		return errors.Invalid(field, "must contain at least one multi-byte character")
	},

	"number": func(field string, param string) error {
		return errors.Invalid(field, "must be a number")
	},

	"numeric": func(field string, param string) error {
		return errors.Invalid(field, "must contain only numeric characters")
	},

	"printascii": func(field string, param string) error {
		return errors.Invalid(field, "must contain only printable ASCII characters")
	},

	"startsnotwith": func(field string, param string) error {
		return errors.Invalid(field, "must not start with %s", param)
	},

	"startswith": func(field string, param string) error {
		return errors.Invalid(field, "must start with %s", param)
	},

	"uppercase": func(field string, param string) error {
		return errors.Invalid(field, "must contain only uppercase characters")
	},

	"base64": func(field string, param string) error {
		return errors.Invalid(field, "is not a valid Base64 string")
	},

	"base64url": func(field string, param string) error {
		return errors.Invalid(field, "is not a valid Base64URL string")
	},

	"bic": func(field string, param string) error {
		return errors.Invalid(field, "is not a valid Business Identifier Code (BIC)")
	},

	"bcp47_language_tag": func(field string, param string) error {
		return errors.Invalid(field, "is not a valid language tag (BCP 47)")
	},

	"btc_addr": func(field string, param string) error {
		return errors.Invalid(field, "is not a valid Bitcoin address")
	},

	"btc_addr_bech32": func(field string, param string) error {
		return errors.Invalid(field, "is not a valid Bitcoin Bech32 address (segwit)")
	},

	"credit_card": func(field string, param string) error {
		return errors.Invalid(field, "is not a valid credit card number")
	},

	"datetime": func(field string, param string) error {
		return errors.Invalid(field, "is not a valid datetime")
	},

	"e164": func(field string, param string) error {
		return errors.Invalid(field, "is not a valid e164 formatted phone number")
	},

	"email": func(field string, param string) error {
		return errors.Invalid(field, "is not a valid e-mail address")
	},

	"eth_addr": func(field string, param string) error {
		return errors.Invalid(field, "is not a valid Ethereum address")
	},

	"hexadecimal": func(field string, param string) error {
		return errors.Invalid(field, "is not a valid hexadecimal string")
	},

	"hexcolor": func(field string, param string) error {
		return errors.Invalid(field, "is not a valid hexcolor string")
	},

	"hsl": func(field string, param string) error {
		return errors.Invalid(field, "is not a valid HSL string")
	},

	"hsla": func(field string, param string) error {
		return errors.Invalid(field, "is not a valid HSLA string")
	},

	"html": func(field string, param string) error {
		return errors.Invalid(field, "contains invalid HTML tags or attributes")
	},

	"html_encoded": func(field string, param string) error {
		return errors.Invalid(field, "is not a valid HTML-encoded string")
	},

	"iban": func(field string, param string) error {
		return errors.Invalid(
			field,
			"is not a valid International Standard Book Number (ISBN)",
		)
	},

	"isbn10": func(field string, param string) error {
		return errors.Invalid(
			field,
			"is not a valid International Standard Book Number 10 (ISBN-10)",
		)
	},

	"isbn13": func(field string, param string) error {
		return errors.Invalid(
			field,
			"is not a valid International Standard Book Number 13 (ISBN-13)",
		)
	},

	"iso3166_1_alpha2": func(field string, param string) error {
		return errors.Invalid(
			field,
			"is not a valid two-letter country code (ISO 3166-1 alpha-2)",
		)
	},

	"iso3166_1_alpha3": func(field string, param string) error {
		return errors.Invalid(
			field,
			"is not a valid three-letter country code (ISO 3166-1 alpha-3)",
		)
	},

	"iso3166_1_alpha_numeric": func(field string, param string) error {
		return errors.Invalid(
			field,
			"is not a valid numeric country code (ISO 3166-1 numeric)",
		)
	},

	"iso3166_2": func(field string, param string) error {
		return errors.Invalid(field, "is not a valid country subdivision code (ISO 3166-2)")
	},

	"iso4217": func(field string, param string) error {
		return errors.Invalid(field, "is not a valid currency code (ISO 4217)")
	},

	"json": func(field string, param string) error {
		return errors.Invalid(field, "is not a valid JSON string")
	},

	"jwt": func(field string, param string) error {
		return errors.Invalid(field, "is not a valid JSON Web Token (JWT)")
	},

	"latitude": func(field string, param string) error {
		return errors.Invalid(field, "is not a valid latitude")
	},

	"longitude": func(field string, param string) error {
		return errors.Invalid(field, "is not a valid longitude")
	},

	"postcode_iso3166_alpha2": func(field string, param string) error {
		return errors.Invalid(field, "is not a valid postcode (ISO 3166-1 alpha-2)")
	},

	"postcode_iso3166_alpha2_field": func(field string, param string) error {
		return errors.Invalid(field, "is not a valid postcode (ISO 3166-1 alpha-2)")
	},

	"rgb": func(field string, param string) error {
		return errors.Invalid(field, "is not a valid RGB string")
	},

	"rgba": func(field string, param string) error {
		return errors.Invalid(field, "is not a valid RGBA string")
	},

	"ssn": func(field string, param string) error {
		return errors.Invalid(field, "is not a valid Social Security Number (SSN)")
	},

	"timezone": func(field string, param string) error {
		return errors.Invalid(field, "is not a valid timezone")
	},

	"uuid": func(field string, param string) error {
		return errors.Invalid(field, "is not a valid UUID")
	},

	"uuid3": func(field string, param string) error {
		return errors.Invalid(field, "is not a valid UUID v3")
	},

	"uuid3_rfc4122": func(field string, param string) error {
		return errors.Invalid(field, "is not a valid UUID v3 (RFC 4122)")
	},

	"uuid4": func(field string, param string) error {
		return errors.Invalid(field, "is not a valid UUID v4")
	},

	"uuid4_rfc4122": func(field string, param string) error {
		return errors.Invalid(field, "is not a valid UUID v4 (RFC 4122)")
	},

	"uuid5": func(field string, param string) error {
		return errors.Invalid(field, "is not a valid UUID v5")
	},

	"uuid5_rfc4122": func(field string, param string) error {
		return errors.Invalid(field, "is not a valid UUID v5 (RFC 4122)")
	},

	"uuid_rfc4122": func(field string, param string) error {
		return errors.Invalid(field, "is not a valid UUID (RFC 4122)")
	},

	"md4": func(field string, param string) error {
		return errors.Invalid(field, "is not a valid MD4 hash")
	},

	"md5": func(field string, param string) error {
		return errors.Invalid(field, "is not a valid MD5 hash")
	},

	"sha256": func(field string, param string) error {
		return errors.Invalid(field, "is not a valid SHA256 hash")
	},

	"sha384": func(field string, param string) error {
		return errors.Invalid(field, "is not a valid SHA384 hash")
	},

	"sha512": func(field string, param string) error {
		return errors.Invalid(field, "is not a valid SHA512 hash")
	},

	"ripemd128": func(field string, param string) error {
		return errors.Invalid(field, "is not a valid RIPEMD-128 hash")
	},

	"ripemd160": func(field string, param string) error {
		return errors.Invalid(field, "is not a valid RIPEMD-160 hash")
	},

	"tiger128": func(field string, param string) error {
		return errors.Invalid(field, "is not a valid TIGER128 hash")
	},

	"tiger160": func(field string, param string) error {
		return errors.Invalid(field, "is not a valid TIGER160 hash")
	},

	"tiger192": func(field string, param string) error {
		return errors.Invalid(field, "is not a valid TIGER192 hash")
	},

	"semver": func(field string, param string) error {
		return errors.Invalid(field, "is not a valid semantic version")
	},

	"ulid": func(field string, param string) error {
		return errors.Invalid(field, "is not a valid ULID")
	},

	"eq": func(field string, param string) error {
		return errors.Invalid(field, "is not equal to %s", param)
	},

	"gt": func(field string, param string) error {
		return errors.Invalid(field, "is not greater than %s", param)
	},

	"gte": func(field string, param string) error {
		return errors.Invalid(
			field,
			fmt.Sprintf("is not greater than or equal to %s", param),
		)
	},

	"lt": func(field string, param string) error {
		return errors.Invalid(field, "is not less than %s", param)
	},

	"lte": func(field string, param string) error {
		return errors.Invalid(field, "is not less than or equal to %s", param)
	},

	"ne": func(field string, param string) error {
		return errors.Invalid(field, "is not not equal to %s", param)
	},

	"dir": func(field string, param string) error {
		return errors.Invalid(field, "is not a valid directory")
	},

	"file": func(field string, param string) error {
		return errors.Invalid(field, "is not a valid file path")
	},

	"isdefault": func(field string, param string) error {
		return errors.Invalid(field, "is not the default value")
	},

	"len": func(field string, param string) error {
		return errors.Invalid(field, "is not the correct length")
	},

	"max": func(field string, param string) error {
		return errors.Invalid(field, "is not less than or equal to %s", param)
	},

	"min": func(field string, param string) error {
		return errors.Invalid(
			field,
			fmt.Sprintf("is not greater than or equal to %s", param),
		)
	},

	"oneof": func(field string, param string) error {
		return errors.Invalid(field, "is not one of the allowed values")
	},

	"required": func(field string, param string) error {
		return errors.Invalid(field, "is required")
	},

	"required_if": func(field string, param string) error {
		return errors.Invalid(field, "%s is required", param)
	},

	"required_unless": func(field string, param string) error {
		return errors.Invalid(field, "is required")
	},

	"required_with": func(field string, param string) error {
		return errors.Invalid(field, "is required with %s", param)
	},

	"required_with_all": func(field string, param string) error {
		return errors.Invalid(field, "is required with %s", param)
	},

	"required_without": func(field string, param string) error {
		return errors.Invalid(field, "is required without %s", param)
	},

	"required_without_all": func(field string, param string) error {
		return errors.Invalid(field, "is required without %s", param)
	},

	"excluded_if": func(field string, param string) error {
		return errors.Invalid(field, "is not allowed when %s is present", param)
	},

	"excluded_unless": func(field string, param string) error {
		return errors.Invalid(
			field,
			fmt.Sprintf("is not allowed when %s is  not present", param),
		)
	},

	"excluded_with": func(field string, param string) error {
		return errors.Invalid(field, "is not allowed when %s is present", param)
	},

	"excluded_with_all": func(field string, param string) error {
		return errors.Invalid(field, "is not allowed when %s is present", param)
	},

	"excluded_without": func(field string, param string) error {
		return errors.Invalid(
			field,
			fmt.Sprintf("is not allowed when %s is not present", param),
		)
	},

	"excluded_without_all": func(field string, param string) error {
		return errors.Invalid(
			field,
			fmt.Sprintf("is not allowed when %s is not present", param),
		)
	},

	"unique": func(field string, param string) error {
		return errors.Invalid(field, "is not unique")
	},

	"iscolor": func(field string, param string) error {
		return errors.Invalid(field, "is not a valid color")
	},

	"country_code": func(field string, param string) error {
		return errors.Invalid(field, "is not a valid country code")
	},
}

func defaultError(field string, param string) error {
	return errors.Invalid(field, "unknown error")
}
