package a2p

import validation "github.com/go-ozzo/ozzo-validation/v4"

type CustomerProfileData struct {
	SID            string `json:"customer_profile_sid"`
	FriendlyName   string `json:"friendly_name"`
	Email          string `json:"email"`
	StatusCallback string `json:"status_callback"`
	PolicySid      string `json:"policy_sid"`
}

type BusinessInfoData struct {
	SID                        string `json:"end_user_sid"`
	BusinessName               string `json:"business_name"`
	SocialMediaProfileUrls     string `json:"social_media_profile_urls"`
	WebsiteUrl                 string `json:"website_url"`
	BusinessRegionsOfOperation string `json:"business_regions_of_operation"`
	BusinessType               string `json:"business_type"`
	BusinessRegistrationId     string `json:"business_registration_identifier"`
	BusinessIdentity           string `json:"business_identity"`
	BusinessIndustry           string `json:"business_industry"`
	BusinessRegistrationNumber string `json:"business_registration_number"`
}

type EndUserAuthorizedRep1BusinessInfoData struct {
	SID                            string `json:"end_user_rep1_sid"`
	Rep1BusinessName               string `json:"end_user_rep1_business_name"`
	Rep1SocialMediaProfileUrls     string `json:"end_user_rep1_social_media_profile_urls"`
	Rep1WebsiteUrl                 string `json:"end_user_rep1_website_url"`
	Rep1BusinessRegionsOfOperation string `json:"end_user_rep1_business_regions_of_operation"`
	Rep1BusinessType               string `json:"end_user_rep1_business_type"`
	Rep1BusinessRegistrationId     string `json:"end_user_rep1_business_registration_identifier"`
	Rep1BusinessIdentity           string `json:"end_user_rep1_business_identity"`
	Rep1BusinessIndustry           string `json:"end_user_rep1_business_industry"`
	Rep1BusinessRegistrationNumber string `json:"end_user_rep1_business_registration_number"`
}

type EndUserAssignmentData struct {
	SID                string `json:"assignment_resource_sid"`
	CustomerProfileSid string `json:"customer_profile_sid"`
	EndUserSid         string `json:"end_user_sid"`
}

type AddressData struct {
	SID             string `json:"address_sid"`
	PathAccountSid  string `json:"path_account_sid"`
	CustomerName    string `json:"customer_name"`
	Street          string `json:"street"`
	City            string `json:"city"`
	Region          string `json:"region"`
	PostalCode      string `json:"postal_code"`
	IsoCountry      string `json:"iso_country"`
	FriendlyName    string `json:"friendly_name"`
	StreetSecondary string `json:"street_secondary"`
}

type SupportingDocumentData struct {
	SID          string `json:"supporting_document_sid"`
	FriendlyName string `json:"friendly_name"`
	AddressSid   string `json:"address_sid"`
}

type TrustProductData struct {
	SID            string `json:"trust_product_sid"`
	FriendlyName   string `json:"friendly_name"`
	Email          string `json:"email"`
	PolicySid      string `json:"policy_sid"`
	StatusCallback string `json:"status_callback"`
}

type EndUserMessagingProfileData struct {
	SID           string `json:"end_user_messaging_profile_sid"`
	CompanyType   string `json:"company_type"`
	StockExchange string `json:"stock_exchange"`
	StockTicker   string `json:"stock_ticker"`
}

/*
Note :
The customer_profile_bundle_sid is the SID of your customer's Secondary Customer Profile.
The a2p_profile_bundle_sid is the SID of the TrustProduct created SID.
Skip_automatic_sec_vet is an optional Boolean.
*/
type BrandRegistrationData struct {
	CustomerProfileBundleSid string `json:"customer_profile_bundle_sid"`
	A2PProfileBundleSid      string `json:"a2p_profile_bundle_sid"`
}

type MessagingServiceData struct {
	SID               string `json:"messaging_service_sid"`
	FriendlyName      string `json:"friendly_name"`
	InboundRequestUrl string `json:"inbound_request_url"`
	FallbackUrl       string `json:"fallback_url"`
}

type CampaignData struct {
	SID                  string   `json:"campaign_sid"`
	Description          string   `json:"description"`
	Usecase              string   `json:"usecase"`
	CampaignStatus       string   `json:"campaign_status"`
	HasEmbeddedLinks     bool     `json:"has_embedded_links"`
	HasEmbeddedPhone     bool     `json:"has_embedded_phone"`
	MessageSamples       []string `json:"message_samples"`
	MessageFlow          string   `json:"message_flow"`
	BrandRegistrationSid string   `json:"brand_registration_sid"`
}

// Optional : uncomment for step 4.2
type MessagingServiceAdditional struct {
	SID                   string `json:"messaging_service_sid"`
	FriendlyName          string `json:"friendly_name"`
	InboundRequestUrl     string `json:"inbound_request_url"`
	FallbackUrl           string `json:"fallback_url"`
	StatusCallback        string `json:"status_callback"`
	StickySender          bool   `json:"sticky_sender"`
	SmartEncoding         bool   `json:"smart_encoding"`
	MmsConverter          bool   `json:"mms_converter"`
	FallbackToLongCode    bool   `json:"fallback_to_long_code"`
	ScanMessageContent    string `json:"scan_message_content"`
	AreaCodeGeomatch      bool   `json:"area_code_geomatch"`
	ValidityPeriod        int    `json:"validity_period"`
	SynchronousValidation bool   `json:"synchronous_validation"`
	Usecase               string `json:"usecase"`
}

