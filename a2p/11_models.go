package a2p

type CustomerProfileData struct {
	FriendlyName   string
	Email          string
	StatusCallback string
	PolicySid      string
}

type BusinessInfoData struct {
	BusinessName               string
	SocialMediaProfileUrls     string
	WebsiteUrl                 string
	BusinessRegionsOfOperation string
	BusinessType               string
	BusinessRegistrationId     string
	BusinessIdentity           string
	BusinessIndustry           string
	BusinessRegistrationNumber string
}

type EndUserAssignmentData struct {
	CustomerProfileSid string
	EndUserSid         string
}

type AddressData struct {
	PathAccountSid  string
	CustomerName    string
	Street          string
	City            string
	Region          string
	PostalCode      string
	IsoCountry      string
	FriendlyName    string
	StreetSecondary string
}

type SupportingDocumentData struct {
	FriendlyName string
	AddressSid   string
}

type TrustProductData struct {
	FriendlyName   string
	Email          string
	PolicySid      string
	StatusCallback string
}

type EndUserMessagingProfileData struct {
	CompanyType   string
	StockExchange string
	StockTicker   string
}

/*
Note :
The customer_profile_bundle_sid is the SID of your customer's Secondary Customer Profile.
The a2p_profile_bundle_sid is the SID of the TrustProduct created SID.
Skip_automatic_sec_vet is an optional Boolean.
*/
type BrandRegistrationData struct {
	CustomerProfileBundleSid string
	A2PProfileBundleSid      string
}

type MessagingServiceData struct {
	FriendlyName      string
	InboundRequestUrl string
	FallbackUrl       string
}

type CampaignData struct {
	Description          string
	Usecase              string
	HasEmbeddedLinks     bool
	HasEmbeddedPhone     bool
	MessageSamples       []string
	MessageFlow          string
	BrandRegistrationSid string
}

// Optional : uncomment for step 4.2
type MessagingServiceAdditional struct {
	FriendlyName          string
	InboundRequestUrl     string
	FallbackUrl           string
	StatusCallback        string
	StickySender          bool
	SmartEncoding         bool
	MmsConverter          bool
	FallbackToLongCode    bool
	ScanMessageContent    string
	AreaCodeGeomatch      bool
	ValidityPeriod        int
	SynchronousValidation bool
	Usecase               string
}

type MessageStatusData struct {
	MessageSid string
}

type MessageErrorData struct {
	MessageSid string
}

type TwilioOnboardingData struct {
	CustomerName                  string
	CustomerEmail                 string
	CustomerPhoneNumber           string
	BusinessName                  string
	WebsiteURL                    string
	BusinessIndustry              string
	BusinessType                  string
	RegionOfOperation             string
	RegistrationNumber            string
	AuthorizedRepresentativeName  string
	AuthorizedRepresentativeTitle string
	AuthorizedRepresentativeEmail string
	AuthorizedRepresentativePhone string
	FriendlyName                  string
	UseCase                       string
}

type CampaignInfo struct {
	Sid                  string
	Description          string
	UseCase              string
	CampaignStatus       string
	BrandRegistrationSid string
}

type AssociatePhoneNumberData struct {
	PhoneNumberSid      string
	MessagingServiceSid string
}

type FinalizeMessagingServiceConfigData struct {
	MessagingServiceSid   string
	StatusCallback        string
	StickySender          bool
	SmartEncoding         bool
	MmsConverter          bool
	FallbackToLongCode    bool
	ScanMessageContent    string
	AreaCodeGeomatch      bool
	ValidityPeriod        int
	SynchronousValidation bool
}
