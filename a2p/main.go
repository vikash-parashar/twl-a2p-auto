package a2p

import (
	"fmt"

	messaging "github.com/twilio/twilio-go/rest/messaging/v1"
)

func (s *A2PService) OnboardCustomer(params FullA2POnboardingParams) (FullA2POnboardingResponse, error) {
	err := params.Validate()
	if err != nil {
		return FullA2POnboardingResponse{}, fmt.Errorf("invalid onboarding data: %w", err)
	}

	// Stage 1.0: Automate Subaccount Creation - If the SubAccountSID and SubAccountAuthToken are not provided, create a new subaccount
	if params.SubAccountSID == "" || params.SubAccountAuthToken == "" {
		sid, token, err := s.CreateSubAccount(SubAccountData{
			FriendlyName: params.CustomerName,
		})

		if err != nil {
			return FullA2POnboardingResponse{}, fmt.Errorf("error at stage 1.0 : %w", err)
		}
		params.SubAccountSID = sid
		params.SubAccountAuthToken = token
	}

	// Stage 2.1: Create a secondary customer profile
	customerProfileSid, err := s.CreateSecondaryCustomerProfile(CustomerProfileData{
		FriendlyName:   params.FriendlyName,
		Email:          params.Email,
		PolicySid:      "RNdfbf3fae0e1107f8aded0e7cead80bf5",
		StatusCallback: "www.demo.com/callback/status",
	})
	if err != nil {
		return FullA2POnboardingResponse{}, fmt.Errorf("error at stage 2.1: %w", err)
	}

	// Stage 2.2: Create an EndUser Business Information resource
	endUserBusinessInfoSID, err := s.CreateEndUserBusinessInfo(BusinessInfoData{
		BusinessName:               params.BusinessName,
		SocialMediaProfileUrls:     params.SocialMediaProfileURLs,
		WebsiteUrl:                 params.WebsiteURL,
		BusinessRegionsOfOperation: params.RegionOfOperation,
		BusinessType:               params.BusinessType,
		BusinessRegistrationId:     params.BusinessRegistrationId,
		BusinessIdentity:           params.BusinessIdentity,
		BusinessIndustry:           params.BusinessIndustry,
		BusinessRegistrationNumber: params.BusinessRegistrationNumber,
	})
	if err != nil {
		return FullA2POnboardingResponse{}, fmt.Errorf("error at stage 2.2: %w", err)
	}

	// Stage 2.3: Attach EndUser to the Secondary Customer Profile
	// attachEndUserToProfileSID
	_, err = s.AttachEndUserToProfile(EndUserAssignmentData{
		CustomerProfileSid: customerProfileSid,
		EndUserSid:         endUserBusinessInfoSID,
	})
	if err != nil {
		return FullA2POnboardingResponse{}, fmt.Errorf("error at stage 2.3: %w", err)
	}

	// Step 2.4. Create an EndUser resource of type: authorized_representative_1

	endUserAuthorizedRep1SID, err := s.CreateEndUserAuthorizedRep1(EndUserAuthorizedRep1BusinessInfoData{
		Rep1BusinessName:               params.Rep1BusinessName,
		Rep1SocialMediaProfileUrls:     params.Rep1SocialMediaProfileUrls,
		Rep1WebsiteUrl:                 params.Rep1WebsiteUrl,
		Rep1BusinessRegionsOfOperation: params.Rep1BusinessRegionsOfOperation,
		Rep1BusinessType:               params.Rep1BusinessType,
		Rep1BusinessRegistrationId:     params.Rep1BusinessRegistrationId,
		Rep1BusinessIdentity:           params.Rep1BusinessIdentity,
		Rep1BusinessIndustry:           params.Rep1BusinessIndustry,
		Rep1BusinessRegistrationNumber: params.Rep1BusinessRegistrationNumber,
	})
	if err != nil {
		return FullA2POnboardingResponse{}, fmt.Errorf("error at stage 2.4: %w", err)
	}

	// Step 2.5: Attach EndUser to the Secondary Customer Profile
	// attachEndUserToProfileSID
	_, err = s.AttachEndUserAuthorizedRep1ToProfile(EndUserAssignmentData{
		CustomerProfileSid: customerProfileSid,
		EndUserSid:         endUserAuthorizedRep1SID,
	})
	if err != nil {
		return FullA2POnboardingResponse{}, fmt.Errorf("error at stage 2.5: %w", err)
	}

	// Step 2.6 Create An Address Resource and returns address sid
	addressSID, err := s.CreateAddressResource(AddressData{
		PathAccountSid: params.ESid,
		CustomerName:   params.CustomerName,
		Street:         params.Street,
		City:           params.City,
		Region:         params.Region,
		PostalCode:     params.PostalCode,
		IsoCountry:     params.IsoCountry,
		FriendlyName:   fmt.Sprintf("%s - Address Resource", params.CustomerName),
	})
	if err != nil {
		return FullA2POnboardingResponse{}, fmt.Errorf("error at stage 2.6: %w", err)
	}

	// Step 2.7 Create a supporting document resource and returns supporting_document_sid
	supportingDocumentSID, err := s.CreateSupportingDocument(SupportingDocumentData{
		FriendlyName: fmt.Sprintf("%s - Business License Document", params.CustomerName),
		AddressSid:   addressSID,
	})
	if err != nil {
		return FullA2POnboardingResponse{}, fmt.Errorf("error at stage 2.7: %w", err)
	}

	// Step 2.8 Attach the supporting document to the Secondary Customer Profile
	//attachSupportingDocumentToProfileSID
	_, err = s.AttachSupportingDocumentToProfile(customerProfileSid, &supportingDocumentSID)
	if err != nil {
		return FullA2POnboardingResponse{}, fmt.Errorf("error at stage 2.8: %w", err)
	}

	// Step 2.9. Evaluate the Secondary Customer Profile
	//evaluateSecondaryCustomerProfileSID
	_, err = s.EvaluateSecondaryCustomerProfile(customerProfileSid)
	if err != nil {
		return FullA2POnboardingResponse{}, fmt.Errorf("error at stage 2.9: %w", err)
	}

	// Step 2.10. Submit the Secondary Customer Profile for review  - status must be set to pending-review
	// submitSecondaryCustomerProfileForReviewSID
	_, err = s.SubmitSecondaryCustomerProfileForReview(customerProfileSid)
	if err != nil {
		return FullA2POnboardingResponse{}, fmt.Errorf("error at stage 2.10: %w", err)
	}

	// Step 3.1: Create a TrustProduct Resource
	trustProductSID, err := s.CreateTrustProduct(TrustProductData{
		FriendlyName:   params.FriendlyName,
		PolicySid:      "RNdfbf3fae0e1107f8aded0e7cead80bf5",
		Email:          params.Email,
		StatusCallback: "www.demo.com/callback/status",
	})
	if err != nil {
		return FullA2POnboardingResponse{}, fmt.Errorf("error at stage 3.1: %w", err)
	}

	// Step 3.2: Create an EndUser Resource of Type us_a2p_messaging_profile_information
	endUserMessagingProfileSID, err := s.CreateEndUserMessagingProfile(EndUserMessagingProfileData{
		CompanyType:   params.BusinessType,
		StockExchange: "",
		StockTicker:   "",
	})
	if err != nil {
		return FullA2POnboardingResponse{}, fmt.Errorf("error at stage 3.2: %w", err)
	}

	// Step 3.3: Attach the EndUser to the TrustProduct
	//attachEndUserToTrustProductSID
	_, err = s.AttachEndUserToTrustProduct(trustProductSID, endUserMessagingProfileSID)
	if err != nil {
		return FullA2POnboardingResponse{}, fmt.Errorf("error at stage 3.3: %w", err)
	}

	// Step 3.4: Attach the Secondary Customer Profile to the TrustProduct
	//attachSecondaryCustomerProfileToTrustProductSID
	_, err = s.AttachSecondaryCustomerProfileToTrustProduct(trustProductSID, customerProfileSid)
	if err != nil {
		return FullA2POnboardingResponse{}, fmt.Errorf("error at stage 3.4: %w", err)
	}

	// Step 3.5: Evaluate the TrustProduct
	// evaluateTrustProductSID
	_, err = s.EvaluateTrustProduct(trustProductSID, "RNdfbf3fae0e1107f8aded0e7cead80bf5")
	if err != nil {
		return FullA2POnboardingResponse{}, fmt.Errorf("error at stage 3.5: %w", err)
	}

	// Step 3.6: Submit the TrustProduct for Review  - status must be set to pending-review
	// submitTrustProductForReviewSID
	_, err = s.SubmitTrustProductForReview(trustProductSID)
	if err != nil {
		return FullA2POnboardingResponse{}, fmt.Errorf("error at stage 3.6: %w", err)
	}

	// Step 4.1: Create a BrandRegistration
	brandRegistrationSID, brandRegistrationStatus, err := s.CreateBrandRegistration(BrandRegistrationData{
		CustomerProfileBundleSid: customerProfileSid,
		A2PProfileBundleSid:      trustProductSID,
	})

	if err != nil {
		return FullA2POnboardingResponse{}, fmt.Errorf("error at stage 4.1: %w", err)
	}

	// Step 5.1: Create a MessagingService Resource - This will return MessageServiceSID
	messagingServiceSID, err := s.CreateMessagingService(MessagingServiceData{
		FriendlyName:      params.FriendlyName,
		InboundRequestUrl: "www.demo.com/inbound",
		FallbackUrl:       "www.demo.com/fallback",
	})

	if err != nil {
		return FullA2POnboardingResponse{}, fmt.Errorf("error at stage 5.1: %w", err)
	}

	if brandRegistrationStatus != "APPROVED" {
		// Step 6.1: Create the A2P Campaign
		// campaignSID
		_, err := s.CreateA2PCampaign(messagingServiceSID, CampaignData{
			BrandRegistrationSid: brandRegistrationSID,
		})

		if err != nil {
			return FullA2POnboardingResponse{}, fmt.Errorf("error at stage 6.2: %w", err)
		}

		// Step 7.1: Add a Phone Number to the Messaging Service , once you have a phone number, you can associate it with the messaging service
		// response
		_, err = s.AddPhoneNumberToMessagingService(messagingServiceSID, &messaging.CreatePhoneNumberParams{
			PhoneNumberSid: &params.TwilioPurchasedPhoneNumberSID,
		})

		if err != nil {
			return FullA2POnboardingResponse{}, fmt.Errorf("error at stage 7.3: %w", err)
		}

		return FullA2POnboardingResponse{
			Message: "A2P Onboarding Request Submitted Successful",
			data:    &params,
		}, nil
	} else {

		// Step 7.3: Add a Phone Number to the Messaging Service , once you have a phone number, you can associate it with the messaging service
		// response
		_, err := s.AddPhoneNumberToMessagingService(messagingServiceSID, &messaging.CreatePhoneNumberParams{
			PhoneNumberSid: &params.TwilioPurchasedPhoneNumberSID,
		})

		if err != nil {
			return FullA2POnboardingResponse{}, fmt.Errorf("error at stage 7.3: %w", err)
		}

		return FullA2POnboardingResponse{
			Message: "failed to create A2P campaign: BrandRegistration status is not approved yet .",
			data:    &params,
		}, nil
	}

}