type MessageStatusData struct {
	MessageSid string `json:"message_sid"`
}

type MessageErrorData struct {
	MessageSid string `json:"message_sid"`
}

type AssociatePhoneNumberData struct {
	PhoneNumberSid      string `json:"phone_number_sid"`
	MessagingServiceSid string `json:"messaging_service_sid"`
}

type FinalizeMessagingServiceConfigData struct {
	MessagingServiceSid   string `json:"messaging_service_sid"`
	StatusCallback        string `json:"status_callback"`
	StickySender          bool   `json:"sticky_sender"`
	SmartEncoding         bool   `json:"smart_encoding"`
	MmsConverter          bool   `json:"mms_converter"`
	FallbackToLongCode    bool   `json:"fallback_to_long_code"`
	ScanMessageContent    string `json:"scan_message_content"`
	AreaCodeGeomatch      bool   `json:"area_code_geomatch"`
	ValidityPeriod        int    `json:"validity_period"`
	SynchronousValidation bool   `json:"synchronous_validation"`
}

type FullA2POnboardingParams struct {
	ESid                           string `json:"sid"`
	EToken                         string `json:"token"`
	SubAccountSID                  string `json:"subaccount_sid"`
	SubAccountAuthToken            string `json:"subaccount_auth_token"`
	CustomerName                   string `json:"customer_name"`
	Email                          string `json:"customer_email"`
	PhoneNumber                    string `json:"customer_phone_number"`
	Street                         string `json:"street"`
	City                           string `json:"city"`
	Region                         string `json:"region"`
	PostalCode                     string `json:"postal_code"`
	IsoCountry                     string `json:"iso_country"`
	SocialMediaProfileURLs         string `json:"social_media_profile_urls"`
	WebsiteURL                     string `json:"website_url"`
	BusinessName                   string `json:"business_name"`
	BusinessIndustry               string `json:"business_industry"`
	BusinessType                   string `json:"business_type"`
	BusinessRegistrationId         string `json:"business_registration_identifier"`
	BusinessIdentity               string `json:"business_identity"`
	BusinessRegistrationNumber     string `json:"business_registration_number"`
	RegionOfOperation              string `json:"region_of_operation"`
	TwilioPurchasedPhoneNumber     string `json:"twilio_purchased_phone_number"`
	TwilioPurchasedPhoneNumberSID  string `json:"twilio_purchased_phone_number_sid"`
	AuthorizedRepresentativeName   string `json:"authorized_representative_name"`
	AuthorizedRepresentativeTitle  string `json:"authorized_representative_title"`
	AuthorizedRepresentativeEmail  string `json:"authorized_representative_email"`
	AuthorizedRepresentativePhone  string `json:"authorized_representative_phone"`
	Rep1BusinessName               string `json:"end_user_rep1_business_name"`
	Rep1SocialMediaProfileUrls     string `json:"end_user_rep1_social_media_profile_urls"`
	Rep1WebsiteUrl                 string `json:"end_user_rep1_website_url"`
	Rep1BusinessRegionsOfOperation string `json:"end_user_rep1_business_regions_of_operation"`
	Rep1BusinessType               string `json:"end_user_rep1_business_type"`
	Rep1BusinessRegistrationId     string `json:"end_user_rep1_business_registration_identifier"`
	Rep1BusinessIdentity           string `json:"end_user_rep1_business_identity"`
	Rep1BusinessIndustry           string `json:"end_user_rep1_business_industry"`
	Rep1BusinessRegistrationNumber string `json:"end_user_rep1_business_registration_number"`
	FriendlyName                   string `json:"friendly_name"`
	UseCase                        string `json:"use_case"`
	AreaCode                       string `json:"area_code"`
}

func (f *FullA2POnboardingParams) Validate() error {
	return validation.ValidateStruct(&f,
		validation.Field(&f.CustomerName, validation.Required),
		validation.Field(&f.Email, validation.Required),
		validation.Field(&f.PhoneNumber, validation.Required),
		validation.Field(&f.Street, validation.Required),
		validation.Field(&f.City, validation.Required),
		validation.Field(&f.Region, validation.Required),
		validation.Field(&f.PostalCode, validation.Required),
		validation.Field(&f.IsoCountry, validation.Required),
		validation.Field(&f.SocialMediaProfileURLs, validation.Required),
		validation.Field(&f.WebsiteURL, validation.Required),
		validation.Field(&f.BusinessName, validation.Required),
		validation.Field(&f.BusinessIndustry, validation.Required),
		validation.Field(&f.BusinessType, validation.Required),
		validation.Field(&f.BusinessRegistrationId, validation.Required),
		validation.Field(&f.BusinessIdentity, validation.Required),
		validation.Field(&f.BusinessRegistrationNumber, validation.Required),
		validation.Field(&f.RegionOfOperation, validation.Required),
		validation.Field(&f.TwilioPurchasedPhoneNumber, validation.Required),
		validation.Field(&f.AuthorizedRepresentativeName, validation.Required),
		validation.Field(&f.AuthorizedRepresentativeTitle, validation.Required),
		validation.Field(&f.AuthorizedRepresentativeEmail, validation.Required),
		validation.Field(&f.AuthorizedRepresentativePhone, validation.Required),
		validation.Field(&f.FriendlyName, validation.Required),
		validation.Field(&f.UseCase, validation.Required),
		validation.Field(&f.AreaCode, validation.Required),
	)

}

type FullA2POnboardingResponse struct {
	Message string `json:"message"`
	data    *FullA2POnboardingParams
}
