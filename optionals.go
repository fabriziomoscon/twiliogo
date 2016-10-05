package twiliogo

import (
	"fmt"
	"strconv"
)

type Optional interface {
	GetParam() (string, string)
}

type Callback string

func (callback Callback) GetParam() (string, string) {
	return "Url", string(callback)
}

type ApplicationSid string

func (applicationSid ApplicationSid) GetParam() (string, string) {
	return "ApplicationSid", string(applicationSid)
}

type VoiceApplicationSid string

func (voiceApplicationSid VoiceApplicationSid) GetParam() (string, string) {
	return "VoiceApplicationSid", string(voiceApplicationSid)
}

type SmsApplicationSid string

func (smsApplicationSid SmsApplicationSid) GetParam() (string, string) {
	return "SmsApplicationSid", string(smsApplicationSid)
}

type Method string

func (method Method) GetParam() (string, string) {
	return "Method", string(method)
}

type FallbackUrl string

func (fallbackUrl FallbackUrl) GetParam() (string, string) {
	return "FallbackUrl", string(fallbackUrl)
}

type FallbackMethod string

func (fallbackMethod FallbackMethod) GetParam() (string, string) {
	return "FallbackMethod", string(fallbackMethod)
}

type StatusCallback string

func (statusCallback StatusCallback) GetParam() (string, string) {
	return "StatusCallback", string(statusCallback)
}

type StatusCallbackMethod string

func (statusCallbackMethod StatusCallbackMethod) GetParam() (string, string) {
	return "StatusCallbackMethod", string(statusCallbackMethod)
}

type SendDigits string

func (sendDigits SendDigits) GetParam() (string, string) {
	return "SendDigits", string(sendDigits)
}

type IfMachine string

func (ifMachine IfMachine) GetParam() (string, string) {
	return "IfMachine", string(ifMachine)
}

type Timeout string

func (timeout Timeout) GetParam() (string, string) {
	return "Timeout", string(timeout)
}

type Record string

func (record Record) GetParam() (string, string) {
	return "Record", string(record)
}

type To string

func (to To) GetParam() (string, string) {
	return "To", string(to)
}

type From string

func (from From) GetParam() (string, string) {
	return "From", string(from)
}

type Status string

func (status Status) GetParam() (string, string) {
	return "Status", string(status)
}

type StartTime string

func (startTime StartTime) GetParam() (string, string) {
	return "StartTime", string(startTime)
}

type ParentCallSid string

func (parentCallSid ParentCallSid) GetParam() (string, string) {
	return "ParentCallSid", string(parentCallSid)
}

type DateSent string

func (dateSent DateSent) GetParam() (string, string) {
	return "DateSent", string(dateSent)
}

type Body string

func (body Body) GetParam() (string, string) {
	return "Body", string(body)
}

type MediaUrl string

func (mediaUrl MediaUrl) GetParam() (string, string) {
	return "MediaUrl", string(mediaUrl)
}

type FriendlyName string

func (friendlyName FriendlyName) GetParam() (string, string) {
	return "FriendlyName", string(friendlyName)
}

type PhoneNumber string

func (phoneNumber PhoneNumber) GetParam() (string, string) {
	return "PhoneNumber", string(phoneNumber)
}

type AreaCode string

func (areaCode AreaCode) GetParam() (string, string) {
	return "AreaCode", string(areaCode)
}

//For AvaliablePhoneNumbers
type Contains string

func (contains Contains) GetParam() (string, string) {
	return "Contains", string(contains)
}

type SmsEnabled bool

func (smsEnabled SmsEnabled) GetParam() (string, string) {
	return "SmsEnabled", strconv.FormatBool(bool(smsEnabled))
}

type MmsEnabled bool

func (mmsEnabled MmsEnabled) GetParam() (string, string) {
	return "MmsEnabled", strconv.FormatBool(bool(mmsEnabled))
}

type VoiceEnabled bool

func (voiceEnabled VoiceEnabled) GetParam() (string, string) {
	return "VoiceEnabled", strconv.FormatBool(bool(voiceEnabled))
}

type ExcludeAllAddressRequired bool

func (excludeAllAddressRequired ExcludeAllAddressRequired) GetParam() (string, string) {
	return "ExcludeAllAddressRequired", strconv.FormatBool(bool(excludeAllAddressRequired))
}

type ExcludeLocalAddressRequired bool

func (excludeLocalAddressRequired ExcludeLocalAddressRequired) GetParam() (string, string) {
	return "ExcludeLocalAddressRequired", strconv.FormatBool(bool(excludeLocalAddressRequired))
}

type ExcludeForiegnAddressRequired bool

func (excludeForiegnAddressRequired ExcludeForiegnAddressRequired) GetParam() (string, string) {
	return "ExcludeForiegnAddressRequired", strconv.FormatBool(bool(excludeForiegnAddressRequired))
}

type Beta bool

func (beta Beta) GetParam() (string, string) {
	return "Beta", strconv.FormatBool(bool(beta))
}

type NearNumber string

func (nearNumber NearNumber) GetParam() (string, string) {
	return "NearNumber", string(nearNumber)
}

type NearLatLong struct {
	Latitude  float64
	Longitude float64
}

func (nearLatLong NearLatLong) GetParam() (string, string) {
	return "NearLatLong", fmt.Sprintf("%.5f,%.5f",
		nearLatLong.Latitude,
		nearLatLong.Longitude,
	)
}

type Distance uint16

func (distance Distance) GetParam() (string, string) {
	return "Distance", fmt.Sprintf("%d", distance)
}

type InPostalCode string

func (inPostalCode InPostalCode) GetParam() (string, string) {
	return "InPostalCode", string(inPostalCode)
}

type InRegion string

func (inRegion InRegion) GetParam() (string, string) {
	return "InRegion", string(inRegion)
}

type InRateCenter string

func (inRateCenter InRateCenter) GetParam() (string, string) {
	return "InRateCenter", string(inRateCenter)
}

type InLata string

func (inLata InLata) GetParam() (string, string) {
	return "InLata", string(inLata)
}

// Application optionals
type VoiceUrl string

func (voiceUrl VoiceUrl) GetParam() (string, string) {
	return "VoiceUrl", string(voiceUrl)
}

type VoiceUrlMethod string

func (voiceUrlMethod VoiceUrlMethod) GetParam() (string, string) {
	return "VoiceUrlMethod", string(voiceUrlMethod)
}

type VoiceFallbackUrl string

func (voiceFallbackUrl VoiceFallbackUrl) GetParam() (string, string) {
	return "VoiceFallbackUrl", string(voiceFallbackUrl)
}

type VoiceFallbackMethod string

func (voiceFallbackMethod VoiceFallbackMethod) GetParam() (string, string) {
	return "VoiceFallbackMethod", string(voiceFallbackMethod)
}

type SmsUrl string

func (smsUrl SmsUrl) GetParam() (string, string) {
	return "SmsUrl", string(smsUrl)
}

type SmsUrlMethod string

func (smsUrlMethod SmsUrlMethod) GetParam() (string, string) {
	return "SmsUrlMethod", string(smsUrlMethod)
}

type SmsFallbackUrl string

func (smsFallbackUrl SmsFallbackUrl) GetParam() (string, string) {
	return "SmsFallbackUrl", string(smsFallbackUrl)
}

type SmsFallbackMethod string

func (smsFallbackMethod SmsFallbackMethod) GetParam() (string, string) {
	return "SmsFallbackMethod", string(smsFallbackMethod)
}

type SmsStatusCallback string

func (smsStatusCallback SmsStatusCallback) GetParam() (string, string) {
	return "SmsStatusCallback", string(smsStatusCallback)
}

type MessageStatusCallback string

func (messageStatusCallback MessageStatusCallback) GetParam() (string, string) {
	return "MessageStatusCallback", string(messageStatusCallback)
}

// for Address
type CustomerName string

func (customerName CustomerName) GetParam() (string, string) {
	return "CustomerName", string(customerName)
}

type Street string

func (street Street) GetParam() (string, string) {
	return "Street", string(street)
}

type City string

func (city City) GetParam() (string, string) {
	return "City", string(city)
}

type Region string

func (region Region) GetParam() (string, string) {
	return "Region", string(region)
}

type PostalCode string

func (postalcode PostalCode) GetParam() (string, string) {
	return "PostalCode", string(postalcode)
}

type IsoCountry string

func (isoCountry IsoCountry) GetParam() (string, string) {
	return "IsoCountry", string(isoCountry)
}
